package model

import "time"

type Invoice struct {
	Id               int       `gorm:"primaryKey;autoIncreament" json:"id"`
	MotelId          int       `json:"motel_id"`
	InvoiceDate      string    `json:"invoice_date"` //Hoá đơn này cho thangs napf, format: 2022-12-01
	RentalPrice      int       `json:"rental_price"`
	ElecRate         int       `json:"elec_rate"`
	WaterRate        int       `json:"water_rate"`
	ServiceFee       int       `json:"service_fee"`
	GarbageFee       int       `json:"garbage_fee"`
	ParkingFee       int       `json:"parking_fee"`
	ElecIndexBefore  int       `json:"elec_index_before"`
	ElecIndexAfter   int       `json:"elec_index_after"`
	WaterIndexBefore int       `json:"water_index_before"`
	WaterIndexAfter  int       `json:"water_index_after"`
	DueDate          string    `json:"due_date"`
	PayStatus        int       `json:"pay_status"` //0 là chưa tt. 1 là đã tt 2 là quá hạn.
	CreateAt         time.Time `json:"create_at" gorm:"autoCreateTime"`
	Note             string    `json:"note"`
}

func (Invoice) TableName() string {
	return "invoices"
}
