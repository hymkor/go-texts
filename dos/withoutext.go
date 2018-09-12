package dos

import (
	"path/filepath"
)

func WithoutExt(fname string) string {
	ext := filepath.Ext(fname)
	return fname[:len(fname)-len(ext)]
}
