package program

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

const NUMCELLS = 30000

// This is sort of a cheat.  The parser is also kinda the interpreter, because
// I'm not going to an IL.  Oops.

type Program struct {
	tok   []byte
	cells [NUMCELLS]byte
}

func New(input string) *Program {
	p := &Program{tok: []byte(input)}
	return p
}

func (p *Program) Run(in io.Reader, out io.Writer) {
	err := p.Parse(0, in, out)
	if err != nil {
		fmt.Printf("parse error: %v\n", err)
	} else {
		// print cells maybe?
	}
}

func (p *Program) restore() {
	for i := 0; i < len(p.tok); i++ {
		if p.tok[i] == '}' {
			p.tok[i] = ']'
		} else if p.tok[i] == '{' {
			p.tok[i] = '['
		}
	}
}

func (p *Program) Parse(pos int, in io.Reader, out io.Writer) error {
	cellptr := 0
	for i := pos; i < len(p.tok); {
		switch p.tok[i] {
		case '>':
			cellptr++
			if cellptr == NUMCELLS {
				cellptr = 0
			}
			i++
		case '<':
			cellptr--
			if cellptr < 0 {
				cellptr = NUMCELLS - 1
			}
			i++
		case '+':
			p.cells[cellptr]++
			i++
		case '-':
			p.cells[cellptr]--
			i++
		case ',':
			reader := bufio.NewReader(in)
			input, _ := reader.ReadString('\n')
			p.cells[cellptr] = []byte(input)[0]
			i++
		case '.':
			fmt.Fprintf(out, "%c", p.cells[cellptr])
			i++
		case '[':
			ride := -1
			tmp := i + 1
			for tmp < len(p.tok) {
				if p.tok[tmp] == ']' {
					ride = tmp
				}
				tmp++
			}
			if ride == -1 {
				return errors.New("[ without ]")
			}
			if p.cells[cellptr] == 0 {
				p.tok[i] = '{'
				i = ride + 1
				p.tok[ride] = '}'
			} else {
				i++
			}
		case ']':
			p.tok[i] = '}' // Error?  This means ] without [, which shouldn't happen...
		default:
			i++
		}
	}
	p.restore()
	return nil
}
