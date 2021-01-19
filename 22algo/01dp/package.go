package dp

/**
背包问题：金矿问题
有一个国家发现了5座金矿，每座金矿的黄金储量不同，
需要参与挖掘的工人数也不同。参与挖矿工人的总数是10人。
每座金矿要么全挖，要么不挖，不能派出一半人挖取一半金矿。
要求用程序求解出，要想得到尽可能多的黄金，应该选择挖取哪几座金矿？
*/

/**
workNum - 现有的工人数量
minePrices - 每个矿的价格
mineWorkload - 每个矿需要的人数
*/
func solve1(workerNum int, minePrices, mineWorkload []int) int {
	return recursive(len(minePrices)-1, workerNum, minePrices, mineWorkload)
}

/**
dp
*/
func solve2(workerNum int, minePrices, mineWorkload []int) int {
	dp := make([][]int, len(minePrices))

	length := len(minePrices)

	// 这里写第0行以下的记录
	for i := 0; i < length; i++ {
		dp[i] = make([]int, workerNum+1)

		for j := 0; j <= workerNum; j++ {
			// 如果当前的人数不满足
			if mineWorkload[i] > j {
				dp[i][j] = 0
				continue
			}

			if i == 0 {
				dp[i][j] = minePrices[i]
				continue
			}

			dp[i][j] = max(dp[i-1][j], dp[i-1][j-mineWorkload[i]]+minePrices[i])
		}
	}

	return dp[length-1][workerNum]
}

func recursive(i, workerNum int, minePrices, mineWorkload []int) int {
	if mineWorkload[i] > workerNum {
		return 0
	}

	if i == 0 {
		if mineWorkload[i] >= workerNum {
			return minePrices[i]
		}
		return 0
	}

	return max(recursive(i-1, workerNum, minePrices, mineWorkload), recursive(i-1, workerNum-mineWorkload[i], minePrices, mineWorkload)+minePrices[i])
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
