package service

import (
	"stripe-go-payment/app/repository"
	"stripe-go-payment/models"
)

type ChargeService interface {
	SaveCharge(charge *models.Charge) (*models.Charge, error)
}
type chargeService struct {
	chargeRepository repository.ChargeRepository
}

func NewChargeService(c repository.ChargeRepository) ChargeService {
	return &chargeService{
		chargeRepository: c,
	}
}

func (cs chargeService) SaveCharge(charge *models.Charge) (*models.Charge, error) {
	return cs.chargeRepository.Save(charge)
}
