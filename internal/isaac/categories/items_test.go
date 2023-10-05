package categories

import (
	"os"
	"testing"
)

func TestCreateItemsCsv(t *testing.T) {

	err := CreateItemsCsv()

	if err != nil {
		t.Error(err)
	}

	os.RemoveAll("./isaac")

}
