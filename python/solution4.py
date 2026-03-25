def solve():
	product = 0
	largestPalindrome = 0
	iMax = 0
	jMax = 0

	# run through all products 1-999, and 1-999
	for i in range(999):
		for j in range(999):
			product = i * j
			# Check size first to save time
			if product > largestPalindrome:
				# Then do string conversion and comparison
				productStr = str(product)
				productStrReverse = productStr[::-1]
				if productStr == productStrReverse:
					largestPalindrome = product
					iMax = i
					jMax = j

	print("The largest palindrome is " + str(largestPalindrome) + ", which is a product of " + str(iMax) + " and " + str(jMax))