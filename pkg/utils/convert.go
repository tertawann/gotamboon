package utils

import "strconv"

func ConvertStringToInt(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return num, nil
}
