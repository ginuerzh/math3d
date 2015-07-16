package m32

import (
	"math"
)

const (
	RadPerDeg = math.Pi / 180.0
	DegPerRad = 180.0 / math.Pi
)

// returns a view matrix using the opengl lookAt style. COLUMN ORDER.
func LookAt(eye, target, up Vec3) Mat4 {
	// inverse translation
	p := Ident4()
	p = p.Translate(eye.Negate())
	// distance vector
	d := target.Sub(eye)
	// forward vector
	f := d.Normalize()
	// right vector
	r := f.Cross(up).Normalize()
	// real up vector
	u := r.Cross(f).Normalize()

	ori := Ident4()
	ori[0] = r[0]
	ori[4] = r[1]
	ori[8] = r[2]
	ori[1] = u[0]
	ori[5] = u[1]
	ori[9] = u[2]
	ori[2] = -f[0]
	ori[6] = -f[1]
	ori[10] = -f[2]

	return ori.Mul4(p)
}

// returns a perspective function mimicking the opengl projection style.
func Perspective(fovy, aspect float32, near, far float32) Mat4 {
	fovRad := float64(fovy) * RadPerDeg
	r := float32(math.Tan(fovRad/2.0)) * near
	sx := (2.0 * near) / (r*aspect + r*aspect)
	sy := near / r
	sz := -(far + near) / (far - near)
	pz := -(2.0 * far * near) / (far - near)
	// make sure bottom-right corner is zero
	m := ZeroMat4()
	m[0] = sx
	m[5] = sy
	m[10] = sz
	m[14] = pz
	m[11] = -1.0

	return m
}
