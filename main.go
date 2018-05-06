package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/pcherednichenko/geoDistance/data/csv"
	"github.com/pcherednichenko/geoDistance/geo"
)

func main() {
	var (
		useGoroutines = flag.Bool("useGoroutines", false, "count of goroutines")
		filename      = flag.String("filename", "", "count of goroutines")
		lat           = 51.925146
		lon           = 4.478617
	)
	flag.Parse()
	if len(*filename) == 0 {
		panic("empty flag filename")
	}
	data := csv.New(*filename)
	start := time.Now()
	result, err := geo.SortByDistance(lat, lon, data, *useGoroutines)
	if err != nil {
		panic(err)
	}

	// Print result
	fmt.Println("Top 5 closest:")
	for _, r := range result[:5] {
		fmt.Printf("City ID: %d, distance: %f m, coordinates: (%f, %f), url: %s \n",
			r.CityID, r.Distance, r.Lat, r.Lon, r.URL,
		)
	}
	fmt.Println("Top 5 furthest:")
	for _, r := range result[len(result)-5:] {
		fmt.Printf("City ID: %d, distance: %f m, lat: %f, lon: %f, url: %s \n",
			r.CityID, r.Distance, r.Lat, r.Lon, r.URL,
		)
	}
	fmt.Printf("Running time: %s \n", time.Since(start).String())
}
