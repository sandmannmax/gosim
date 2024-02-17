package engine

type vector3 struct {
	x float64
	y float64
	z float64
}

func (v *vector3) mul(c float64) {
	v.x *= c
	v.y *= c
	v.z *= c
}

func (v1 *vector3) add(v2 vector3) {
	v1.x += v2.x
	v1.y += v2.y
	v1.z += v2.z
}

func (v1 *vector3) addScaled(v2 vector3, c float64) {
	v1.x += v2.x * c
	v1.y += v2.y * c
	v1.z += v2.z * c
}
