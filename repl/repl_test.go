package repl

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	h := []byte("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
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
	Start(tmpfile, &b)

	t.Logf("GOT %s", b.String())

	if err := tmpfile.Close(); err != nil {
		t.Fatalf("close tmp file failed: %v", err)
	}
}
