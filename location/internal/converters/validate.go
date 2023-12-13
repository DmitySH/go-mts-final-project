package converters

func isCorrectLat(lat float64) bool {
	return lat >= 0 && lat <= 90
}

func isCorrectLng(lng float64) bool {
	return lng >= 0 && lng <= 180
}
