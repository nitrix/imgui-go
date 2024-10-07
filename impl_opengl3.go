package imgui

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -DCIMGUI_USE_OPENGL3 -Idist/include
// #cgo windows LDFLAGS: -Ldist/windows
// #cgo linux LDFLAGS: -Ldist/linux
// #cgo windows LDFLAGS: -lcimgui -limm32 -luser32 -lkernel32 -lgdi32 -lshell32 -static -lstdc++
// #cgo linux LDFLAGS: -lcimgui -lm -lc++
// #cgo darwin,amd64 LDFLAGS: -Ldist/macos/amd64
// #cgo darwin,arm64 LDFLAGS: -Ldist/macos/arm64
// #cgo darwin LDFLAGS: -lcimgui -framework CoreFoundation -lc++ -framework OpenGL -framework Cocoa -framework IOKit -framework QuartzCore
// #include "cimgui/cimgui.h"
// #include "cimgui/cimgui_impl.h"
import "C"

import (
	"github.com/nitrix/glfw-go"
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

func ImplOpenGL3_RenderDrawData(d *C.ImDrawData) {
	C.ImGui_ImplOpenGL3_RenderDrawData(d)
}