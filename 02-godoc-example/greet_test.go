package example

import "fmt"

func ExampleHello() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}

func ExamplePage() {
	checkIns := map[string]bool{
		"Bob":  true,
		"Jon":  false,
		"Gary": false,
		"Eve":  true,
	}

	Page(checkIns)

	// Unordered Output:
	// Paging Jon: please see the front desk to check in
	// Paging Gary: please see the front desk to check in
}

func ExampleHello_spanish() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}

func Demo_Hey() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}
