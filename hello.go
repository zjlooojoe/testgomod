package hello

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "noName"
	}
	return fmt.Sprintf("hello, %s", name)
}
