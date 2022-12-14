package model

import "time"

// type AdminRequest struct {
// 	Id          int
// 	Creator     int //Creator
// 	MotelId     int
// 	RequestType int //Thuee phong//tra phong
// 	Title       string
// 	DueDate     time.Time
// 	Content     string
// 	CreateAt    time.Time
// }

// type ClientRequest struct {
// 	Id          int
// 	UserId      int
// 	RequestType int //Thuee phong//tra phong//Sua chua
// 	// Title       string
// 	// DueDate     time.Time
// 	MotelId  int
// 	Content  string
// 	Status   int // Complete/In-Complete
// 	CreateAt time.Time
// 	Creator  int
// 	//Ngày đã hoàn thành?
// }

type Request struct {
	Id          int       `json:"id" gorm:"primaryKey; autoIncreament"`
	Creator     string    `json:"creator"` //Creator
	IsFromAdmin bool      `json:"is_from_admin"`
	MotelId     int       `json:"motel_id"`
	RequestType int       `json:"type"`   // Loại yêu cầu: 1: thuê phòng, 2 trả phòng, 3 sửa chữa
	Status      int       `json:"status"` //chưa hoàn thành 1, đã hoàn thành 2 nghe
	Title       string    `json:"title"`
	DueDate     string    `json:"due_date"`
	Content     string    `json:"content"`
	CreateAt    time.Time `json:"create_at" gorm:"autoCreateTime"`
}

func (Request) TableName() string {
	return "request"
}
