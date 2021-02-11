package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userstorage"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// // User struct
// type User struct {
// 	UserID    int    `json:"user_id" form:"user_id" gorm:"user_id"`
// 	Username  string `json:"username" form:"username" gorm:"username"`
// 	Fullname  string `json:"fullname" form:"fullname" gorm:"fullname"`
// 	Role      string `json:"role" form:"role" gorm:"role"`
// 	ProfileID int    `json:"profile_id" form:"profile_id" gorm:"profile_id"`
// 	Avatar    JSON   `json:"avatar" form:"avatar" gorm:"avatar"`

// 	CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at" form:"updated_at" gorm:"updated_at"`
// 	Status    int       `json:"status" form:"status" gorm:"status"`
// }

// Location struct
type Location struct {
	LocationID int `json:"id" form:"id" gorm:"id"`

	City   string `json:"city" form:"city" gorm:"city"`
	Ward   string `json:"ward" form:"ward" gorm:"ward"`
	Street string `json:"street" form:"street" gorm:"street"`

	CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" gorm:"updated_at"`
	Status    int       `json:"status" form:"status" gorm:"status"`
}

// Restaurant struct
type Restaurant struct {
	RestaurantID   int    `json:"id" form:"id" gorm:"id"`
	RestaurantName string `json:"restaurant_name" form:"restaurant_name" gorm:"restaurant_name"`
	LocationID     int    `json:"location_id" form:"location_id" gorm:"location_id"`
	OpenHour       string `json:"open_hour" form:"open_hour" gorm:"open_hour" `
	CloseHour      string `json:"close_hour" form:"close_hour" gorm:"close_hour" `
	Description    string `json:"description" form:"description" gorm:"description"`

	CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" gorm:"updated_at"`
	Status    int       `json:"status" form:"status" gorm:"status"`
}

//Note struct
// type Note struct {
// 	ID      int  `json:"id,omitempty"`
// 	Embeded User `json:"embed"`
// }

func main() {

	// Connect to db
	dsn := "admin:admin@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to db", db)

	// Router
	router := gin.Default()
	// Test new Res
	routerV1 := router.Group("/v1")
	resRoute := routerV1.Group("/restaurants")
	{
		resRoute.GET("/:id", func(c *gin.Context) {
			var res Restaurant
			id, _ := strconv.Atoi(c.Param("id"))
			log.Println(id)
			if err := db.Table("restaurants").Where("id=?", id).First(&res).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Cannot found in database"})
				return
			}
			log.Println(res)
			c.JSON(http.StatusOK, gin.H{"data": res})
		})
		resRoute.POST("", func(c *gin.Context) {
			var res Restaurant
			if err := c.ShouldBind(&res); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error format data"})
				return
			}
			if err := db.Table("restaurants").Create(&res).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data cannot store in database. Wrong format"})
				return
			}
			log.Println(res)
			c.JSON(http.StatusOK, gin.H{"data": 1})
		})
	}
	userRoute := routerV1.Group("/users")
	{
		userRoute.DELETE("/:user-id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("user-id"))
			store := userstorage.NewSQLStore(db)
			biz := userbiz.NewDeleteUserBiz(store)
			if err := biz.DeleteUser(id); err != nil {
				c.JSON(http.StatusGone, gin.H{
					"error": err.Error(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"data": 1,
			})
		})

	}

	router.Run(":8080")
}
