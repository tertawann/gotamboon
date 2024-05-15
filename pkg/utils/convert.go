package utils

import (
	"strconv"

	"github.com/gotamboon/modules/entities"
)

func ConvertStringToInt(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func ConvertMapToSlice(maps map[string]*entities.DonatorRanking, size int) []*entities.DonatorRanking {
	slices := make([]*entities.DonatorRanking, 0, size)

	for _, val := range maps {
		slices = append(slices, val)
	}

	return slices
}
