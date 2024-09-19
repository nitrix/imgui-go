# imgui-go

A no-binding approach to imgui in Go.

## Preview

![preview.png](preview.png)

## Generation

The binaries inside of `dist` are compiled using CMake with:

```
mkdir build
cd build
cmake ..
ninja
cd ..
```

then moved to their appropriate location with:

```
go generate
```

## Usage

See the [example/example.go](example/example.go) file for how to use the library.

You're limited to OpenGL and GLFW. The glfw bindings must be the ones provided in this repo, `github.com/go-gl/glfw` isn't recent enough and will panic at runtime on unknown enums.

## License

This is free and unencumbered software released into the public domain. See the [UNLICENSE](UNLICENSE) file for more details.