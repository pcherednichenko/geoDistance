package data

// Coordinates store all points data
type Coordinates map[int][2]float64 // [city_id][lat][lon]

// Data main interface to get data of points
type Data interface {
	Data() (Coordinates, error)
}
