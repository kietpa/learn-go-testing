package example

import "fmt"

type Demo struct{}

func (d Demo) Hey() {}

func Hello(name string) (string, error) {
	return fmt.Sprintln("Hello,", name), nil
}

// unordered output example

// checks who hasn't checked in
func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s: please see the front desk to check in\n", name)
		}
	}
}
