package applications

import (
	"reflect"
	"testing"
)

func TestValidator_withReference_withSuccessIndex_withChannels_isSuccess(t *testing.T) {
	script := `
		%rootToken;
		-space;
		-endOfLine;
		-star <openParenthesis;
		-star <smallerThan;
		-plus >closeParenthesis;

		rootToken : .openParenthesis .rootToken .closeParenthesis
				  | .five .smallerThan .five
				  ;

		openParenthesis: $40;
		closeParenthesis: $41;
		five: $53;
		smallerThan: $60;
		space: $32;
		endOfLine: $10;
		star: $42;
		plus: $43;
	`

	data := []byte("(*( 5 *< 5 )+) 567")
	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := NewApplication().Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 14 {
		t.Errorf("the cursor was expected to be %d, %d returned", 14, cursor)
		return
	}

	if !result.Token().IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"rootToken", "openParenthesis", "rootToken", "openParenthesis", "rootToken", "five", "smallerThan", "five", "closeParenthesis", "closeParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withReference_withSuccessIndex_isSuccess(t *testing.T) {
	script := `
		%rootToken;
		rootToken : .openParenthesis .rootToken .closeParenthesis
				  | .five .smallerThan .five
				  ;

		openParenthesis: $40;
		closeParenthesis: $41;
		five: $53;
		smallerThan: $60;
	`

	data := []byte("((5<5))567")
	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 7 {
		t.Errorf("the cursor was expected to be %d, %d returned", 7, cursor)
		return
	}

	if !result.Token().IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"rootToken", "openParenthesis", "rootToken", "openParenthesis", "rootToken", "five", "smallerThan", "five", "closeParenthesis", "closeParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withReference_withSuccessIndex_notEnoughData_cannotHavePrefix_isNotSuccess(t *testing.T) {
	script := `
		%rootToken;
		rootToken : .openParenthesis .rootToken .closeParenthesis
					| .five .smallerThan .five
					;

		openParenthesis: $40;
		closeParenthesis: $41;
		five: $53;
		smallerThan: $60;
	`

	data := []byte("((5<5)")
	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, data, false)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 0 {
		t.Errorf("the index was expected to be %d,%d returned", 0, index)
		return
	}

	cursor := result.Cursor()
	if cursor != 0 {
		t.Errorf("the cursor was expected to be %d, %d returned", 0, cursor)
		return
	}

	if result.Token().IsSuccess() {
		t.Errorf("the result was expected to NOT be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"rootToken"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withReference_withSuccessIndex_notEnoughData_withPrefix_isSuccess(t *testing.T) {
	script := `
		%rootToken;
		rootToken : .openParenthesis .rootToken .closeParenthesis
					| .five .smallerThan .five
					;

		openParenthesis: $40;
		closeParenthesis: $41;
		five: $53;
		smallerThan: $60;
	`

	data := []byte("((5<5)")
	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 1 {
		t.Errorf("the index was expected to be %d,%d returned", 1, index)
		return
	}

	cursor := result.Cursor()
	if cursor != 6 {
		t.Errorf("the cursor was expected to be %d, %d returned", 6, cursor)
		return
	}

	if !result.Token().IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"rootToken", "openParenthesis", "rootToken", "five", "smallerThan", "five", "closeParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withReference_isInfiniteRecursive_isNotSuccess(t *testing.T) {
	script := `
		%rootToken;
		rootToken : .rootToken;
	`

	data := []byte("((5<5))")
	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 0 {
		t.Errorf("the cursor was expected to be %d, %d returned", 0, cursor)
		return
	}

	if result.Token().IsSuccess() {
		t.Errorf("the result was expected to NOT be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"rootToken"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withOneLine_withSpecificCardinality_withSubTokens_withSuccessIndex_isSuccess(t *testing.T) {
	script := `
		%rootToken;
		rootToken : .openParenthesis .hyphen .closeParenthesis;
		openParenthesis: $40;
		hyphen: $45;
		closeParenthesis: $41;
	`

	data := []byte("(-)345")
	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 3 {
		t.Errorf("the cursor was expected to be %d, %d returned", 3, cursor)
		return
	}

	if !result.Token().IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"rootToken", "openParenthesis", "hyphen", "closeParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}

}

func TestValidator_withOneLine_withSpecificCardinality_withByte_withoutSuccessIndex_isSuccess(t *testing.T) {
	script := `
		%openParenthesis;
		openParenthesis : $40;
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 1 {
		t.Errorf("the cursor was expected to be %d, %d returned", 1, cursor)
		return
	}

	if !result.Token().IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"openParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withOneLine_withMinimumCardinality_withByte_withExactlyMinOccurences_isSuccess(t *testing.T) {

	script := `
		%openParenthesis;
		openParenthesis : $40[2,];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("(("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 2 {
		t.Errorf("the cursor was expected to be %d, %d returned", 2, cursor)
		return
	}

	if !result.Token().IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"openParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withOneLine_withMinimumCardinality_withByte_withMinimumPlusOccurences_isSuccess(t *testing.T) {
	script := `
		%openParenthesis;
		openParenthesis : $40[2,];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("((("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 3 {
		t.Errorf("the cursor was expected to be %d, %d returned", 3, cursor)
		return
	}

	if !result.Token().IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"openParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withOneLine_withMinimumCardinality_withByte_withLessThanMinimum_isNotSuccess(t *testing.T) {
	script := `
		%openParenthesis;
		openParenthesis : $40[2,];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 0 {
		t.Errorf("the cursor was expected to be %d, %d returned", 0, cursor)
		return
	}

	if result.Token().IsSuccess() {
		t.Errorf("the result was expected to NOT be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"openParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withOneLine_withRangeCardinality_withByte_withMaximumExcceeded_withPrefix_isSuccess(t *testing.T) {
	script := `
		%openParenthesis;
		openParenthesis : $40[2,5];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("(((((("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 0 {
		t.Errorf("the index was expected to be %d, %d returned", 0, index)
		return
	}

	cursor := result.Cursor()
	if cursor != 5 {
		t.Errorf("the cursor was expected to be %d, %d returned", 5, cursor)
		return
	}

	if !result.Token().IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"openParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withOneLine_withRangeCardinality_withByte_withExactlyMaximumOccurences_isSuccess(t *testing.T) {
	script := `
		%openParenthesis;
		openParenthesis : $40[2,5];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("((((("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 5 {
		t.Errorf("the cursor was expected to be %d, %d returned", 5, cursor)
		return
	}

	if !result.Token().IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"openParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestValidator_withOneLine_withRangeCardinality_withByte_withinRangeOccurences_isSuccess(t *testing.T) {
	script := `
		%openParenthesis;
		openParenthesis : $40[2,5];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("(((("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 4 {
		t.Errorf("the cursor was expected to be %d, %d returned", 4, cursor)
		return
	}

	if !result.Token().IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Token().Path()
	expectedPath := []string{"openParenthesis"}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}
