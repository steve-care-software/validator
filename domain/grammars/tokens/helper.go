package tokens

import (
	"encoding/binary"
	"errors"
	"fmt"
)

func parseUintData(data []byte) (*uint8, *uint16, *uint32, *uint64, []byte, error) {
	switch data[0] {
	case 8:
		if len(data) <= 1 {
			return nil, nil, nil, nil, nil, errors.New("the data was expected to contain at least 1 byte in order to convert it to a uint8")
		}

		value := uint8(data[1])
		return &value, nil, nil, nil, data[2:], nil
	case 16:

		if len(data) <= 10 {
			return nil, nil, nil, nil, nil, errors.New("the data was expected to contain at least 10 bytes in order to convert it to a uint16")
		}

		value := binary.BigEndian.Uint16(data[1:9])
		return nil, &value, nil, nil, data[9:], nil
	case 32:

		if len(data) <= 10 {
			return nil, nil, nil, nil, nil, errors.New("the data was expected to contain at least 10 bytes in order to convert it to a uint32")
		}

		value := binary.BigEndian.Uint32(data[1:9])
		return nil, nil, &value, nil, data[9:], nil
	case 64:

		if len(data) <= 10 {
			return nil, nil, nil, nil, nil, errors.New("the data was expected to contain at least 10 bytes in order to convert it to a uint64")
		}

		value := binary.BigEndian.Uint64(data[1:9])
		return nil, nil, nil, &value, data[9:], nil
	}

	str := fmt.Sprintf("the referenced element was expected to contain one of these: [8, 16, 32, 64] in its data at index %d, %d provided", 1, data[1])
	return nil, nil, nil, nil, nil, errors.New(str)
}
