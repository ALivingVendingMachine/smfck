package main

import (
	"fmt"
	"github.com/alivingvendingmachine/smfck/repl"
	"os"
)

func main() {
	fmt.Printf("I hesitated\nbefore untying the bow\nthat bound this book together.\n")
	repl.Start(os.Stdin, os.Stdout)
}
