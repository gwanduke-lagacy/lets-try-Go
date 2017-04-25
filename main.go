package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

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
