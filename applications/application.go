package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/validator/domain/grammars"
	"github.com/steve-care-software/validator/domain/grammars/channels"
	"github.com/steve-care-software/validator/domain/grammars/tokens"
	"github.com/steve-care-software/validator/domain/results"
)

type application struct {
	grammarAdapter                      grammars.Adapter
	resultBuilder                       results.Builder
	resultTokenBuilder                  results.TokenBuilder
	resultBlockBuilder                  results.BlockBuilder
	resultLineBuilder                   results.LineBuilder
	resultElementWithCardinalityBuilder results.ElementWithCardinalityBuilder
	resultElementBuilder                results.ElementBuilder
}

func createApplication(
	grammarAdapter grammars.Adapter,
	resultBuilder results.Builder,
	resultTokenBuilder results.TokenBuilder,
	resultBlockBuilder results.BlockBuilder,
	resultLineBuilder results.LineBuilder,
	resultElementWithCardinalityBuilder results.ElementWithCardinalityBuilder,
	resultElementBuilder results.ElementBuilder,
) Application {
	out := application{
		grammarAdapter:                      grammarAdapter,
		resultBuilder:                       resultBuilder,
		resultTokenBuilder:                  resultTokenBuilder,
		resultBlockBuilder:                  resultBlockBuilder,
		resultLineBuilder:                   resultLineBuilder,
		resultElementWithCardinalityBuilder: resultElementWithCardinalityBuilder,
		resultElementBuilder:                resultElementBuilder,
	}

	return &out
}

// Compile compiles a script to a validator grammar
func (app *application) Compile(script string) (grammars.Grammar, error) {
	return app.grammarAdapter.ToGrammar(script)
}

// Execute executes the application
func (app *application) Execute(grammar grammars.Grammar, data []byte, canHavePrefix bool) (results.Result, error) {
	token := grammar.Root()
	channels := grammar.Channels()
	externals := grammar.Externals()
	if canHavePrefix {
		index := uint(0)
		reaminingData := data
		for {
			if len(reaminingData) <= 0 {
				break
			}

			retResultToken, err := app.executeOnce(token, externals, channels, reaminingData, index)
			if err == nil && retResultToken.IsSuccess() {
				return app.resultBuilder.Create().WithIndex(index).WithToken(retResultToken).Now()
			}

			reaminingData = reaminingData[1:]
			index++
		}
	}

	retResultToken, err := app.executeOnce(token, externals, channels, data, 0)
	if err != nil {
		return nil, err
	}

	return app.resultBuilder.Create().WithIndex(0).WithToken(retResultToken).Now()
}

func (app *application) executeOnce(
	currentToken tokens.Token,
	externals grammars.Externals,
	chans channels.Channels,
	data []byte,
	index uint,
) (results.Token, error) {
	channelsList := []channels.Channel{}
	if chans != nil {
		channelsList = chans.List()
	}

	retResultToken, _, err := app.executeToken(currentToken, externals, channelsList, data, map[string]*tokenData{})
	if err != nil {
		return nil, err
	}

	return retResultToken, nil
}

func (app *application) executeChannels(
	externals grammars.Externals,
	channelsList []channels.Channel,
	currentData []byte,
	prevTokenData map[string]*tokenData,
) ([]byte, error) {
	executePreviousFn := func(
		input []byte,
		token tokens.Token,
		previous tokens.Token,
	) ([]byte, error) {
		retResToken, _, err := app.executeToken(token, externals, []channels.Channel{}, input, prevTokenData)
		if err != nil {
			return nil, err
		}

		retRemaining := retResToken.Block().Remaining()
		_, _, err = app.executeToken(previous, externals, []channels.Channel{}, retRemaining, prevTokenData)
		if err != nil {
			return nil, err
		}

		return retRemaining, nil
	}

	executeNextFn := func(
		input []byte,
		token tokens.Token,
		next tokens.Token,
	) ([]byte, error) {
		retResToken, _, err := app.executeToken(next, externals, []channels.Channel{}, input, prevTokenData)
		if err != nil {
			return nil, err
		}

		afterNextRemaining := retResToken.Block().Remaining()
		retResTokenSecond, _, err := app.executeToken(token, externals, []channels.Channel{}, afterNextRemaining, prevTokenData)
		if err != nil {
			return nil, err
		}

		retRemainingAmount := len(retResTokenSecond.Block().Remaining())
		amountKept := len(input) - len(afterNextRemaining)
		amountRemoved := len(afterNextRemaining) - retRemainingAmount
		return append(input[:amountKept], input[amountKept+amountRemoved:]...), nil
	}

	executeFn := func(
		input []byte,
	) ([]byte, error) {
		fn := func(data []byte) ([]byte, error) {
			if len(data) <= 0 {
				return []byte{}, nil
			}

			for _, oneChannel := range channelsList {
				if !oneChannel.HasCondition() {
					token := oneChannel.Token()
					retResToken, _, err := app.executeToken(token, externals, []channels.Channel{}, data, prevTokenData)
					if err != nil {
						continue
					}

					block := retResToken.Block()
					data = block.Remaining()
					continue
				}

				token := oneChannel.Token()
				condition := oneChannel.Condition()
				if condition.IsNext() {
					next := condition.Next()
					retRemaining, err := executeNextFn(data, token, next)
					if err != nil {
						continue
					}

					data = retRemaining
					continue
				}

				if condition.IsPrevious() {
					previous := condition.Previous()
					retRemaining, err := executePreviousFn(data, token, previous)
					if err != nil {
						continue
					}

					data = retRemaining
					continue
				}
			}

			return data, nil
		}

		remaining, err := fn(input)
		if err != nil {
			return nil, err
		}

		if len(remaining) == len(input) {
			return remaining, nil
		}

		return fn(remaining)
	}

	return executeFn(currentData)
}

