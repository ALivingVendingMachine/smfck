package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/alivingvendingmachine/smfck/repl"
	"github.com/alivingvendingmachine/smfck/program"
	"github.com/jroimartin/gocui"
	"io/ioutil"
	"os"
)

var filename string
var numcells int

func main() {
	replFlag := flag.Bool("repl", false, "run in repl mode")
	cellsFlag := flag.Int("cells", 64, "number of cells")
	numcells = *cellsFlag
	flag.Parse()
	filename = flag.Arg(0)
	if *replFlag {
		fmt.Printf("I hesitated\nbefore untying the bow\nthat bound this book together.\n")
		repl.Start(os.Stdin, os.Stdout, numcells)
		return
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		panic(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("header", 0, 0, maxX - 1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Title = "brainfuck"
	}
	if v, err := g.SetView("inst", 0, 3, maxX/2 - 2, maxY - 4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Title = "instructions"
		if filename == "" {
			fmt.Fprintf(v, "-------------------------\n")
		} else {
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}
			fmt.Fprintf(v, "%s", file)
			v.Editable = true
			v.Wrap = true
			if _, err := g.SetCurrentView("inst"); err != nil {
				return err
			}
		}
		return nil
		// MAKE EDITOR
	}
	if v, err := g.SetView("data", maxX/2 - 1, 3, maxX - 1, maxY - 4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		lineSize := ((maxX - 1) - (maxX/2 - 1)) / 5
		numLines := (maxY - 4) - 3
		k := 0
		v.Frame = true
		v.Title = "data"
		v.Wrap = false
		for i := 0; i < numLines; i++ {
			if i < 10 {
				fmt.Fprintf(v, "%d     ", i * lineSize)
			} else if i > 10 && i < 100 {
				fmt.Fprintf(v, "%d    ", i * lineSize)
			} else if i > 100 && i < 1000 {
				fmt.Fprintf(v, "%d   ", i * lineSize)
			} else if i > 1000 && i < 10000 {
				fmt.Fprintf(v, "%d   ", i * lineSize)
			}
			for j := 0; j < lineSize; j++ {
				if k < numcells {
					fmt.Fprintf(v, "0x0 ")
				} else {
					fmt.Fprintf(v, "0x. ")
				}
				k++
			}
			fmt.Fprintf(v, "\n")
		}
		// WRITE CELLS
	}
	if v, err := g.SetView("footer", 0, maxY - 3, maxX - 1, maxY - 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "e - edit \t^c - quit\t^g - run\tesc - options\n")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("inst", gocui.KeyCtrlG, gocui.ModNone, run); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}
	return nil
}

func run(g *gocui.Gui, v *gocui.View) error {
	p := program.New(v.Buffer(), numcells)
	buf := bytes.NewBufferString("")
	ret := p.Run(os.Stdin, buf)
	fmt.Printf("%s", string(ret))
	if dataV, err := g.View("data"); err == nil {
		dataV.Clear()
		for i := 0; i < numcells; i++ {
			fmt.Fprintf(dataV, "0x%X ", byte(string(ret)[i]))
		}
	} else {
		return err
	}

	return nil
}
