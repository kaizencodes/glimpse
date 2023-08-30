# Glimpse

![](examples/marbles.png)

Glimpse is a ray tracer written in Go. It is an implementation of the [Ray tracer challenge](https://pragprog.com/titles/jbtracer/the-ray-tracer-challenge/) by Jamis Buck.
This is a personal project to learn Go and ray tracing. It is not meant to be used in production.

It is a work in progress, I will be adding more features such as textures, soft shadows, and anti-aliasing.

## Features

It can render primitives such as 
- spheres 
- planes 
- cubes
- cylinders. 

It can render shadows, reflections, and refractions. It can render a scene with multiple light sources.

## Installation

[Install go](https://go.dev/dl/)

Install dependencies, [cue](https://cuelang.org) which is used for yaml parsing.

```
go mod download
go get -u golang.org/x/lint/golint
go install cuelang.org/go/cmd/cue@latest
```

## Usage

After compilation run the executable like so `./glimpse -f example.yml`. The `-f` flag is used to specify the input file. The input file describes the scene to be rendered.

The format is the following:

```
# describe the camera
camera:
  width: 800                    # width of the image
  height: 400                   # height of the image
  fov: 0.785                    # field of view in radians
  from: [8, 6, -8]              # camera position, x,y,z coordinates
  to: [0, 3, 0]                 # camera target (where it is looking at)
  up: [0, 1, 0]                 # camera up vector

# describe the light source
lights:
  - position: [0, 6.9, -5]        # light position, x,y,z coordinates
    intensity: [1, 1, 0.9]        # light intensity, r,g,b values between 0 and 1

# describe the objects in the scene
objects:
  - type: sphere                # type of object
    # By default the object is rendered in the center of the scene.
    # You can transform the object by applying a series of transformations.
    transform:  
      - type: "translate"       # translates the object in the 3d space
        values: [0, 1, 0]
      - type: "scale"           # scales the object in the 3d space
        values: [20, 7, 20]
```

You can see complete scenes in the [examples](examples) directory.

-o flag is used to specify the output file. The output file is a pmm image. The default output folder is [renders](renders).
 

## With Docker

Alternatively you can use Docker. In my tests it was 2x slower than running it natively.

You can build the image with `docker build -t glimpse .` 

For a quick render you can run `docker run -it --mount type=bind,source="$(pwd)"/,target=/glimpse glimpse:latest ./glimpse -f examples/marbles.yml -o renders/marbles`

For ongoing development however it is preferable to maintain a container. There is a docker compose that builds the image and starts a container that run `sleeps infinity`.

Start the container with `docker-compose up -d`
Then you can run commands in the container with `docker exec -it glimpse bash`

The source code is mounted so this can be used to recompile the code after changes `go build .`, to run the tests `go test -gcflags=-l -v ./...`, and to run render which will be saved in the mounted host directory.
