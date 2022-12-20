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

#scale: ["scale", number, number, number]
#translate: ["translate", number, number, number]
#rotateX: ["rotateX", number]
#rotateY: ["rotateY", number]
#rotateZ: ["rotateZ", number]

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
  sphere: {
    transform: #transform
    material?: #material
  }
}

#Cube: {
  cube: {
    transform: #transform
    material?: #material 
  }
}

#Plane: {
  plane: {
    transform: #transform
    material?: #material 
  }
}

#Cylinder: {
  cylinder: {
    transform: #transform
    minimum: number
    maximum: number
    closed: bool
    material?: #material
    
  }
}

#Model: {
  model: {
    path: string
    transform?: #transform
  }
}

#Shapes: {
[...#Sphere | #Cube | #Plane | #Cylinder | #Model]
}

camera: #Camera
light: #Light
shapes: #Shapes
