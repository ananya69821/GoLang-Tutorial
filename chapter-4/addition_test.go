package main

import "testing"

func Add(x, y int) int {
    return x + y
}

// go test cli command
func TestAdd(t *testing.T) {
    sum := Add(2, 3)
    if sum != 5 {
        t.Errorf("Expected 5, but got %v", sum)
    } else {
		t.Log("Hooray")
	}
}