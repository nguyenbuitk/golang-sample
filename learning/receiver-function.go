package main

import (
	"fmt"
)

// define Todo Struct
type Todo struct {
	name string
	completed bool
}

// define Todos Type
type Todos []Todo

// t sẽ triển khai Complete method: t.Complete(1)
// có thể hình dung Todos như 1 class (bởi vì trong c++, struct có thể coi như là class theo một khía cạnh nào đó)
func (t *Todos) Complete(index int) error {
	if index < 0 || index >= len(*t) {
		return fmt.Errorf("index out of range")
	}

	(*t)[index].completed = true
	return nil
}

func main() {
	// todos = Class Todos
	todos := Todos {
		Todo{"Task 1", false},
		Todo{"Task 2", false},
		Todo{"Task 3", false},
	}

	fmt.Println("Before completing: ", todos)

	if err := todos.Complete(1); err != nil {
		fmt.Println("Error completing task: ", err)
	}

	fmt.Println("After completing: ", todos)
}
