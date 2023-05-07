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
	Type      string
	Transform []Transform
	Material  Material
}

type Transform struct {
	Type   string
	Values []float64
}

type Material struct {
	Color                                                                            []float64
	Ambient, Diffuse, Specular, Shininess, Reflective, Transparency, RefractiveIndex float64
}