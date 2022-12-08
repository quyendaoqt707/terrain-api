package model

// "golang.org/x/crypto/bcrypt"

type User struct {
	// Id          int    `gorm:"primaryKey; autoIncreament"` //row_id
	Email       string `json:"email" gorm:"primaryKey"`
	IsAdmin     bool   `json:"is_admin"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	DateOfBirth string `json:"date_of_birth"`
	CidNumber   string `json:"cid_number"`
	AvatarId    int
}

func (User) TableName() string {
	return "user"
}

// type Admin struct {

// }
