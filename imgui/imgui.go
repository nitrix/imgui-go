package imgui

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -I../dist/include
// #cgo windows LDFLAGS: -L../dist/windows
// #cgo linux LDFLAGS: -L../dist/linux
// #cgo windows LDFLAGS: -lcimgui -limm32 -static -lc++ -lc++abi
// #cgo linux LDFLAGS: -lcimgui -lm -lc++
// #cgo darwin,amd64 LDFLAGS: -L../dist/macos/amd64
// #cgo darwin,arm64 LDFLAGS: -L../dist/macos/arm64
// #cgo darwin LDFLAGS: -lcimgui -framework CoreFoundation -lc++ -lc++abi -framework OpenGL -framework Cocoa -framework IOKit -framework QuartzCore
// #include "cimgui/cimgui.h"
import "C"

import (
	"fmt"
)

type Context C.ImGuiContext
type DrawData C.ImDrawData

func CreateContext() (*Context, error) {
	ctx := C.igCreateContext(nil)
	if ctx == nil {
		return nil, fmt.Errorf("failed to create context")
	}
	return (*Context)(ctx), nil
}

func (ctx *Context) Destroy() {
	C.igDestroyContext((*C.ImGuiContext)(ctx))
}

func NewFrame() {
	C.igNewFrame()
}

func ShowDemoWindow() {
	C.igShowDemoWindow(nil)
}

func GetDrawData() *DrawData {
	return (*DrawData)(C.igGetDrawData())
}

func Render() {
	C.igRender()
}
