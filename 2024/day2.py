def safelevel(nums):
    diffs = [nums[i + 1] - nums[i] for i in range(len(nums) - 1)]
    if not all(diff < 0 for diff in diffs) and not all(diff > 0 for diff in diffs):
        return False

    diffs = [abs(diff) for diff in diffs]
    if not all (1 <= diff <= 3 for diff in diffs):
        return False
    return True


with open('day2input.txt') as fp:
    
    safenum = 0
    for line in fp:
        nums = [int(num) for num in line.split(' ')]

        if safelevel(nums):
            safenum += 1
            continue
        
        for i in range(len(nums)):
            removed = nums[:i] + nums[i + 1:]
            if safelevel(removed):
                safenum += 1
                break


    print(safenum)