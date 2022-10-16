package data

import "testing"

func TestCheck(t *testing.T) {
	p := Product{
		Name:  "nics",
		Price: 1.00,
		SKU:   "asf",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
