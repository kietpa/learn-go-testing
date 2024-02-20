package underscore

import "testing"

// func TestCamel(t *testing.T) {
// 	testCases := []struct {
// 		arg  string
// 		want string
// 	}{
// 		{"camelCaseString", "camel_case_string"},
// 		{"camel case", "camel case"},
// 		{"endsWithA", "ends_with_a"},
// 	}

// 	for _, tc := range testCases {
// 		got := Camel(tc.arg)
// 		if got != tc.want {
// 			t.Errorf("Camel(%q) = %q; want = %q", tc.arg, got, tc.want)
// 		}
// 	}

// }

func TestCamel(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct { // can also use map
		name string
		args args
		want string
	}{
		{"simple", args{"camelCaseString"}, "camel_case_string"},
		{"spaces", args{"camel case"}, "camel case"},
		{"ends with capital", args{"endsWithA"}, "ends_with_a"},
	}
	for _, tt := range tests {
		// subtest
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel(tt.args.str); got != tt.want {
				// when using t.fatalf doesn't stop entire test, only subtest
				t.Errorf("Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
