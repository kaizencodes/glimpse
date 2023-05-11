#Tuple: 3 * [number]

#Camera: {
	width: number
	height: number
	fov: number
	from: #Tuple
  to: #Tuple
  up: #Tuple
}
#Light: {
  position: #Tuple
  intensity: #Tuple
}

#scale: {
  type: "scale" 
  values: #Tuple
}
#translate: {
  type: "translate"
  values: #Tuple
}
#rotateX: {
  type: "rotate-x"
  values: [number]
}
#rotateY: {
  type: "rotate-y"
  values: [number]
}
#rotateZ: {
  type: "rotate-z"
  values: [number]
}

#transform: [...#scale | #translate | #rotateX | #rotateY | #rotateZ]

#pattern: {
  type: string
  colors: [...#Tuple]
}

#material: {
  color?: #Tuple
  pattern?: #pattern
  ambient: number
  diffuse: number
  specular: number
  shininess: number
  reflective: number
  transparency: number
  refractive_index: number
}

#Sphere: {
  type: string & "sphere"
  transform: #transform
  material?: #material
}

#Cube: {
  type: "cube"
  transform: #transform
  material?: #material 
}

#Plane: {
  type: "plane"
  transform: #transform
  material?: #material 
}

#Cylinder: {
  type: "cylinder"
  transform: #transform
  minimum: number
  maximum: number
  closed: bool
  material?: #material
}

#Model: {
  type: "model"
  path: string
  transform?: #transform
}

#Object: {
[...#Sphere | #Cube | #Plane | #Cylinder | #Model]
}

camera: #Camera
light: #Light
objects: #Object