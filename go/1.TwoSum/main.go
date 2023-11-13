package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for a, b := range nums {
		if b >= target {continue}
		for i := a+1; i < len(nums); i++ {
			if b+nums[i] == target{return []int{a,i}}
		}
	}
	return []int{0,0}
}

func main() {
	a:=[]int{1,2,3,4}
	fmt.Println(twoSum(a,8))
}