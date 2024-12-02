package main

func cutText(text string) string {
    var result string
    for _, char := range text {
        if char != ' ' {
            result += string(char)
        }
    }
    return result
}