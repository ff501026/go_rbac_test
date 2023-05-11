package util

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func ValidateUUID(input string) (uuid.UUID, error) {
	return uuid.Parse(input)
}

func BydUuidStringToUuid(input string) string {
	index1, index2, index3, index4 := 8, 12, 16, 20

	return input[:index1] + "-" + input[index1:index2] + "-" + input[index2:index3] + "-" + input[index3:index4] + "-" + input[index4:]
}
