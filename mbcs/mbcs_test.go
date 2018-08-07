package mbcs

import (
	"bytes"
	"testing"
)

func TestUtoAc(t *testing.T) {
	ansi, err := UtoA("AAAA", ACP, true)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	if !bytes.Equal(ansi, []byte("AAAA")) {
		t.Fatal("UtoAc failed")
		return
	}
}

func TestUtoAz(t *testing.T) {
	ansiz, err := UtoA("BBBB", ACP, false)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	if !bytes.Equal(ansiz, []byte("BBBB\000")) {
		t.Fatal("UtoAz failed")
	}
}

func TestUtoA(t *testing.T) {
	ansiz, err := UtoA("CCCC", ACP, false)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	if !bytes.Equal(ansiz, []byte("CCCC\000")) {
		t.Fatal("UtoA failed")
	}
}

const Japanese = "あいうえお"

func TestAtoU(t *testing.T) {
	ansi, err := UtoA(Japanese, ACP, true)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	utf8, err := AtoU(ansi, ACP)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	if utf8 != Japanese {
		t.Fatal("AtoU(UtoA) failed")
		return
	}
}
