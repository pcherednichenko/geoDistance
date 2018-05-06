package haversin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type data struct {
	Lat       float64
	Lon       float64
	SearchLat float64
	SearchLon float64

	ExpectedResult float64
}

func testData() []data {
	return []data{
		{
			Lat:       24.223,
			Lon:       15.99,
			SearchLat: 0.2,
			SearchLon: -24.99,

			ExpectedResult: 5162738.571,
		},
		{
			Lat:       -23.33,
			Lon:       0.0,
			SearchLat: 94.111,
			SearchLon: -40.33,

			ExpectedResult: 12961462.765,
		},
		{
			Lat:       64.99,
			Lon:       11.4,
			SearchLat: -42.55,
			SearchLon: -11.88,

			ExpectedResult: 12141584.829,
		},
		{
			Lat:       -30.50,
			Lon:       -60.9,
			SearchLat: -3.22,
			SearchLon: -9.44,

			ExpectedResult: 6192873.084,
		},
	}
}

func TestCalculateDistance(t *testing.T) {
	for _, d := range testData() {
		equals, message := equalDistance(d.ExpectedResult, Distance(d.Lat, d.Lon, d.SearchLat, d.SearchLon))
		assert.True(t, equals, message)
	}
}

func equalDistance(expected, actual float64) (bool, string) {
	if actual < expected-0.1 {
		return false, fmt.Sprintf("actual %f not equal expected %f", actual, expected)
	}
	if actual > expected+0.1 {
		return false, fmt.Sprintf("actual %f not equal expected %f", actual, expected)
	}
	return true, ""
}
