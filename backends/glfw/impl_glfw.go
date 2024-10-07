package glfw

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -DCIMGUI_USE_GLFW -I ../../dist/include
// #cgo windows LDFLAGS: -L../../dist/windows
// #cgo linux LDFLAGS: -L../../dist/linux
// #cgo windows LDFLAGS: -lcimgui -limm32 -luser32 -lkernel32 -lgdi32 -lshell32 -static -lstdc++
// #cgo linux LDFLAGS: -lcimgui -lm -lc++
// #cgo darwin,amd64 LDFLAGS: -L../../dist/macos/amd64
// #cgo darwin,arm64 LDFLAGS: -L../../dist/macos/arm64
// #cgo darwin LDFLAGS: -lcimgui -framework CoreFoundation -lc++ -framework OpenGL -framework Cocoa -framework IOKit -framework QuartzCore
// #include "cimgui/cimgui.h"
// #include "cimgui/cimgui_impl.h"
import "C"

import (
	"github.com/nitrix/glfw-go"
)

func Init(window *glfw.Window) {
	C.ImGui_ImplGlfw_InitForOpenGL((*C.GLFWwindow)(window.Handle()), C.bool(true))
}

func NewFrame() {
	C.ImGui_ImplGlfw_NewFrame()
}

func Shutdown() {
	C.ImGui_ImplGlfw_Shutdown()
}
