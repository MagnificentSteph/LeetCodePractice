// Time: O(nk), Space: O(nk)
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// degrade to: best-time-to-buy-and-sell-stock-ii
	if k > len(prices)/2 {
		profit := 0
		for i := 1; i < len(prices); i++ {
			profit += Max(0,prices[i]-prices[i-1])
		}
		return profit
	}

	// row   : transactions
	// column: days
	profits := make([][]int, k+1)
	for t := 0; t < len(profits); t++ {
		profits[t] = make([]int, len(prices))
	}

	for t := 1; t < k+1; t++ {
		maxThusFar := math.MinInt32
		for d := 1; d < len(prices); d++ {
			// the max forfit before day 'd'.
			maxThusFar = Max(maxThusFar, profits[t-1][d-1]-prices[d-1])
			// 1. If we don't sell on day 'd' then our max profit will be the previous day which is profits[t][d-1].
			// 2. If we sell on day 'd' then our max profit will be the price we sold plus the max profit before day 'd'.
			profits[t][d] = Max(profits[t][d-1], maxThusFar+prices[d])
		}
	}

	return profits[len(profits)-1][len(profits[0])-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}