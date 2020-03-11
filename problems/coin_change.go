/**
 * 322. 零钱兑换
 * https://leetcode-cn.com/problems/coin-change/
 **/

package problems

import (
	"math"
	"sort"
)

func CoinChange(coins []int, amount int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	result := math.MaxInt32
	calculateCoins(amount, coins, 0, 0, &result)

	if result == math.MaxInt32 {
		return -1
	}

	return result
}

func calculateCoins(amount int, coins []int, coinIndex int, count int, result *int) {

	if amount == 0 {
		*result = int(math.Min(float64(count), float64(*result)))
		return
	}
	if coinIndex >= len(coins) {
		return
	}

	for i := amount / coins[coinIndex]; i >= 0 && i+count < *result; i-- {
		remainder := i * coins[coinIndex]

		calculateCoins(amount-remainder, coins, coinIndex+1, count+i, result)

	}
}
