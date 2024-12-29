left = []
right = []

with open('day1input.txt') as fp:
    for line in fp:
        l, r = line.split('   ')
        print('left:', l, 'right:', r)
        left.append(int(l))
        right.append(int(r))

left.sort()
right.sort()
total = 0
for i in range(len(left)):
    total += abs(left[i] - right[i])

print(total)

# counts
counts = {}
for r in right:
    if r not in counts:
        counts[r] = 0
    counts[r] += 1

similarity = 0
for l in left:
    if l in counts:
        similarity += l * counts[l]

print(similarity)