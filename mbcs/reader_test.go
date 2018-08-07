package mbcs

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func TestReader2(t *testing.T) {
	fd, err := os.Open("japan.txt")
	if err != nil {
		t.Fatalf("%s\n", err.Error())
		return
	}
	defer fd.Close()

	reader := NewReader(fd)
	sc := bufio.NewScanner(reader)
	buffer := make([]string, 0, 3)
	for sc.Scan() {
		buffer = append(buffer, sc.Text())
	}
	if err := sc.Err(); err != nil && err != io.EOF {
		println(err.Error())
	}
	if len(buffer) != 3 {
		t.Fatalf("lines too few\n")
		return
	}
	if buffer[0] != "日本語サンプル" {
		t.Fatalf("line[0]=\"%s\"\n", buffer[0])
	}
	println(buffer[0])
	println(buffer[1])
	println(buffer[2])
}
