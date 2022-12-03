package model

// "golang.org/x/crypto/bcrypt"

type User struct {
	Id          int `gorm:"primary"` //row_id
	IsAdmin     bool
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	DateOfBirth string `json:"date_of_birth"`
	CidNumber   bool   `json:"cid_number"`
	AvatarId    int
}

func (User) TableName() string {
	return "user"
}

// type Admin struct {

// }
