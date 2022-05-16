package function

import "math"

// EarthDistance 两点之间的距离,不分前后[默认单位:米]
func EarthDistance(lat1, lng1, lat2, lng2 float64, units ...string) float64 {
	var transform float64 = 1
	if len(units) > 0 {
		if units[0] == "km" || units[0] == "KM" {
			transform = 1000
		} else if units[0] == "m" || units[0] == "M" {
			transform = 1
		}
	}
	var radius float64 = 6371000 // 6378137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius / transform
}
