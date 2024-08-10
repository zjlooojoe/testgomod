package hello_test

import (
	"fmt"
	hello "github.com/zjlooojoe/testgomod"
	"testing"
)

func TestHello(t *testing.T) {
	data := "joe"
	expected := fmt.Sprintf("hello %s", data)
	result := hello.Hello(data)
	if result != expected {
		t.Fatalf("expected resule %s, but got %s", expected, result)
	}
}
