/**
 * 1160. 拼写单词
 * https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/
 **/

package problems

func CountCharacters(words []string, chars string) int {
	var ans int

	charsMap := makeMap(chars)

	for _, word := range words {
		wordMap := makeMap(word)
		if canSpell(charsMap, wordMap) {
			ans += len(word)
		}
	}

	return ans
}

func makeMap(str string) map[string]int {
	theMap := map[string]int{}
	for _, char := range str {
		theMap[string(char)]++
	}

	return theMap
}

func canSpell(charsMap, wordMap map[string]int) bool {
	for k, v := range wordMap {
		if v > charsMap[k] {
			return false
		}
	}

	return true
}
