package repl

import (
	"bufio"
	"fmt"
	"github.com/alivingvendingmachine/smfck/program"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		p := program.New(line)
		p.Run(in, out)
		fmt.Printf("\n")
	}
}
