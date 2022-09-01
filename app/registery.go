package app

import "stripe-go-payment/models"

// create model struct for multiple models type using empty interface

type Model struct {
	m interface{}
}

func RegisterModels() []Model {
	return []Model{
		{m: models.Charge{}},
		{m: models.User{}},
	}
}
