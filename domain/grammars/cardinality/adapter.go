package cardinality

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/steve-care-software/validator/domain/utils"
)

type adapter struct {
	builder           Builder
	nonZeroMultiple   byte
	zeroMultiple      byte
	optional          byte
	rangePrefix       byte
	rangeSuffix       byte
	rangeSeparator    byte
	numbersCharacters []byte
}

func createAdapter(
	builder Builder,
	nonZeroMultiple byte,
	zeroMultiple byte,
	optional byte,
	rangePrefix byte,
	rangeSuffix byte,
	rangeSeparator byte,
	numbersCharacters []byte,
) Adapter {
	out := adapter{
		builder:           builder,
		nonZeroMultiple:   nonZeroMultiple,
		zeroMultiple:      zeroMultiple,
		optional:          optional,
		rangePrefix:       rangePrefix,
		rangeSuffix:       rangeSuffix,
		rangeSeparator:    rangeSeparator,
		numbersCharacters: numbersCharacters,
	}

	return &out
}

// ToCardinality converts a script to a cardinality instance
func (app *adapter) ToCardinality(script string) (Cardinality, []byte, error) {
	input := []byte(script)
	builder := app.builder.Create()
	if len(input) <= 0 {
		ins, err := builder.WithMinimum(1).WithMaximum(1).Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, []byte{}, nil
	}

	remaining := input
	if input[0] == app.nonZeroMultiple {
		builder.WithMinimum(1)
		remaining = input[1:]
	}

	if input[0] == app.zeroMultiple {
		builder.WithMinimum(0)
		remaining = input[1:]
	}

	if input[0] == app.optional {
		builder.WithMinimum(0).WithMaximum(1)
		remaining = input[1:]
	}

	if input[0] == app.rangePrefix {
		pMin, pMax, retRemaining, err := app.fetchRange(input[1:])
		if err != nil {
			return nil, nil, err
		}

		builder.WithMinimum(*pMin)
		if pMax != nil {
			builder.WithMaximum(*pMax)
		}

		remaining = retRemaining
	}

	ins, err := builder.Now()
	if err != nil {
		ins, err = builder.WithMinimum(1).WithMaximum(1).Now()
		if err != nil {
			return nil, nil, err
		}
	}

	return ins, remaining, nil
}

func (app *adapter) fetchRange(input []byte) (*uint8, *uint8, []byte, error) {
	pFirstNumber, isSpecific, retRemaining, err := app.fetchFirstNumberInRange(input)
	if err != nil {
		return nil, nil, nil, err
	}

	if isSpecific {
		return pFirstNumber, nil, retRemaining, nil
	}

	pSecondNumber, _, retRemainingAfterMax, _ := app.fetchFirstNumberInRange(retRemaining)
	return pFirstNumber, pSecondNumber, retRemainingAfterMax, nil
}

func (app *adapter) fetchFirstNumberInRange(input []byte) (*uint8, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, false, nil, errors.New("the input was NOT expected to be empty while fetching the element's cardinality range number (min/max)")
	}

	isSpecific := true
	numberBytes := []byte{}
	for _, oneInputByte := range input {
		if oneInputByte == app.rangeSeparator {
			isSpecific = false
			break
		}

		if oneInputByte == app.rangeSuffix {
			break
		}

		if !utils.IsBytePresent(oneInputByte, app.numbersCharacters) {
			return nil, false, nil, errors.New("the input elements within a range must be numbers")
		}

		numberBytes = append(numberBytes, oneInputByte)
	}

	if len(numberBytes) <= 0 {
		return nil, false, input[1:], nil
	}

	intNumber, err := strconv.Atoi(string(numberBytes))
	if err != nil {
		return nil, false, nil, err
	}

	if intNumber >= 256 {
		str := fmt.Sprintf("the elements of a cardinality (range, specific) must contain a maximum value of 256, %d provided", intNumber)
		return nil, false, nil, errors.New(str)
	}

	casted := uint8(intNumber)
	return &casted, isSpecific, input[len(numberBytes)+1:], nil
}
