package json

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseObject(t *testing.T) {
	tests := []struct {
		name           string
		input          []interface{}
		expectedResult interface{}
	}{

		// corresponds to {"foo": 5}
		{"Basic JSON Object", []interface{}{'{', "foo", ":", 5, '}'}, map[string]interface{}{"foo": 5}},
		// corresponds to {}
		{"Empty JSON Object", []interface{}{'{', '}'}, map[string]interface{}{}},
		// // corresponds to {"foo": {"bar": 5}}
		{"Nested JSON Object 2", []interface{}{'{', "foo", ":", map[string]interface{}{"bar": 5}, '}'}, map[string]interface{}{"foo": map[string]interface{}{"bar": 5}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, _, err := ParseObject(tc.input)

			if err != nil {
				t.Errorf("Expected ParseObject to return no error, but got %v", err)
			}

			if !cmp.Equal(result, tc.expectedResult) {
				t.Logf(cmp.Diff(result, tc.expectedResult))
				t.Errorf("Expected tokens for %s to be %v, but got %v", tc.name, tc.expectedResult, result)
			}
		})
	}
}

func TestParseArray(t *testing.T) {
	tests := []struct {
		name           string
		input          []interface{}
		expectedResult interface{}
	}{

		// corresponds to [1,2,3]
		{"Basic JSON Array", []interface{}{'[', 1, ',', 2, ',', 3, ']'}, []interface{}{1, 2, 3}},
		// corresponds to []
		{"Empty JSON Array", []interface{}{'[', ']'}, []interface{}{}},
		// // corresponds to [1, [2, 3]]
		{"Nested JSON Array", []interface{}{'[', 1, ',', '[', 2, ',', 3, ']', ']'}, []interface{}{1, []interface{}{2, 3}}},
		// // corresponds to [{"foo": 5}]
		{"Array with Object", []interface{}{'[', map[string]interface{}{"foo": 5}, ']'}, []interface{}{map[string]interface{}{"foo": 5}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, _, err := ParseArray(tc.input)

			if err != nil {
				t.Errorf("Expected ParseArray to return no error, but got %v", err)
			}

			if !cmp.Equal(result, tc.expectedResult) {
				t.Logf(cmp.Diff(result, tc.expectedResult))
				t.Errorf("Expected tokens for %s to be %v, but got %v", tc.name, tc.expectedResult, result)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name           string
		input          []interface{}
		expectedResult interface{}
	}{
		// corresponds to {"foo": 5}
		{"Basic JSON Object", []interface{}{'{', "foo", ":", 5, '}'}, map[string]interface{}{"foo": 5}},
		// corresponds to {}
		{"Empty JSON Object Simple", []interface{}{'{', '}'}, map[string]interface{}{}},
		// corresponds to {"foo": {"bar": 5}}
		{"Nested JSON Object Complex", []interface{}{'{', "foo", ":", map[string]interface{}{"bar": 5}, '}'}, map[string]interface{}{"foo": map[string]interface{}{"bar": 5}}},
		// corresponds to [1,2,3]
		{"Basic JSON Array", []interface{}{'[', 1, ',', 2, ',', 3, ']'}, []interface{}{1, 2, 3}},
		// corresponds to []
		{"Empty JSON Array Simple", []interface{}{'[', ']'}, []interface{}{}},
		// corresponds to [1, [2, 3]]
		{"Nested JSON Array Complex", []interface{}{'[', 1, ',', '[', 2, ',', 3, ']', ']'}, []interface{}{1, []interface{}{2, 3}}},
		// corresponds to [{"foo": 5}]
		{"Array with Object", []interface{}{'[', map[string]interface{}{"foo": 5}, ']'}, []interface{}{map[string]interface{}{"foo": 5}}},
		// corresponds to {"foo": [1, 2, {"bar": 2}]
		{"Nested JSON Object", []interface{}{'{', "foo", ":", []interface{}{1, 2, map[string]interface{}{"bar": 2}}, '}'}, map[string]interface{}{"foo": []interface{}{1, 2, map[string]interface{}{"bar": 2}}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, _, err := Parse(tc.input)

			if err != nil {
				t.Errorf("Expected Parse to return no error, but got %v", err)
			}

			if !cmp.Equal(result, tc.expectedResult) {
				t.Logf(cmp.Diff(result, tc.expectedResult))
				t.Errorf("Expected result for %s to be %v, but got %v", tc.name, tc.expectedResult, result)
			}
		})
	}
}
