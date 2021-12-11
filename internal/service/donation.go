package service

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"opn-challenge/internal/models"
	"time"
)

type DonationService interface {
	Donate(record models.DonationRecord) error
}

type DefaultDonationService struct {
	paymentClient *omise.Client
}

func NewDefaultDonationService(paymentClient *omise.Client) *DefaultDonationService {
	return &DefaultDonationService{paymentClient: paymentClient}
}

func (d *DefaultDonationService) Donate(record models.DonationRecord) error {
	card, createToken := &omise.Card{}, &operations.CreateToken{
		Name:            record.Name,
		Number:          record.CCNumber,
		ExpirationMonth: time.Month(record.ExpMonth),
		ExpirationYear:  int(record.ExpYear),
	}

	if err := d.paymentClient.Do(card, createToken); err != nil {
		return err
	}

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   record.AmountSubunits,
		Currency: "thb",
		Card:     card.ID,
	}
	if err := d.paymentClient.Do(charge, createCharge); err != nil {
		return err
	}
	return nil
}
