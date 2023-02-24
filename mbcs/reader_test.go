package mbcs

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	oldfilter "github.com/hymkor/go-texts/filter/old"
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

func OldAtoUReader(r io.Reader, cp uintptr) io.Reader {
	return oldfilter.New(r, func(line []byte) ([]byte, error) {
		text, err := AtoU(line, cp)
		return []byte(text), err
	})
}

func BenchmarkFilterOld(b *testing.B) {
	data, err := ioutil.ReadFile("reader.go")
	if err != nil {
		b.Fatalf("%s\n", err.Error())
		return
	}
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(data)
		ioutil.ReadAll(OldAtoUReader(r, ACP))
	}
}

func BenchmarkFilter(b *testing.B) {
	data, err := ioutil.ReadFile("reader.go")
	if err != nil {
		b.Fatalf("%s\n", err.Error())
		return
	}
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(data)
		ioutil.ReadAll(NewAtoUReader(r, ACP))
	}
}
