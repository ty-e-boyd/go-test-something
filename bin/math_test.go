package calculator

import "testing"

func TestAdd(t *testing.T) {
	r := Add([]int{3, 3})

	if r != 6 {
		t.Errorf("Result was incorrect. Got %v, wanted %v", r, 6)
	}
}
