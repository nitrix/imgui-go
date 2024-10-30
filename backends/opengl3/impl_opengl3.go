package opengl3

// #cgo CXXFLAGS: -std=c++17
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #define CIMGUI_USE_OPENGL3
// #include "../../dist/cimgui/cimgui.h"
// #include "../../dist/cimgui/cimgui_impl.h"
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
