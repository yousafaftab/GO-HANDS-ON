package main

import "fmt"

// Setting alises for Celcius and Fahrenheit type
type Celsius float32
type Fahrenheit float32

// main function call
func main() {
	// calling the conversion function in the main function
	temp := toFahrenheit(32)
	// prnting the converted value of the temperature in Fahrenheit
	fmt.Println("temperature in Fahrenheit is ", temp)
}

// Function that will convert temperature from celcius to Fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {
	var fahrenheit Fahrenheit
	fahrenheit = Fahrenheit(t*9/5) + 32
	return fahrenheit
}
