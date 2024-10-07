package imgui

// #include "cimgui/cimgui.h"
// void wrap_igText(const char *fmt);
import "C"
import gofmt "fmt"
import "github.com/go-gl/mathgl/mgl32"
import "unsafe"

func CreateContext(shared_font_atlas *FontAtlas) *Context {
	return (*Context)(C.igCreateContext((*C.ImFontAtlas)(shared_font_atlas)))
}

func DestroyContext(ctx *Context) {
	C.igDestroyContext((*C.ImGuiContext)(ctx))
}

func Render() {
	C.igRender()
}

func GetDrawData() *DrawData {
	return (*DrawData)(C.igGetDrawData())
}

func ShowDemoWindow(p_open *bool) {
	C.igShowDemoWindow((*C.bool)(p_open))
}

func Begin(name string, p_open *bool, flags WindowFlags) bool {
	return (bool)(C.igBegin(stringPool.StoreCString(name), (*C.bool)(p_open), (C.ImGuiWindowFlags)(flags)))
}

func End() {
	C.igEnd()
}

func Text(fmt string, args ...interface{}) {
	variadic := stringPool.StoreCString(gofmt.Sprintf(fmt, args...))
	C.wrap_igText(variadic)
}

func Button(label string, size mgl32.Vec2) bool {
	return (bool)(C.igButton(stringPool.StoreCString(label), C.ImVec2{C.float(size.X()), C.float(size.Y())}))
}

func Checkbox(label string, v *bool) bool {
	return (bool)(C.igCheckbox(stringPool.StoreCString(label), (*C.bool)(v)))
}

func DockSpaceOverViewport(dockspace_id ID, viewport *Viewport, flags DockNodeFlags, window_class *WindowClass) ID {
	return (ID)(C.igDockSpaceOverViewport((C.ImGuiID)(dockspace_id), (*C.ImGuiViewport)(viewport), (C.ImGuiDockNodeFlags)(flags), (*C.ImGuiWindowClass)(window_class)))
}

func UpdatePlatformWindows() {
	C.igUpdatePlatformWindows()
}

func RenderPlatformWindowsDefault(platform_render_arg unsafe.Pointer, renderer_render_arg unsafe.Pointer) {
	C.igRenderPlatformWindowsDefault(platform_render_arg, renderer_render_arg)
}

