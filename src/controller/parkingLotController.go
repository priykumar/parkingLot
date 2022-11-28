package controller

import (
	dtos "github.com/priykumar/parkingLot/src/dto"
	service "github.com/priykumar/parkingLot/src/service"
)

type ParkingLotController struct {
	Service service.ParkingLotService
}

func InitController() ParkingLotController {
	return ParkingLotController{
		Service: service.InitService(),
	}
}

func (c *ParkingLotController) CreateParkingLot(request dtos.ParkingLotRequest) {
	c.Service = service.ParkingLotService{}

}
