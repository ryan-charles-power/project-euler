package solutions

import (
	"fmt"
	"math/big"
	"strconv"
)

func Solve20() {
	n := 100
	sum := 0

	nFactorial := big.NewInt(1)

	for i := 1; i <= n; i++ {
		bigI := big.NewInt(int64(i))
		nFactorial.Mul(nFactorial, bigI)
	}

	nFactorialString := nFactorial.String()

	for i := 0; i < len(nFactorialString); i++ {
		digit, _ := strconv.Atoi(string(nFactorialString[i]))
		sum += digit
	}

	fmt.Printf("The sum of the digits in the solution to %d! is %d", n, sum)
}
