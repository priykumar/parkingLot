package service

import (
	"fmt"

	"github.com/priykumar/parkingLot/src/model"
	"github.com/priykumar/parkingLot/src/repository"
)

type ParkingLotService struct {
	Repo *repository.ParkingLotRepository
}

func InitService() *ParkingLotService {
	return &ParkingLotService{
		Repo: repository.InitialiseDatabase(),
	}
}

func (s *ParkingLotService) CreateParkingLot(address string) model.ParkingLot {
	fmt.Println("Creating Parking lot")
	//s.Repo.CreateNewParkingLot(address)
	return model.ParkingLot{
		Address: address,
	}
}
