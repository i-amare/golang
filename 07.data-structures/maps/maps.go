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
	fmt.Println(websites)

	for key, val := range websites {
		fmt.Println(key, val)
	}
}
