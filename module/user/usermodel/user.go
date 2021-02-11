package usermodel

import "github.com/thanhdat1902/restapi/food_deli/common"

// EntityName variable
const EntityName = "User"

// JSON type
type JSON []byte

//User struct
type User struct {
	common.SQLModel `json:",inline"`

	UserID    int    `json:"user_id" form:"user_id" gorm:"user_id"`
	Username  string `json:"username" form:"username" gorm:"username"`
	Fullname  string `json:"fullname" form:"fullname" gorm:"fullname"`
	Role      string `json:"role" form:"role" gorm:"role"`
	ProfileID int    `json:"profile_id" form:"profile_id" gorm:"profile_id"`
	Avatar    JSON   `json:"avatar" form:"avatar" gorm:"avatar"`
}

// TableName of user
func (User) TableName() string {
	return "Users"
}