func (app *application) executeReference(
	refName string,
	externals grammars.Externals,
	channels []channels.Channel,
	currentData []byte,
	prevTokenData map[string]*tokenData,
) (results.Token, map[string]*tokenData, error) {
	if tokenData, ok := prevTokenData[refName]; ok {
		prevData := tokenData.Data()
		if len(currentData) == len(prevData) {
			str := fmt.Sprintf("the referenced token (name: %s) is an infinite recursive token", refName)
			return nil, prevTokenData, errors.New(str)
		}

		token := tokenData.Token()
		return app.executeToken(token, externals, channels, currentData, prevTokenData)
	}

	str := fmt.Sprintf("the referenced token (name: %s) is NOT declared", refName)
	return nil, prevTokenData, errors.New(str)
}

func (app *application) executeToken(
	currentToken tokens.Token,
	externals grammars.Externals,
	channels []channels.Channel,
	currentData []byte,
	prevTokenData map[string]*tokenData,
) (results.Token, map[string]*tokenData, error) {
	amountChannels := 0
	if len(channels) > 0 {
		remainingData, err := app.executeChannels(externals, channels, currentData, prevTokenData)
		if err != nil {
			return nil, nil, err
		}

		amountChannels += len(currentData) - len(remainingData)
		currentData = remainingData
	}

	// add the data to the previous token data map:
	name := currentToken.Name()
	prevTokenData[name] = createTokenData(currentToken, currentData)

	lines := currentToken.Lines()
	resultBlock, retTokenData, err := app.executeLines(lines, externals, channels, currentData, prevTokenData)
	if err != nil {
		return nil, nil, err
	}

	if len(channels) > 0 {
		remaining := resultBlock.Remaining()
		remDataAfterChannels, err := app.executeChannels(externals, channels, remaining, prevTokenData)
		if err != nil {
			return nil, nil, err
		}

		amountChannels += len(remaining) - len(remDataAfterChannels)
	}

	resultToken, err := app.resultTokenBuilder.Create().WithName(name).WithBlock(resultBlock).WithChannels(uint(amountChannels)).Now()
	if err != nil {
		return nil, nil, err
	}

	return resultToken, retTokenData, nil
}

func (app *application) executeLines(
	lines tokens.Lines,
	externals grammars.Externals,
	channels []channels.Channel,
	currentData []byte,
	prevTokenData map[string]*tokenData,
) (results.Block, map[string]*tokenData, error) {
	list := lines.List()
	remainingData := currentData
	resultLines := []results.Line{}
	for idx, oneLine := range list {
		retElements, retTokenData, err := app.executeLine(oneLine, externals, channels, remainingData, prevTokenData)
		if err != nil {
			continue
		}

		line, err := app.resultLineBuilder.Create().WithIndex(uint(idx)).WithElements(retElements).Now()
		if err != nil {
			return nil, nil, err
		}

		prevTokenData = retTokenData
		resultLines = append(resultLines, line)
		if line.IsSuccess() {
			remainingData = line.Remaining()
		}
	}

	block, err := app.resultBlockBuilder.Create().WithList(resultLines).WithInput(currentData).Now()
	if err != nil {
		return nil, nil, err
	}

	return block, prevTokenData, nil
}

