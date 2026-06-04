package solutions

import (
	"fmt"
)

var monthToDays = map[int]int{
	1:  31, // Jan
	2:  28, // Feb
	3:  31, // March
	4:  30, // April
	5:  31, // May
	6:  30, // June
	7:  31, // July
	8:  31, // August
	9:  30, // September
	10: 31, // October
	11: 30, // November
	12: 31, // December
}

func Solve19() {
	numSundays := 0
	year := 1900
	month := 1
	day := 7

	// Only care until 2000
	for year <= 2000 {

		// Go to next sunday
		day += 7

		// Check the number of days in the current month
		currentMonthDays := monthToDays[month]

		// If it is february, check for leap year
		if currentMonthDays == 28 {
			if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
				// Increment 28 to 29 on a leap year
				currentMonthDays += 1
			}
		}

		// If we have moved into a new month
		if day > currentMonthDays {

			// Reset the day and increment the month
			day -= currentMonthDays
			month++

			// If we have moved into a new year
			if month > 12 {

				// Reset the month to January and increment the year
				month = 1
				year++
			}
		}

		// Only start caring in 1901
		if year >= 1901 && year <= 2000 {

			// If the day is the first, it is a Sunday since we always increment by 7
			if day == 1 {
				numSundays++
			}
		}

	}

	fmt.Printf("The number of Sundays on the first of the month from 1 Jan 1901 to 31 Dec 2000 is %d", numSundays)
}
