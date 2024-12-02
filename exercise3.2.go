package main

import (
    "strconv"
)
func countChar(num int) int {
    count := 0
    for i := 1; i <= num; i++ {
        numStr := strconv.Itoa(i)
        count += checkChar(numStr, "9")
    }
	return count
}

func checkChar(s string, char string) int {
    count := 0
    for _, c := range s {
        if string(c) == char {
            count++
        }
    }
    return count
}
