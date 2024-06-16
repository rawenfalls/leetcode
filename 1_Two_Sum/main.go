package main

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	for i,j := range nums{
		if idx, ok := m[target - j]; ok {
			return []int{idx, i}
		}
		m[j] = i
	}
	return nil
}