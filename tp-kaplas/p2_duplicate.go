package main

func DuplicateV1(pile []int) int {
	for i := 0; i < len(pile); i++ {
		for j := i + 1; j < len(pile); j++ {
			if pile[i] == pile[j] {
				return pile[i]
			}
		}
	}
	return -1
}

func DuplicateV2(pile []int) int {
	n := len(pile) - 1
	vu := make([]bool, n+1)
	for _, k := range pile {
		if vu[k] {
			return k
		}
		vu[k] = true
	}
	return -1
}

func DuplicateV3(pile []int) int {
	n := len(pile) - 1
	sommeAttendue := n * (n + 1) / 2
	sommeReelle := 0
	for _, k := range pile {
		sommeReelle += k
	}
	return sommeReelle - sommeAttendue
}
