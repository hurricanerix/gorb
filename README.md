Go Red Book
===========

Examples from
[OpenGL Programming Guide,  Eighth Edition](http://www.amazon.com/OpenGL-Programming-Guide-Official-Learning/dp/0321773039/)
ported to the [Go Programming Language](https://golang.org/).

In addition to being ported from C++ to Go, the following additional changes
have been made:

* code initializing OpenGL/creating windows, has been moved to the [utils package](utils/README.md) in
order to reduce duplicating code that does not change between examples.
* where appropriate the examples have been modified to have
a more consistent flow from one example to the next.  This involves some of the following:
  * using the same style in all shaders (not usually a problem with Go).
  * consistent naming of variables across all examples.  If it is called "MCvertex" in one
  example, it should be called "MCvertex" in all examples.
* any modifications specific to the examples, should be documented in the "Notes" section
of the example.

#### Chapter 1: Introduction to OpenGL
  * [Triangles](./01/triangles/README.md)

#### Chapter 3: Drawing with OpenGL
  * [Draw Commands](./03/drawcommands/README.md)
  * [Primitive Restart](./03/primitive-restart/README.md)

#### Chapter 4: Color, Pixels, and Framebuffers
  * [Gouraud](./04/gouraud/README.md)

# Running Examples

First see "Installing Examples" if you have not done so already.

```
$ cd $GOPATH/src/github.com/hurricanerix/gorb
$ make
mkdir -p bin
go build -o bin/ch01-triangles 01/triangles/main.go
go build -o bin/ch03-drawcommands 03/drawcommands/main.go
...
$ ./ch01-triagnles
...
```


# Installing Examples

1. [Install Go](https://golang.org/doc/install)

2. Platform specific stuff (see sections below)

3. go get

From your terminal run the following command:

```
$ go get github.com/hurricanerix/gorb/...
```

## Linux

```
$ sudo apt-get install git-core libgl1-mesa-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev
```

## OSX

```
$ brew install glfw3
```

### Windows

TODO: Figure this out...
