package main

func containsDuplicate(nums []int) bool {
    m := make(map[int]struct{})
	for _,j := range nums{
		if _, ok := m[j]; ok{
			return true
		}
		m[j] = struct{}{}
	}
	return false
}