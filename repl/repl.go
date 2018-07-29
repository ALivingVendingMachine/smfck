package repl

import (
	"bufio"
	"fmt"
	"github.com/alivingvendingmachine/smfck/program"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, cells int) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		p := program.New(line, cells)
		p.Run(in, out)
		fmt.Printf("\n")
	}
}
