package main

import "fmt"

func main() {
	merge([]int{1}, 1, []int{}, 0)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var tmp1, tmp2, start int
	if len(nums2) == 0 {
		return
	}
	if nums1[0] > nums2[0] {
		start = nums2[0] - 1
	} else {
		start = nums1[0] - 1
	}
	for {
		if n+m-1 < 0 {
			break
		}
		if m > 0 {
			tmp1 = nums1[m-1]
		} else {
			tmp1 = start
		}
		if n > 0 {
			tmp2 = nums2[n-1]
		} else {
			tmp2 = start
		}
		if tmp1 > tmp2 {
			fmt.Println(m + n - 1)
			nums1[m+n-1] = tmp1
			if m > 0 {
				m--
			}
		} else {
			fmt.Println(m + n - 1)
			nums1[m+n-1] = tmp2
			if n > 0 {
				n--
			}
		}
	}
	fmt.Println(nums1)
}
