package main

import "fmt"

func starStairs(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	fmt.Println()
	for  i := rows; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}