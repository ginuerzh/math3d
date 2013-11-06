// math3d_test.go
package math3d

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	m := Translate(1, 2, 3)
	t.Log(m)
}
