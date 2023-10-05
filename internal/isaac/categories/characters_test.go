package categories

import (
	"os"
	"testing"
)

func TestCreateCharactersCsv(t *testing.T) {

	err := CreateCharactersCsv()

	if err != nil {
		t.Error(err)
	}

	os.RemoveAll("./isaac")

}
