package models

import (
	"gorm.io/gorm"
)

type Ingredient struct {
	IngredientID     uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         *string `gorm:"type:varchar(255)" json:"name"`
	Calories     int     `gorm:"type:int" json:"calories"`
	ExpiryDate   string  `gorm:"type:varchar(255)" json:"expiry_date"`
	Quantity     int     `gorm:"type:int" json:"quantity"`
}

func MigrateIngredient(db *gorm.DB) error {
	err := db.AutoMigrate(&Ingredient{})
	return err
}
