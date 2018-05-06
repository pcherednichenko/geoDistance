package url

import "strconv"

// FromCoordinates generate url to google maps by coordinates
func FromCoordinates(lat, lon float64) string {
	return "https://www.google.com/maps/?q=" +
		strconv.FormatFloat(lat, 'f', int(8), 64) + "," +
		strconv.FormatFloat(lon, 'f', int(8), 64)
}
