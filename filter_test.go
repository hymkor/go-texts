package mbcs

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestNewFilter(t *testing.T) {
	from := `abcdefg
hijklmn
opqrstu
vwxyz
`

	to := `ABCDEFG
HIJKLMN
OPQRSTU
VWXYZ
`

	toUpperReader := NewFilter(strings.NewReader(from),
		func(data []byte) ([]byte, error) {
			return []byte(strings.ToUpper(string(data))), nil
		},
	)

	bytes, err := ioutil.ReadAll(toUpperReader)
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	text := string(bytes)
	if text != to {
		t.Fatalf("NG: `%s` != `%s`\n", text, to)
		return
	}
}
