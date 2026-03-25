def solve():
    i = 1
    j = 1
    sum = 0
    temp = 0

    while j < 4000000:
        if j%2 == 0:
            sum += j
        temp = j
        j = i + j
        i = temp

    print("The sum of all even fibonacci numbers less than 4,000,000 is " + str(sum))