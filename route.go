package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const weatherUndergroundAPIKey string = "964f63783709f6d0"

// Person 사람
type Person struct {
	Name   string
	Age    int
	Height int
	Weight float64
}

// StartRoute 라우팅을 시작합니다.
func StartRoute() {
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
		weatherUnderground{apiKey: weatherUndergroundAPIKey},
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
