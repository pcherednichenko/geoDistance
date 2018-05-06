package csv

import (
	"encoding/csv"
	"os"

	"fmt"
	"strconv"

	"github.com/pcherednichenko/geoDistance/geo/data"
)

// CSV structure for load csv file, store filename and implement interface Data
type CSV struct {
	filepath string
}

// New create CSV structure with filename
func New(filepath string) CSV {
	return CSV{
		filepath: filepath,
	}
}

// Data read by lines all data from CSV file by id, lat, lon format (see geoData.csv)
func (c CSV) Data() (data.Coordinates, error) {
	csvFile, err := os.Open(c.filepath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(csvFile)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(lines)-1 <= 0 {
		return nil, fmt.Errorf("empty csv file")
	}
	return readFromLines(lines)
}

// readFromLines read by line by format geoData.csv
func readFromLines(lines [][]string) (data.Coordinates, error) {
	coordinates := make(data.Coordinates, len(lines)-1)
	for i, line := range lines {
		if i == 0 {
			continue
		}
		if len(line) != 3 {
			return nil, fmt.Errorf("wrong line: %v", line)
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, fmt.Errorf("wrong id, should be integer, we get: %s, error: %v", line[0], err)
		}
		lat, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return nil, fmt.Errorf("wrong lat, should be integer, we get: %s, error: %v", line[1], err)
		}
		lon, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			return nil, fmt.Errorf("wrong lon, should be integer, we get: %s, error: %v", line[2], err)
		}
		coordinates[id] = [2]float64{lat, lon}
	}
	return coordinates, nil
}
