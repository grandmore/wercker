package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// CitiesResponse holds our cities array
type CitiesResponse struct {
	Cities []string `json:"cities"`
}

//CityHandler processes http response and returns array
func CityHandler(res http.ResponseWriter, req *http.Request) {

	citiesResponse := &CitiesResponse{

		Cities: []string{
			"San Francisco",
			"Amsterdam",
			"Berlin",
			"New York",
			"Tokyo",
			"Kyoto",
			"Osaka",
			"Nagasaki",
			"Naha",
			"London",
			"Paris",
			"Seoul",
			"Austin",
			"London",
		},
	}
	data, _ := json.MarshalIndent(citiesResponse, "", "  ")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	log.Println("Listening on this host: http://localhost:5001")

	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal("Unable to listen on :5001: ", err)
	}
}
