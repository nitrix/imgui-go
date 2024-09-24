# imgui-go

A binding for imgui in Go.

## Example

![example.png](example.png)

Try the example with `go run github.com/nitrix/imgui-go/example@latest`.

The sources for it are located here [example/example.go](example/example.go).

## Progress

The bindings are written manually on an "as-needed" basis. 
The GLFW API is kept intentionally compatible with `github.com/go-gl/glfw`.

## Important

You're currently limited to using the OpenGL and GLFW backends. The GLFW window must be created with this library
because the backends rely on recently added enums that not all other popular bindings support.

## License

This is free and unencumbered software released into the public domain. See the [UNLICENSE](UNLICENSE) file for more details.
