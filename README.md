# imgui-go

A binding for imgui in Go.

## Preview

![preview.png](preview.png)

## Status

Still very much a proof-of-concept. It's missing the GitHub actions + binaries for linux and mac.

## Usage

Try the example with `go run github.com/nitrix/imgui-go/example`.

The sources for it are here [example/example.go](example/example.go).

You're limited to using OpenGL and GLFW. The glfw has must be the bindings provided in this repo, `github.com/go-gl/glfw` isn't recent enough and will panic at runtime on unknown enums.

## Maintainers

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

## License

This is free and unencumbered software released into the public domain. See the [UNLICENSE](UNLICENSE) file for more details.