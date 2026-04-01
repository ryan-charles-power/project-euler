def solve():
    try:
        with open("../project-euler/resources/problem11grid.txt", "r") as f:
            # Parse the grid into a 2D list of integers
            matrix = [
                [int(num) for num in line.split()] 
                for line in f if line.strip()
            ]
    except FileNotFoundError:
        print("File not found.")
        return

    max_product = 0
    max_factors = []

    rows = len(matrix)
    cols = len(matrix[0])

    def check(product, factors):
        nonlocal max_product, max_factors
        if product > max_product:
            max_product = product
            max_factors = factors

    for i in range(rows):
        for j in range(cols):
            # North
            if j >= 3:
                factors = [matrix[i][j], matrix[i][j-1], matrix[i][j-2], matrix[i][j-3]]
                check(factors[0] * factors[1] * factors[2] * factors[3], factors)
            
            # East
            if i <= rows - 4:
                factors = [matrix[i][j], matrix[i+1][j], matrix[i+2][j], matrix[i+3][j]]
                check(factors[0] * factors[1] * factors[2] * factors[3], factors)
                
            # North-East
            if i <= rows - 4 and j >= 3:
                factors = [matrix[i][j], matrix[i+1][j-1], matrix[i+2][j-2], matrix[i+3][j-3]]
                check(factors[0] * factors[1] * factors[2] * factors[3], factors)
                
            # North-West
            if i >= 3 and j >= 3:
                factors = [matrix[i][j], matrix[i-1][j-1], matrix[i-2][j-2], matrix[i-3][j-3]]
                check(factors[0] * factors[1] * factors[2] * factors[3], factors)

    print(f"The max product is {max_product} made of the products {max_factors}")