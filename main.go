package main

import (
	"fmt"

	"brainfuck/program"
)

func usage() {
	fmt.Println("go run main.go [brainfuck source file]")
}

func main() {
	//	flag.Parse()
	//	if len(flag.Args()) == 0 {
	//		usage()
	//		os.Exit(1)
	//	}
	//
	//	fp, err := os.Open(flag.Arg(0))
	//	if err != nil {
	//		usage()
	//		os.Exit(1)
	//	}
	//
	//	reader := bufio.NewReader(fp)
	//	b, e := reader.ReadByte()
	//	for e == nil {
	//		fmt.Println(b)
	//	}
	//
	//	fmt.Println("Done!")

	p, err := program.New("+++[>++]", 8)
	if err != nil {
		return
	}
	p.Run()
}
