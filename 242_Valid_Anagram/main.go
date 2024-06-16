package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
   	}

	arr := [26]int{}
	for i:=range s{
		arr[s[i]- 'a']++
		arr[t[i]- 'a']--
	}

	for _,j:=range arr{
		if j!=0{
			return false
		}
	}
	return true
}