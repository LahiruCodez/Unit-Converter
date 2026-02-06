package main

import (
	"Unit-Converter/utils"
	"fmt"
)

func main() {
	var value float64
	var choice int
	var category int

	fmt.Println("==========================================")
	fmt.Println("----- Welcome to the Unit Converter! -----")
	fmt.Println("==========================================")
	fmt.Println()
	fmt.Println("** Select the category **")
	for {
		fmt.Println()
		fmt.Println("=================================")
		fmt.Println("| 1.Length Conversions: 	|")
		fmt.Println("| 2.Time Conversions: 		|")
		fmt.Println("| 3.Volume Conversions: 	|")
		fmt.Println("| 4.Temperature Conversions: 	|")
		fmt.Println("=================================")
		fmt.Scanln(&category)

		switch category {
		case 1:
			fmt.Println()
			fmt.Println("1. Centimeters to Meters")
			fmt.Println("2. Meters to Centimeters")
			fmt.Println("3. Kilometers to Meters")
			fmt.Println("4. Meters to Kilometers")
			fmt.Println("5. Kilometers to Centimeters")
			fmt.Println("6. Centimeters to Kilometers")

			fmt.Println()
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice)

			fmt.Println()
			fmt.Print("Enter the value to convert: ")
			fmt.Scanln(&value)

		case 2:
			fmt.Println()
			fmt.Println("1. Seconds to Minutes")
			fmt.Println("2. Minutes to Seconds")
			fmt.Println("3. Minutes to Hours")
			fmt.Println("4. Hours to Minutes")
			fmt.Println("5. Seconds to Hours")
			fmt.Println("6. Hours to Seconds")

			fmt.Println()
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice)

			fmt.Println()
			fmt.Print("Enter the value to convert: ")
			fmt.Scanln(&value)
		case 3:
			fmt.Println()
			fmt.Println("1. Liters to Milliliters")
			fmt.Println("2. Milliliters to Liters")
			fmt.Println("3. Liters to Gallons")
			fmt.Println("4. Gallons to Liters")
			fmt.Println("5. Milliliters to Gallons")
			fmt.Println("6. Gallons to Milliliters")

			fmt.Println()
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice)

			fmt.Println()
			fmt.Print("Enter the value to convert: ")
			fmt.Scanln(&value)
		case 4:
			fmt.Println()
			fmt.Println("1. Celsius to Fahrenheit")
			fmt.Println("2. Fahrenheit to Celsius")
			fmt.Println("3. Celsius to Kelvin")
			fmt.Println("4. Kelvin to Celsius")
			fmt.Println("5. Fahrenheit to Kelvin")
			fmt.Println("6. Kelvin to Fahrenheit")

			fmt.Println()
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice)

			fmt.Println()
			fmt.Print("Enter the value to convert: ")
			fmt.Scanln(&value)
		default:
			fmt.Println("Invalid category")
			fmt.Println()
			continue
		}

		var result float64

		switch category {
		case 1:
			switch choice {
			case 1:
				result = utils.CmToM(value)
				fmt.Println()
				fmt.Printf("%.2f cm is equal to %.2f m\n", value, result)
				fmt.Println("------------------------------")
			case 2:
				result = utils.MToCm(value)
				fmt.Println()
				fmt.Printf("%.2f m is equal to %.2f cm\n", value, result)
				fmt.Println("------------------------------")
			case 3:
				result = utils.KmToM(value)
				fmt.Println()
				fmt.Printf("%.2f km is equal to %.2f m\n", value, result)
				fmt.Println("------------------------------")
			case 4:
				result = utils.MToKm(value)
				fmt.Println()
				fmt.Printf("%.2f m is equal to %.2f km\n", value, result)
				fmt.Println("------------------------------")
			case 5:
				result = utils.KmToCm(value)
				fmt.Println()
				fmt.Printf("%.2f km is equal to %.2f cm\n", value, result)
				fmt.Println("------------------------------")
			case 6:
				result = utils.CmToKm(value)
				fmt.Println()
				fmt.Printf("%.2f cm is equal to %.2f km\n", value, result)
				fmt.Println("------------------------------")
			default:
				fmt.Println("Invalid choice")
				fmt.Println()
				continue
			}
		case 2:
			switch choice {
			case 1:
				result = utils.SToMin(value)
				fmt.Println()
				fmt.Printf("%.2f seconds is equal to %.2f minutes\n", value, result)
				fmt.Println("------------------------------")
			case 2:
				result = utils.MinToS(value)
				fmt.Println()
				fmt.Printf("%.2f minutes is equal to %.2f seconds\n", value, result)
				fmt.Println("------------------------------")
			case 3:
				result = utils.MinToH(value)
				fmt.Println()
				fmt.Printf("%.2f minutes is equal to %.2f hours\n", value, result)
				fmt.Println("------------------------------")
			case 4:
				result = utils.HToMin(value)
				fmt.Println()
				fmt.Printf("%.2f hours is equal to %.2f minutes\n", value, result)
				fmt.Println("------------------------------")
			case 5:
				result = utils.SToH(value)
				fmt.Println()
				fmt.Printf("%.2f seconds is equal to %.2f hours\n", value, result)
				fmt.Println("------------------------------")
			case 6:
				result = utils.HToS(value)
				fmt.Println()
				fmt.Printf("%.2f hours is equal to %.2f seconds\n", value, result)
				fmt.Println("------------------------------")
			default:
				fmt.Println("Invalid choice")
				fmt.Println()
				continue
			}

		case 3:
			switch choice {
			case 1:
				result = utils.LToMl(value)
				fmt.Println()
				fmt.Printf("%.2f liters is equal to %.2f milliliters\n", value, result)
				fmt.Println("------------------------------")
			case 2:
				result = utils.MlToL(value)
				fmt.Println()
				fmt.Printf("%.2f milliliters is equal to %.2f liters\n", value, result)
				fmt.Println("------------------------------")
			case 3:
				result = utils.LToGal(value)
				fmt.Println()
				fmt.Printf("%.2f liters is equal to %.2f gallons\n", value, result)
				fmt.Println("------------------------------")
			case 4:
				result = utils.GalToL(value)
				fmt.Println()
				fmt.Printf("%.2f gallons is equal to %.2f liters\n", value, result)
				fmt.Println("------------------------------")
			case 5:
				result = utils.MlToGal(value)
				fmt.Println()
				fmt.Printf("%.2f milliliters is equal to %.2f gallons\n", value, result)
				fmt.Println("------------------------------")
			case 6:
				result = utils.GalToMl(value)
				fmt.Println()
				fmt.Printf("%.2f gallons is equal to %.2f milliliters\n", value, result)
				fmt.Println("------------------------------")
			default:
				fmt.Println("Invalid choice")
				fmt.Println()
				continue
			}
		case 4:
			switch choice {
			case 1:
				result = utils.CToF(value)
				fmt.Println()
				fmt.Printf("%.2f degrees Celsius is equal to %.2f degrees Fahrenheit\n", value, result)
				fmt.Println("------------------------------")
			case 2:
				result = utils.FToC(value)
				fmt.Println()
				fmt.Printf("%.2f degrees Fahrenheit is equal to %.2f degrees Celsius\n", value, result)
				fmt.Println("------------------------------")
			case 3:
				result = utils.CToK(value)
				fmt.Println()
				fmt.Printf("%.2f degrees Celsius is equal to %.2f degrees Kelvin\n", value, result)
				fmt.Println("------------------------------")
			case 4:
				result = utils.KToC(value)
				fmt.Println()
				fmt.Printf("%.2f degrees Kelvin is equal to %.2f degrees Celsius\n", value, result)
				fmt.Println("------------------------------")
			case 5:
				result = utils.FToK(value)
				fmt.Println()
				fmt.Printf("%.2f degrees Fahrenheit is equal to %.2f degrees Kelvin\n", value, result)
				fmt.Println("------------------------------")
			case 6:
				result = utils.KToF(value)
				fmt.Println()
				fmt.Printf("%.2f degrees Kelvin is equal to %.2f degrees Fahrenheit\n", value, result)
				fmt.Println("------------------------------")
			default:
				fmt.Println("Invalid choice")
				fmt.Println()
				continue
			}
		default:
			fmt.Println("Invalid category")
			fmt.Println()
			continue
		}
	}
}
