package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#000000",
	}

	Print(colors)
}

// function print map of colors
func Print(colors map[string]string) {
    for k, v := range colors {
        fmt.Println("hex code for", k, "is",  v)
    }
}

// function remove colors from
