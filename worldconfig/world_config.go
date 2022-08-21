package worldconfig

type WorldConfig struct {
	Camera
}

type Camera struct {
	Width              int64
	Height             int64
	Fov                float64
	ViewTransformation `yaml:"view_transformation"`
}

type ViewTransformation struct {
	From []float64
	To   []float64
	Up   []float64
}
