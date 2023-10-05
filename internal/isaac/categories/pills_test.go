package categories

import (
	"os"
	"testing"
)

func TestCreatePillsCsv(t *testing.T) {

	err := CreatePillsCsv()

	if err != nil {
		t.Error(err)
	}

	os.RemoveAll("./isaac")

}
