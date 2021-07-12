package main

import (
	"fmt"
	"io"
	"os"
)

func MainOutput(out io.Writer) {
	
	fmt.Fprintln(out, "Hello World!")

}

func main() {
	MainOutput(os.Stdout)
}
