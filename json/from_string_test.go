package json

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedResult interface{}
	}{
		// corresponds to {"foo": 5}
		{"Basic JSON Object", `{"foo": 5}`, map[string]interface{}{"foo": 5}},
		// corresponds to {}
		{"Empty JSON Object", `{}`, map[string]interface{}{}},
		// corresponds to {"foo": {"bar": 5}}
		{"Nested JSON Object 2", `{"foo": {"bar": 5}}`, map[string]interface{}{"foo": map[string]interface{}{"bar": 5}}},
		// corresponds to [1,2,3]
		{"Basic JSON Array", `[1,2,3]`, []interface{}{1, 2, 3}},
		// corresponds to []
		{"Empty JSON Array", `[]`, []interface{}{}},
		// corresponds to [1, [2, 3]]
		{"Nested JSON Array", `[1, [2, 3]]`, []interface{}{1, []interface{}{2, 3}}},
		// corresponds to [{"foo": 5}]
		{"Array with Object", `[{"foo": 5}]`, []interface{}{map[string]interface{}{"foo": 5}}},
		// corresponds to {"foo": [1,2,{"bar": 2}]}
		{"Nested JSON", `{"foo": [1, 2, {"bar": 2}]}`, map[string]interface{}{"foo": []interface{}{1, 2, map[string]interface{}{"bar": 2}}}},
		// corresponds to ["foo", 50, true, null]
		{"Array with String, Number, Boolean, and Null", `["foo", 50, true, null]`, []interface{}{"foo", 50, true, nil}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := FromString(tc.input)

			if err != nil {
				t.Errorf("Expected FromString to return no error, but got %v", err)
			}

			if !cmp.Equal(result, tc.expectedResult) {
				t.Logf(cmp.Diff(result, tc.expectedResult))
				t.Errorf("Expected result for %s to be %v, but got %v", tc.name, tc.expectedResult, result)
			}
		})
	}
}
