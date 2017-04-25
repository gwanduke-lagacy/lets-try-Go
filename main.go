package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Person 사람
type Person struct {
	Name   string
	Age    int
	Height int
	Weight float64
}

func main() {
	// CSV
	ReadCSV()
	WriteCSV()

	// Loggers
	fpLog, err := PrepareFileToLogging("tmp/gwanduke_log.txt")
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	myFileLogger = log.New(fpLog, "MY_FILE_LOGGER: ", log.Ldate|log.Ltime|log.Lshortfile)
	myStdLogger = log.New(os.Stdout, "STANDARD_LOGGER: ", log.LstdFlags)

	// Prepare
	mwp := multiWeatherProvider{
		openWeatherMap{},
		weatherUnderground{apiKey: "964f63783709f6d0"},
	}

	WriteDoubleLogging(fpLog, "라우팅 시작")

	// 라우팅: 요청된 Request Path에 어떤 Request 핸들러를 사용할지 지정
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			myFileLogger.Println("GET 요청이 호출 되었습니다.")
			myStdLogger.Println("GET 요청이 호출 되었습니다. (커스텀 로거)")
			// log.SetFlags(0)
			log.Println("GET 요청이 호출 되었습니다. (표준 로거)")

			person := Person{"gwanduke", 28, 182, 79}
			jsonBytes, err := json.Marshal(person)

			if err != nil {
				panic(err)
			}

			jsonString := string(jsonBytes)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(jsonString))
		case "POST":
			var person Person
			jsonString := []byte("{\"name\":\"gwanduke\", \"age\":28, \"height\": 182, \"weight\": 73}")
			err := json.Unmarshal(jsonString, &person)

			if err != nil {
				panic(err)
			}

			jsonBytes, err := json.Marshal(person)
			w.Header().Add("Content-Type", "application/text")
			w.Write([]byte(string(jsonBytes)))
		// case "PUT":
		// case "DELETE":
		default:
			w.Write([]byte("Not Found"))
		}
	})

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		temp, err := mwp.temperature(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"city": city,
			"temp": temp,
			"took": time.Since(begin).String(),
		})
	})

	// 테스트 핸들러
	// Handle(path, Handler)
	// type Handler interface {
	//   ServeHTTP(ResponseWriter, *Request)
	// }
	// http.Handle("/", new(testHandler))
	http.Handle("/", http.FileServer(http.Dir("public")))

	WriteDoubleLogging(fpLog, "라우팅 종료")

	// 지정된 포트에 웹서버를 열고 클라이언트 Request를 받아들여 새 Go루틴에 작업 할당
	// ListenAndServe(:포트, ServeMux(default: DefaultServeMux))
	// ServeMux는 기본적으로 HTTP Request Router 혹은 Multiplexor 인데,
	// 개발자가 별도로 지정하여 라우팅 제어 가능하다.
	// 기본값을 사용할 경우 Handle(), HandleFunc() 사용해 라우팅 패턴 추가
	http.ListenAndServe(":8080", nil)
}

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	localPath := "wwwroot" + req.URL.Path
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}

	return contentType
}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}

// http.ResponseWriter: HTTP Response에 내용 작성
// http.Request: 입력된 Request 요청을 검토
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=ec7d4673c2832ea731212b751d9f0896&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}

// {
//     "name": "Tokyo",
//     "coord": {
//         "lon": 139.69,
//         "lat": 35.69
//     },
//     "weather": [
//         {
//             "id": 803,
//             "main": "Clouds",
//             "description": "broken clouds",
//             "icon": "04n"
//         }
//     ],
//     "main": {
//         "temp": 296.69,
//         "pressure": 1014,
//         "humidity": 83,
//         "temp_min": 295.37,
//         "temp_max": 298.15
//     }
// }

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

type weatherProvider interface {
	temperature(city string) (float64, error) // in Kelvin, naturally
}

type openWeatherMap struct{}

func (w openWeatherMap) temperature(city string) (float64, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=ec7d4673c2832ea731212b751d9f0896&q=" + city)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var d struct {
		Main struct {
			Kelvin float64 `json:"temp"`
		} `json:"main"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}

	log.Printf("openWeatherMap: %s: %.2f", city, d.Main.Kelvin)
	return d.Main.Kelvin, nil
}

type weatherUnderground struct {
	apiKey string
}

func (w weatherUnderground) temperature(city string) (float64, error) {
	resp, err := http.Get("http://api.wunderground.com/api/" + w.apiKey + "/conditions/q/" + city + ".json")
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var d struct {
		Observation struct {
			Celsius float64 `json:"temp_c"`
		} `json:"current_observation"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}

	kelvin := d.Observation.Celsius + 273.15
	log.Printf("weatherUnderground: %s: %.2f", city, kelvin)

	return kelvin, nil
}

func temperature(city string, providers ...weatherProvider) (float64, error) {
	sum := 0.0

	for _, provider := range providers {
		k, err := provider.temperature(city)
		if err != nil {
			return 0, err
		}

		sum += k
	}

	return sum / float64(len(providers)), nil
}

type multiWeatherProvider []weatherProvider

func (w multiWeatherProvider) temperature(city string) (float64, error) {
	// 온도를 위한 채널과 에러를 위한 채널을 만든다
	// 각 provider는 오직 하나로 푸쉬 할 것이다.
	temps := make(chan float64, len(w))
	errs := make(chan error, len(w))

	// 각 provider는 익명함수와 함께 고루틴을 일어나게 하기위해
	// 그 함수는 온도 메서드를 호출할 것이며 응답을 포워딩 할 것이다.
	for _, provider := range w {
		go func(p weatherProvider) {
			k, err := p.temperature(city)
			if err != nil {
				errs <- err
				return
			}
			temps <- k
		}(provider)
	}

	sum := 0.0

	// 온도를 수집하거나 각 provider로 부터 에러를 수집한다.
	for i := 0; i < len(w); i++ {
		select {
		case temp := <-temps:
			sum += temp
		case err := <-errs:
			return 0, err
		}
	}

	// 평균을 리턴, 이전과 같이
	return sum / float64(len(w)), nil
}
