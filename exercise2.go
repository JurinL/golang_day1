package main

import ( "fmt" )

func main() {
    x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
	minX := minNum(x);
	maxX := maxNum(x);
	fmt.Printf("Min number is %d\n", minX)
	fmt.Printf("Max number is %d\n", maxX)
}
func minNum(x []int) int {
	min := x[0]
	for _, value := range x {
		if value < min {
			min = value
		}
	}
	return min
}
func maxNum(x []int) int {
	max := x[0]
	for _, value := range x {
		if value > max {
			max = value
		}
	}
	return max
}