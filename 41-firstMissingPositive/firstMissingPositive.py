from typing import List

# first


def firstMissingPositive(nums: List[int]) -> int:
    smallestPositive = 1

    nums.sort()

    for num in nums:
        if(num <= 0 or smallestPositive != num):
            continue

        smallestPositive = num+1

    return smallestPositive


nums = [int(num) for num in input().split()]

print(firstMissingPositive(nums))
