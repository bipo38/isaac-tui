package downloads

import (
	"os"
	"testing"

	"github.com/bipo38/isaac-tui/config"
)

func TestImage(t *testing.T) {

	cases := []struct {
		url, fp, fn string
		expected    string
	}{
		{
			url:      "https://static.wikia.nocookie.net/bindingofisaacre_gamepedia/images/9/91/Collectible_Abel_icon.png/revision/latest?cb=20210821083135",
			fp:       config.Item["imgFolder"],
			fn:       "Collectible_Abel_icon.png",
			expected: "isaac/items/images/Collectible_Abel_icon.png",
		},
	}

	for _, tt := range cases {

		t.Run("Test Download Image", func(t *testing.T) {

			got, err := Image(tt.url, tt.fp, tt.fn)

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
