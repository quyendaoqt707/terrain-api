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
	user0 := new(model.User)
	user0.Email = "00000000@gmail.com"
	user0.Phone = "11111111"
	user0.Password = "827ccb0eea8a706c4c34a16891f84e7b" //12345
	user0.FullName = "Lê Văn Tám"
	user0.DateOfBirth = "2001-06-13"
	user0.CidNumber = "00000000"
	user0.IsAdmin = false
	user0.AvatarUrl = "img_16710903510.jpg"
	user0.MotelId = 1
	DB.Create(user0)

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

	user8 := new(model.User)
	user8.Email = "hongvan@gmail.com"
	user8.Phone = "88888888"
	user8.Password = "827ccb0eea8a706c4c34a16891f84e7b" //12345
	user8.FullName = "Nguyễn Thị Hồng Vân"
	user8.DateOfBirth = "1988-07-20"
	user8.CidNumber = "88888888"
	user8.IsAdmin = true
	user8.AvatarUrl = "no-image.jpg"
	DB.Create(user7)
}

func initMotel() {
	//GROUP 1 có 3 phòng trọ
	motel0 := new(model.Motel)
	motel0.Id = 1
	motel0.Area = 40
	motel0.Description = `* Cách trung tâm Q1, Q2, Q3, Phú Nhuận 5 - 10 phút di chuyển (đường rộng, không bị kẹt xe, không ngập nước). Đường nội bộ xe tải, đối diện chung cư Wilton Tower, luôn mát mẻ. Khu dân cư an ninh, yên tĩnh, dân trí cao. Trước cửa nhà là cửa hàng tiện lợi 24h GS 25, Circle K, nhà thuốc Pharmacity. 200m tới Pearl Plaza - siêu thị Coopmart Extra, Bách Hóa Xanh, 1km tới Vinhomes Landmark 81. Xung quanh toàn quán cà phê (The Coffee House, Highland, Starbuck, KOI The... ), ăn uống. Nhiều phòng tập gym xung quanh: California Fitness and Yoga,... Thông qua đường D2, Ung Văn Khiêm đến các trường Đại học: Hutech, Giao thông vận tải, Ngoại thương, Hồng Bàng,`
	motel0.ElecRate = 2500
	motel0.Floor = 1
	motel0.GarbageFee = 100000
	motel0.GroupId = 1
	motel0.MaxSlot = 4
	motel0.Name = "PEARL PLAZA - D1"
	motel0.ParkingFee = 120000
	motel0.Status = 2
	motel0.RentalPrice = 6000000
	motel0.ServiceFee = 50000
	motel0.WaterRate = 5000
	motel0.ShortDesc = "1,2,4,5,6"
	motel0.Images = "4.jpg,img_167109035007.jpg,5.jpg"
	DB.Create(motel0)

	motel1 := new(model.Motel)
	motel1.Id = 2
	motel1.Area = 40
	motel1.Description = `* Cách trung tâm Q1, Q2, Q3, Phú Nhuận 5 - 10 phút di chuyển (đường rộng, không bị kẹt xe, không ngập nước). Đường nội bộ xe tải, đối diện chung cư Wilton Tower, luôn mát mẻ. Khu dân cư an ninh, yên tĩnh, dân trí cao. Trước cửa nhà là cửa hàng tiện lợi 24h GS 25, Circle K, nhà thuốc Pharmacity. 200m tới Pearl Plaza - siêu thị Coopmart Extra, Bách Hóa Xanh, 1km tới Vinhomes Landmark 81. Xung quanh toàn quán cà phê (The Coffee House, Highland, Starbuck, KOI The... ), ăn uống. Nhiều phòng tập gym xung quanh: California Fitness and Yoga,... Thông qua đường D2, Ung Văn Khiêm đến các trường Đại học: Hutech, Giao thông vận tải, Ngoại thương, Hồng Bàng,`
	motel1.ElecRate = 2500
	motel1.Floor = 2
	motel1.GarbageFee = 100000
	motel1.GroupId = 1
	motel1.MaxSlot = 4
	motel1.Name = "PEARL PLAZA - D2"
	motel1.ParkingFee = 120000
	motel1.Status = 2
	motel1.RentalPrice = 5000000
	motel1.ServiceFee = 50000
	motel1.WaterRate = 5000
	motel1.ShortDesc = "1,2,4,5,6"
	motel1.Images = "img_167109035008.jpg,img_167109035007.jpg,img_167109035006.jpg"
	DB.Create(motel1)

	motel11 := new(model.Motel)
	motel11.Area = 40
	motel11.Id = 3
	motel11.Description = `* Cách trung tâm Q1, Q2, Q3, Phú Nhuận 5 - 10 phút di chuyển (đường rộng, không bị kẹt xe, không ngập nước). Đường nội bộ xe tải, đối diện chung cư Wilton Tower, luôn mát mẻ. Khu dân cư an ninh, yên tĩnh, dân trí cao. Trước cửa nhà là cửa hàng tiện lợi 24h GS 25, Circle K, nhà thuốc Pharmacity. 200m tới Pearl Plaza - siêu thị Coopmart Extra, Bách Hóa Xanh, 1km tới Vinhomes Landmark 81. Xung quanh toàn quán cà phê (The Coffee House, Highland, Starbuck, KOI The... ), ăn uống. Nhiều phòng tập gym xung quanh: California Fitness and Yoga,... Thông qua đường D2, Ung Văn Khiêm đến các trường Đại học: Hutech, Giao thông vận tải, Ngoại thương, Hồng Bàng,`
	motel11.ElecRate = 2500
	motel11.Floor = 3
	motel11.GarbageFee = 100000
	motel11.GroupId = 1
	motel11.MaxSlot = 4
	motel11.Name = "PEARL PLAZA - D3"
	motel11.ParkingFee = 120000
	motel11.Status = 2
	motel11.RentalPrice = 4000000
	motel11.ServiceFee = 50000
	motel11.WaterRate = 5000
	motel11.ShortDesc = "1,,5,6"
	motel11.Images = "img_167109035008.jpg,img_167109035007.jpg,img_167109035006.jpg"
	DB.Create(motel11)

	motel111 := new(model.Motel)
	motel111.Area = 40
	motel111.Description = `* Cách trung tâm Q1, Q2, Q3, Phú Nhuận 5 - 10 phút di chuyển (đường rộng, không bị kẹt xe, không ngập nước). Đường nội bộ xe tải, đối diện chung cư Wilton Tower, luôn mát mẻ. Khu dân cư an ninh, yên tĩnh, dân trí cao. Trước cửa nhà là cửa hàng tiện lợi 24h GS 25, Circle K, nhà thuốc Pharmacity. 200m tới Pearl Plaza - siêu thị Coopmart Extra, Bách Hóa Xanh, 1km tới Vinhomes Landmark 81. Xung quanh toàn quán cà phê (The Coffee House, Highland, Starbuck, KOI The... ), ăn uống. Nhiều phòng tập gym xung quanh: California Fitness and Yoga,... Thông qua đường D2, Ung Văn Khiêm đến các trường Đại học: Hutech, Giao thông vận tải, Ngoại thương, Hồng Bàng,`
	motel111.ElecRate = 2500
	motel111.Floor = 4
	motel111.GarbageFee = 100000
	motel111.GroupId = 1
	motel111.MaxSlot = 4
	motel111.Name = "PEARL PLAZA - D4"
	motel111.ParkingFee = 120000
	motel111.Status = 2
	motel111.RentalPrice = 4000000
	motel111.ServiceFee = 50000
	motel111.WaterRate = 5000
	motel111.ShortDesc = "1,,5,6"
	motel111.Images = "1.jpg,2.jpg,3.jpg"
	DB.Create(motel111)

	//GROUP 2:

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

	motel3 := new(model.Motel)
	motel3.Area = 35
	motel3.Description = `Cho thuê căn hộ dịch vụ siêu mới, rộng rãi và thoáng mát ngay trung tâm quận Bình Thạnh, cách quận 1 chỉ 3 phút.
	Đặc biệt: Giảm ngay 500k cho tháng đầu tiên nếu ký HĐ từ 1 - 15/12.
	Giá chỉ từ 5,8tr - 6,8tr.
	ĐC: 59/4 Đinh Bộ Lĩnh, P26, Bình Thạnh.
	Liên hệ xem nhà 24/24h Ms Vân chính chủ.
	0903.396.059 zalo/call.`
	motel3.ElecRate = 2500
	motel3.Floor = 1
	motel3.GarbageFee = 100000
	motel3.GroupId = 3
	motel3.MaxSlot = 4
	motel3.Name = "Hồng Vân - P1"
	motel3.ParkingFee = 120000
	motel3.Status = 2
	motel3.RentalPrice = 5800000
	motel3.ServiceFee = 50000
	motel3.WaterRate = 5000
	motel3.ShortDesc = "1,2,4,5,6"
	motel3.Images = "1.jpg,2.jpg,3.jpg"
	DB.Create(motel3)

	motel32 := new(model.Motel)
	motel32.Area = 35
	motel32.Description = `Cho thuê căn hộ dịch vụ siêu mới, rộng rãi và thoáng mát ngay trung tâm quận Bình Thạnh, cách quận 1 chỉ 3 phút.
	Đặc biệt: Giảm ngay 500k cho tháng đầu tiên nếu ký HĐ từ 1 - 15/12.
	Giá chỉ từ 5,8tr - 6,8tr.
	ĐC: 59/4 Đinh Bộ Lĩnh, P26, Bình Thạnh.
	Liên hệ xem nhà 24/24h Ms Vân chính chủ.
	0903.396.059 zalo/call.`
	motel32.ElecRate = 2500
	motel32.Floor = 2
	motel32.GarbageFee = 100000
	motel32.GroupId = 3
	motel32.MaxSlot = 2
	motel32.Name = "Hồng Vân - P2"
	motel32.ParkingFee = 120000
	motel32.Status = 2
	motel32.RentalPrice = 5800000
	motel32.ServiceFee = 50000
	motel32.WaterRate = 5000
	motel32.ShortDesc = "1,2,4,5,6"
	motel32.Images = "img_167109035011.jpg,5.jpg,3.jpg"
	DB.Create(motel3)

	motel33 := new(model.Motel)
	motel33.Area = 35
	motel33.Description = `Cho thuê căn hộ dịch vụ siêu mới, rộng rãi và thoáng mát ngay trung tâm quận Bình Thạnh, cách quận 1 chỉ 3 phút.
	Đặc biệt: Giảm ngay 500k cho tháng đầu tiên nếu ký HĐ từ 1 - 15/12.
	Giá chỉ từ 5,8tr - 6,8tr.
	ĐC: 59/4 Đinh Bộ Lĩnh, P26, Bình Thạnh.
	Liên hệ xem nhà 24/24h Ms Vân chính chủ.
	0903.396.059 zalo/call.`
	motel33.ElecRate = 2500
	motel33.Floor = 3
	motel33.GarbageFee = 100000
	motel33.GroupId = 3
	motel33.MaxSlot = 4
	motel33.Name = "Hồng Vân - P3"
	motel33.ParkingFee = 120000
	motel33.Status = 2
	motel33.RentalPrice = 5800000
	motel33.ServiceFee = 50000
	motel33.WaterRate = 5000
	motel33.ShortDesc = "1,2,4,5,6"
	motel33.Images = "4.jpg,5.jpg,3.jpg"
	DB.Create(motel3)

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

	gr3 := new(model.MotelGroup)
	gr3.Address = "59 ĐINH BỘ LĨNH, BÌNH THẠNH"
	gr3.OwnerId = "88888888"
	gr3.GroupName = "Hồng Vân Motel"
	DB.Create(gr3)
}

