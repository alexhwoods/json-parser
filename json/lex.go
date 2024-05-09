package json

import (
	"errors"
	"strconv"
	"strings"
)

const JSON_WHITESPACE = " \t\n"
const JSON_SYNTAX = "{}[],:"
const JSON_QUOTE = '"'
const JSON_COMMA = ','

func LexString(input string) (string, string, bool) {
	jsonString := ""
	s := input

	if len(s) == 0 {
		return "", "", false
	}

	if input[0] != JSON_QUOTE {
		return "", input, false
	} else {
		s = s[1:]
	}

	for _, c := range s {
		if c == JSON_QUOTE {
			s = s[len(jsonString)+1:]
			return jsonString, s, true
		}

		jsonString += string(c)
	}

	panic("unreachable")
}

const NUMBER_SYNTAX = "-0123456789."
const NUMBER_DOT = '.'

func LexNumber(input string) (interface{}, string, bool) {
	numberString := ""
	s := input

	if len(s) == 0 {
		return "", "", false
	}

	for _, c := range s {
		if c == JSON_COMMA {
			break
		}

		if !strings.ContainsRune(NUMBER_SYNTAX, c) && len(numberString) == 0 {
			return "", input, false
		} else if !strings.ContainsRune(NUMBER_SYNTAX, c) {
			break
		} else {
			numberString += string(c)
		}
	}

	s = s[len(numberString):]

	if strings.ContainsRune(numberString, NUMBER_DOT) {
		value, err := strconv.ParseFloat(numberString, 64)

		if err != nil {
			return "", input, false
		}

		return value, s, true
	}

	value, err := strconv.Atoi(numberString)

	if err != nil {
		return "", input, false
	}

	return value, s, true
}

func LexBoolean(input string) (interface{}, string, bool) {
	if strings.HasPrefix(input, "true") {
		return true, input[4:], true
	}

	if strings.HasPrefix(input, "false") {
		return false, input[5:], true
	}

	return "", input, false
}

func LexNull(input string) (interface{}, string, bool) {
	if strings.HasPrefix(input, "null") {
		return nil, input[4:], true
	}

	return "", input, false
}

func lex(input string) ([]interface{}, error) {
	tokens := make([]interface{}, 0)

	s := input

	for len(s) > 0 {
		jsonString, remainder, success := LexString(s)
		if success {
			tokens = append(tokens, jsonString)
			s = remainder
			continue
		}

		jsonNumber, remainder, success := LexNumber(s)
		if success {
			tokens = append(tokens, jsonNumber)
			s = remainder
			continue
		}

		jsonBoolean, remainder, success := LexBoolean(s)
		if success {
			tokens = append(tokens, jsonBoolean)
			s = remainder
			continue
		}

		jsonNull, remainder, success := LexNull(s)
		if success {
			tokens = append(tokens, jsonNull)
			s = remainder
			continue
		}

		c := rune(s[0])

		if strings.ContainsRune(JSON_WHITESPACE, c) {
			s = s[1:]
		} else if strings.ContainsRune(JSON_SYNTAX, c) {
			tokens = append(tokens, c)
			s = s[1:]
		} else {
			return nil, errors.New("Unexpected character: " + string(c))
		}
	}

	return tokens, nil
}
