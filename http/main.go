package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// custom type logWriter
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// 	The io.Copy function is called, which copies the response body from resp.Body to lw. The io.Copy function reads data from the source (resp.Body) and writes it to the destination (lw). In this case, the destination is an instance of logWriter.

	// The logWriter type implements the Write method of the io.Writer interface. This method is automatically invoked by io.Copy for each chunk of data read from the response body. In this implementation, it prints the received bytes as a string and the number of bytes received
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
