import numpy as np

matrix = []

with open('day4input.txt') as fp:
    matrix = [[c for c in line.strip()] for line in fp.readlines()]

transposed = []

for i in range(len(matrix[0])):
    transposed.append([row[i] for row in matrix])

# print(matrix)
# print(transposed)
# print(np.diagonal(matrix, -8))

mirrored = []
for row in matrix:
    mirrored.append(row[::-1])

# print(np.diagonal(mirrored))

# diagonal - num of rows each side, mirrored num of rows each side

def findXmas(arr):
    if len(arr) < 4:
        return 0
    print('findxmas', arr)
    count = 0
    
    xmas = 0
    samx = 0
    
    for char in arr:
        # print('char:', char)
        if char == 'XMAS'[xmas]:
            # print('increment xmas:', xmas)
            
            if xmas == 3:
                count += 1
                xmas = 0
            xmas += 1
        elif char == 'X':
            # print('start xmas', 0)
            xmas = 1
        else:
            # print('reset xmas')
            xmas = 0
        if char == 'SAMX'[samx]:
            # print('increment samx:', xmas)
            
            if samx == 3:
                count += 1
                samx = 0
            samx += 1
        elif char == 'S':
            # print('start samx', 0)
            xmas = 1
        else:
            # print('reset samx')
            samx = 0
    return count

totalcount = 0
for row in matrix:
    totalcount += findXmas(row)

for row in transposed:
    totalcount += findXmas(row)

height = len(matrix)

totalcount += findXmas(np.diagonal(matrix))
for i in range(1, height):
    totalcount += findXmas(np.diagonal(matrix, -i))
    totalcount += findXmas(np.diagonal(matrix, i))

totalcount += findXmas(np.diagonal(mirrored))
for i in range(1, height):
    totalcount += findXmas(np.diagonal(mirrored, -i))
    totalcount += findXmas(np.diagonal(mirrored, i))

print(totalcount)