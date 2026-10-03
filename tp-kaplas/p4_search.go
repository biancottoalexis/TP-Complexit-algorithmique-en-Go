package main

func SearchV1(ligne []int, v int) int {
	for i, k := range ligne {
		if k == v {
			return i
		}
	}
	return -1
}

func SearchV2(ligne []int, v int) int {
	gauche, droite := 0, len(ligne)-1
	for gauche <= droite {
		milieu := (gauche + droite) / 2
		switch {
		case ligne[milieu] == v:
			return milieu
		case ligne[milieu] < v:
			gauche = milieu + 1
		default:
			droite = milieu - 1
		}
	}
	return -1
}
