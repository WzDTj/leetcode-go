package problems

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		idx, ok := m[complement]
		if ok {
			return []int{idx, i}
		} else {
			m[v] = i
		}
	}

	return nil
}
