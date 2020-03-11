/**
 * 121. 买卖股票的最佳时机
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
 **/

package problems

import (
	"math"
)

/**
 * 递推：
 * 1、依次计算最大售价
 * 2、依次计算最大利润
 **/
func MaxProfit(prices []int) int {
	var maxProfit, maxSellPrice float64

	for i := len(prices) - 2; i >= 0; i-- {
		maxSellPrice = math.Max(maxSellPrice, float64(prices[i+1]))
		maxProfit = math.Max(maxProfit, maxSellPrice-float64(prices[i]))
	}

	return int(maxProfit)
}
