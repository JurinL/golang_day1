package main

import "fmt"

func main() {
    var myWords = "ine t"
    nospace := cutText(myWords)
    fmt.Println(nospace)
}

func cutText(text string) string {
    var result string
    for _, char := range text {
        if char != ' ' {
            result += string(char)
        }
    }
    return result
}