package imgui

// #include "cimgui/cimgui.h"
// void wrap_igText(const char *fmt);
import "C"
import gofmt "fmt"
import "github.com/go-gl/mathgl/mgl32"
import "unsafe"

func CreateContext(shared_font_atlas *C.ImFontAtlas) *C.ImGuiContext {
	return C.igCreateContext(shared_font_atlas)
}

func DestroyContext(ctx *C.ImGuiContext) {
	C.igDestroyContext(ctx)
}

func Render() {
	C.igRender()
}

func GetDrawData() *C.ImDrawData {
	return C.igGetDrawData()
}

func ShowDemoWindow(p_open *bool) {
	C.igShowDemoWindow((*C.bool)(p_open))
}

func Begin(name string, p_open *bool, flags C.ImGuiWindowFlags) bool {
	return (bool)(C.igBegin(stringPool.StoreCString(name), (*C.bool)(p_open), flags))
}

func End() {
	C.igEnd()
}

func Text(fmt string, args ...interface{}) {
	s := stringPool.StoreCString(gofmt.Sprintf(fmt, args...))
	C.wrap_igText(s)
}

func Button(label string, size mgl32.Vec2) bool {
	return (bool)(C.igButton(stringPool.StoreCString(label), C.ImVec2{C.float(size.X()), C.float(size.Y())}))
}

func Checkbox(label string, v *bool) bool {
	return (bool)(C.igCheckbox(stringPool.StoreCString(label), (*C.bool)(v)))
}

func DockSpaceOverViewport(dockspace_id C.ImGuiID, viewport *C.ImGuiViewport, flags C.ImGuiDockNodeFlags, window_class *C.ImGuiWindowClass) C.ImGuiID {
	return C.igDockSpaceOverViewport(dockspace_id, viewport, flags, window_class)
}

func UpdatePlatformWindows() {
	C.igUpdatePlatformWindows()
}

func RenderPlatformWindowsDefault(platform_render_arg unsafe.Pointer, renderer_render_arg unsafe.Pointer) {
	C.igRenderPlatformWindowsDefault(platform_render_arg, renderer_render_arg)
}

