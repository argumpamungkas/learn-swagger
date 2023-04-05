package routers

import (
	"DTS/Chapter-2/Sesi/sesi-5-swagger/controllers"

	"github.com/gin-gonic/gin"

	_ "DTS/Chapter-2/Sesi/sesi-5-swagger/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

func StartServer() (r *gin.Engine) {
	r = gin.Default()

	// Create
	r.POST("/cars", controllers.CreateCar)

	// Read All
	r.GET("/cars", controllers.GetAllCars)

	// Read
	r.GET("/cars/:carID", controllers.GetCarsById)

	// Update
	r.PUT("/cars/:carID", controllers.UpdateCar)

	// Delete
	r.DELETE("/cars/:carID", controllers.DeleteCar)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
