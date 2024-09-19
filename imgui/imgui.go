package imgui

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -DCIMGUI_USE_GLFW -DCIMGUI_USE_OPENGL3
// #cgo LDFLAGS: -L../build -lcimgui -static -lc++ -lc++abi
// #include "../thirdparty/cimgui/cimgui.h"
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
