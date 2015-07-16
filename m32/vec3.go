package m32

import (
	"math"
)

type Vec3 [3]float32

var (
	zero3 Vec3
)

func NewVec3(x, y, z float32) Vec3 {
	return Vec3{x, y, z}
}

// Magnitude
func (v Vec3) Len() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
}

func (v Vec3) Normalize() Vec3 {
	if v == zero3 {
		return v
	}
	return v.Scale(1 / v.Len())
}

func (v Vec3) Negate() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

func (v Vec3) Scale(n float32) Vec3 {
	return Vec3{v[0] * n, v[1] * n, v[2] * n}
}

func (v Vec3) Add(v3 Vec3) Vec3 {
	return Vec3{v[0] + v3[0], v[1] + v3[1], v[2] + v3[2]}
}

func (v Vec3) Sub(v3 Vec3) Vec3 {
	return Vec3{v[0] - v3[0], v[1] - v3[1], v[2] - v3[2]}
}

func (v Vec3) Distance(v3 Vec3) float32 {
	return v3.Sub(v).Len()
}

func (v Vec3) Dot(v3 Vec3) float32 {
	return v[0]*v3[0] + v[1]*v3[1] + v[2]*v3[2]
}

func (v Vec3) Cross(v3 Vec3) Vec3 {
	return Vec3{
		v[1]*v3[2] - v[2]*v3[1],
		v[2]*v3[0] - v[0]*v3[2],
		v[0]*v3[1] - v[1]*v3[0],
	}
}

func (v Vec3) Angle(v3 Vec3) float32 {
	if v == zero3 || v3 == zero3 {
		return 0
	}
	return float32(math.Acos(float64(v.Dot(v3) / (v.Len() * v3.Len()))))
}
