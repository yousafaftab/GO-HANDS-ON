package main

import "fmt"

/*
Write a function Season that has as input-parameter of a month number
and which returns the name of the season to which this month belongs (disregard the day in the month).
Make sure you follow the following criteria of month names and their values:

January: 1
February: 2
March: 3
April: 4
May: 5
June: 6
July: 7
August: 8
September: 9
October: 10
November: 11
December: 12

Another thing to note is the seasons of the months. Look at the following mapping:
Winter: 1, 2, and 12
Spring: 3, 4, and 5
Summer: 6, 7, and 8
Autumn: 9, 10, and 11
*/

func main() {
	var month string
	month = season(5)
	fmt.Printf("this is %s Season \n", month)
	month = season(6)
	fmt.Printf("this is %s Season \n", month)
	month = season(10)
	fmt.Printf("this is %s Season \n", month)
	month = season(1)
	fmt.Printf("this is %s Season \n", month)
	month = season(100)
	fmt.Printf("this is %s Season \n", month)

}

// Function that will take Season number as Input and return Season name
func season(month int) string {
	switch month {
	case 1, 2, 12:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	default:
		return "Unknown"
	}
}
