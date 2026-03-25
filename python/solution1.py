def solve():
    sum = 0

    for i in range(1000):
        if i%3==0 or i%5==0:
            sum += i

    print(str(sum) + " is the sum of all multiples of 3 or 5")