package underscore

import "testing"

func TestCamel(t *testing.T) {
	testCases := []struct {
		arg  string
		want string
	}{
		{"camelCaseString", "camel_case_string"},
		{"camel case", "camel case"},
		{"endsWithA", "ends_with_a"},
	}

	for _, tc := range testCases {
		got := Camel(tc.arg)
		if got != tc.want {
			t.Errorf("Camel(%q) = %q; want = %q", tc.arg, got, tc.want)
		}
	}

}
