package imgui

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -Idist/include
// #include "cimgui/cimgui.h"
// #include "imgui_wrappers.h"
// #include <stdbool.h>
import "C"
import "fmt"
import "unsafe"
import "github.com/go-gl/mathgl/mgl32"

func Begin(name string, p_open *bool, flags WindowFlags) bool {
	return (bool)(C.igBegin((*C.char)(stringPool.StoreCString(name)), (*C.bool)(p_open), (C.ImGuiWindowFlags)(flags)))
}

func Button(label string, size mgl32.Vec2) bool {
	return (bool)(C.igButton((*C.char)(stringPool.StoreCString(label)), (C.ImVec2)(mglVec2ToImVec2(size))))
}

func Checkbox(label string, v *bool) bool {
	return (bool)(C.igCheckbox((*C.char)(stringPool.StoreCString(label)), (*C.bool)(v)))
}

func CreateContext(shared_font_atlas *FontAtlas) *Context {
	return (*Context)(C.igCreateContext((*C.ImFontAtlas)(shared_font_atlas)))
}

func DestroyContext(ctx *Context) {
	C.igDestroyContext((*C.ImGuiContext)(ctx))
}

func DockSpaceOverViewport(dockspace_id ID, viewport *Viewport, flags DockNodeFlags, window_class *WindowClass) ID {
	return (ID)(C.igDockSpaceOverViewport((C.ImGuiID)(dockspace_id), (*C.ImGuiViewport)(viewport), (C.ImGuiDockNodeFlags)(flags), (*C.ImGuiWindowClass)(window_class)))
}

func End() {
	C.igEnd()
}

func GetDrawData() *DrawData {
	return (*DrawData)(C.igGetDrawData())
}

func Render() {
	C.igRender()
}

func RenderPlatformWindowsDefault(platform_render_arg unsafe.Pointer, renderer_render_arg unsafe.Pointer) {
	C.igRenderPlatformWindowsDefault((unsafe.Pointer)(platform_render_arg), (unsafe.Pointer)(renderer_render_arg))
}

func ShowDemoWindow(p_open *bool) {
	C.igShowDemoWindow((*C.bool)(p_open))
}

func Text(_fmt string, _args ...interface{}) {
	_vfmt := fmt.Sprintf(_fmt, _args...)
	C.wrap_igText((*C.char)(stringPool.StoreCString(_vfmt)))
}

func UpdatePlatformWindows() {
	C.igUpdatePlatformWindows()
}

