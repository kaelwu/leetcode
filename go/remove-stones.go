package main

import "fmt"

func main() {
	fmt.Println(removeStones([][]int{[]int{0, 0}, []int{1, 2}}))
}

func removeStones(stones [][]int) int {
	remove := 0
	if len(stones) == 0 {
		return 0
	}
	m := make([][]int, 0)
	var maxX, maxY int
	for _, v := range stones {
		if v[0] > maxX {
			maxX = v[0]
		}
		if v[1] > maxY {
			maxY = v[1]
		}
	}
	maxX += 1
	maxY += 1
	for i := 0; i < maxX; i++ {
		tmp := make([]int, maxY)
		m = append(m, tmp)
	}
	for _, v1 := range stones {
		x := v1[0]
		y := v1[1]
		m[x][y] = 1
	}
	for i := 0; i < maxX; i++ {
		line := 0
		for j := 0; j < maxY; j++ {
			if m[i][j] == 1 {
				line++
			}
		}
		if line >= 1 {
			remove += line - 1
		}
	}
	for j := 0; j < maxY; j++ {
		list := 0
		for i := 0; i < maxX; i++ {
			if m[i][j] == 1 {
				list++
			}
		}
		if list >= 1 {
			remove += list - 1
		}
	}
	if remove >= 1 {
		return remove - 1
	}
	return 0
}
