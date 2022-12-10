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
	initInvoice()
	initRequest()

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

func initInvoice() {
	iv1 := new(model.Invoice)
	iv1.Id = 1
	iv1.MotelId = 1
	iv1.RentalPrice = 4000000
	iv1.ElecRate = 2500
	// iv1.DueDate, _ = time.Parse("2006-01-02", "2022-12-31")
	iv1.DueDate = "2022-12-31"

	iv1.GarbageFee = 100000
	iv1.InvoiceDate = "2022-12"
	iv1.ElecIndexBefore = 100
	iv1.ElecIndexAfter = 180
	iv1.WaterIndexBefore = 100
	iv1.WaterIndexAfter = 120
	iv1.ParkingFee = 120000
	iv1.PayStatus = 0
	iv1.RentalPrice = 4000000
	iv1.ServiceFee = 50000
	iv1.WaterRate = 5000
	DB.Create(iv1)

	iv2 := new(model.Invoice)
	iv2.Id = 2
	iv2.MotelId = 1
	iv2.RentalPrice = 4000000
	iv2.ElecRate = 2500
	// iv2.DueDate, _ = time.Parse("2006-01-02", "2022-11-31")
	iv2.DueDate = "2022-11-31"

	iv2.GarbageFee = 100000
	iv2.InvoiceDate = "2022-11"
	iv2.ElecIndexBefore = 100
	iv2.ElecIndexAfter = 200
	iv2.WaterIndexBefore = 100
	iv2.WaterIndexAfter = 120
	iv2.ParkingFee = 120000
	iv2.PayStatus = 0
	iv2.RentalPrice = 4000000
	iv2.ServiceFee = 50000
	iv2.WaterRate = 5000
	DB.Create(iv2)

}

func initRequest() {
	rq1 := new(model.Request)
	rq1.Creator = "abc@gmail.com"
	rq1.IsFromAdmin = false
	rq1.MotelId = 1
	rq1.Title = "Đăng kí dịch vụ wifi"
	rq1.Content = "Lorem isume"
	rq1.Status = 0
	rq1.RequestType = 2 //Dịch vụ
	DB.Create(rq1)

	rq2 := new(model.Request)
	rq2.Creator = "admin@gmail.com"
	rq2.IsFromAdmin = true
	rq2.MotelId = 1
	rq2.Title = "Thông báo thay đổi phương thức thanh toán"
	rq2.Content = "Lorem isume"
	rq2.Status = 0
	rq2.RequestType = 0 //Thông báo từ admin tới user
	DB.Create(rq2)
}
