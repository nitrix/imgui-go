package opengl3

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -DCIMGUI_USE_OPENGL3 -I../../dist/include
// #cgo windows LDFLAGS: -L../../dist/windows
// #cgo linux LDFLAGS: -L../../dist/linux
// #cgo windows LDFLAGS: -lcimgui -limm32 -luser32 -lkernel32 -lgdi32 -lshell32 -static -lstdc++
// #cgo linux LDFLAGS: -lcimgui -lm -lc++
// #cgo darwin,amd64 LDFLAGS: -L../../dist/macos/amd64
// #cgo darwin,arm64 LDFLAGS: -L../../dist/macos/arm64
// #cgo darwin LDFLAGS: -lcimgui -framework CoreFoundation -lc++ -framework OpenGL -framework Cocoa -framework IOKit -framework QuartzCore
// #include "cimgui/cimgui.h"
// #include "cimgui/cimgui_impl.h"
// #include <stdlib.h>
import "C"

import (
	"unsafe"

	"github.com/nitrix/glfw-go"
	"github.com/nitrix/imgui-go"
)

func Init(window *glfw.Window) {
	s := C.CString("#version 330 core")
	C.ImGui_ImplOpenGL3_Init(s)
	C.free(unsafe.Pointer(s))
}

func NewFrame() {
	C.ImGui_ImplOpenGL3_NewFrame()
}

func Shutdown() {
	C.ImGui_ImplOpenGL3_Shutdown()
}

func RenderDrawData(d *imgui.DrawData) {
	wd := (*C.ImDrawData)(unsafe.Pointer(d))
	C.ImGui_ImplOpenGL3_RenderDrawData(wd)
}
