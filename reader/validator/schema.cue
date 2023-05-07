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
  values: [number, number, number]
}
#translate: {
  type: "translate"
  values: [number, number, number]
}
#rotateX: {
  type: "rotate-x"
  values: number
}
#rotateY: {
  type: "rotate-y"
  values: number
}
#rotateZ: {
  type: "rotate-z"
  values: number
}

#transform: [...#scale | #translate | #rotateX | #rotateY | #rotateZ]

#material: {
    ambient: number
    diffuse: number
    specular: number
    shininnes: number
    reflective: number
    transparency: number
    refractiveIndex: number
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