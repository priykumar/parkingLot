package service

import "fmt"

type ParkingLotService struct {
}

func InitService() ParkingLotService {
	return ParkingLotService{}
}

func (s *ParkingLotService) CreateParkingLot() {
	fmt.Println("Creating Parking lot")
}
