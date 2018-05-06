package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
)

const n = 3000000

// Generator for big count of geo points for tests
func main() {
	file, err := os.Create("bigGeoData.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"id", "lat", "lng"})
	if err != nil {
		panic(err)
	}
	for i := 0; i <= n; i++ {
		lat := strconv.FormatFloat(rand.Float64()*100-50, 'f', 8, 64)
		lon := strconv.FormatFloat(rand.Float64()*100-50, 'f', 8, 64)
		err := writer.Write([]string{strconv.Itoa(i), lat, lon})
		if err != nil {
			panic(err)
		}
	}
}
