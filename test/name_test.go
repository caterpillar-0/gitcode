package name

import (
	"testing"
)

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Jerry"
	if output != expectOutput {
		t.Errorf("expect : %s actual : %s", expectOutput, output)
	}
}
