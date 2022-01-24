from typing import List


def fib(n: int) -> int:
    """
    509. 斐波那契数
    :param n:
    :return:
    """
    dp = [0, 1]
    for i in range(2, n + 1):
        dp.append(dp[i - 1] + dp[i - 2])
    return dp[-1]


def test_fib():
    print(fib(10))


def tribonacci(n: int) -> int:
    """
    1137. 第 N 个泰波那契数
    :param n:
    :return:
    """
    dp = [0, 1, 1]
    for i in range(3, n + 1):
        dp.append(dp[i - 3] + dp[i - 2] + dp[i - 1])
    return dp[-1]


def test_tribonacci():
    print(tribonacci(4))


def climbStairs(n: int) -> int:
    """
    70. 爬楼梯
    :param n:
    :return:
    """
    dp = [1, 2]
    for i in range(2, n):
        dp.append(dp[i - 1] + dp[i - 2])
    return dp[-1]


def test_climbStairs():
    print(climbStairs(3))


def minCostClimbingStairs(cost: List[int]) -> int:
    """
    746. 使用最小花费爬楼梯
    :param cost:
    :return:
    """
    dp = [0, 0]
    for i in range(2, len(cost) + 1):
        dp.append(min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]))
    return dp[-1]


def test_minCostClimbingStairs():
    print(minCostClimbingStairs([1, 100, 1, 1, 1, 100, 1, 1, 100, 1]))


def rob(nums: List[int]) -> int:
    """
    198. 打家劫舍
    :param nums:
    :return:
    """
    if len(nums) <= 1:
        return nums[0]
    dp = [nums[0], nums[1]]
    for i in range(2, len(nums)):
        dp.append(max(dp[:i - 1]) + nums[i])
    return max(dp)


def test_rob():
    print(rob([2]))


def rob2(nums: List[int]) -> int:
    """
    213. 打家劫舍 II
    :param nums:
    :return:
    """


def deleteAndEarn(nums: List[int]) -> int:
    """
    740. 删除并获得点数
    :param nums:
    :return:
    """
    convert_nums = [0] * (max(nums) + 1)
    for num in nums:
        convert_nums[num] += 1

    dp = [0, convert_nums[1]]
    for i in range(2, len(convert_nums)):
        dp.append(max(dp[:i - 1]) + i * convert_nums[i])
    return max(dp)


def test_deleteAndEarn():
    print(deleteAndEarn([2, 2, 3, 3, 3, 4]))


def canJump(nums: List) -> bool:
    """
    55. 跳跃游戏
    :param nums:
    :return:
    """
    # TODO: 需要复习


