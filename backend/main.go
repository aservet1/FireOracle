package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ErrorStruct struct {
	ErrorStr string `json:"error"`
}

type FireForecastStruct struct {
	Fireforecast []bool `json:"fireforecast"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Entered")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	county := vars["county"]
	forecast, err := getFireForecast(county)
	if err != nil {
		var e ErrorStruct
		fmt.Println(err)
		e.ErrorStr = "Something went wrong"
		json.NewEncoder(w).Encode(e)
		return
	}
	var f FireForecastStruct
	f.Fireforecast = forecast
	json.NewEncoder(w).Encode(f)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/forestFireForecast/{county}", homePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
