package glfw

// #cgo CPPFLAGS: -I../../dist
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #define CIMGUI_USE_GLFW
// #include "../../dist/cimgui.h"
// #include "../../dist/cimgui_impl.h"
import "C"

import (
	"github.com/nitrix/glfw-go"
	_ "github.com/nitrix/imgui-go/dist"
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
