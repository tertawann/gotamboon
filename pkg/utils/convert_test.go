package utils

import (
	"testing"

	r "github.com/stretchr/testify/require"
)

func TestConvertStringToInt(t *testing.T) {

	result, err := ConvertStringToInt("1")
	r.NoError(t, err, "Can")
	r.Equal(t, 1, result, "Should be integer")
	r.NotEqual(t, "1", result, "Should be string")
}