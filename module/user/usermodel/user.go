package usermodel

import "time"

// EntityName variable
const EntityName = "User"

// JSON type
type JSON []byte

//User struct
type User struct {
	ID        int        `json:"id" form:"id" gorm:"primaryKey"`
	Username  string     `json:"username" form:"username"`
	Fullname  string     `json:"fullname" form:"fullname"`
	Role      string     `json:"role" form:"role"`
	ProfileID int        `json:"profile_id" form:"profile_id"`
	Avatar    JSON       `json:"avatar" form:"avatar"`
	CreatedAt *time.Time `json:"created_at" form:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at"`
	Status    int        `json:"status" form:"status" gorm:"default:1"`
}

// TableName of user
func (User) TableName() string {
	return "users"
}
