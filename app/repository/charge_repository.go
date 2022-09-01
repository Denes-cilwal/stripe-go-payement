package repository

import (
	"gorm.io/gorm"
	"stripe-go-payment/models"
)

type ChargeRepository interface {
	Save(charge *models.Charge) (*models.Charge, error)
	Migrate() error
}

type chargeRepository struct {
	DB *gorm.DB
}

func NewChargeRepository(db *gorm.DB) ChargeRepository {
	return &chargeRepository{
		DB: db,
	}
}
func (c *chargeRepository) Migrate() error {
	c.DB.AutoMigrate(&models.Charge{})

	return nil
}

func (c *chargeRepository) Save(charge *models.Charge) (*models.Charge, error) {
	return charge, c.DB.Create(&charge).Error
}
