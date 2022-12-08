package model

type Motel struct {
	Id          int    `json:"id" gorm:"primaryKey; autoIncreament"`
	Name        string `json:"name"`
	GroupId     int    `json:"group_id"`
	Status      int    `json:"status" gorm:"default:1"` //1: Trống 2: Có người thuê nhưng chưa full 3 : full
	MaxSlot     int    `json:"max_slot"`
	Floor       int    `json:"floor"`
	Description string `json:"description"`
	Area        int    `json:"area"`
	RentalPrice int    `json:"rental_price"`
	ElecRate    int    `json:"elec_rate"`
	WaterRate   int    `json:"water_rate"`
	ServiceFee  int    `json:"service_fee"`
	GarbageFee  int    `json:"garbage_fee"`
	ParkingFee  int    `json:"parking_fee"`
	// Images      []int `gorm:"type:integer[]"`
}

func (Motel) TableName() string {
	return "motel"
}

type MotelGroup struct {
	Id        int    `json:"id" gorm:"primaryKey; autoIncreament"`
	OwnerId   string `json:"owner_id"`
	GroupName string `json:"group_name"`
	Address   string `json:"address"`
}

func (MotelGroup) TableName() string {
	return "motel_group"
}
