package models

import "time"

type VehicleType string
type GateType string
type GateStatus string
type ParkingSpotStatus string
type InvoicePaidStatus string
type PaymentStatus string
type PaymentMode string

const (
	Small    VehicleType = "SMALL"
	Medium   VehicleType = "MEDIUM"
	Large    VehicleType = "LARGE"
	Electric VehicleType = "ELECTRIC"
)

const (
	Entry GateType = "ENTRY"
	Exit  GateType = "EXIT"
)

const (
	Open   GateStatus = "OPEN"
	Closed GateStatus = "CLOSED"
)

const (
	Available   ParkingSpotStatus = "AVAILABLE"
	Unavailable ParkingSpotStatus = "UNAVAILABLE"
)

const (
	Paid   InvoicePaidStatus = "PAID"
	UnPaid InvoicePaidStatus = "UNPAID"
)

const (
	Success PaymentStatus = "SUCCESS"
	Pending PaymentStatus = "PENDING"
	Failure PaymentStatus = "FAILURE"
)

const (
	Cash       PaymentMode = "CASH"
	DebitCard  PaymentMode = "DEBITE_CARD"
	CreditCard PaymentMode = "CREDIT_CARD"
	NetBanking PaymentMode = "NET_BANKING"
	UPI        PaymentMode = "UPI"
)

type BaseClass struct {
	Id int64
}

type ParkingLot struct {
	Id                  BaseClass
	Address             string
	Gates               []Gate
	Floors              []ParkingFloor
	VehicleTypePriceMap map[VehicleType]int
}

type Gate struct {
	Id       BaseClass
	Number   int
	Type     GateType
	Status   GateStatus
	Operator Operator
}

type ParkingFloor struct {
	Id          BaseClass
	Spots       []ParkingSpot
	FloorNumber int
}

type ParkingSpot struct {
	Id          BaseClass
	SpotNumber  int
	VehicleType VehicleType
	Status      ParkingSpotStatus
}

type ElectricParkingSpot struct {
	ParkingSpot     ParkingSpot
	ElectricCharger ElectricCharger
}

type ElectricCharger struct {
	Id          BaseClass
	Consumption int
}

type Operator struct {
	Id   BaseClass
	Name string
}

type Invoice struct {
	Id                BaseClass
	Ticket            Ticket
	Operator          Operator
	Amount            int
	ExitTime          time.Time
	InvoicePaidStatus InvoicePaidStatus
	Payment           []Payment
}

type Payment struct {
	Id            BaseClass
	RefNumber     int
	PaymentTime   time.Time
	Amount        int
	Invoice       Invoice
	PaymentStatus PaymentStatus
	PaymentMode   PaymentMode
}

type Ticket struct {
	Id          BaseClass
	EntryTime   time.Time
	EntryGate   Gate
	GeneratedBy Operator
	Vehicle     Vehicle
	ParkingSpot ParkingSpot
	Parkinglot  ParkingLot
}

type Vehicle struct {
	Id            BaseClass
	VehicleNumber string
	VehicleType   VehicleType
}
