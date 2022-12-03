package model

type Motel struct {
	Id          int
	Name        string
	GroupId     int
	Status      int //1: Trống 2: Có người thuê nhưng chưa full 3 : full
	MaxSlot     int
	Floor       int
	Description string
	Area        string
	RentalPrice int
	ElecRate    int
	WaterRate   int
	ServiceFee  int
	GarbageFee  int
	ParkingFee  int
	Images      []int
}

type MotelGroup struct {
	Id        int
	OwnerId   int
	GroupName string
	Address   string
}
