package parsers

import (
	"testing"
)

func TestParseExtension(t *testing.T) {

	cases := []struct {
		extension string
		expected  string
	}{
		{
			extension: "Added in Afterbirth †",
			expected:  "Afterbirth Plus",
		},
		{
			extension: "Afterbirth †",
			expected:  "Afterbirth Plus",
		},
		{
			extension: "Added in Afterbirth",
			expected:  "Afterbirth",
		},
		{
			extension: "Afterbirth",
			expected:  "Afterbirth",
		},
		{
			extension: "Added in Rebirth",
			expected:  "Rebirth",
		},
		{
			extension: "Rebirth",
			expected:  "Rebirth",
		},
		{
			extension: "",
			expected:  "Rebirth",
		},
	}

	for _, tt := range cases {

		t.Run("Test Parse Extension", func(t *testing.T) {

			got := ParseExtension(tt.extension)

			if got != tt.expected {
				t.Errorf("got %s, want %s", got, tt.expected)
			}
		})
	}

}
