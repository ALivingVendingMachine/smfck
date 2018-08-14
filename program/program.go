package program

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Program represents a complete brainfuck program, with the commands and memory
// cells
type Program struct {
	commands string
	cells    []byte
	numCells int
}

// New returns a new brainfuck program
func New(comIn string, numCells int) (Program, error) {
	if numCells < 1 {
		return Program{}, fmt.Errorf("number of cells cannot be zero or negative")
	}
	s := make([]byte, numCells)
	return Program{commands: comIn, cells: s, numCells: numCells}, nil
}

// Run runs a given brainfuck program
func (p *Program) Run() (string, error) {
	var tmp int
	var ip int
	var dp int
	var ret string
	var err error

	for ip < len(p.commands) {
		if string(p.commands[ip]) == "," {
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadByte()
			p.cells[dp] = input
		} else if string(p.commands[ip]) == "." {
			fmt.Printf("%s", string(p.cells[dp]))
			ret = ret + string(p.cells[dp])
		} else if string(p.commands[ip]) == ">" {
			dp++
		} else if string(p.commands[ip]) == "<" {
			dp--
		} else if string(p.commands[ip]) == "+" {
			p.cells[dp]++
		} else if string(p.commands[ip]) == "-" {
			p.cells[dp]--
		} else if string(p.commands[ip]) == "[" {
			if p.cells[dp] == 0 {
				ip++
				for tmp > 0 || string(p.commands[ip]) != "]" {
					if string(p.commands[ip]) == "[" {
						tmp++
					}
					if string(p.commands[ip]) == "]" {
						tmp--
					}
					ip++
				}
			}
		} else if string(p.commands[ip]) == "]" {
			if p.cells[dp] != 0 {
				ip--
				for tmp > 0 || string(p.commands[ip]) != "[" {
					if string(p.commands[ip]) == "]" {
						tmp++
					}
					if string(p.commands[ip]) == "[" {
						tmp--
					}
					ip--
				}
				ip--
			}
		} else {
			fmt.Fprintf(os.Stderr, "trivia byte: %c", p.commands[ip])
			err = errors.New("non-token character encountered")
		}
		ip++
	}
	return ret, err
}
