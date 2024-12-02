package main

func CheckChar(s string, char string) int {
    count := 0
    for _, c := range s {
        if string(c) == char {
            count++
        }
    }
    return count
}