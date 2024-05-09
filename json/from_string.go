package json

func FromString(input string) (interface{}, error) {
	tokens, err := Lex(input)

	if err != nil {
		return nil, err
	}

	result, _, err := Parse(tokens)

	return result, err
}
