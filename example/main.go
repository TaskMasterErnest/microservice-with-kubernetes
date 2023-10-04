package main

import (
	"flag"
	"fmt"
)

// writing a temperature converter code.

// func to change from Fahrenheit to Celcius
func fahrenheit2celcius(t float64) float64 {
	return (t - 32) * 5.0 / 9.0
}

// func to change from Celcius to Fahrenheit
func celcius2fahrenheit(t float64) float64 {
	return (9.0 / 5.0 * t) + 32
}

// create the CLI functions to take care of these functions
func main() {
	// specify the flag usage for the temperature converter
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Temperature converter tool\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage Information:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "This program converts temperatures between Celcius and Fahrenheit\n")
		fmt.Fprintf(flag.CommandLine.Output(), "'mode' is either 'c2f' or 'f2c'\n")
		fmt.Fprintf(flag.CommandLine.Output(), "'temperature' is a floating point number to be converted according to mode\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Eg. convert -temperature=<value> -mode=<value>\n\n")
		// os.Exit(1)
		flag.PrintDefaults()
	}

	var mode string
	var temperature float64

	// add flags
	flag.StringVar(&mode, "mode", "c2f", "Conversion mode: 'c2f' for Celcius to Fahrenheit, 'f2c' for Fahrenheit to Celcius")
	flag.Float64Var(&temperature, "temperature", 0.0, "Temperature to be converted")
	flag.Parse()

	// switch case to work with temperature recorded
	switch mode {
	case "c2f":
		result := celcius2fahrenheit(temperature)
		fmt.Printf("%.2f Celcius is %.2f Fahrenheit\n", temperature, result)

	case "f2c":
		result := fahrenheit2celcius(temperature)
		fmt.Printf("%.2f Fahrenheit is %.2f Celcius\n", temperature, result)

	default:
		fmt.Println("Invalid mode. Please use 'c2f' or 'f2c'.")
	}

}
