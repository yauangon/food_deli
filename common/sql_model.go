package common

import "time"

// SQLModel type
type SQLModel struct {
	ID        int        `json:"id" gorm:"id"`
	CreatedAt *time.Time `json:"created_at" form:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at" gorm:"updated_at"`
	Status    int        `json:"status" form:"status" gorm:"status"`
}
