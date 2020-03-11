/**
 * 1103. 分糖果 II
 * https://leetcode-cn.com/problems/distribute-candies-to-people/
 **/

package problems

import (
	// "fmt"
	"math"
)

func DistributeCandies(candies int, num_people int) []int {
	// 计算完整给予的次数。通过等差数列求和，然后解二次函数可得。
	times := int(math.Ceil((-1 + math.Sqrt(float64(1+8*candies))) / 2))
	// 计算给予得轮次，向上取整。
	rounds := int(math.Ceil(float64(times) / float64(num_people)))
	// 计算剩余多少。
	remainder := 0
	completedTimesCandies := (times*times + times) / 2
	if completedTimesCandies > candies {
		noRemainingCandies := (int(math.Pow(float64(times-1), 2.0)) + times - 1) / 2
		remainder = candies - noRemainingCandies
	}

	// fmt.Println("times: ", times)
	// fmt.Println("rounds: ", rounds)
	// fmt.Println("remainder: ", remainder)

	var peoples []int
	for i := 0; i < num_people; i++ {
		finalTime := num_people*(rounds-1) + (i + 1)
		var sum int
		switch {
		case finalTime < times:
			fallthrough
		case finalTime == times && remainder == 0:
			sum = rounds*(i+1) + num_people*(rounds*rounds-rounds)/2
		case finalTime == times:
			round := rounds - 1
			sum = round*(i+1) + num_people*(round*round-round)/2 + remainder
		case finalTime > times:
			round := rounds - 1
			sum = round*(i+1) + num_people*(round*round-round)/2
		}
		// fmt.Println("peoples #", i+1, " got ", sum, " candies.")

		peoples = append(peoples, sum)
	}
	return peoples
}
