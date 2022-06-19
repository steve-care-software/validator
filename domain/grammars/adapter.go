package grammars

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/steve-care-software/validator/domain/grammars/cardinality"
	"github.com/steve-care-software/validator/domain/grammars/channels"
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

type adapter struct {
	grammarBuilder                Builder
	channelsBuilder               channels.Builder
	channelBuilder                channels.ChannelBuilder
	conditionBuilder              channels.ConditionBuilder
	tokenBuilder                  tokens.TokenBuilder
	linesBuilder                  tokens.LinesBuilder
	lineBuilder                   tokens.LineBuilder
	elementWithCardinalityBuilder tokens.ElementWithCardinalityBuilder
	elementBuilder                tokens.ElementBuilder
	cardinalityBuilder            cardinality.Builder
	rootPrefix                    byte
	rootSuffix                    byte
	channelPrefix                 byte
	channelSuffix                 byte
	channelConditionPrevious      byte
	channelConditionNext          byte
	channelConditionAnd           byte
	tokenNamePrefix               byte
	bytePrefix                    byte
	linesPrefix                   byte
	linesSuffix                   byte
	lineDelimiter                 byte
	cardinalityNonZeroMultiple    byte
	cardinalityZeroMultiple       byte
	cardinalityOptional           byte
	cardinalityRangePrefix        byte
	cardinalityRangeSuffix        byte
	cardinalityRangeSeparator     byte
	commentPrefix                 byte
	commentSuffix                 byte
	numbersCharacters             []byte
	tokenNameCharacters           []byte
	channelCharacters             []byte
}

func createAdapter(
	grammarBuilder Builder,
	channelsBuilder channels.Builder,
	channelBuilder channels.ChannelBuilder,
	conditionBuilder channels.ConditionBuilder,
	tokenBuilder tokens.TokenBuilder,
	linesBuilder tokens.LinesBuilder,
	lineBuilder tokens.LineBuilder,
	elementWithCardinalityBuilder tokens.ElementWithCardinalityBuilder,
	elementBuilder tokens.ElementBuilder,
	cardinalityBuilder cardinality.Builder,
	rootPrefix byte,
	rootSuffix byte,
	channelPrefix byte,
	channelSuffix byte,
	channelConditionPrevious byte,
	channelConditionNext byte,
	channelConditionAnd byte,
	tokenNamePrefix byte,
	bytePrefix byte,
	linesPrefix byte,
	linesSuffix byte,
	lineDelimiter byte,
	cardinalityNonZeroMultiple byte,
	cardinalityZeroMultiple byte,
	cardinalityOptional byte,
	cardinalityRangePrefix byte,
	cardinalityRangeSuffix byte,
	cardinalityRangeSeparator byte,
	commentPrefix byte,
	commentSuffix byte,
	numbersCharacters []byte,
	tokenNameCharacters []byte,
	channelCharacters []byte,
) Adapter {
	out := adapter{
		grammarBuilder:                grammarBuilder,
		channelsBuilder:               channelsBuilder,
		channelBuilder:                channelBuilder,
		conditionBuilder:              conditionBuilder,
		tokenBuilder:                  tokenBuilder,
		linesBuilder:                  linesBuilder,
		lineBuilder:                   lineBuilder,
		elementWithCardinalityBuilder: elementWithCardinalityBuilder,
		elementBuilder:                elementBuilder,
		cardinalityBuilder:            cardinalityBuilder,
		rootPrefix:                    rootPrefix,
		rootSuffix:                    rootSuffix,
		channelPrefix:                 channelPrefix,
		channelSuffix:                 channelSuffix,
		channelConditionPrevious:      channelConditionPrevious,
		channelConditionNext:          channelConditionNext,
		channelConditionAnd:           channelConditionAnd,
		tokenNamePrefix:               tokenNamePrefix,
		bytePrefix:                    bytePrefix,
		linesPrefix:                   linesPrefix,
		linesSuffix:                   linesSuffix,
		lineDelimiter:                 lineDelimiter,
		cardinalityNonZeroMultiple:    cardinalityNonZeroMultiple,
		cardinalityZeroMultiple:       cardinalityZeroMultiple,
		cardinalityOptional:           cardinalityOptional,
		cardinalityRangePrefix:        cardinalityRangePrefix,
		cardinalityRangeSuffix:        cardinalityRangeSuffix,
		cardinalityRangeSeparator:     cardinalityRangeSeparator,
		commentPrefix:                 commentPrefix,
		commentSuffix:                 commentSuffix,
		numbersCharacters:             numbersCharacters,
		tokenNameCharacters:           tokenNameCharacters,
		channelCharacters:             channelCharacters,
	}

	return &out
}

