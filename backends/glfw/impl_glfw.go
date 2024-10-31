package glfw

// #cgo CXXFLAGS: -std=c++17 -O3
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #define CIMGUI_USE_GLFW
// #include "../../dist/cimgui/cimgui.h"
// #include "../../dist/cimgui/cimgui_impl.h"
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
