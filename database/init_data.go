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
	//-- CREATE NORMAL USER
	user1 := new(model.User)
	user1.Email = "anguyen@gmail.com"
	user1.Phone = "11111111"
	user1.Password = "827ccb0eea8a706c4c34a16891f84e7b" //12345
	user1.FullName = "Nguyễn Văn A"
	user1.DateOfBirth = "1997-06-13"
	user1.CidNumber = "111111111"
	user1.IsAdmin = false
	user1.AvatarUrl = "img_16710903510.jpg"
	user1.MotelId = 1
	DB.Create(user1)

	user2 := new(model.User)
	user2.Email = "huanhoahoe@gmail.com"
	user2.Phone = "22222222"
	user2.Password = "827ccb0eea8a706c4c34a16891f84e7b" //12345
	user2.FullName = "Huấn Hoa Hoè"
	user2.DateOfBirth = "2001-08-21"
	user2.CidNumber = "22222222"
	user2.IsAdmin = false
	user2.AvatarUrl = "img_167109035002.png"
	user2.MotelId = 2
	DB.Create(user2)

	user3 := new(model.User)
	user3.Email = "chubathong@gmail.com"
	user3.Phone = "33333333"
	user3.Password = "827ccb0eea8a706c4c34a16891f84e7b" //12345
	user3.FullName = "Chu Bá Thông"
	user3.DateOfBirth = "1997-08-21"
	user3.CidNumber = "33333333"
	user3.IsAdmin = false
	user3.AvatarUrl = "img_167109035009.jpg"
	user3.MotelId = 2
	DB.Create(user3)

	user4 := new(model.User)
	user4.Email = "lane@gmail.com"
	user4.Phone = "44444444"
	user4.Password = "827ccb0eea8a706c4c34a16891f84e7b" //12345
	user4.FullName = "Lan Đức Ê"
	user4.DateOfBirth = "1990-02-11"
	user4.CidNumber = "44444444"
	user4.IsAdmin = false
	user4.AvatarUrl = "img_167109035006.jpg"
	// user4.MotelId = 1
	DB.Create(user4)

	user5 := new(model.User)
	user5.Email = "hien@gmail.com"
	user5.Phone = "55555555"
	user5.Password = "827ccb0eea8a706c4c34a16891f84e7b" //12345
	user5.FullName = "Nguyễn Ngọc Hiển"
	user5.DateOfBirth = "1990-02-11"
	user5.CidNumber = "55555555"
	user5.IsAdmin = false
	user5.AvatarUrl = "img_167109035007.jpg"

	DB.Create(user5)

	//-- CREATE ADMIN
	user6 := new(model.User)
	user6.Email = "admin1@gmail.com"
	user6.Phone = "66666666"
	user6.Password = "827ccb0eea8a706c4c34a16891f84e7b" //12345
	user6.FullName = "Admin 1"
	user6.DateOfBirth = "1980-04-11"
	user6.CidNumber = "66666666"
	user6.IsAdmin = true
	user6.AvatarUrl = "img_167109035008.jpg"

	DB.Create(user6)

	user7 := new(model.User)
	user7.Email = "admin2@gmail.com"
	user7.Phone = "77777777"
	user7.Password = "827ccb0eea8a706c4c34a16891f84e7b" //12345
	user7.FullName = "Admin 2"
	user7.DateOfBirth = "1991-01-25"
	user7.CidNumber = "77777777"
	user7.IsAdmin = true
	user7.AvatarUrl = "no-image.jpg"

	DB.Create(user7)
}

