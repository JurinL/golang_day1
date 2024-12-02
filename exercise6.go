package main

import "fmt"

type Company struct {
	Name, Address, Email, Tel, Facebook string
}

func main() {
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
}
