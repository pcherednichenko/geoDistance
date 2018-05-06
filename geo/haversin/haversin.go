package haversin

import "math"

const earthRadius = 6378100

// Distance calculate distance between points, using haversine formula
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	var la1, lo1, la2, lo2 float64
	la1 = degreesToRadians(lat1)
	lo1 = degreesToRadians(lon1)
	la2 = degreesToRadians(lat2)
	lo2 = degreesToRadians(lon2)

	h := sinpow(la2-la1) + math.Cos(la1)*math.Cos(la2)*sinpow(lo2-lo1)
	return 2 * earthRadius * math.Asin(math.Sqrt(h))
}

func sinpow(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}