func initMotel() {
	motel1 := new(model.Motel)
	motel1.Area = 40
	motel1.Description = `* Cách trung tâm Q1, Q2, Q3, Phú Nhuận 5 - 10 phút di chuyển (đường rộng, không bị kẹt xe, không ngập nước). Đường nội bộ xe tải, đối diện chung cư Wilton Tower, luôn mát mẻ. Khu dân cư an ninh, yên tĩnh, dân trí cao. Trước cửa nhà là cửa hàng tiện lợi 24h GS 25, Circle K, nhà thuốc Pharmacity. 200m tới Pearl Plaza - siêu thị Coopmart Extra, Bách Hóa Xanh, 1km tới Vinhomes Landmark 81. Xung quanh toàn quán cà phê (The Coffee House, Highland, Starbuck, KOI The... ), ăn uống. Nhiều phòng tập gym xung quanh: California Fitness and Yoga,... Thông qua đường D2, Ung Văn Khiêm đến các trường Đại học: Hutech, Giao thông vận tải, Ngoại thương, Hồng Bàng,`
	motel1.ElecRate = 2500
	motel1.Floor = 1
	motel1.GarbageFee = 100000
	motel1.GroupId = 1
	motel1.MaxSlot = 4
	motel1.Name = "PEARL PLAZA - D1"
	motel1.ParkingFee = 120000
	motel1.Status = 2
	motel1.RentalPrice = 6000000
	motel1.ServiceFee = 50000
	motel1.WaterRate = 5000
	motel1.ShortDesc = "1,2,4,5,6"
	motel1.Images = "img_167109035008.jpg,img_167109035007.jpg,img_167109035006.jpg"

	DB.Create(motel1)

	motel2 := new(model.Motel)
	motel2.Area = 40
	motel2.Description = `Bcons Garden tọa lạc ngay vị trí đắc địa trung tâm hành chính của thị xã Dĩ An, Bình Dương. Thuận tiện di chuyển đến các địa điểm: Liền kề trung tâm hành chính Dĩ An; Cách Thủ Đức, HCM: 2 km. Ga Bình Chiểu: 3 km. Khu du lịch Suối Tiên, khu công nghệ cao TP. HCM: 7 km. Quận 2: 10 km. Ngã ba Vũng Tàu: 10km.`
	motel2.ElecRate = 4500
	motel2.Floor = 2
	motel2.GarbageFee = 120000
	motel2.GroupId = 2
	motel2.MaxSlot = 3
	motel2.Name = "BCONS-211"
	motel2.ParkingFee = 140000
	motel2.Status = 2
	motel2.RentalPrice = 4000000
	motel2.ServiceFee = 150000
	motel2.WaterRate = 3000
	motel2.ShortDesc = "1,2,3,4,5,6"
	motel2.Images = "img_167109035011.jpg,img_167109035008.jpg,img_167109035007.jpg"

	DB.Create(motel2)

}

func initMotelGroup() {
	gr1 := new(model.MotelGroup)
	gr1.Address = " 71/1/ Nguyễn Văn Thương (D1 Cũ), P25, Q Bình Thạnh."
	gr1.OwnerId = "66666666"
	gr1.GroupName = "PEARL PLAZA - BÌNH THẠNH"
	DB.Create(gr1)

	gr2 := new(model.MotelGroup)
	gr2.Address = "Khu phố 6, Đông Hoà, Dĩ An, Bình Dương"
	gr2.OwnerId = "77777777"
	gr2.GroupName = "Bcons Garden"
	DB.Create(gr2)
}

func initInvoice() {
	iv1 := new(model.Invoice)
	iv1.Id = 1
	iv1.MotelId = 1
	iv1.RentalPrice = 6000000
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
	iv2.MotelId = 2
	iv2.RentalPrice = 4000000
	iv2.ElecRate = 4500
	// iv2.DueDate, _ = time.Parse("2006-01-02", "2022-11-31")
	iv2.DueDate = "2022-11-31"

	iv2.GarbageFee = 120000
	iv2.InvoiceDate = "2022-11"
	iv2.ElecIndexBefore = 100
	iv2.ElecIndexAfter = 200
	iv2.WaterIndexBefore = 100
	iv2.WaterIndexAfter = 120
	iv2.ParkingFee = 140000
	iv2.PayStatus = 1
	iv2.RentalPrice = 4000000
	iv2.ServiceFee = 150000
	iv2.WaterRate = 3000
	DB.Create(iv2)

}

func initRequest() {
	rq1 := new(model.Request)
	rq1.Creator = "11111111"
	rq1.IsFromAdmin = false
	rq1.MotelId = 1
	rq1.Title = "Đăng kí dịch vụ wifi"
	rq1.Content = "Cho phòng e đk dịch vụ Wifi"
	rq1.Status = 0
	rq1.RequestType = 3 //Dịch vụ
	DB.Create(rq1)

	rq2 := new(model.Request)
	rq2.Creator = "77777777"
	rq2.IsFromAdmin = true
	rq2.MotelId = 1
	rq2.Title = "Thông báo thay đổi phương thức thanh toán"
	rq2.Content = "Thay đổi phương thức thanh toán"
	rq2.Status = 0
	rq2.RequestType = 0 //Thông báo từ admin tới user
	DB.Create(rq2)

	rq3 := new(model.Request)
	rq3.Creator = "22222222"
	rq3.IsFromAdmin = false
	rq3.MotelId = 2
	rq3.Title = "Thuê phòng"
	rq3.Content = "Thuê phòng"
	rq3.Status = 0
	rq3.RequestType = 1 //ĐK thuê phòng
	DB.Create(rq3)
}
