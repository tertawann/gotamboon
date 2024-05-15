package utils

import (
	"errors"
	"strconv"

	"github.com/gotamboon/modules/entities"
)

func ConvertStringToInt(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("can't convert string to int")
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
