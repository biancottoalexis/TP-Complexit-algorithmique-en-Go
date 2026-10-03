package main

import "fmt"

func main() {
	n := Shuffled(10)
	fmt.Println(n)
	fmt.Println(SmallestV1(n))

	d := []int{7, 2, 9, 6, 1, 10, 4, 3, 6, 8, 5}
	fmt.Println(DuplicateV1(d), DuplicateV2(d), DuplicateV3(d))

	fmt.Println(TowerHeightV1(10), TowerHeightV2(10))

	ligne := []int{2, 4, 5, 8, 9, 10}
	fmt.Println(SearchV1(ligne, 8), SearchV2(ligne, 8))

	fmt.Println(CountV1([]int{3, 1, 3, 2}, 3))
	fmt.Println(CountV2([]int{3, 1, 3, 2}, 3))
}
