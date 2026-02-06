package utils

// Celsius to Fahrenheit
func CToF(c float64) float64 {
	return (c * 9 / 5) + 32
}

// Fahrenheit to Celsius
func FToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

// Celsius to Kelvin
func CToK(c float64) float64 {
	return c + 273.15
}

// Kelvin to Celsius
func KToC(k float64) float64 {
	return k - 273.15
}

// Fahrenheit to Kelvin
func FToK(f float64) float64 {
	return (f-32)*5/9 + 273.15
}

// Kelvin to Fahrenheit
func KToF(k float64) float64 {
	return (k-273.15)*9/5 + 32
}
