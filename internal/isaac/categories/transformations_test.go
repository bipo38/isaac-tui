package isaac

import (
	"os"
	"testing"
)

func TestCreateTransformationsCsv(t *testing.T) {

	err := CreateTransformationsCsv()

	if err != nil {
		t.Error(err)
	}

	os.RemoveAll("./isaac")

}
