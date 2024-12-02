package main

import "fmt"

func main() {
    result := powerNum(20, 3)
    fmt.Printf("20^3 = %d\n", result)
}

func powerNum(num, pow int) int {
    result := 1
    for i := 0; i < pow; i++ {
        result *= num
    }
    return result
}
