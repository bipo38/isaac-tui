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
