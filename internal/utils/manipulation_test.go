package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetHeaders(t *testing.T) {

	type example struct {
		id, name string
	}

	var test example

	got := GetHeaders(test)
	want := []string{"id", "name"}

	fmt.Print(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestParserFileName(t *testing.T) {

	cases := []struct {
		name      string
		extension string
		expected  string
	}{
		{
			name:      "test",
			extension: "csv",
			expected:  "test.csv",
		},
		{
			name:      "test",
			extension: ".csv",
			expected:  "test.csv",
		},
	}

	for i, tt := range cases {

		t.Run("Test"+string(rune(i)), func(t *testing.T) {

			got := ParserFileName(tt.name, tt.extension)

			if got != tt.expected {
				t.Errorf("got %s, want %s", got, tt.expected)
			}
		})
	}
}
