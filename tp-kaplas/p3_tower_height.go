package main

func TowerHeightV1(n int) int {
	somme := 0
	for i := 1; i <= n; i++ {
		somme += i
	}
	return somme
}

func TowerHeightV2(n int) int {
	return n * (n + 1) / 2
}
