package geo

import (
	"runtime"
	"sort"

	"github.com/pcherednichenko/geoDistance/geo/data"
	"github.com/pcherednichenko/geoDistance/geo/haversin"
	"github.com/pcherednichenko/geoDistance/geo/url"
)

// Result of calculation distance between two points
type Result struct {
	CityID   int
	Distance float64
	Lat      float64
	Lon      float64
	URL      string
}

// SortByDistance sort geo points by distance from original point
func SortByDistance(lat, lon float64, d data.Data, useGoroutines bool) (result []Result, err error) {
	if useGoroutines {
		result, err = sortByDistanceGoroutine(lat, lon, d)
	} else {
		result, err = sortByDistanceDefault(lat, lon, d)
	}
	return
}

// sortByDistanceDefault default search and sort without goroutines
func sortByDistanceDefault(lat, lon float64, d data.Data) ([]Result, error) {
	coordinates, err := d.Data()
	if err != nil {
		return nil, err
	}
	result := make([]Result, len(coordinates))
	// use t++ to avoid memory relocation during append() method
	t := 0
	for id, c := range coordinates {
		result[t] = Result{
			CityID:   id,
			Distance: haversin.Distance(lat, lon, c[0], c[1]),
			Lat:      c[0],
			Lon:      c[1],
			URL:      url.FromCoordinates(c[0], c[1]),
		}
		t++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Distance < result[j].Distance
	})
	return result, nil
}

// sortByDistanceGoroutine use goroutines for sorting by distance
func sortByDistanceGoroutine(lat, lon float64, d data.Data) ([]Result, error) {
	coordinates, err := d.Data()
	if err != nil {
		return nil, err
	}
	results := make([]Result, len(coordinates))
	coord := make(chan dataChan)
	done := make(chan struct{})
	result := make(chan Result)
	for i := 0; i <= runtime.NumCPU()-1; i++ {
		go func() {
			for {
				select {
				case c := <-coord:
					result <- Result{
						CityID:   c.cityID,
						Distance: haversin.Distance(lat, lon, c.lat, c.lon),
						Lat:      c.lat,
						Lon:      c.lon,
						URL:      url.FromCoordinates(c.lat, c.lon),
					}
				case <-done:
					return
				}
			}
		}()
	}
	go func() {
		for id, c := range coordinates {
			coord <- dataChan{
				searchLat: lat,
				searchLon: lon,
				cityID:    id,
				lat:       c[0],
				lon:       c[1],
			}
		}
	}()

	t := 0
	for range coordinates {
		results[t] = <-result
		t++
	}
	close(done)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})
	return results, nil
}

type dataChan struct {
	searchLat float64
	searchLon float64
	cityID    int
	lat       float64
	lon       float64
}
