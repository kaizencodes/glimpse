package object

type Touple struct {
	x, y, z, w float64
}

func (t *Touple) IsPoint() bool {
	return t.w == 1.0
}

func (t *Touple) IsVector() bool {
	return t.w == 0.0
}

func NewVector(x, y, z float64) *Touple {
	return &Touple{x, y, z, 0.0}
}

func NewPoint(x, y, z float64) *Touple {
	return &Touple{x, y, z, 1.0}
}
