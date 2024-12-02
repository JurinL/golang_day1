package main

import (
    "fmt"
    "strconv"
)

func main() {
	num := 10000
    count := countChar(num)
    fmt.Printf("The digit 9 appears %d times in numbers from 1 to %d\n", count,num)
}
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
