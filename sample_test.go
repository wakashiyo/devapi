package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	result := calc(1, 2)
	if result != 3 {
		t.Fatal("failed test")
	}
}
