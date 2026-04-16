package solutions

import "fmt"

var digitOneMap = map[int]int{
	0: 0,
	1: 3, // one
	2: 3, // two
	3: 5, // three
	4: 4, // four
	5: 4, // five
	6: 3, // six
	7: 5, // seven
	8: 5, // eight
	9: 4, // nine
}

var digitTwoMap = map[int]int{
	0: 0,
	// TENS DEALT WITH MANUALLY
	2: 6, // twenty
	3: 6, // thirty
	4: 5, // forty
	5: 5, // fifty
	6: 5, // sixty
	7: 7, // seventy
	8: 6, // eighty
	9: 6, // ninety
}

var tensMap = map[int]int{
	0: 3, // ten
	1: 6, // eleven
	2: 6, // twelve
	3: 8, // thirteen
	4: 8, // fourteen
	5: 7, // fifteen
	6: 7, // sixteen
	7: 9, // seventeen
	8: 8, // eighteen
	9: 8, // nineteen
}

func Solve17() {
	var numLetters int

	numAsArray := [3]int{0, 0, 1}

	// Go through every number
	for i := 1; i <= 999; i++ {
		// go through every digit of the number
		for j := 0; j <= 2; j++ {
			// HANDLE DIGIT ONE
			switch j {
			// Hundreds
			case 0:
				// if i > 99
				if numAsArray[j] > 0 {
					// one... two... etc
					numLetters += digitOneMap[numAsArray[j]]

					// hundred
					numLetters += 7

					if (numAsArray[j+1] != 0) || (numAsArray[j+2] != 0) {
						numLetters += 3
					}
				}
			// Tens
			case 1:
				if numAsArray[j] == 1 {
					// ten... eleven... etc
					numLetters += tensMap[numAsArray[j+1]]
				} else {
					// twenty... thirty... etc
					numLetters += digitTwoMap[numAsArray[j]]
				}

			// Ones
			case 2:
				// As long as not already handled (ten, eleven, twelve, etc)
				if numAsArray[j-1] != 1 {
					// One... Two... etc
					numLetters += digitOneMap[numAsArray[j]]
				}
			}
		}

		// Update current number
		if numAsArray[0] == 9 {
			// We are are at 900 to 999
			if numAsArray[1] == 9 {
				// We are at 990 to 999
				if numAsArray[2] == 9 {
					// We are at 999
					// DO NOTHING AS WE HANDLE 1000 ALONE
				} else {
					// We are at 99x
					numAsArray[2]++
				}
			} else {
				// We are at 900 to 989
				if numAsArray[2] == 9 {
					// We are at 9x9
					numAsArray[1]++
					numAsArray[2] = 0
				} else {
					// We are at 9x0 to 9x8
					numAsArray[2]++
				}
			}
		} else {
			// We are at 000 to 899
			if numAsArray[1] == 9 {
				// We are at x90 to x99
				if numAsArray[2] == 9 {
					// We are at x99
					numAsArray[0]++
					numAsArray[1] = 0
					numAsArray[2] = 0
				} else {
					// We are at x9x
					numAsArray[2]++
				}
			} else {
				// We are at x00 to x89
				if numAsArray[2] == 9 {
					// We are at xx9
					numAsArray[1]++
					numAsArray[2] = 0
				} else {
					// We are at 9x0 to 9x8
					numAsArray[2]++
				}
			}
		}
	}

	// one thousand ... 11
	numLetters += 11

	fmt.Printf("The sum of all letters used to count from 1 to 1000 is %d", numLetters)
}
