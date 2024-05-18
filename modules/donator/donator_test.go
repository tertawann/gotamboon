package donator

import (
	"fmt"
	"testing"
	"time"

	"github.com/gotamboon/modules/entities"
	"github.com/gotamboon/modules/omisetor"
	"github.com/stretchr/testify/assert"
)

func setupDonator() *Donator {

	_donator, err := NewDonator()
	if err != nil {
		panic(fmt.Sprintf("Failed to instance donator: %v", err))
	}

	return _donator

}

func setupOmisetor(pkey string, skey string) *omisetor.Omise {

	_omisetor, err := omisetor.NewOmiseClient(pkey, skey)

	if err != nil {
		panic(fmt.Sprintf("Failed to instance omisetor: %v", err))
	}

	return _omisetor

}

func TestGetDonationList(t *testing.T) {
	_donator := setupDonator()

	t.Run("success", func(t *testing.T) {

		exDonationList := []*entities.Donation{
			{
				Name:           "test",
				AmountSubunits: 15000,
				CCNumber:       "595995",
				CVV:            "7750",
				ExpMonth:       time.Month(5),
				ExpYear:        2025,
			},
			{
				Name:           "test",
				AmountSubunits: 30000,
				CCNumber:       "16000000",
				CVV:            "7750",
				ExpMonth:       time.Month(2),
				ExpYear:        2025,
			},
		}

		_donator.donationList = exDonationList

		donationList := _donator.GetDonationList()

		assert.Equal(t, exDonationList, donationList)
	})
}

func TestSplitDonationList(t *testing.T) {

	_donator := setupDonator()

	t.Run("success", func(t *testing.T) {
		exDecryptedFile := `Name,AmountSubunits,CCNumber,CVV,ExpMonth,ExpYear
		Mr. Grossman R Oldbuck,2879410,5375543637862918,488,11,2025`

		err := _donator.SplitDonationList(exDecryptedFile)

		assert.Equal(t, nil, err)
	})

	t.Run("can not convert amount string to int", func(t *testing.T) {
		exDecryptedFile := `Name,AmountSubunits,CCNumber,CVV,ExpMonth,ExpYear
		Mr. Grossman R Oldbuck,qweqwe,5375543637862918,488,11,2025`

		err := _donator.SplitDonationList(exDecryptedFile)

		assert.Equal(t, "can not convert amount string to int", err.Error())
	})

	t.Run("can not convert month string to int", func(t *testing.T) {
		exDecryptedFile := `Name,AmountSubunits,CCNumber,CVV,ExpMonth,ExpYear
		Mr. Grossman R Oldbuck,155000,5375543637862918,488,qweqe,2025`

		err := _donator.SplitDonationList(exDecryptedFile)

		assert.Equal(t, "can not convert month string to int", err.Error())
	})

	t.Run("can not convert year string to int", func(t *testing.T) {
		exDecryptedFile := `Name,AmountSubunits,CCNumber,CVV,ExpMonth,ExpYear
		Mr. Grossman R Oldbuck,2879410,5375543637862918,488,11,aa`

		err := _donator.SplitDonationList(exDecryptedFile)

		assert.Equal(t, "can not convert year string to int", err.Error())
	})
}

func TestPerformDonations(t *testing.T) {

	_donator := setupDonator()
	_omistor := setupOmisetor("pkey_test_5zq13eohc3ecomf54em", "skey_test_5zq13g2oiynactycmdt")

	t.Run("success", func(t *testing.T) {

		exDonationList := &entities.Donation{
			Name:           "test",
			AmountSubunits: 15000,
			CCNumber:       "5555 5555 5555 4444",
			CVV:            "7750",
			ExpMonth:       time.Month(5),
			ExpYear:        2025,
		}

		err := _donator.PerformDonations(_omistor, exDonationList)

		assert.Equal(t, nil, err)
	})

	t.Run("failed to create token for donator", func(t *testing.T) {

		exDonationList := &entities.Donation{
			Name:           "test",
			AmountSubunits: 15000,
			CCNumber:       "111",
			CVV:            "7750",
			ExpMonth:       time.Month(5),
			ExpYear:        2022,
		}

		err := _donator.PerformDonations(_omistor, exDonationList)

		assert.Equal(t, "failed to create token for donator", err.Error())
	})

}
