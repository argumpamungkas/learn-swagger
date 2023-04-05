package controllers

import (
	"DTS/Chapter-2/Sesi/sesi-5-swagger/models"
	"DTS/Chapter-2/Sesi/sesi-5-swagger/repo"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCars godoc
// @Summary Post create car
// @Description Post create a new car, NOTE : id auto increment, so in body car_id is deleted
// @Tags Create Car
// @Accept json
// @Produce application/json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars [post]
func CreateCar(ctx *gin.Context) {
	db := repo.GetDB()
	var newCar models.Car
	var lastID int

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintln("error because ", err.Error()),
		})
		return
	}

	_ = db.Select("car_id").Last(&newCar).Scan(&lastID)

	newCar = models.Car{
		CarID: uint(lastID) + 1,
		Brand: newCar.Brand,
		Model: newCar.Model,
		Price: newCar.Price,
	}

	err := db.Create(&newCar).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    newCar,
		"message": "Created new car is successfully",
	})

}

// GetAllCars godoc
// @Summary Get All cars
// @Description Get details of all car
// @Tags Get Car
// @Accept json
// @Produce application/json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCars(ctx *gin.Context) {
	db := repo.GetDB()

	var cars []models.Car
	err := db.Order("car_id asc").Find(&cars).Error
	if err != nil {
		fmt.Println("Error getting car data", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": cars,
	})
}

// GetCarbyID godoc
// @Summary Get details for a given car_id
// @Description Get details of car corresponding to the input car_id
// @Tags Get Car
// @Accept json
// @Produce application/json
// @Param carID path integer true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{carID} [get]
func GetCarsById(ctx *gin.Context) {
	db := repo.GetDB()
	carID := ctx.Param("carID")
	var car models.Car

	idCar, err := strconv.Atoi(carID)
	if err != nil {
		return
	}

	err = db.First(&car, "car_id = ?", idCar).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("cars with id %v not found", idCar),
		})
		return
	}

	ctx.JSON(http.StatusOK, car)
}

// UpdateCar godoc
// @Summary Updated car data with car_id
// @Description Update data car
// @Tags Update Car
// @Accept json
// @Produce json
// @Param carID path integer true "car_id of the car to be updated"
// @Param models.Car body models.Car true "updated car"
// @Success 200 {object} models.Car
// @Router /cars/{carID} [put]
func UpdateCar(ctx *gin.Context) {
	db := repo.GetDB()
	carID := ctx.Param("carID")
	var updateCar, findCar models.Car

	idCar, err := strconv.Atoi(carID)
	if err != nil {
		return
	}

	err = db.First(&findCar, "car_id = ?", idCar).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("cars with id %v not found", idCar),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	// Validate input
	updateCar = models.Car{
		CarID: uint(idCar),
		Brand: updateCar.Brand,
		Model: updateCar.Model,
		Price: updateCar.Price,
	}

	err = db.Model(&updateCar).Where("car_id = ?", idCar).Updates(&updateCar).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    updateCar,
		"message": "Updated data car is successfully",
	})
}

// DeleteCar godoc
// @Summary Deleted car data with car_id
// @Description Deleted data car
// @Tags Deleted Car
// @Accept json
// @Produce json
// @Param carID path integer true "car_id of the car to be deleted"
// @Success 200 {object} models.Car
// @Router /cars/{carID} [delete]
func DeleteCar(ctx *gin.Context) {
	db := repo.GetDB()
	carID := ctx.Param("carID")
	var findCar models.Car

	idCar, err := strconv.Atoi(carID)
	if err != nil {
		log.Println("Invalid Convert id")
		return
	}

	err = db.Where("car_id = ?", idCar).First(&findCar).Delete(&findCar).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("cars with id %v not found", idCar),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carID),
	})
}
