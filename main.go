package main

import (
	"fmt"
	"strconv"
)
func main() {
	// Exercise 0
	switchIf()
	fmt.Println("-----------------------------------------------------------------------------")
	
	// Exercise 1.1
	j := 0
    for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			j++
		}
	}
	fmt.Printf("There is %d numbers between 1-100 that can divide 3 perfectly.",j)
	fmt.Println("-----------------------------------------------------------------------------")
	
	// Exercise 1.2
	result := powerNum(20, 3)
    fmt.Printf("20^3 = %d\n", result)
	fmt.Println("-----------------------------------------------------------------------------")

	// Exercise 2
	x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
	minX := minNum(x);
	maxX := maxNum(x);
	fmt.Printf("Min number is %d\n", minX)
	fmt.Printf("Max number is %d\n", maxX)
	fmt.Println("-----------------------------------------------------------------------------")

	// Exercise 3.1
	count1 := 0
    
    for i := 1; i <= 1000; i++ {
        numStr := strconv.Itoa(i)
        count1 += CheckChar(numStr, "9")
    }
    
    fmt.Printf("The digit 9 appears %d times in numbers from 1 to 1000\n", count1)
	fmt.Println("-----------------------------------------------------------------------------")

	// Exercise 3.2
	num := 10000
    count2 := countChar(num)
    fmt.Printf("The digit 9 appears %d times in numbers from 1 to %d\n", count2,num)
	fmt.Println("-----------------------------------------------------------------------------")

	// Exercise 4.1
	var myWords1 = "AW SOME GO!"
    var nospace1 string
    for _, char := range myWords1 {
        if char != ' ' {
            nospace1 += string(char)
        }
    }
	fmt.Println(nospace1)
	fmt.Println("-----------------------------------------------------------------------------")

	// Exercise 4.2
	var myWords2 = "ine t"
    nospace2 := cutText(myWords2)
    fmt.Println(nospace2)
	fmt.Println("-----------------------------------------------------------------------------")

	// Exercise 5
	createPeople()
	fmt.Println("-----------------------------------------------------------------------------")

	// Exercise 6
	INET := Company{
		"บริษัท อินเทอร์เน็ตประเทศไทย จำกัด (มหาชน)",
		"1768 อาคารไทยซัมมิท ทาวเวอร์ ชั้น 10-12 และ\nชั้น IT ถ.เพชรบุรีตัดใหม่ แขวงบางกะปิ เขตห้วยขวาง กรุงเทพมหานคร 10310",
		"info@inet.co.th",
		"02 257 7000 (Call Center), 02 257 7111 (NOC)",
		"INET Cloud",
	}

	fmt.Printf("\nCompany Name: %s\n", INET.Name)
	fmt.Printf("Address: %s\n", INET.Address)
	fmt.Printf("Email: %s\n", INET.Email)
	fmt.Printf("Tel.: %s\n", INET.Tel)
	fmt.Printf("Facebook: %s\n\n", INET.Facebook)
	fmt.Println("-----------------------------------------------------------------------------")

	// Extra
	var rows int
	fmt.Printf("Enter the number of rows: ")
	fmt.Scanf("%d", &rows)
	starStairs(rows);
	fmt.Println("-----------------------------------------------------------------------------")
}