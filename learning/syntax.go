package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"io/ioutil"
	"log"
)

func new_function() {
	// var a *int
	var b *bool
	var s *string
	var arr *[10]int

	// cấp phát bộ nhớ cho biến
	a := new(int)
	b = new(bool)
	s = new(string)
	arr = new([10]int)

	fmt.Println(a, " ", *a)
	fmt.Println(s, " ", *s)
	fmt.Println(b, " ", *b)
	fmt.Println(arr, " ", *arr)
}

func append_function() {
	my_slice := []string{}
	//my_slice_2 := []string{}
	my_slice_3 := []string{}
	my_slice = []string{"0", "1", "2","3","4"}
	fmt.Println(my_slice)


	//my_slice_2 = append(my_slice, "Jane", "Maya")
	// fmt.Println(my_slice_2)


	index := 2
	my_slice_3 = append(my_slice[:index-1], my_slice[index:]...)

	fmt.Println(my_slice_3)
	// fmt.Println(my_slice_3)
}


func bufio_package() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type something: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed: %v\n", input)
}


func flag_package() {

    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}


func read_file() {
	content, err := ioutil.ReadFile("/home/buinguyen/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

func pointer(){
	var x int = 54
	var p *int
	p = &x
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Adderss of x = ", &x)
	fmt.Println("Value stored in  p = ", p)
	fmt.Println("Value stored in  *p = ", *p)
	fmt.Println("Value stored in  &p = ", &p)
}
