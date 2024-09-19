package imgui

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -DCIMGUI_USE_GLFW -I ../dist/include
// #cgo windows LDFLAGS: -L../dist/windows
// #cgo LDFLAGS: -lcimgui -static -lc++ -lc++abi
// #include "cimgui/cimgui.h"
// #include "cimgui/cimgui_impl.h"
import "C"

import (
	"github.com/nitrix/imgui-go/glfw"
)

func ImplGlfw_Init(window *glfw.Window) {
	C.ImGui_ImplGlfw_InitForOpenGL((*C.GLFWwindow)(window.Handle()), C.bool(true))
}

func ImplGlfw_NewFrame() {
	C.ImGui_ImplGlfw_NewFrame()
}

func ImplGlfw_Shutdown() {
	C.ImGui_ImplGlfw_Shutdown()
}
