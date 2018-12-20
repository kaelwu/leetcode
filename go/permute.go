package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	tmp := nums[len(nums)-1]
	tmpNums := permute(nums[:len(nums)-1])
	var newNums [][]int
	for _, v := range tmpNums {
		v = append(v, tmp)
		tmpV := make([]int, len(v))
		copy(tmpV, v)
		for k := range tmpV {
			//a1 := tmp[K]
			//a2 := tmpV[len(tmpV)-1]
			tmpV[k], tmpV[len(tmpV)-1] = tmpV[len(tmpV)-1], tmpV[k]
			insertData := make([]int, len(tmpV))
			copy(insertData, tmpV)
			newNums = append(newNums, insertData)
			// 交换回来
			tmpV[len(tmpV)-1], tmpV[k] = tmpV[k], tmpV[len(tmpV)-1]
		}
	}
	return newNums
}
