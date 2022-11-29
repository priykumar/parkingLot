package dto

type ParkingLotRequest struct {
	Address string
}

type ParkingLotResponse struct {
}

func SetParkingLotRequest1(address string) ParkingLotRequest {
	return ParkingLotRequest{
		Address: address,
	}
}