func initInvoice() {

	// Hoá đơn của phòng group 1
	iv1 := new(model.Invoice)
	// iv1.Id = 1
	iv1.MotelId = 1
	iv1.RentalPrice = 6000000
	iv1.ElecRate = 2500
	// iv1.DueDate, _ = time.Parse("2006-01-02", "2022-12-31")
	iv1.DueDate = "2022-12-31"
	iv1.GarbageFee = 100000
	iv1.InvoiceDate = "2022-12"
	iv1.ElecIndexBefore = 1100
	iv1.ElecIndexAfter = 1180
	iv1.WaterIndexBefore = 100
	iv1.WaterIndexAfter = 120
	iv1.ParkingFee = 120000
	iv1.PayStatus = 0 //Chưa thanh toán
	iv1.RentalPrice = 4000000
	iv1.ServiceFee = 50000
	iv1.WaterRate = 5000
	DB.Create(iv1)

	iv0 := new(model.Invoice)
	// iv1.Id = 1
	iv0.MotelId = 1
	iv0.RentalPrice = 6000000
	iv0.ElecRate = 2500
	// iv1.DueDate, _ = time.Parse("2006-01-02", "2022-12-31")
	iv0.DueDate = "2022-11-31"
	iv0.GarbageFee = 100000
	iv0.InvoiceDate = "2022-11"
	iv0.ElecIndexBefore = 1104
	iv0.ElecIndexAfter = 1197
	iv0.WaterIndexBefore = 86
	iv0.WaterIndexAfter = 129
	iv0.ParkingFee = 120000
	iv0.PayStatus = 1 //Đã thanh toán
	iv0.RentalPrice = 4000000
	iv0.ServiceFee = 50000
	iv0.WaterRate = 5000
	DB.Create(iv0)

	iv00 := new(model.Invoice)
	// iv1.Id = 1
	iv00.MotelId = 2
	iv00.RentalPrice = 6000000
	iv00.ElecRate = 2500
	// iv1.DueDate, _ = time.Parse("2006-01-02", "2022-12-31")
	iv00.DueDate = "2022-12-31"
	iv00.GarbageFee = 100000
	iv00.InvoiceDate = "2022-12"
	iv00.ElecIndexBefore = 864
	iv00.ElecIndexAfter = 1097
	iv00.WaterIndexBefore = 86
	iv00.WaterIndexAfter = 129
	iv00.ParkingFee = 120000
	iv00.PayStatus = 1 //Đã thanh toán
	iv00.RentalPrice = 4000000
	iv00.ServiceFee = 50000
	iv00.WaterRate = 5000
	DB.Create(iv00)

	iv11 := new(model.Invoice)
	// iv1.Id = 1
	iv11.MotelId = 3
	iv11.RentalPrice = 6000000
	iv11.ElecRate = 2500
	// iv1.DueDate, _ = time.Parse("2006-01-02", "2022-12-31")
	iv11.DueDate = "2022-12-31"
	iv11.GarbageFee = 100000
	iv11.InvoiceDate = "2022-12"
	iv11.ElecIndexBefore = 864
	iv11.ElecIndexAfter = 1097
	iv11.WaterIndexBefore = 86
	iv11.WaterIndexAfter = 129
	iv11.ParkingFee = 120000
	iv11.PayStatus = 1 //Đã thanh toán
	iv11.RentalPrice = 4000000
	iv11.ServiceFee = 50000
	iv11.WaterRate = 5000
	DB.Create(iv11)

	iv12 := new(model.Invoice)
	// iv1.Id = 1
	iv12.MotelId = 4
	iv12.RentalPrice = 6000000
	iv12.ElecRate = 2500
	// iv1.DueDate, _ = time.Parse("2006-01-02", "2022-12-31")
	iv12.DueDate = "2022-12-31"
	iv12.GarbageFee = 100000
	iv12.InvoiceDate = "2022-12"
	iv12.ElecIndexBefore = 864
	iv12.ElecIndexAfter = 1097
	iv12.WaterIndexBefore = 86
	iv12.WaterIndexAfter = 129
	iv12.ParkingFee = 120000
	iv12.PayStatus = 1 //Đã thanh toán
	iv12.RentalPrice = 4000000
	iv12.ServiceFee = 50000
	iv12.WaterRate = 5000
	DB.Create(iv12)

	iv13 := new(model.Invoice)
	// iv1.Id = 1
	iv13.MotelId = 5
	iv13.RentalPrice = 6000000
	iv13.ElecRate = 2500
	// iv1.DueDate, _ = time.Parse("2006-01-02", "2022-12-31")
	iv13.DueDate = "2022-12-31"
	iv13.GarbageFee = 100000
	iv13.InvoiceDate = "2022-12"
	iv13.ElecIndexBefore = 864
	iv13.ElecIndexAfter = 1097
	iv13.WaterIndexBefore = 86
	iv13.WaterIndexAfter = 129
	iv13.ParkingFee = 120000
	iv13.PayStatus = 1 //Đã thanh toán
	iv13.RentalPrice = 4000000
	iv13.ServiceFee = 50000
	iv13.WaterRate = 5000
	DB.Create(iv12)

	// Hoá đơn của phòng 2 tháng 11/2022
	iv2 := new(model.Invoice)
	// iv2.Id = 2
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

	//NOTE: Tháng 12
	iv3 := new(model.Invoice)
	// iv3.Id = 2
	iv3.MotelId = 4
	iv3.RentalPrice = 4000000
	iv3.ElecRate = 4500
	// iv2.DueDate, _ = time.Parse("2006-01-02", "2022-11-31")
	iv3.DueDate = "2022-12-31"
	iv3.GarbageFee = 120000
	iv3.InvoiceDate = "2022-12"
	iv3.ElecIndexBefore = 1231
	iv3.ElecIndexAfter = 1345
	iv3.WaterIndexBefore = 2324
	iv3.WaterIndexAfter = 2455
	iv3.ParkingFee = 140000
	iv3.PayStatus = 0
	iv3.RentalPrice = 4000000
	iv3.ServiceFee = 150000
	iv3.WaterRate = 3000
	DB.Create(iv3)

}

func initRequest() {
	rq1 := new(model.Request)
	rq1.Creator = "11111111"
	rq1.IsFromAdmin = false
	rq1.MotelId = 1
	rq1.Title = "Đăng kí dịch vụ wifi (11111111)"
	rq1.Content = "Cho phòng e đk dịch vụ Wifi"
	rq1.Status = 0
	rq1.RequestType = 3 //Dịch vụ
	DB.Create(rq1)

	rq12 := new(model.Request)
	rq12.Creator = "22222222"
	rq12.IsFromAdmin = false
	rq12.MotelId = 1
	rq12.Title = "Yêu cầu gửi từ user 22222222"
	rq12.Content = "Yêu cầu gửi từ user 22222222"
	rq12.Status = 0
	rq12.RequestType = 3 //Dịch vụ
	DB.Create(rq1)

	rq13 := new(model.Request)
	rq13.Creator = "33333333"
	rq13.IsFromAdmin = false
	rq13.MotelId = 1
	rq13.Title = "Yêu cầu gửi từ user 33333333"
	rq13.Content = "Yêu cầu gửi từ user 33333333"
	rq13.Status = 0
	rq13.RequestType = 3 //Dịch vụ
	DB.Create(rq13)

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
