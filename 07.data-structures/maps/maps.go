package main

import (
	"fmt"
)

func main() {
	// websites := map[string]string{
	// 	"Google":              "https://google.com",
	// 	"Amazon Web Services": "https://aws.com",
	// }
	// fmt.Println(websites)
	
	// websites["Air BnB"] = "https//:airbnb.com"
	// delete(websites, "Google")	
	// fmt.Println(websites)

	arr := make([]int, 0, 10)

	arr = append(arr, 1)
	arr = append(arr, 2)

	fmt.Println("Array Values: ", arr)
	fmt.Println("Array Length: ", len(arr), "Array Cap: ", cap(arr))
}
