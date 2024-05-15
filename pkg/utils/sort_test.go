package utils

import (
	"testing"

	"github.com/gotamboon/modules/entities"
	"github.com/stretchr/testify/assert"
)

func TestSortDescDonatorsByTotal(t *testing.T) {
	rankingMaps := make(map[string]*entities.DonatorRanking)
	rankingMaps["donator1"] = &entities.DonatorRanking{
		Name:  "Donator 1",
		Total: 500.00,
	}
	rankingMaps["donator2"] = &entities.DonatorRanking{
		Name:  "Donator 2",
		Total: 200.00,
	}
	rankingMaps["donator3"] = &entities.DonatorRanking{
		Name:  "Donator 3",
		Total: 300.00,
	}

	expectedDesc := []*entities.DonatorRanking{
		{
			Name:  "Donator 1",
			Total: 500.00,
		},
		{
			Name:  "Donator 3",
			Total: 300.00,
		},
		{
			Name:  "Donator 2",
			Total: 200.00,
		},
	}

	result := SortDescDonatorsByTotal(rankingMaps)

	assert.Equal(t, expectedDesc, result, "Sorted slice should match expected slice sort desc")
}
