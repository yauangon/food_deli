package restaurantmodel

import "time"

// Restaurant type
type Restaurant struct {
	ID             int        `json:"id,omitempty" form:"id" gorm:"primaryKey"`
	RestaurantName string     `json:"restaurant_name" form:"restaurant_name"`
	LocationID     *int       `json:"location_id" form:"location_id" gorm:"index"`
	OpenHour       *string    `json:"open_hour" form:"open_hour" `
	CloseHour      *string    `json:"close_hour" form:"close_hour"`
	Description    string     `json:"description" form:"description"`
	CreatedAt      *time.Time `json:"created_at" form:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" form:"updated_at"`
	Status         int        `json:"status" form:"status" gorm:"default:1"`
}

// Entity name
var Entity = "Restaurant"

// TableName of restaurants
func (Restaurant) TableName() string {
	return "restaurants"
}
