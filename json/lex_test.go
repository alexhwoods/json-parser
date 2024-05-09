package json

import (
	"reflect"
	"testing"
)

func TestLexString(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		expectedValue     string
		expectedRemainder string
		expectSuccess     bool
	}{
		{"Simple Strings", `"foo"`, "foo", "", true},
		{"Empty Strings", `""`, "", "", true},
		{"Strings with spaces", `"foo bar"`, "foo bar", "", true},
		{"Numbers", `123`, "", "123", false},
		{"Empty", ``, "", "", false},
		{"Boolean", `true`, "", "true", false},
		{"Number with extra", `5}`, "", "5}", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value, remainder, isString := LexString(tc.input)
			if isString != tc.expectSuccess {
				t.Errorf("Expected success flag for %s to be %v, got %v", tc.name, tc.expectSuccess, isString)
			}
			if remainder != tc.expectedRemainder {
				t.Errorf("Expected remainder for %s to be %q, but got %q", tc.name, tc.expectedRemainder, remainder)
			}
			if value != tc.expectedValue {
				t.Errorf("Expected value for %s to be %q, but got %q", tc.name, tc.expectedValue, value)
			}
		})
	}
}

func TestLexNumber(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		expectedValue     interface{}
		expectedRemainder string
		expectSuccess     bool
	}{
		// {"Integer With Single Digit", `5`, 5, "", true},
		// {"Integer With Multiple Digits", `21`, 21, "", true},
		// {"Floating Point Number", `23.4`, 23.4, "", true},
		// {"String", `"hello"`, "", `"hello"`, false},
		// {"Empty", ``, "", "", false},
		// {"Boolean", `true`, "", "true", false},
		{"Number with extra", `5}`, 5, "}", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value, remainder, success := LexNumber(tc.input)
			if success != tc.expectSuccess {
				t.Errorf("Expected success flag for %s to be %v, got %v", tc.name, tc.expectSuccess, success)
			}
			if remainder != tc.expectedRemainder {
				t.Errorf("Expected remainder for %s to be %q, but got %q", tc.name, tc.expectedRemainder, remainder)
			}
			if value != tc.expectedValue {
				t.Errorf("Expected value for %s to be %q, but got %q", tc.name, tc.expectedValue, value)
			}
		})
	}
}

func TestLexBoolean(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		expectedValue     interface{}
		expectedRemainder string
		expectSuccess     bool
	}{
		{"true", `true`, true, "", true},
		{"false", `false`, false, "", true},
		{"Floating Point Number", `23.4`, "", "23.4", false},
		{"String", `"hello"`, "", `"hello"`, false},
		{"Empty", ``, "", "", false},
		{"null", "null", "", "null", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value, remainder, success := LexBoolean(tc.input)
			if success != tc.expectSuccess {
				t.Errorf("Expected success flag for %s to be %v, got %v", tc.name, tc.expectSuccess, success)
			}
			if remainder != tc.expectedRemainder {
				t.Errorf("Expected remainder for %s to be %q, but got %q", tc.name, tc.expectedRemainder, remainder)
			}
			if value != tc.expectedValue {
				t.Errorf("Expected value for %s to be %q, but got %q", tc.name, tc.expectedValue, value)
			}
		})
	}
}

func TestLexNull(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		expectedValue     interface{}
		expectedRemainder string
		expectSuccess     bool
	}{
		{"null", "null", nil, "", true},
		{"null with extra", "null, 5", nil, ", 5", true},
		{"not null", `"hello"`, "", `"hello"`, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value, remainder, success := LexNull(tc.input)
			if success != tc.expectSuccess {
				t.Errorf("Expected success flag for %s to be %v, got %v", tc.name, tc.expectSuccess, success)
			}
			if remainder != tc.expectedRemainder {
				t.Errorf("Expected remainder for %s to be %q, but got %q", tc.name, tc.expectedRemainder, remainder)
			}
			if value != tc.expectedValue {
				t.Errorf("Expected value for %s to be %q, but got %q", tc.name, tc.expectedValue, value)
			}
		})
	}
}

func TestLex(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedTokens interface{}
	}{
		{"Basic JSON", `{"foo": 5}`, []interface{}{'{', "foo", ',', ':', '5', '}'}},
		{"Nested JSON", `{"foo": [1, 2, {"bar": 2}]}`, []interface{}{'{', "foo", ':', '[', '1', ',', '2', ',', '{', "bar", ':', '2', '}', ']'}},
		{"Empty JSON", `{}`, []interface{}{'{', '}'}},
		{"Empty Array", `[]`, []interface{}{'[', ']'}},
		{"Array with one element", `[1]`, []interface{}{'[', '1', ']'}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := Lex(tc.input)

			if err != nil {
				t.Errorf("Expected lex to return no error, but got %v", err)
			}

			if reflect.DeepEqual(tokens, []interface{}{tc.expectedTokens}) {
				t.Errorf("Expected tokens for %s to be %v, but got %v", tc.name, tc.expectedTokens, tokens)
			}
		})
	}
}
