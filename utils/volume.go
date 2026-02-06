package utils

// Liters to Milliliters
func LToMl(l float64) float64 {
	return l * 1000
}

// Milliliters to Liters
func MlToL(ml float64) float64 {
	return ml / 1000
}

// Liters to Gallons
func LToGal(l float64) float64 {
	return l * 0.264172
}

// Gallons to Liters
func GalToL(gal float64) float64 {
	return gal / 0.264172
}

// Milliliters to Gallons
func MlToGal(ml float64) float64 {
	return ml * 0.000264172
}

// Gallons to Milliliters
func GalToMl(gal float64) float64 {
	return gal / 0.000264172
}
