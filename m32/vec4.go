package m32

import (
	"math"
)

type Vec4 [4]float32

var (
	zero4 Vec4
)

func NewVec4(x, y, z, w float32) Vec4 {
	return Vec4{x, y, z, w}
}

// Magnitude
func (v Vec4) Len() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])))
}

func (v Vec4) Normalize() Vec4 {
	if v == zero4 {
		return v
	}
	return v.Scale(1 / v.Len())
}

func (v Vec4) Negate() Vec4 {
	return Vec4{-v[0], -v[1], -v[2], -v[3]}
}

func (v Vec4) Scale(n float32) Vec4 {
	return Vec4{v[0] * n, v[1] * n, v[2] * n, v[3] * n}
}

func (v Vec4) Add(v4 Vec4) Vec4 {
	return Vec4{v[0] + v4[0], v[1] + v4[1], v[2] + v4[2], v[3] + v4[3]}
}

func (v Vec4) Sub(v4 Vec4) Vec4 {
	return Vec4{v[0] - v4[0], v[1] - v4[1], v[2] - v4[2], v[3] - v4[3]}
}

func (v Vec4) Distance(v4 Vec4) float32 {
	return v4.Sub(v).Len()
}

func (v Vec4) Dot(v4 Vec4) float32 {
	return v[0]*v4[0] + v[1]*v4[1] + v[2]*v4[2] + v[3]*v4[3]
}
