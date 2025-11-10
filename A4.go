/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-10
 * @fileoverview Display and calculate the average of 56.9, 89.7, and 90.2
 */

package main

import "fmt"

func main() {

	// store 56.9 as number1
	const number1 float64 = 56.9

	// store 89.7 as number2
	const number2 float64 = 89.7

	// store 90.2 as number3
	const number3 float64 = 90.2

	// add numbers
	const addNumbers float64 = number1 + number2 + number3

	// calculation of average
	const average float64 = addNumbers / 3

	// display average
	fmt.Println("The average of", number1, ",", number2, ",", number3, "is", average)

	fmt.Println("\nDone.")

}
