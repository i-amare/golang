package main

import (
	"fmt"
)

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	
	websites["Air BnB"] = "https//:airbnb.com"
	delete(websites, "Google")	
	fmt.Println(websites)
}
