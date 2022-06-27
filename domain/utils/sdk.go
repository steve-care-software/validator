package utils

import (
	"errors"
	"strconv"
)

// IsBytePresent returns true if the byte is presents false otherwise
func IsBytePresent(value byte, data []byte) bool {
	isPresent := false
	for _, oneChanByte := range data {
		if value == oneChanByte {
			isPresent = true
			break
		}
	}

	return isPresent
}

// FetchNumber fetches a number from data
func FetchNumber(input []byte) (*uint, []byte, error) {
	numberCharacters := []byte("0123456789")
	indexBytes := []byte{}
	for _, oneInputByte := range input {
		if !IsBytePresent(oneInputByte, numberCharacters) {
			break
		}

		indexBytes = append(indexBytes, oneInputByte)
	}

	if len(indexBytes) <= 0 {
		return nil, nil, errors.New("the input does not contain a number")
	}

	intNumber, err := strconv.Atoi(string(indexBytes))
	if err != nil {
		return nil, nil, err
	}

	casted := uint(intNumber)
	return &casted, input[len(indexBytes):], nil
}
