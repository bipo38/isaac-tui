package creates

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var globalPath = "test/subtest"

func TestCreateDirs(t *testing.T) {

	got := Dirs(globalPath)

	if got != nil {
		t.Errorf("got %v, want %v", got, nil)
	}

	removeAll()
}

func TestCreateFile(t *testing.T) {

	got, err := File(fmt.Sprintf("%s/test.rs", globalPath))

	if err != nil {
		t.Errorf("got %v, want %v", got, nil)
	}

	if got == nil {
		t.Errorf("got %v, want *os.File", got)
	}

	removeAll()
}

func TestExistPath(t *testing.T) {

	cases := []struct {
		path     string
		expected bool
	}{
		{
			path:     globalPath,
			expected: true,
		},
		{
			path:     "testing/substest",
			expected: false,
		},
	}

	Dirs(cases[0].path)

	for _, tt := range cases {

		t.Run("Test Parser File Name", func(t *testing.T) {

			got, err := ExistPath(tt.path)

			if err != nil {
				t.Errorf("got %v, want %v", got, nil)
			}

			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}

	removeAll()
}

func removeAll() {

	splitted := strings.Split(globalPath, "/")

	os.RemoveAll(fmt.Sprintf("./%s", splitted[0]))

}
