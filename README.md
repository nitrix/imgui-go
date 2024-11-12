# imgui-go

A binding library of imgui for Go.

## Example

![example-1.91.4.png](example-1.91.4.png)

Try the example with `go run github.com/nitrix/imgui-go/example@latest`.

The sources for it are located here [example/example.go](example/example.go).

## Compatibility

You're currently limited to using the OpenGL and GLFW backends. The GLFW window must be created by [this other library](https://github.com/nitrix/glfw-go)
because the backends rely on recently added enums that not all the other popular bindings support.

## Credits

See [this repo](https://github.com/cimgui/cimgui) for the bindings C library and [this repo](https://github.com/ocornut/imgui) for the original C++ library.

## License

This is free and unencumbered software released into the public domain. See the [UNLICENSE](UNLICENSE) file for more details.
