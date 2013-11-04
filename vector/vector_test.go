// vector_test.go
package vector

import (
	"testing"
)

const (
	dim = 4
)

var (
	v1 Vector = NewVector(dim, 1.0, 0.0, 0.0)
	v2 Vector = NewVector(dim, 0.0, 0.0, -1.0)
)

func TestAdd(t *testing.T) {
	r, err := Add(v1, v2)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r)
}

func TestSub(t *testing.T) {
	r, err := Sub(v1, v2)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r)
}

func TestMulti(t *testing.T) {
	r := Multi(v1, 3)
	t.Log(r)
}

func TestDiv(t *testing.T) {
	r := Div(v1, 3)
	t.Log(r)
}

func TestDot(t *testing.T) {
	r, err := Dot(v1, v2)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r)
}

func TestCross(t *testing.T) {
	r, err := Cross(v1, v2)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r)

	r, _ = Cross(v2, v1)
	t.Log(r)
}

func TestDistance(t *testing.T) {
	r, err := Distance(v1, v2)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r)
}

func TestNormalize(t *testing.T) {
	r := Normalize(v1)
	t.Log(r)
	r = Normalize(v2)
	t.Log(r)
}

func TestAngle(t *testing.T) {
	r, err := Angle(v1, v2)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r)
}

func TestReflect(t *testing.T) {
	r, err := Reflect(v1, NewVector(dim, 0, 0, 1.0))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r)
}

func TestParallel(t *testing.T) {
	r, err := v1.Parallel(v2)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r)
}

func TestPerp(t *testing.T) {
	r, err := v1.Perp(v2)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r)
}
