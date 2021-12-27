//go:build !foo

package p

import "fmt"

func Bar() {
	fmt.Println("Hello world")
}