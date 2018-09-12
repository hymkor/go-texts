package dos

import (
	"testing"
)

func TestWithoutExt(t *testing.T) {
	const source = `c:\foo\bar.hoge\ahaha.txt`
	result := WithoutExt(source)
	if result != `c:\foo\bar.hoge\ahaha` {
		t.Fatalf("`%s` != `%s`\n", source, result)
	}
	const source2 = `c:\foo\bar.hoge\ahaha`
	result = WithoutExt(source)
	if result != `c:\foo\bar.hoge\ahaha` {
		t.Fatalf("`%s` != `%s`\n", source2, result)
	}
}
