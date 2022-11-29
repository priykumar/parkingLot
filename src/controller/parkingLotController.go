package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/priykumar/parkingLot/src/dto"
	"github.com/priykumar/parkingLot/src/service"
)

type ParkingLotController struct {
	Service *service.ParkingLotService
}

func InitController() *ParkingLotController {
	return &ParkingLotController{
		Service: service.InitService(),
	}
}

// func SetParkingLotRequest(w http.ResponseWriter, r *http.Request) {
// 	var newParkingLot dto.ParkingLotRequest
// 	w.Header().Set("Content-Type", "Application/json")
// 	json.NewDecoder(r.Body).Decode(&newParkingLot)

// 	CreateParkingLot(newParkingLot)

// }
func (c *ParkingLotController) CreateParkingLot(w http.ResponseWriter, r *http.Request) {
	var newParkingLot dto.ParkingLotRequest
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&newParkingLot)

	p := c.Service.CreateParkingLot(newParkingLot.Address)
	fmt.Printf("%#v", p)
}
