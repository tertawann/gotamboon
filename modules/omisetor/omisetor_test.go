package omisetor

import (
	"fmt"
	"testing"
	"time"

	"github.com/gotamboon/modules/entities"
	"github.com/stretchr/testify/assert"
)

func setupOmisetor(pkey string, skey string) *Omise {

	_omisetor, err := NewOmiseClient(pkey, skey)

	if err != nil {
		panic(fmt.Sprintf("Failed to instance omisetor: %v", err))
	}

	return _omisetor

}

func TestCreateChargeByToken(t *testing.T) {

	_omistor := setupOmisetor("pkey_test_5zq13eohc3ecomf54em", "skey_test_5zq13g2oiynactycmdt")

	t.Run("failed", func(t *testing.T) {

		exDonationList := &entities.Donation{
			Name:           "test",
			AmountSubunits: 15000,
			CCNumber:       "5555 5555 5555 4444",
			CVV:            "7750",
			ExpMonth:       time.Month(5),
			ExpYear:        2025,
		}

		_, err := _omistor.CreateChargeByToken(exDonationList, "testset")

		assert.Equal(t, "(404/not_found) token testset was not found", err.Error())
	})
}
