package model

// "golang.org/x/crypto/bcrypt"

type User struct {
	// Id          int    `gorm:"primaryKey; autoIncreament"` //row_id
	Phone       string `json:"phone" gorm:"primaryKey"`
	Email       string `json:"email"`
	IsAdmin     bool   `json:"is_admin"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	DateOfBirth string `json:"date_of_birth"`
	CidNumber   string `json:"cid_number"`
	AvatarUrl   string `json:"avatar_url"`
}

func (User) TableName() string {
	return "user"
}

// type Admin struct {

// }
