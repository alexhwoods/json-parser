package json

import (
	"errors"
)

func ParseObject(tokens []interface{}) (map[string]interface{}, []interface{}, error) {
	m := make(map[string]interface{})

	if len(tokens) == 0 {
		return nil, nil, errors.New("empty object")
	}

	if tokens[0] != '{' {
		return nil, nil, errors.New("expected opening bracket")
	}

	if tokens[1] == '}' {
		return m, tokens[2:], nil
	}

	tokens = tokens[1:]

	for {
		key := tokens[0]
		if _, ok := key.(string); ok {
			tokens = tokens[1:]
		} else {
			return nil, nil, errors.New("expected string key")
		}

		if tokens[0] != ':' {
			return nil, nil, errors.New("expected colon")
		}

		tokens = tokens[1:]

		var value interface{}
		var err error
		value, tokens, err = Parse(tokens)

		if err != nil {
			return nil, nil, err
		}

		m[key.(string)] = value

		c := tokens[0]

		if c == '}' {
			return m, tokens[1:], nil
		} else if c == ',' {
			tokens = tokens[1:]
		} else {
			return nil, nil, errors.New("expected comma or closing bracket")
		}
	}
}

func ParseArray(tokens []interface{}) ([]interface{}, []interface{}, error) {
	m := make([]interface{}, 0)

	if len(tokens) == 0 {
		return nil, nil, errors.New("empty object")
	}

	if tokens[0] != '[' {
		return nil, nil, errors.New("expected opening square bracket")
	}

	if tokens[1] == ']' {
		return m, tokens[2:], nil
	}

	tokens = tokens[1:]

	for {
		var value interface{}
		var err error
		value, tokens, err = Parse(tokens)

		if err != nil {
			return nil, nil, err
		}

		m = append(m, value)

		c := tokens[0]

		if c == ']' {
			return m, tokens[1:], nil
		} else if c == ',' {
			tokens = tokens[1:]
		} else if c == '[' {
			value, tokens, err = ParseArray(tokens)

			if err != nil {
				return nil, nil, err
			}

			m = append(m, value)
		} else {
			return nil, nil, errors.New("expected comma or closing square bracket")
		}
	}
}

func Parse(tokens []interface{}) (interface{}, []interface{}, error) {
	if len(tokens) == 0 {
		return nil, nil, errors.New("empty input")
	}

	t := tokens[0]

	if t == '{' {
		return ParseObject(tokens)
	} else if t == '[' {
		return ParseArray(tokens)
	} else {
		return t, tokens[1:], nil
	}
}
