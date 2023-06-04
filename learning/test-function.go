package main

import "fmt"

type arr struct {
    x int
    y int
}

func sum(x int, y int, a arr) int {
    return a.x + a.y + x + y
}

// func main() {
//     a := arr{x: 1, y: 2}
//     result := sum(3, 4, a)
//     fmt.Println(result) // Kết quả là 10
// }
