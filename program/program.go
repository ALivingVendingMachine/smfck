package program

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

// This is sort of a cheat.  The parser is also kinda the interpreter, because
// I'm not going to an IL.  Oops.

type Program struct {
	tok   []byte
	cells []byte
	numcells int
}

func New(input string, cellSize int) *Program {
	s := make([]byte, cellSize)
	p := &Program{tok: []byte(input), cells: s, numcells: cellSize}
	return p
}

func (p *Program) Run(in io.Reader, out io.Writer) []byte {
	err := p.Parse(0, in, out)
	if err != nil {
		fmt.Printf("parse error: %v\n", err)
	}
	return p.cells
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
			if cellptr == p.numcells {
				cellptr = 0
			}
			i++
		case '<':
			cellptr--
			if cellptr < 0 {
				cellptr = p.numcells - 1
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
			if p.cells[cellptr] >= 33 && p.cells[cellptr] <= 126 {
				fmt.Fprintf(out, "%c", p.cells[cellptr])
			} else {
				fmt.Fprintf(out, "%d", p.cells[cellptr])
			}
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
