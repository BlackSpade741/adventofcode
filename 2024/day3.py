import re

input = ''
with open('day3input.txt') as fp:
    input = fp.read()
print(input)

mulregex = r'mul\([0-9]+,[0-9]+\)'
doregex = r'do\(\)'
dontregex = r'don\'t\(\)'
regex = rf'({mulregex}|{doregex}|{dontregex})'

res = 0
on = True
for match in re.findall(regex, input):
    print(match)
    if match.startswith('don'):
        on = False
    elif match.startswith('do'):
        on = True
    elif on:
        mul = match[4:-1]
        x, y = mul.split(',')
        res += int(x) * int(y)
print(res)
