package restaurantmodel

// Restaurant type
type UpdateRestaurant struct {
	ID             int     `json:"id,omitempty" form:"id" gorm:"primaryKey"`
	RestaurantName string  `json:"restaurant_name" form:"restaurant_name"`
	LocationID     *int    `json:"location_id" form:"location_id" gorm:"index"`
	OpenHour       *string `json:"open_hour" form:"open_hour" `
	CloseHour      *string `json:"close_hour" form:"close_hour"`
	Description    *string `json:"description" form:"description"`
}

// TableName of restaurants
func (UpdateRestaurant) TableName() string {
	return "restaurants"
}
