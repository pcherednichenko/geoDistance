package geo

import (
	"testing"

	"fmt"

	"github.com/pcherednichenko/geoDistance/geo/data"
	"github.com/stretchr/testify/assert"
)

type testParams struct {
	lat       float64
	lon       float64
	data      data.Coordinates
	dataError error

	expectedResult []Result
	expectError    bool
}

func testData() []testParams {
	return []testParams{
		{
			lat: 10.22,
			lon: 70.22,
			data: data.Coordinates{
				12: [2]float64{12.12, 21.21},
			},
			expectedResult: []Result{
				{
					CityID:   12,
					Distance: 5.34970642035235e+06,
					Lat:      12.12,
					Lon:      21.21,
					URL:      "https://www.google.com/maps/?q=12.12000000,21.21000000",
				},
			},
			expectError: false,
		},
		{
			lat: 29.22,
			lon: 91.00,
			data: data.Coordinates{
				12: [2]float64{12.12, 21.21},
				39: [2]float64{30.99, 90.88},
			},
			expectedResult: []Result{
				{
					CityID:   39,
					Distance: 197372.9080581844,
					Lat:      30.99,
					Lon:      90.88,
					URL:      "https://www.google.com/maps/?q=30.99000000,90.88000000",
				},
				{
					CityID:   12,
					Distance: 7.412937943450556e+06,
					Lat:      12.12,
					Lon:      21.21,
					URL:      "https://www.google.com/maps/?q=12.12000000,21.21000000",
				},
			},
			expectError: false,
		},
		{
			lat: 10.22,
			lon: 70.22,
			data: data.Coordinates{
				12: [2]float64{12.12, 21.21},
			},
			dataError:      fmt.Errorf("some error"),
			expectedResult: nil,
			expectError:    true,
		},
	}
}

func TestGeoSort(t *testing.T) {
	for _, d := range testData() {
		m := mock{
			coordinates: d.data,
			error:       d.dataError,
		}
		result, err := sortByDistanceGoroutine(d.lat, d.lon, m)
		assert.Equal(t, d.expectedResult, result)
		if d.expectError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

type mock struct {
	coordinates data.Coordinates
	error       error
}

func (m mock) Data() (data.Coordinates, error) {
	return m.coordinates, m.error
}
