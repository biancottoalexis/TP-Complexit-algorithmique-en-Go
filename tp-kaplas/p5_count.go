package main

func CountV1(pile []int, plafond int) []int {
	compte := make([]int, plafond+1)
	for k := 1; k <= plafond; k++ {
		n := 0
		for _, kapla := range pile {
			if kapla == k {
				n++
			}
		}
		compte[k] = n
	}
	return compte
}

func CountV2(pile []int, plafond int) []int {
	compte := make([]int, plafond+1)
	for _, k := range pile {
		compte[k]++
	}
	return compte
}
