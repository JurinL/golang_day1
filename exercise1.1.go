package main

import(
	"fmt"
)

func main() {
	j := 0
    for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			j++
		}
	}
	fmt.Printf("There is %d numbers between 1-100 that can divide 3 perfectly.",j)
}
