def solve():
    i = 0
    j = 1
    primeNumber = 0
    while i < 10001:
        if isPrimeNumber(j):
            primeNumber = j
            i+=1
        j+=1
    print("The 10001st prime number is: " + str(primeNumber))

import math

def isPrimeNumber(n):
    if n <= 1:
        return False
    
    if n <= 3:
        return True
    
	# Eliminate even numbers and multiples of 3 immediately
    if n%2 == 0 or n%3 == 0:
        return False
    
	# Check divisibility starting from 5, skipping even numbers
	# Every prime > 3 is of the form 6k ± 1
    limit = int(math.sqrt(n))
    for i in range(5, limit + 1, 6):
        if n % i == 0 or n % (i + 2) == 0:
            return False
    
    return True