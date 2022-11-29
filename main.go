package main

import (
	"fmt"
	"net/http"

	mux "github.com/gorilla/mux"
	"github.com/priykumar/parkingLot/src/controller"
)

func main() {
	c := controller.InitController()

	router := mux.NewRouter()
	router.HandleFunc("/createParkingLot", c.CreateParkingLot).Methods("POST")

	done := make(chan bool)
	go http.ListenAndServe(":6000", router)
	fmt.Println("Listening to [:6000]")
	<-done
}
