package mbcs

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestA2UTransformer(t *testing.T) {
	const teststr = "あい\nうえお\nかきくけ\nこ"

	mbcs, err := UtoA(teststr, ACP, true)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	utf8, err := ioutil.ReadAll(NewA2UReader(bytes.NewReader(mbcs), ACP))
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if string(utf8) != teststr {
		t.Fatalf("[%s] != [%s]", utf8, teststr)
	}
}
