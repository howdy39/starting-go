package foo

import (
	"testing"
)

func TestIsOne(t *testing.T) {
	n := 1
	b := IsOne(n)

	if b != true {
		t.Error("%d is not one", n)
	}

	c := IsOne(2)

	if c != false {
		t.Error("%d is not one", n)
	}
}
