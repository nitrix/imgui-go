package imgui

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -I../dist/include
// #cgo windows LDFLAGS: -L../dist/windows
// #cgo linux LDFLAGS: -L../dist/linux
// #cgo windows LDFLAGS: -lcimgui -limm32 -luser32 -lkernel32 -lgdi32 -lshell32 -static -lstdc++
// #cgo linux LDFLAGS: -lcimgui -lm -lc++
// #cgo darwin,amd64 LDFLAGS: -L../dist/macos/amd64
// #cgo darwin,arm64 LDFLAGS: -L../dist/macos/arm64
// #cgo darwin LDFLAGS: -lcimgui -framework CoreFoundation -lc++ -framework OpenGL -framework Cocoa -framework IOKit -framework QuartzCore
// #include "cimgui/cimgui.h"
// #include <stdlib.h>
import "C"

var stringPool StringPool

func NewFrame() {
	stringPool.Reset()
	C.igNewFrame()
}

// func ShowDemoWindow() {
// 	C.igShowDemoWindow(nil)
// }

// func Begin(name string, open *bool, flags int) bool {
// 	s := stringPool.StoreCString(name)
// 	return (bool)(C.igBegin(s, (*C.bool)(open), (C.ImGuiWindowFlags)(flags)))
// }

// func End() {
// 	C.igEnd()
// }

// func DockSpaceOverViewport(dockspaceID C.ImGuiID, viewport *C.ImGuiViewport, flags C.ImGuiDockNodeFlags, windowClass *C.ImGuiWindowClass) {
// 	C.igDockSpaceOverViewport(dockspaceID, viewport, flags, windowClass)
// }

// func UpdatePlatformWindows() {
// 	C.igUpdatePlatformWindows()
// }

// func RenderPlatformWindowsDefault(platformArg unsafe.Pointer, rendererArg unsafe.Pointer) {
// 	C.igRenderPlatformWindowsDefault(platformArg, rendererArg)
// }

// func GetDrawData() *DrawData {
// 	return (*DrawData)(C.igGetDrawData())
// }

// func Render() {
// 	C.igRender()
// }
