package imgui

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -DCIMGUI_USE_GLFW -DCIMGUI_USE_OPENGL3
// #cgo LDFLAGS: -L../build -lcimgui -static -lc++ -lc++abi
// #include "../thirdparty/cimgui/cimgui.h"
// #include "../thirdparty/cimgui/generator/output/cimgui_impl.h"
import "C"

import (
	"github.com/nitrix/imgui-go/glfw"
)

func ImplOpenGL3_Init(window *glfw.Window) {
	C.ImGui_ImplOpenGL3_Init(C.CString("#version 330 core"))
}

func ImplOpenGL3_NewFrame() {
	C.ImGui_ImplOpenGL3_NewFrame()
}

func ImplOpenGL3_Shutdown() {
	C.ImGui_ImplOpenGL3_Shutdown()
}

func ImplOpenGL3_RenderDrawData(d *DrawData) {
	C.ImGui_ImplOpenGL3_RenderDrawData((*C.ImDrawData)(d))
}
