package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type UserInfo struct {
	Id       int        `gorm:"primaryKey, index" json:"id"`
	UserName string     `json:"user_name" gorm:"unique"`
	Language string     `json:"language" gorm:"default:'vi'"`
	Theme    int        `json:"theme" gorm:"default:1"`
	CreateAt time.Time  `gorm:"type:timestamp; autoCreateTime"`
	UpdateAt *time.Time `gorm:"type:timestamp; default:null; autoUpdateTime"`
	// DeletedAt *time.Time `json:"-"`
}

func (UserInfo) TableName() string {
	return "h_user_info"
}

type ErrorResponseUserInfo struct {
	Field string
	Tag   string
	Value string
}

func ValidateUserInfo(userInfo UserInfo) []*ErrorResponseUserInfo {
	var errors []*ErrorResponseUserInfo
	validate := validator.New()
	err := validate.Struct(userInfo)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponseUserInfo
			element.Field = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
