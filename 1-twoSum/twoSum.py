from typing import Dict, List


def twoSum(nums: List[int], target: int) -> List[int]:
    mapper: Dict[int, int] = {}

    for i, num in enumerate(nums):
        remain = target - num

        if(remain not in mapper):
            mapper[num] = i
            continue
        return [mapper[remain], i]


nums = [int(item) for item in input("Enter the list items : ").split()]
target = int(input())

ans = twoSum(nums, target)
print(ans)
