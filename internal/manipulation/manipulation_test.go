package manipulation

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bipo38/isaac-tui/config"
)

func TestGetHeaders(t *testing.T) {

	type example struct {
		id, name string
	}

	var test example

	got := GetHeaders(test)
	want := []string{"id", "name"}

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

	for _, tt := range cases {

		t.Run("Test Parser File Name", func(t *testing.T) {

			got := ParserFileName(tt.name, tt.extension)

			if got != tt.expected {
				t.Errorf("got %s, want %s", got, tt.expected)
			}
		})
	}
}

func TestRouteParser(t *testing.T) {

	s := config.Default["routeStart"]
	r := "transformations/images"
	fName := "isaac.png"

	gotRoute, gotFileName := RouteParser(r, fName)

	wantRoute := fmt.Sprintf("%s/%s/", s, r)
	wantFileRoute := fmt.Sprintf("%s/%s/%s", s, r, fName)

	if gotRoute != wantRoute {
		t.Errorf("got %v want %v", gotRoute, wantRoute)

	}

	if gotFileName != wantFileRoute {
		t.Errorf("got %v want %v", gotFileName, wantFileRoute)

	}
}
