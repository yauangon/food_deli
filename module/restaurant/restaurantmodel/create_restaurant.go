package restaurantmodel

import (
	"food_deli/common"
)

// Restaurant type
type CreateRestaurant struct {
	ID          int           `json:"id,omitempty" form:"id" gorm:"primaryKey"`
	Name        string        `json:"name" form:"name" gorm:"column:name"`
	CityID      *int          `json:"city_id" form:"city_id" gorm:"column:city_id"`
	Addr        *string       `json:"addr" form:"addr" gorm:"column:addr"`
	Lat         *float64      `json:"lat" form:"lat" gorm:"column:lat"`
	Long        *float64      `json:"long" form:"long" gorm:"column:long"`
	Logo        *common.Image `json:"logo" gorm:"column:logo"`
	OpenHour    *string       `json:"open_hour" form:"open_hour" gorm:"column:open_hour"`
	CloseHour   *string       `json:"close_hour" form:"close_hour" gorm:"column:close_hour"`
	Description string        `json:"description" form:"description" gorm:"column:description"`
}

// TableName of restaurants
func (CreateRestaurant) TableName() string {
	return "restaurants"
}
