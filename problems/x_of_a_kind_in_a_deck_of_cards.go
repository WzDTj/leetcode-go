/**
 * 914. 卡牌分组
 * https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/
 **/

package problems

func HasGroupsSizeX(deck []int) bool {
	cardsMap := map[int]int{}
	for _, v := range deck {
		cardsMap[v]++
	}

	var gcd int
	for _, v := range cardsMap {
		switch {
		case gcd == 0:
			gcd = v
		default:
			gcd = findGCD(gcd, v)
		}
	}

	return gcd >= 2
}

func findGCD(x, y int) int {
	if x == 0 {
		return y
	}

	return findGCD(y%x, x)
}
