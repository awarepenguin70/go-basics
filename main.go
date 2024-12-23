// main package - every go program starts with this
package main

// importing necessary packages
import (
	"fmt"
)

// main function - the program starts execution here
func main() {
	fmt.Println("Welcome to Golang Basics!")

	constantsAndVariables()
	functionsAndControlStructures()
	collectionsAndLoops()
	stringsRunesAndBytes()
	structsAndInterfaces()
	pointers()
}

// Constants, Variables, and Basic Data Types
func constantsAndVariables() {
	fmt.Println("\n--- Constants, Variables, and Basic Data Types ---")

	// constants
	const pi = 3.14
	fmt.Println("Value of constant pi:", pi)

	// variables
	var name string = "Rahul" // explicitly typed variable
	age := 20                 // shorthand declaration
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// basic data types
	var isStudent bool = true
	var score float64 = 95.5
	fmt.Println("Am I a student?", isStudent)
	fmt.Println("My score is:", score)
}

// Functions and Control Structures
func functionsAndControlStructures() {
	fmt.Println("\n--- Functions and Control Structures ---")

	// calling a function
	sum := add(10, 15)
	fmt.Println("Sum of 10 and 15 is:", sum)

	// if-else
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// switch-case
	day := "Saturday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println(day, "is a weekend.")
	default:
		fmt.Println(day, "is a weekday.")
	}
}

// a simple function to add two numbers
func add(a, b int) int {
	return a + b
}

// Arrays, Slices, Maps, and Loops
func collectionsAndLoops() {
	fmt.Println("\n--- Arrays, Slices, Maps, and Loops ---")

	// arrays
	var arr = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// slices
	slice := []int{4, 5, 6}
	fmt.Println("Slice:", slice)

	// maps
	countries := map[string]string{"India": "New Delhi", "USA": "Washington"}
	fmt.Println("Map of countries:", countries)

	// for loop
	fmt.Println("Looping through numbers 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// range loop
	fmt.Println("Looping through slice:")
	for idx, val := range slice {
		fmt.Printf("Index %d has value %d\n", idx, val)
	}
}

// Strings, Runes, and Bytes
func stringsRunesAndBytes() {
	fmt.Println("\n--- Strings, Runes, and Bytes ---")

	// strings
	str := "Hello, Go!"
	fmt.Println("String:", str)

	// converting string to bytes
	bytes := []byte(str)
	fmt.Println("Bytes:", bytes)

	// iterating through runes
	fmt.Println("Runes in the string:")
	for i, r := range str {
		fmt.Printf("Index %d: %c\n", i, r)
	}
}

// Structs and Interfaces
func structsAndInterfaces() {
	fmt.Println("\n--- Structs and Interfaces ---")

	// struct example
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Rahul", Age: 21}
	fmt.Println("Person struct:", person)

	// interface example
	cat := Cat{Name: "Whiskers"}
	dog := Dog{Name: "Buddy"}

	animals := []Animal{cat, dog}
	fmt.Println("Animals speaking:")
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

// defining an interface
type Animal interface {
	Speak() string
}

// defining a struct for Cat
type Cat struct {
	Name string
}

// implementing the Speak method for Cat
func (c Cat) Speak() string {
	return c.Name + " says meow!"
}

// defining a struct for Dog
type Dog struct {
	Name string
}

// implementing the Speak method for Dog
func (d Dog) Speak() string {
	return d.Name + " says woof!"
}

// Pointers
func pointers() {
	fmt.Println("\n--- Pointers ---")

	num := 42
	ptr := &num // pointer to num
	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", ptr)

	// modifying value using pointer
	*ptr = 100
	fmt.Println("Modified value of num:", num)
}
