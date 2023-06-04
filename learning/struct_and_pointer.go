package main

import (
	"fmt"
)

type item struct {
	Description string
	Completed bool
}

type Todos []item

// This function has a value receiver (t Todos), meaning that it operates on a copy of the original Todos struct. Any changes made to the struct within the function will only affect the copy and not the original struct.
func (t Todos) Add (description string){
	t = append(t, item{Description: description})
}

// This function has a pointer receiver (t *Todos), meaning that it operates on a pointer to the original Todos struct. Any changes made to the struct within the function will affect the original struct.
// In general, it's recommended to use pointer receivers for methods that need to modify the underlying struct,
func (t *Todos) AddWithPointer (description string){
	*t = append(*t, item{Description: description})
}

// func main(){
// 	i := item{Description: "wash dish", Completed: false}
// 	var todos Todos
// 	todos = append(todos, i)
// 	fmt.Printf("Task %v is %v\n", todos[0].Description, todos[0].Completed)
// }
