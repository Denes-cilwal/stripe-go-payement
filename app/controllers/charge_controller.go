package controller

import (
	"github.com/pkg/errors"
	"net/http"
	service "stripe-go-payment/app/services"
	"stripe-go-payment/models"
	"stripe-go-payment/responses"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

type chargeController struct {
	chargeServices service.ChargeService
}
type ChargeController interface {
	AddPayment(c *gin.Context)
}

func NewChargeController(s service.ChargeService) ChargeController {
	return &chargeController{
		chargeServices: s,
	}
}

func (ch chargeController) AddPayment(c *gin.Context) {

	stripe.Key = "sk_test_51Ld3zTEc0DwYJw2aFW03bxZjrZDRs9YaACq5GPGd0shkpGcqPaU8S0mfqrQhKRjFymCpQIN3BBrx49JC8hDcxOow002PaDeCWD"
	var payment models.Charge

	if err := c.ShouldBindJSON(&payment); err != nil {
		err = errors.Errorf("Bad Request Error")
		responses.HandleError(c, err)
		return
	}

	_, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(payment.Amount),
		Currency:     stripe.String(string(stripe.CurrencyJPY)),
		Source:       &stripe.SourceParams{Token: stripe.String(payment.CardToken)},
		ReceiptEmail: stripe.String(payment.ReceiptEmail),
		Description:  stripe.String(payment.Name),
	})
	if err != nil {
		responses.HandleError(c, err)
		return
	}

	//payment.UserID = c.MustGet("uid").(string)
	_, err = ch.chargeServices.SaveCharge(&payment)
	if err != nil {
		err = errors.Errorf("Payment unsuccessful")
		responses.HandleError(c, err)
		return
	}
	c.String(http.StatusCreated, "Payment successful")
}
