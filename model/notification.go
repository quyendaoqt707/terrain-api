package model

import "time"

type Notification struct {
	Id           int
	IsFromAdmin  bool
	Severity     int //1 2 3
	Creator      int //Creator
	ModelGroupId int
	MotelId      []int
	RequestType  int //Thuee phong//tra phong
	Title        string
	DueDate      time.Time //Ngày hết hạn thông báo
	Content      string
	CreateAt     time.Time
}
