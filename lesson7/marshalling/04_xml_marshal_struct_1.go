package main

import (
	"encoding/xml"
	"fmt"
)

// User1 struct represent one item to be marshalled into XML
type User1 struct {
	id      uint32
	name    string
	surname string
}

// User2 struct represent one item to be marshalled into XML
type User2 struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	user1 := User1{
		1,
		"Pepek",
		"Vyskoč"}

	user2 := User2{
		1,
		"Pepek",
		"Vyskoč"}

	user1asXML, _ := xml.Marshal(user1)
	fmt.Println(string(user1asXML))

	fmt.Println()

	user2asXML, _ := xml.Marshal(user2)
	fmt.Println(string(user2asXML))
}