// ToGrammar converts a script to grammar
func (app *adapter) ToGrammar(script string) (Grammar, error) {
	// convert to bytes:
	bytes := []byte(script)

	// remove channel characters:
	remainingAfterChans := app.removeChannelCharacters(bytes)

	// retrieve the root token:
	rootTokenName, remaining, err := app.fetchRootTokenName(remainingAfterChans)
	if err != nil {
		return nil, err
	}

	// retrieve the channel bytes:, if any:
	channelsBytes, remainingAfterChannels, err := app.fetchChannelsBytes(remaining)
	if err != nil {
		return nil, err
	}

	// retrieve the script tokens:
	tokensMap, err := app.toScriptTokens(remainingAfterChannels)
	if err != nil {
		return nil, err
	}

	// converts the channelsBytes and tokensMap to channels:
	channels, err := app.toChannels(channelsBytes, tokensMap)
	if err != nil {
		return nil, err
	}

	// convert the script tokens and rootTokenName to a root Token instance:
	rootToken, err := app.toToken(rootTokenName, tokensMap, []string{})
	if err != nil {
		return nil, err
	}

	builder := app.grammarBuilder.Create().WithRoot(rootToken)
	if channels != nil {
		builder.WithChannels(channels)
	}

	return builder.Now()
}

func (app *adapter) toChannels(channelBytes map[int][]byte, scriptTokensMap map[string]*scriptToken) (channels.Channels, error) {
	list := []channels.Channel{}
	for _, oneChannelBytes := range channelBytes {
		chanIns, err := app.fetchChannel(oneChannelBytes, scriptTokensMap)
		if err != nil {
			return nil, err
		}

		list = append(list, chanIns)
	}

	if len(list) <= 0 {
		return nil, nil
	}

	return app.channelsBuilder.Create().WithList(list).Now()
}

