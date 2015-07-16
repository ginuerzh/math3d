package m32

import (
	"math"
)

type Vec2 [2]float32

var (
	zero2 Vec2
)

func NewVec2(x, y float32) Vec2 {
	return Vec2{x, y}
}

// Length or magnitude
func (v Vec2) Len() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1])))
}

func (v Vec2) Normalize() Vec2 {
	if v == zero2 {
		return v
	}
	return v.Scale(1 / v.Len())
}

func (v Vec2) Negate() Vec2 {
	return Vec2{-v[0], -v[1]}
}

func (v Vec2) Scale(n float32) Vec2 {
	return Vec2{v[0] * n, v[1] * n}
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v[0] + v2[0], v[1] + v2[1]}
}

func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v[0] - v2[0], v[1] - v2[1]}
}

func (v Vec2) Distance(v2 Vec2) float32 {
	return v2.Sub(v).Len()
}

func (v Vec2) Dot(v2 Vec2) float32 {
	return v[0]*v2[0] + v[1]*v2[1]
}

func (v Vec2) Angle(v2 Vec2) float32 {
	if v == zero2 || v2 == zero2 {
		return 0
	}
	return float32(math.Acos(float64(v.Dot(v2) / (v.Len() * v2.Len()))))
}
