package omisetor

import (
	"errors"

	"github.com/gotamboon/modules/entities"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type Omise struct {
	publicKey   string
	secretKey   string
	omiseClient *omise.Client
	card        *omise.Card
	charge      *omise.Charge
}

func NewOmiseClient(publicKey, secretKey string) (*Omise, error) {
	client, err := omise.NewClient(publicKey, secretKey)
	if err != nil {
		return nil, errors.New("can't instance omise")
	}

	return &Omise{
		publicKey:   publicKey,
		secretKey:   secretKey,
		omiseClient: client,
		card:        &omise.Card{},
		charge:      &omise.Charge{},
	}, nil
}

func (o *Omise) GenerateToken(donator *entities.Donation) (*omise.Card, error) {

	err := o.omiseClient.Do(o.card, &operations.CreateToken{
		Name:            donator.Name,
		Number:          donator.CCNumber,
		ExpirationMonth: donator.ExpMonth,
		ExpirationYear:  2025,
	})

	if err != nil {
		return nil, err
	}

	return o.card, nil
}

func (o *Omise) CreateChargeByToken(donator *entities.Donation, token string) (*omise.Charge, error) {

	err := o.omiseClient.Do(o.charge, &operations.CreateCharge{
		Amount:   donator.AmountSubunits,
		Currency: "thb",
		Card:     token,
	})

	if err != nil {
		return nil, err
	}

	return o.charge, nil
}
