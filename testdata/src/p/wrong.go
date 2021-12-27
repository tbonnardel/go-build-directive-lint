// +build !foo // want `build constraints should use the 'go:build' syntax`

package p

import "fmt"

func Foo() {
	fmt.Println("Hello world")
}