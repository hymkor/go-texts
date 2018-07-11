package mbcs

import (
	"bufio"
	"fmt"
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

	sc := bufio.NewScanner(toUpperReader)
	var text strings.Builder
	for sc.Scan() {
		fmt.Fprintln(&text, sc.Text())
	}
	if text.String() != to {
		t.Fatalf("NG: `%s` != `%s`\n", text.String(), to)
		return
	}
}
