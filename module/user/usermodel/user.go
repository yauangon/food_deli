package usermodel

import "time"

// EntityName variable
const EntityName = "User"

// JSON type
type JSON []byte

//User struct
type User struct {
	ID          int        `json:"id" form:"id" gorm:"primaryKey"`
	Fullname    string     `json:"fullname" form:"fullname" gorm:"default:'User'"`
	Role        string     `json:"role" form:"role" gorm:"default:'USER'"`
	Avatar      JSON       `json:"avatar" form:"avatar"`
	Password    string     `json:"password" form:"password"`
	PhoneNumber string     `json:"phone_number" form:"phone_number"`
	Email       string     `json:"email" form:"email"`
	CreatedAt   *time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" form:"updated_at"`
	Status      int        `json:"status" form:"status" gorm:"default:1"`
}

// TableName of user
func (User) TableName() string {
	return "users"
}
