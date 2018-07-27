package program

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	input := "><[]ii+-"
	p := New(input)
	t.Logf("%q", p.tok)
	if p == nil {
		t.Fatalf("New failed")
	}
}

func TestParse(t *testing.T) {
	input := "+++"

	p := New(input)
	p.Run(os.Stdin, os.Stdout)

	t.Logf("%q", p.tok)
	t.Logf("cells[0] = %d", p.cells[0])

	if p.cells[0] != 3 {
		t.Logf("Cells:\n%q", p.cells)
		t.Fatalf("Parse failed")
	}
}

func TestParses(t *testing.T) {
	input := "+++>++>+>>>>>"

	tests := []struct {
		expectedCells byte
	}{
		{3},
		{2},
		{1},
	}

	p := New(input)
	p.Run(os.Stdin, os.Stdout)

	t.Logf("%q", p.tok)
	for i, tt := range tests {
		t.Logf("cells[%d] = %d\n", i, p.cells[i])
		if p.cells[i] != tt.expectedCells {
			t.Logf("Cells:\n%q", p.cells)
			t.Fatalf("parse failed: expected %d, got %d", tt.expectedCells, p.cells[i])
		}
	}
}

func TestInput(t *testing.T) {
	h := []byte("h")
	tmpfile, err := ioutil.TempFile("", "gofcktest")
	if err != nil {
		t.Fatalf("create tempfile failed: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(h); err != nil {
		t.Fatalf("write to temp file failed: %v", err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatalf("seek failed: %v", err)
	}

	input := "+>++>+++>,"

	tests := []struct {
		exp byte
	}{
		{1},
		{2},
		{3},
		{104},
	}

	p := New(input)
	p.Run(tmpfile, os.Stdout)

	t.Logf("%q", p.tok)
	for i, tt := range tests {
		t.Logf("cells[%d] = %d\n", i, p.cells[i])
		if p.cells[i] != tt.exp {
			t.Logf("Cells:\n%q", p.cells)
			t.Fatalf("parse failed: expected %d, got %d", tt.exp, p.cells[i])
		}
	}

	if err := tmpfile.Close(); err != nil {
		t.Fatalf("close tmp file failed: %v", err)
	}
}

func TestOutput(t *testing.T) {
	h := []byte("o")
	tmpfile, err := ioutil.TempFile("", "gofcktest")
	if err != nil {
		t.Fatalf("create tempfile failed: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(h); err != nil {
		t.Fatalf("write to temp file failed: %v", err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatalf("seek failed: %v", err)
	}

	var b bytes.Buffer

	input := "+>++>+++>,."

	tests := []struct {
		exp byte
	}{
		{1},
		{2},
		{3},
		{'o'},
	}

	p := New(input)
	p.Run(tmpfile, &b)

	t.Logf("%q", p.tok)
	for i, tt := range tests {
		t.Logf("cells[%d] = %d\n", i, p.cells[i])
		if p.cells[i] != tt.exp {
			t.Logf("Cells:\n%q", p.cells)
			t.Fatalf("parse failed: expected %d, got %d", tt.exp, p.cells[i])
		}
	}

	print, err := b.ReadByte()
	if err != nil {
		t.Fatalf("readbyte error: %v", err)
	}
	t.Logf("GOT %c", print)

	if err := tmpfile.Close(); err != nil {
		t.Fatalf("close tmp file failed: %v", err)
	}
}

func TestLoop(t *testing.T) {
	input := "+>++>+++>[++]++++"

	p := New(input)
	p.Run(os.Stdin, os.Stdout)

	tests := []struct {
		exp byte
	}{
		{1},
		{2},
		{3},
		{4},
	}

	t.Logf("%q", p.tok)
	for i, tt := range tests {
		t.Logf("cells[%d] = %d\n", i, p.cells[i])
		if p.cells[i] != tt.exp {
			t.Logf("Cells:\n%q", p.cells)
			t.Fatalf("parse failed: expected %d, got %d", tt.exp, p.cells[i])
		}
	}

	input = "+[++]>++>+"

	p = New(input)
	p.Run(os.Stdin, os.Stdout)

	t.Logf("%q", p.tok)
	tests = []struct {
		exp byte
	}{
		{3},
		{2},
		{1},
	}

	for i, tt := range tests {
		t.Logf("cells[%d] = %d\n", i, p.cells[i])
		if p.cells[i] != tt.exp {
			t.Logf("Cells:\n%q", p.cells)
			t.Fatalf("parse failed: expected %d, got %d", tt.exp, p.cells[i])
		}
	}
}

func TestNestedLoops(t *testing.T) {
	tests := []struct {
		input string
		exp   byte
	}{
		{"+[+[-]+]", 2},
		{"++[+[-]]", 2},
		{"++[[-]+]", 2},
		{"[]", 0},
	}

	for _, tt := range tests {
		p := New(tt.input)
		p.Run(os.Stdin, os.Stdout)
		t.Logf("%q", p.tok)

		if p.cells[0] != tt.exp {
			t.Logf("Cells:\n%q", p.cells)
			t.Fatalf("parse failed: expected %d, got %d", tt.exp, p.cells[0])
		}
	}
}
