num = 600851475143
original = num
divisor = 2

while divisor ^ 2 <= num:
    if num % divisor == 0:
        while num % divisor == 0:
            num /= divisor
    divisor+=1

largest = num
if largest == 1:
    largest = divisor - 1

print("The largest prime factor of " + str(original) + " is " + str(largest))
