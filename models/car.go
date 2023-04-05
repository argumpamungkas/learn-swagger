package models

import (
	"errors"

	"gorm.io/gorm"
)

// Car represents the model for an car
type Car struct {
	CarID uint   `json:"car_id" gorm:"primaryKey"`
	Brand string `json:"brand" gorm:"not null;type:varchar(50)"`
	Model string `json:"model" gorm:"not null;type:varchar(50)"`
	Price int    `json:"price" gorm:"not null"`
}

func (c *Car) TableName() string {
	return "tb_cars"
}

func (c *Car) BeforeCreate(tx *gorm.DB) (err error) {

	if c.Brand == "" || c.Model <= "" || c.Price <= 0 {
		err = errors.New("field cars is not null")
	}

	return
}

func (c *Car) BeforeUpdate(tx *gorm.DB) (err error) {

	if c.Brand == "" || c.Model <= "" || c.Price <= 0 {
		err = errors.New("field cars is not null")
	}

	return
}
