package math

import (
	"encoding/json"
	"net/http"
)

func Sum(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

type Person struct {
	Age        int
	Name       string
	Occupation string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	p := Person{
		Age:        21,
		Name:       "Bob Jones",
		Occupation: "Nurse",
	}
	data, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
