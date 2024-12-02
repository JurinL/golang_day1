package main

func powerNum(num, pow int) int {
    result := 1
    for i := 0; i < pow; i++ {
        result *= num
    }
    return result
}
