package math

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum([]int{10, -2, 3})
	if sum != 11 {
		t.Fatalf("Wanted 11 but received %d", sum)
	}
}

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("http.NewRequest() err = %s", err)
	}
	Handler(w, r)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fatalf("Handler() status = %d; want = %d", resp.StatusCode, 200)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Handler() Content-Type = %q; want = %q", contentType, "application/json")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll err = %s", err)
	}
	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		t.Fatalf("json.Unmarshal err = %s", err)
	}

	if p.Age != 21 {
		t.Errorf("person.Age = %d; want = %d", p.Age, 21)
	}
}
