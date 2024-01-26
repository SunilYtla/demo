// main_test.go
package main

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello, Go!"
    result := hello()
    if result != expected {
        t.Errorf("Expected %s, but got %s", expected, result)
    }
}

func hello() string {
    return "Hello, Go!"
}
