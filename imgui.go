package imgui

//go:generate go run ./generator

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -Idist/include
// #cgo windows LDFLAGS: -Ldist/windows
// #cgo linux LDFLAGS: -Ldist/linux
// #cgo windows LDFLAGS: -lcimgui -limm32 -luser32 -lkernel32 -lgdi32 -lshell32 -static -lstdc++
// #cgo linux LDFLAGS: -lcimgui -lm -lc++
// #cgo darwin,amd64 LDFLAGS: -Ldist/macos/amd64
// #cgo darwin,arm64 LDFLAGS: -Ldist/macos/arm64
// #cgo darwin LDFLAGS: -lcimgui -framework CoreFoundation -lc++ -framework OpenGL -framework Cocoa -framework IOKit -framework QuartzCore
// #include "cimgui/cimgui.h"
// #include <stdlib.h>
import "C"

var stringPool StringPool

func NewFrame() {
	stringPool.Reset()
	C.igNewFrame()
}
