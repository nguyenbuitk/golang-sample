package main

import (
	"flag"
	"fmt"
	"os"
	"todoproject"
	// "github.com/alexeyco/simpletable"
	"github.com/common-nighthawk/go-figure"
)

const (
	fileName = "/home/ubuntu/repository/golang-tutorial/sample-project/.todos.json"
)

func main() {
	myFigure := figure.NewFigure("cli  todos", "", true)
	myFigure.Print()
	fmt.Println()
	addPtr := flag.String("add", "", "add a item of todo")
	deletePtr := flag.Int("delete", 0, "delete an item")
	completePtr := flag.Int("complete", 0, "mark an item as complete")
	listPtr := flag.Bool("list", false, "list item of to do")

	flag.Parse()

	todos := &todo.Todos{}
	todos.Load(fileName)
	switch {
	case *addPtr != "":
		err := todos.Add(*addPtr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = todos.Store(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *deletePtr != 0:
		err := todos.Delete(*deletePtr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = todos.Store(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *completePtr != 0:
		err := todos.Complete(*completePtr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *listPtr:
		todos.Print()

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
	}

	// err := todos.Complete(-2)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err.Error())
	// }
	// fmt.Println("2")
}
