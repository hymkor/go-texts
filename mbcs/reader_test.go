package mbcs

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
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

	reader := NewAutoDetectReader(fd, ACP)
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

func BenchmarkFilter1(b *testing.B) {
	data, err := ioutil.ReadFile("japan.txt")
	if err != nil {
		b.Fatalf("%s\n", err.Error())
		return
	}
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(data)
		ioutil.ReadAll(NewAtoUReader(r, ACP))
	}
}

func BenchmarkFilter2(b *testing.B) {
	data, err := ioutil.ReadFile("japan.txt")
	if err != nil {
		b.Fatalf("%s\n", err.Error())
		return
	}
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(data)
		ioutil.ReadAll(NewA2UReader(r, ACP))
	}
}
