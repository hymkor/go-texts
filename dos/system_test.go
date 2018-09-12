package dos

import (
	"testing"
)

func TestSystem(t *testing.T) {
	err := System("echo \"ahaha ihihi\"")
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}
}
