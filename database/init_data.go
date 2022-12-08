package database

import (
	// "github.com/lib/pq"
	"TerraInnAPI/model"

	"gorm.io/gorm"
)

func init_data(DB *gorm.DB) {
	initUser()
	initMotel()
	initMotelGroup()
}

func initUser() {
	user1 := new(model.User)
	user1.Email = "abc@gmail.com"
	user1.Phone = "0123456789"
	user1.Password = "827ccb0eea8a706c4c34a16891f84e7b" //12345
	user1.FullName = "Nguyễn Văn A"
	user1.DateOfBirth = "2001-06-13"
	user1.CidNumber = "2324322532"
	user1.IsAdmin = false

	DB.Create(user1)
}

func initMotel() {
	motel1 := new(model.Motel)
	motel1.Area = 40
	motel1.Description = " Lorem isum"
	motel1.ElecRate = 2500
	motel1.Floor = 1
	motel1.GarbageFee = 100000
	motel1.GroupId = 1
	motel1.MaxSlot = 4
	motel1.Name = "H1-711"
	motel1.ParkingFee = 120000
	motel1.Status = 2
	motel1.RentalPrice = 4000000
	motel1.ServiceFee = 50000
	motel1.WaterRate = 5000

	DB.Create(motel1)
}

func initMotelGroup() {
	gr1 := new(model.MotelGroup)
	gr1.Address = "44 Khuc Thua Du, Ben Nghe, Quan 1"
	gr1.OwnerId = "admin@gmail.com"
	gr1.GroupName = "Khu trọ An Khang"

	DB.Create(gr1)
}
