package main

import "testing"

/*

main.go
// Function to add two numbers
func Add(a, b int) int {
    return a + b
}


main_test.go
func TestAdd(t *testing.T) {
    result := Add(1, 2)
    expected := 3
    if result != expected {
        t.Errorf("Add(1, 2) = %d; want %d", result, expected)
    }
}


*/

func TestAdd(t *testing.T) {
	result := add(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Add(1, 2) = %d; want %d", result, expected)
	}
}

func TestServerStart(t *testing.T) {

}

func TestCreateUser()
