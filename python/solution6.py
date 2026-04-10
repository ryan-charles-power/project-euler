def solve():
    sum_x = 0
    sum_y = 0

    for i in range(101):
        sum_x += i
        sum_y += i*i
    
    sum_x *= sum_x
    solution = sum_x - sum_y

    print("sum_x: " + str(sum_x) + ", sum_y: " + str(sum_y) + "\nsolution: " + str(solution))
