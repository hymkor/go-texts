package mbcs

import (
	"bytes"
	"testing"
)

func TestUtoAc(t *testing.T) {
	ansi, err := UtoAc("AAAA")
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
	ansiz, err := UtoAz("BBBB")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	if !bytes.Equal(ansiz, []byte("BBBB\000")) {
		t.Fatal("UtoAz failed")
	}
}

func TestUtoA(t *testing.T) {
	ansiz, err := UtoAz("CCCC")
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
	ansi, err := UtoAc(Japanese)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	utf8, err := AtoU(ansi)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	if utf8 != Japanese {
		t.Fatal("AtoU(UtoA) failed")
		return
	}
}
