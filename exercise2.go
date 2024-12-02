package main

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