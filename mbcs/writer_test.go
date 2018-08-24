package mbcs

import (
	"fmt"
	"os"
	"testing"
)

func TestNewWriter(t *testing.T) {
	fd, err := os.Create("MBCSFILE.txt")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	defer fd.Close()
	w := NewWriter(fd, ACP)
	fmt.Fprintln(w, "あいうえお")
}
