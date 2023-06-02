// Package greetings provides a function for saying hello.
package greetings

import "fmt"

// SayHello prints a greeting to the console.
func SayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
