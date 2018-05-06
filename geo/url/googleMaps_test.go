package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUrl(t *testing.T) {
	url := FromCoordinates(23.42, -32.555)
	assert.Equal(t, "https://www.google.com/maps/?q=23.42000000,-32.55500000", url)
}
