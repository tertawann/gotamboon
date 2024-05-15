package utils

import (
	"testing"

	"github.com/gotamboon/modules/entities"
	"github.com/stretchr/testify/assert"
	r "github.com/stretchr/testify/require"
)

func TestConvertStringToInt(t *testing.T) {

	t.Run("success", func(t *testing.T) {

		result, _ := ConvertStringToInt("1")
		r.Equal(t, 1, result)

	})
	t.Run("can't convert string to int", func(t *testing.T) {

		_, err := ConvertStringToInt("abccc")
		r.Equal(t, "can't convert string to int", err.Error())

	})

}
func TestConvertMapToSlice(t *testing.T) {

	maps := make(map[string]*entities.DonatorRanking)
	maps["key"] = &entities.DonatorRanking{
		Name:  "test",
		Total: 100.00,
	}

	slices := make([]*entities.DonatorRanking, 0, len(maps))
	for _, ranking := range maps {
		slices = append(slices, ranking)
	}

	result := ConvertMapToSlice(maps, len(maps))

	assert.Equal(t, slices, result, "Converted slice should match expected slice")
	assert.NotEqual(t, maps, result, "Converted slice should match expected slice but still map")
}
