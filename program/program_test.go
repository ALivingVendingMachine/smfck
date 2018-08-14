package program

import (
	"testing"
)

func TestNewProgram(t *testing.T) {
	r, err := New("+-><[],.", 16)

	if err != nil {
		t.Fatalf("err should not be nil, err is %v", err)
	}
	if r.commands == "" {
		t.Fatal("commands should not be empty")
	}
}

func TestNewProgramError(t *testing.T) {
	r, err := New("+-><[],.", 0)

	if err == nil {
		t.Fatal("err should not have been nil")
	}
	if r.commands != "" {
		t.Fatalf("commands should have been empty, commands are %s", r.commands)
	}
	r, err = New("+-><[],.", -10)

	if err == nil {
		t.Fatal("err should not have been nil")
	}
	if r.commands != "" {
		t.Fatalf("commands should have been empty, commands are %s", r.commands)
	}
}

func TestRunProgram(t *testing.T) {
	var tests = []struct {
		commands string // program commands
		numcells int    // program cell count
		exp      []byte // expected output
	}{
		{"+++", 1, []byte{3}},              // 0
		{"+>++", 2, []byte{1, 2}},          // 1
		{"+>++>+++", 3, []byte{1, 2, 3}},   // 2
		{"+[>++<-]", 3, []byte{0, 2, 0}},   // 3
		{"[+>++>+++]", 3, []byte{0, 0, 0}}, // 4
	}

	for i, tt := range tests {
		prog, err := New(tt.commands, tt.numcells)
		if err != nil {
			t.Fatalf("err is not nil, err is %v", err)
		}
		prog.Run()
		for j := range tt.exp {
			if prog.cells[j] != tt.exp[j] {
				t.Logf("Cells: %q", prog.cells)
				t.Fatalf("test %d failed: expected %d, got %d", i, tt.exp[j], prog.cells[j])
			}
		}
	}
}

func TestHelloWorld(t *testing.T) {
	exp := "Hello World!"
	prog, err := New("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.", 256)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	ret, err := prog.Run()
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	t.Logf("ret: %s\n", ret)
	for i := range ret {
		if ret[i] != exp[i] {
			t.Fatalf("fail: expected %c, got %c", exp[i], ret[i])
		}
	}
}
