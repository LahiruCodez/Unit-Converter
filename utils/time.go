package utils

// Seconds to Minutes
func SToMin(s float64) float64 {
	return s / 60
}

// Seconds to Hours
func MinToS(min float64) float64 {
	return min * 60
}

// Minutes to Hours
func MinToH(min float64) float64 {
	return min / 60
}

// Hours to Minutes
func HToMin(h float64) float64 {
	return h * 60
}

// Seconds to Hours
func SToH(s float64) float64 {
	return s / 3600
}

// Hours to Seconds
func HToS(h float64) float64 {
	return h * 3600
}
