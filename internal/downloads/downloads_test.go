package downloads

import (
	"isaac-scrapper/config"
	"os"
	"testing"
)

func TestImage(t *testing.T) {

	cases := []struct {
		url, fPath, fName string
		expected          string
	}{
		{
			url:      "https://static.wikia.nocookie.net/bindingofisaacre_gamepedia/images/9/91/Collectible_Abel_icon.png/revision/latest?cb=20210821083135",
			fPath:    config.Item["imgFolder"],
			fName:    "Collectible_Abel_icon.png",
			expected: "isaac/items/images/Collectible_Abel_icon.png",
		},
	}

	for _, tt := range cases {

		t.Run("Test Download Image", func(t *testing.T) {

			got, err := Image(tt.url, tt.fPath, tt.fName)

			if err != nil {
				t.Errorf("got %v, want %v", err, nil)
			}

			if got != tt.expected {
				t.Errorf("got %s, want %s", got, tt.expected)
			}
		})
	}

	os.RemoveAll("./isaac")
}
