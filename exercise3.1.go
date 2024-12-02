package main

import (
    "fmt"
    "strconv"
)

func main() {
    count := 0
    
    for i := 1; i <= 1000; i++ {
        numStr := strconv.Itoa(i)
        count += CheckChar(numStr, "9")
    }
    
    fmt.Printf("The digit 9 appears %d times in numbers from 1 to 1000\n", count)
}

func CheckChar(s string, char string) int {
    count := 0
    for _, c := range s {
        if string(c) == char {
            count++
        }
    }
    return count
}