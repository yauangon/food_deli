package restaurantmodel

// Restaurant type
type DeleteRestaurant struct {
	ID int `json:"id,omitempty" form:"id" gorm:"primaryKey"`
}

// TableName of restaurants
func (DeleteRestaurant) TableName() string {
	return "restaurants"
}
