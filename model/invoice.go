package model

import "time"

type Invoice struct {
	Id              int
	MotelId         int
	InvoiceDate     time.Time //Hoá đơn này cho ngày nào
	RentalPrice     int
	ElecRate        int
	WaterRate       int
	ServiceFee      int
	GarbageFee      int
	ParkingFee      int
	ElectCountStart int
	ElectCountEnd   int
	WaterCountStart int
	WaterCountEnd   int
	DueDate         time.Time
	PayStatus       int //0,1
	CreateAt        time.Time
	Note            string
}
