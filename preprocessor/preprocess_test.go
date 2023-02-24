package preprocessor_test

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hymkor/go-texts/preprocessor"
	"github.com/tidwall/transform"
)

func lnum() func() ([]byte, error) {
	fd, err := os.Open("preprocess.go")
	if err != nil {
		panic(err.Error())
	}
	br := bufio.NewReader(fd)
	count := 0
	return func() ([]byte, error) {
		line, err := br.ReadString('\n')
		if err != nil {
			fd.Close()
			return []byte{}, err
		}
		count++
		return []byte(fmt.Sprintf("%d: %s", count, line)), nil
	}
}

func Benchmark_filter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sc := bufio.NewScanner(preprocessor.New(lnum()))
		for sc.Scan() {
			fmt.Fprintln(ioutil.Discard, sc.Text())
		}
		if err := sc.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func Benchmark_transformer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sc := bufio.NewScanner(transform.NewTransformer(lnum()))
		for sc.Scan() {
			fmt.Fprintln(ioutil.Discard, sc.Text())
		}
		if err := sc.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func ioPiper(f func() ([]byte, error)) io.Reader {
	in, out := io.Pipe()
	go func() {
		for {
			data, err := f()
			out.Write(data)
			if err != nil {
				out.CloseWithError(err)
				return
			}
		}
	}()
	return in
}

func Benchmark_iopipe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sc := bufio.NewScanner(ioPiper(lnum()))
		for sc.Scan() {
			fmt.Fprintln(ioutil.Discard, sc.Text())
		}
		if err := sc.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// for example, run `go test -bench . -benchmem`
