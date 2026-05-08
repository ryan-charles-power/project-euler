package solutions

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Solve54() {
	// Read the entire file content into a byte slice.
	fileContent, err := os.ReadFile("../project-euler/resources/problem54poker.txt")
	if err != nil {
		log.Fatal(err) // Handle potential errors, such as file not found.
	}

	// Convert the byte slice to a string.
	str := string(fileContent)

	// Split into lines (rows)
	lines := strings.Split(strings.TrimSpace(str), "\n")

	numWins := 0

	for _, line := range lines {
		cards := strings.Split(strings.TrimSpace(line), " ")
		player1Cards := cards[:5]
		player2Cards := cards[5:]

		if compareHands(player1Cards, player2Cards) {
			numWins++
		}
	}

	fmt.Printf("The number of hands won by player 1 is %d", numWins)
}

/*
In order of increasing value, the possible hands are:
	0: High Card: Highest value card.
	1: One Pair: Two cards of the same value.
	2: Two Pairs: Two different pairs.
	3: Three of a Kind: Three cards of the same value.
	4: Straight: All cards are consecutive values.
	5: Flush: All cards of the same suit.
	6: Full House: Three of a kind and a pair.
	7: Four of a Kind: Four cards of the same value.
	8: Straight Flush: All cards are consecutive values of same suit.
	9: Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.
*/

// Returns true if player 1 wins, false if player 2 wins
func compareHands(player1Cards []string, player2Cards []string) bool {

	player1rank, player1highestCard := rankHand(player1Cards)
	player2rank, player2highestCard := rankHand(player2Cards)

	if player1rank > player2rank {
		return true
	} else if player1rank < player2rank {
		return false
	} else {
		if player1highestCard > player2highestCard {
			return true
		}
		return false
	}
}

func rankHand(cards []string) (int, int) {
	var values [5]string
	var suits [5]string

	for card := range cards {
		// Currently the format is{8C,TS,KC,9H,4S}
		value := string(cards[card][0])
		suit := string(cards[card][1])

		values[card] = value
		suits[card] = suit
	}

	isRoyalFlush := checkRoyalFlush(values, suits)
	if isRoyalFlush {
		return 9, 14
	}
	isStraightFlush := checkStraightFlush(values, suits)
	if isStraightFlush {
		return 8, getHighestCard(values)
	}

	return 0, 0
}

// Checks if the hand is a royal flush (10, J, Q, K, A of the same suit)
func checkRoyalFlush(values [5]string, suits [5]string) bool {
	// A royal flush must have all cards of the same suit
	if suits[0] != suits[1] || suits[0] != suits[2] || suits[0] != suits[3] || suits[0] != suits[4] {
		return false
	}

	// A royal flush must have the values 10, J, Q, K, A
	if numTen := countValue(values, "T"); numTen != 1 {
		return false
	}
	if numJack := countValue(values, "J"); numJack != 1 {
		return false
	}
	if numQueen := countValue(values, "Q"); numQueen != 1 {
		return false
	}
	if numKing := countValue(values, "K"); numKing != 1 {
		return false
	}
	if numAce := countValue(values, "A"); numAce != 1 {
		return false
	}
	return true

}

func checkStraightFlush(values [5]string, suits [5]string) bool {
	// A straight flush must have all cards of the same suit
	if suits[0] != suits[1] || suits[0] != suits[2] || suits[0] != suits[3] || suits[0] != suits[4] {
		return false
	}
	// A straight flush must have consecutive values
	return isConsecutive(values)
}

func countValue(values [5]string, value string) int {
	count := 0
	for _, v := range values {
		if v == value {
			count++
		}
	}
	return count
}

func isConsecutive(values [5]string) bool {
	// Convert values to integers for easier comparison
}

func getHighestCard(values [5]string) int {
	highestValue := 0
	for _, v := range values {
		value := 0
		switch v {
		case "2":
			value = 2
		case "3":
			value = 3
		case "4":
			value = 4
		case "5":
			value = 5
		case "6":
			value = 6
		case "7":
			value = 7
		case "8":
			value = 8
		case "9":
			value = 9
		case "T":
			value = 10
		case "J":
			value = 11
		case "Q":
			value = 12
		case "K":
			value = 13
		case "A":
			value = 14
		}
		if value > highestValue {
			highestValue = value
		}
	}
	return highestValue
}
