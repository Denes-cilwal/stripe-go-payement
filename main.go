package main

import (
	"stripe-go-payment/app"
	infrastructure "stripe-go-payment/app/config"
	"stripe-go-payment/utils"
)

var err error

func main() {
	utils.LoadEnv()
	infrastructure.SetupModels()
	r := app.InitializeRoutes()

	r.Run(":9000")

}
