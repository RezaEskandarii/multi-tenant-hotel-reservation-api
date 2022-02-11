package domain_services

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
	"reservation-api/internal/models"
)

type PaymentService struct {
}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (p *PaymentService) Pay() interface{} {

	var payment models.Charge
	payment.ReceiptEmail = "rezaesskandari@gmail.com"
	apiKey := ""

	stripe.Key = apiKey
	result, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(2111111000000033),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String("hotel reservation"),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String(payment.ReceiptEmail),
	})

	if err != nil {
		return err
	}

	return result
}
