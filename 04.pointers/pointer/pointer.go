package main

func main() {
	age := 32
	agePointer := &age

	println("Age Pointer: ", age)
	println("Age Pointer: ", agePointer)

	changeAge(agePointer, 21)

	println("Age Pointer: ", age)
	println("Age Pointer: ", agePointer)
}

func changeAge(agePointer *int, newAge int) {
	*agePointer = newAge
}

