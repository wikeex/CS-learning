
def fib(n: int) -> int:
    """
    509. 斐波那契数
    :param n:
    :return:
    """
    dp = [0, 1]
    for i in range(2, n+1):
        dp.append(dp[i-1] + dp[i-2])
    return dp[-1]


def test_fib():
    print(fib(100))


def tribonacci(n: int) -> int:
    """
    1137. 第 N 个泰波那契数
    :param n:
    :return:
    """
    dp = [0, 1, 1]
    for i in range(3, n+1):
        dp.append(dp[i-3] + dp[i-2] + dp[i-1])
    return dp[-1]


def test_tribonacci():
    print(tribonacci(4))


