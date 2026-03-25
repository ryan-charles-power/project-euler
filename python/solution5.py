def solve():
	finished = False
	i = 20

	while finished == False:
		isTarget = True
		for j in range(1,20):
			if i%j != 0:
				isTarget = False
				break

		if isTarget:
			result = i
			finished = True
		else:
			i += 20 	# result has to be a multiple of 20

	print(result)