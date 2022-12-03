package model

import "time"

type AdminRequest struct {
	Id          int
	Creator     int //Creator
	MotelId     int
	RequestType int //Thuee phong//tra phong
	Title       string
	DueDate     time.Time
	Content     string
	CreateAt    time.Time
}

type ClientRequest struct {
	Id          int
	UserId      int
	RequestType int //Thuee phong//tra phong//Sua chua
	// Title       string
	// DueDate     time.Time
	MotelId  int
	Content  string
	Status   int // Complete/In-Complete
	CreateAt time.Time
	Creator  int
	//Ngày đã hoàn thành?
}
