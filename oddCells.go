package odd_cell

func oddCells(n int, m int, indices [][]int) int {
	result := 0
	Xmap := make(map[int]int)
	Ymap := make(map[int]int)
	xCount := 0
	yCount := 0
	for _, point := range indices {
		Xmap[point[1]] += 1
		Ymap[point[0]] += 1
	}
	for _, value := range Xmap {
		if value%2 == 1 {
			xCount += 1
		}
	}
	for _, value := range Ymap {
		if value%2 == 1 {
			yCount += 1
		}
	}
	result += yCount * (m - len(Xmap))
	result += xCount * (n - len(Ymap))
	result += xCount*(len(Ymap)-yCount) + yCount*(len(Xmap)-xCount)
	return result
}