func (app *adapter) fetchChannel(input []byte, scriptTokensMap map[string]*scriptToken) (channels.Channel, error) {
	tokenName, remaining, err := app.fetchTokenName(input)
	if err != nil {
		return nil, err
	}

	token, err := app.toToken(tokenName, scriptTokensMap, []string{})
	if err != nil {
		return nil, err
	}

	builder := app.channelBuilder.Create().WithToken(token)
	if len(remaining) > 0 {
		condition, err := app.fetchChannelCondition(remaining, scriptTokensMap)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *adapter) fetchChannelCondition(input []byte, scriptTokensMap map[string]*scriptToken) (channels.Condition, error) {
	if input[0] == app.channelConditionPrevious {
		tokenName, remaining, err := app.fetchTokenName(input[1:])
		if err != nil {
			return nil, err
		}

		token, err := app.toToken(tokenName, scriptTokensMap, []string{})
		if err != nil {
			return nil, err
		}

		if len(remaining) <= 0 {
			return app.conditionBuilder.Create().WithPrevious(token).Now()
		}
	}

	if input[0] == app.channelConditionNext {
		tokenName, remaining, err := app.fetchTokenName(input[1:])
		if err != nil {
			return nil, err
		}

		token, err := app.toToken(tokenName, scriptTokensMap, []string{})
		if err != nil {
			return nil, err
		}

		if len(remaining) <= 0 {
			return app.conditionBuilder.Create().WithNext(token).Now()
		}
	}

	return nil, errors.New("the data could not match any previous or next condition")
}

func (app *adapter) toToken(rootTokenName string, scriptTokensMap map[string]*scriptToken, path []string) (tokens.Token, error) {
	if rootScriptToken, ok := scriptTokensMap[rootTokenName]; ok {
		linesList := []tokens.Line{}
		for _, oneLine := range rootScriptToken.lines {
			elementsList := []tokens.ElementWithCardinality{}
			for _, oneValue := range oneLine.values {
				if oneValue.pByte != nil {
					element, err := app.elementBuilder.Create().WithByte(*oneValue.pByte).Now()
					if err != nil {
						return nil, err
					}

					elementWithCard, err := app.elementWithCardinalityBuilder.Create().WithElement(element).WithCardinality(oneValue.cardinality).Now()
					if err != nil {
						return nil, err
					}

					elementsList = append(elementsList, elementWithCard)
					continue
				}

				isReference := false
				path = append(path, rootTokenName)
				for _, onePrevName := range path {
					if oneValue.tokenName == onePrevName {
						isReference = true
						break
					}
				}

				if isReference {
					if refScriptToken, ok := scriptTokensMap[oneValue.tokenName]; ok {
						element, err := app.elementBuilder.Create().WithReference(refScriptToken.name).Now()
						if err != nil {
							return nil, err
						}

						elementWithCard, err := app.elementWithCardinalityBuilder.Create().WithElement(element).WithCardinality(oneValue.cardinality).Now()
						if err != nil {
							return nil, err
						}

						elementsList = append(elementsList, elementWithCard)
						continue
					}

					str := fmt.Sprintf("the referenced token (name: %s) is undefined", oneValue.tokenName)
					return nil, errors.New(str)

				}

				subToken, err := app.toToken(oneValue.tokenName, scriptTokensMap, path)
				if err != nil {
					return nil, err
				}

				element, err := app.elementBuilder.Create().WithToken(subToken).Now()
				if err != nil {
					return nil, err
				}

				elementWithCard, err := app.elementWithCardinalityBuilder.Create().WithElement(element).WithCardinality(oneValue.cardinality).Now()
				if err != nil {
					return nil, err
				}

				elementsList = append(elementsList, elementWithCard)
				continue
			}

			line, err := app.lineBuilder.Create().WithList(elementsList).Now()
			if err != nil {
				return nil, err
			}

			linesList = append(linesList, line)
		}

		lines, err := app.linesBuilder.Create().WithList(linesList).Now()
		if err != nil {
			return nil, err
		}

		return app.tokenBuilder.Create().WithName(rootScriptToken.name).WithLines(lines).Now()
	}

	str := fmt.Sprintf("the root token (name: %s) is undefined", rootTokenName)
	return nil, errors.New(str)
}

func (app *adapter) fetchChannelsBytes(input []byte) (map[int][]byte, []byte, error) {
	cpt := 0
	isOpen := false
	index := 0
	channelBytes := map[int][]byte{}
	for _, oneInput := range input {
		cpt++
		if !isOpen && (oneInput == app.channelPrefix) {
			isOpen = true
			index = len(channelBytes)
			channelBytes[index] = []byte{}
			continue
		}

		if isOpen && (oneInput == app.channelSuffix) {
			isOpen = false
			continue
		}

		if !isOpen {
			cpt--
			break
		}

		channelBytes[index] = append(channelBytes[index], oneInput)
	}

	return channelBytes, input[cpt:], nil
}

func (app *adapter) fetchRootTokenName(input []byte) (string, []byte, error) {
	if len(input) <= 0 {
		return "", nil, errors.New("the input was NOT expected to be empty while fetching the root token name")
	}

	if input[0] != app.rootPrefix {
		str := fmt.Sprintf("the root prefix (%d) was expected, %d provided", app.rootPrefix, input[0])
		return "", nil, errors.New(str)
	}

	tokenName, remaining, err := app.fetchTokenName(input[1:])
	if err != nil {
		return "", nil, err
	}

	if remaining[0] != app.rootSuffix {
		str := fmt.Sprintf("the root suffix (%d) was expected, %d provided", app.rootSuffix, remaining[0])
		return "", nil, errors.New(str)
	}

	return tokenName, remaining[1:], nil
}

func (app *adapter) toScriptTokens(input []byte) (map[string]*scriptToken, error) {
	remainingInput := input
	tokens := map[string]*scriptToken{}
	for {

		if len(remainingInput) <= 0 {
			break
		}

		scriptToken, remaining, err := app.toScriptToken(remainingInput)
		if err != nil {
			return nil, err
		}

		remainingInput = remaining
		tokens[scriptToken.name] = scriptToken
	}

	return tokens, nil
}

func (app *adapter) toScriptToken(input []byte) (*scriptToken, []byte, error) {
	tokenName, remainingAfterTokenName, err := app.fetchTokenName(input[0:])
	if err != nil {
		return nil, nil, err
	}

	scriptLines, remainingAfterLines, err := app.fetchLines(remainingAfterTokenName)
	if err != nil {
		str := fmt.Sprintf("there was an error while fetching lines in token (name: %s), error: %s", tokenName, err.Error())
		return nil, nil, errors.New(str)
	}

	return &scriptToken{
		name:  tokenName,
		lines: scriptLines,
	}, remainingAfterLines, nil
}

func (app *adapter) fetchLines(input []byte) ([]*scriptLine, []byte, error) {
	if len(input) < 1 {
		return nil, nil, errors.New("the input was NOT expected to be empty while fetching the line values")
	}

	if input[0] != app.linesPrefix {
		str := fmt.Sprintf("the first element of the input was expected to be the linesPrefix (%d), %d provided", app.linesPrefix, input[0])
		return nil, nil, errors.New(str)
	}

	remainingInput := input[1:]
	values := []*scriptValue{}
	lines := []*scriptLine{}
	for {
		value, retRemaining, err := app.fetchValue(remainingInput)
		if err != nil {
			str := fmt.Sprintf("there is an error while fetching the line's element (line: %d, element: %d), error: %s", len(lines), len(values), err.Error())
			return nil, nil, errors.New(str)
		}

		if len(retRemaining) <= 0 {
			str := fmt.Sprintf("the first element of the input was NOT expected to be empty while fetching line (index: %d)", len(lines))
			return nil, nil, errors.New(str)
		}

		values = append(values, value)
		if retRemaining[0] == app.lineDelimiter {
			remainingInput = retRemaining[1:]
			lines = append(lines, &scriptLine{
				values,
			})

			values = []*scriptValue{}
			continue
		}

		if retRemaining[0] == app.linesSuffix {
			remainingInput = retRemaining[1:]
			lines = append(lines, &scriptLine{
				values,
			})

			break
		}

		remainingInput = retRemaining
	}

	return lines, remainingInput, nil
}

func (app *adapter) fetchValue(input []byte) (*scriptValue, []byte, error) {
	pByte, tokenName, remaining, err := app.fetchElement(input)
	if err != nil {
		return nil, nil, err
	}

	cardinality, remainingAfterCardinality, err := app.fetchCardinality(remaining)
	if err != nil {
		return nil, nil, err
	}

	return &scriptValue{
		pByte:       pByte,
		tokenName:   tokenName,
		cardinality: cardinality,
	}, remainingAfterCardinality, nil
}

func (app *adapter) fetchCardinality(input []byte) (cardinality.Cardinality, []byte, error) {
	if len(input) <= 0 {
		return nil, nil, errors.New("the input was NOT expected to be empty while fetching the element's cardinality")
	}

	remaining := input
	builder := app.cardinalityBuilder.Create()
	if input[0] == app.cardinalityNonZeroMultiple {
		builder.WithMinimum(1)
		remaining = input[1:]
	}

	if input[0] == app.cardinalityZeroMultiple {
		builder.WithMinimum(0)
		remaining = input[1:]
	}

	if input[0] == app.cardinalityOptional {
		builder.WithMinimum(0).WithMaximum(1)
		remaining = input[1:]
	}

	if input[0] == app.cardinalityRangePrefix {
		pMin, pMax, retRemaining, err := app.fetchCardinalityRange(input[1:])
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

func (app *adapter) fetchCardinalityRange(input []byte) (*uint8, *uint8, []byte, error) {
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
		if oneInputByte == app.cardinalityRangeSeparator {
			isSpecific = false
			break
		}

		if oneInputByte == app.cardinalityRangeSuffix {
			break
		}

		if !app.isBytePresent(oneInputByte, app.numbersCharacters) {
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

func (app *adapter) fetchElement(input []byte) (*byte, string, []byte, error) {
	if len(input) <= 0 {
		return nil, "", nil, errors.New("the input was NOT expected to be empty while fetching the element")
	}

	if input[0] == app.bytePrefix {
		pUint, remaining, err := app.fetchNumber(input[1:])
		if err != nil {
			return nil, "", nil, err
		}

		if *pUint >= 256 {
			str := fmt.Sprintf("the byte in the given element cannot exceed 256, %d provided", *pUint)
			return nil, "", nil, errors.New(str)
		}

		byteValue := byte(*pUint)
		return &byteValue, "", remaining, nil
	}

	if input[0] == app.tokenNamePrefix {
		str, retRemaining, err := app.fetchTokenName(input[1:])
		if err != nil {
			return nil, "", nil, err
		}

		return nil, str, retRemaining, nil
	}

	str := fmt.Sprintf("the first element of the input was expecting either a bytePrefix (%d) or a tokenNamePrefix (%d) while fetching an element, %d provided", app.bytePrefix, app.tokenNamePrefix, input[0])
	return nil, "", nil, errors.New(str)
}

func (app *adapter) fetchTokenName(input []byte) (string, []byte, error) {
	nameBytes := []byte{}
	for _, oneInputByte := range input {
		if !app.isBytePresent(oneInputByte, app.tokenNameCharacters) {
			break
		}

		nameBytes = append(nameBytes, oneInputByte)
	}

	if len(nameBytes) <= 0 {
		return "", nil, errors.New("the tokenName must contain at least 1 character, none provided")
	}

	return string(nameBytes), input[len(nameBytes):], nil
}

func (app *adapter) fetchNumber(input []byte) (*uint, []byte, error) {
	indexBytes := []byte{}
	for _, oneInputByte := range input {
		if !app.isBytePresent(oneInputByte, app.numbersCharacters) {
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

func (app *adapter) isBytePresent(value byte, data []byte) bool {
	isPresent := false
	for _, oneChanByte := range data {
		if value == oneChanByte {
			isPresent = true
			break
		}
	}

	return isPresent
}

func (app *adapter) removeChannelCharacters(input []byte) []byte {
	output := []byte{}
	for _, oneInputByte := range input {
		if app.isBytePresent(oneInputByte, app.channelCharacters) {
			continue
		}

		output = append(output, oneInputByte)
	}

	return output
}
