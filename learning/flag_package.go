package main


import (
	"flag"
	"fmt"
)

func main() {
	// ./flag_package -word=opt -numb=7 -fork -svar=flag
	// ./flag_package -word=opt
	// ./flag_package -word=opt a1 a2 a3
	// ./flag_package -word=opt a1 a2 a3 -numb=7
	// ./flag_package -h
	// ./flag_package -wat

	// command-line flag called "word"
	//		"word" - Flag name, the user could run it with `myprogram -word=bar`
	//		"foo" - The default value for the flag. If the user does not specify a value for the flag, this value will be used.
	// 		"a string" - A brief description of what the flag does. This text will be displayed to the user when they run the program with the -h or --help flag.
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string

	// first argument pointer to a variable that will hold the values of the '-svar' flag, 'svar' is being passed as a pointer 
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
