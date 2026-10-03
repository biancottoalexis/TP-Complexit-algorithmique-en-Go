package main

func SmallestV1(pile []int) int {
	min := pile[0]
	for i := 1; i < len(pile); i++ {
		if pile[i] < min {
			min = pile[i]
		}
	}
	return min
}
