package categories

import (
	"os"
	"testing"
)

func TestCreateTrinketsCsv(t *testing.T) {

	err := CreateTrinketsCsv()

	if err != nil {
		t.Error(err)
	}

	os.RemoveAll("./isaac")

}
