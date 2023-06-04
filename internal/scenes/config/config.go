package config

type Scene struct {
	Camera  Camera
	Light   Light
	Objects []Object
}

type Camera struct {
	Width  int64
	Height int64
	Fov    float64
	From   []float64
	To     []float64
	Up     []float64
}

type Light struct {
	Position  []float64
	Intensity []float64
}

type Object struct {
	Type             string
	Transform        []Transform
	Material         Material
	Minimum, Maximum float64
	Closed           bool
	File             string
	Children         []Object
}

type Transform struct {
	Type   string
	Values []float64
}

type Material struct {
	Color                                                           []float64
	Pattern                                                         Pattern
	Ambient, Diffuse, Specular, Shininess, Reflective, Transparency float64
	RefractiveIndex                                                 float64 `yaml:"refractive_index"`
}

type Pattern struct {
	Type      string
	Colors    [][]float64
	Transform []Transform
}
