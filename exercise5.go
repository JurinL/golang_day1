package main

import "fmt"

type Person struct {
	Title     string
	FirstName string
	LastName  string
	Age       int
	Address   string
	City      string
	Postcode  string
}

func main() {
	createPeople()
}

func createPeople() {
	peoples := map[string]Person{
		"person1": {
			Title:     "Mr.",
			FirstName: "Jurin",
			LastName:  "Laiyakosit",
			Age:       23,
			Address:   "INET HATYAI",
			City:      "Songkhla",
			Postcode:  "90110",
		},
		"person2": {
			Title:     "Mrs.",
			FirstName: "Jenjira",
			LastName:  "Laiyakosit",
			Age:       26,
			Address:   "Central Festival HATYAI",
			City:      "Songkhla",
			Postcode:  "90110",
		},
		"person3": {
			Title:     "Mr.",
			FirstName: "Bruce",
			LastName:  "Wayne",
			Age:       64,
			Address:   "Wayne Manor",
			City:      "Gotham City",
			Postcode:  "G1 1AA",
		},
		"person4": {
			Title:     "Mr.",
			FirstName: "Clark",
			LastName:  "Kent",
			Age:       32,
			Address:   "Daily Planet",
			City:      "Metropolis",
			Postcode:  "12345",
		},
		"person5": {
			Title:     "Mr.",
			FirstName: "Peter",
			LastName:  "Parker",
			Age:       25,
			Address:   "20 Ingram Street",
			City:      "New York",
			Postcode:  "10001",
		},
	}
	fmt.Println("Result-:")
	for _, person := range peoples {
		fmt.Printf("Name -: %s %s %s (Age: %d)\n",person.Title, person.FirstName, person.LastName, person.Age)
		fmt.Printf("Address -: %s, %s, %s\n\n", person.Address, person.City, person.Postcode)
	}
}