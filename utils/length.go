package utils

// Centimeters to Meters
func CmToM(cm float64) float64 {
	return cm / 100
}

// Meters to Centimeters
func MToCm(m float64) float64 {
	return m * 100
}

// Kilometers to Meters
func KmToM(km float64) float64 {
	return km * 1000
}

// meters to Kilometers
func MToKm(m float64) float64 {
	return m / 1000
}

// Kilometers to Centimeters
func KmToCm(km float64) float64 {
	return km * 100000
}

// Centimeters to Kilometers
func CmToKm(cm float64) float64 {
	return cm / 100000
}