func (app *application) executeLine(
	line tokens.Line,
	externals grammars.Externals,
	channels []channels.Channel,
	currentData []byte,
	prevTokenData map[string]*tokenData,
) ([]results.ElementWithCardinality, map[string]*tokenData, error) {
	list := line.List()
	remainingData := currentData
	elements := []results.ElementWithCardinality{}
	for index, oneElementWithCard := range list {
		retElWithCard, retTokenData, err := app.executeElementWithCardinality(oneElementWithCard, externals, channels, remainingData, prevTokenData)
		if err != nil {
			str := fmt.Sprintf("there was an error while executing line (index: %d): error: %s", index, err.Error())
			return nil, nil, errors.New(str)
		}

		remainingData = retElWithCard.Remaining()
		prevTokenData = retTokenData
		elements = append(elements, retElWithCard)
	}

	return elements, prevTokenData, nil
}

func (app *application) executeElementWithCardinality(
	elementWithCard tokens.ElementWithCardinality,
	externals grammars.Externals,
	channels []channels.Channel,
	currentData []byte,
	prevTokenData map[string]*tokenData,
) (results.ElementWithCardinality, map[string]*tokenData, error) {
	remainingData := currentData
	element := elementWithCard.Element()
	cardinality := elementWithCard.Cardinality()

	cpt := uint(0)
	min := uint(cardinality.Min())
	matches := []results.Element{}
	for {

		if len(remainingData) <= 0 {
			break
		}

		if cardinality.HasMax() {
			pMax := cardinality.Max()
			if cpt >= uint(*pMax) {
				break
			}
		}

		retRemainingData, retResultElement, retTokenData, err := app.executeElement(element, externals, channels, remainingData, prevTokenData)
		if err != nil {
			break
		}

		if len(remainingData) == len(retRemainingData) {
			break
		}

		remainingData = retRemainingData
		prevTokenData = retTokenData
		matches = append(matches, retResultElement)
		cpt++
	}

	missing := uint(0)
	if cpt < min {
		missing = min - cpt
	}

	ins, err := app.resultElementWithCardinalityBuilder.Create().
		WithMissing(missing).
		WithElement(element).
		WithMatches(matches).
		WithRemaining(remainingData).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, prevTokenData, nil
}

func (app *application) executeElement(
	element tokens.Element,
	externals grammars.Externals,
	channels []channels.Channel,
	currentData []byte,
	prevTokenData map[string]*tokenData,
) ([]byte, results.Element, map[string]*tokenData, error) {
	elementBuilder := app.resultElementBuilder.Create()
	if element.IsByte() {
		pByte := element.Byte()
		if len(currentData) > 0 {
			first := currentData[0]
			if *pByte != first {
				str := fmt.Sprintf("the element byte (%d) could not match the first data byte (%d)", *pByte, first)
				return nil, nil, nil, errors.New(str)
			}

			ins, err := elementBuilder.WithValue(*pByte).Now()
			if err != nil {
				return nil, nil, nil, err
			}

			return currentData[1:], ins, prevTokenData, nil
		}

		str := fmt.Sprintf("the byte (%d) could not be found in the data because the remaining data was empty", *pByte)
		return nil, nil, nil, errors.New(str)
	}

	if element.IsToken() {
		token := element.Token()
		resultToken, retTokenData, err := app.executeToken(token, externals, channels, currentData, prevTokenData)
		if err != nil {
			return nil, nil, nil, err
		}

		ins, err := elementBuilder.WithToken(resultToken).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		return resultToken.Block().Remaining(), ins, retTokenData, nil
	}

	if element.IsExternal() {
		name := element.External()
		external, err := externals.Find(name)
		if err != nil {
			return nil, nil, nil, err
		}

		grammar := external.Grammar()
		result, err := app.Execute(grammar, currentData, false)
		if err != nil {
			return nil, nil, nil, err
		}

		resultToken := result.Token()
		resultTokenBlock := resultToken.Block()
		resultTokenChannels := resultToken.Channels()
		token, err := app.resultTokenBuilder.Create().WithName(name).WithBlock(resultTokenBlock).WithChannels(resultTokenChannels).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		ins, err := elementBuilder.WithToken(token).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		return resultTokenBlock.Remaining(), ins, prevTokenData, nil
	}

	reference := element.Reference()
	resultToken, retTokenData, err := app.executeReference(reference, externals, channels, currentData, prevTokenData)
	if err != nil {
		return nil, nil, nil, err
	}

	ins, err := elementBuilder.WithToken(resultToken).Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return resultToken.Block().Remaining(), ins, retTokenData, nil
}
