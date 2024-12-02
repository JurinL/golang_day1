package main

import "fmt"

func main() {
    switchIf()
}
func switchIf() {
    i := 2
    
    fmt.Println("Example-: Switch case condition to If") 

    if(i == 0){
        fmt.Println("Zero")
    }else if(i == 1){
        fmt.Println("One")
    } else if(i == 2){
        fmt.Println("Two")
    } else if(i == 3){
        fmt.Println("Three")
    } else{
        fmt.Println("Your i not in case.")
    }
}