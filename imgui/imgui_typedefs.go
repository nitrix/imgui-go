package imgui

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1
// #include "../thirdparty/cimgui/cimgui.h"
import "C"

type ID C.ImGuiID
type Context C.ImGuiContext
type Viewport C.ImGuiViewport
type DrawData C.ImDrawData
type DockNodeFlags C.ImGuiDockNodeFlags
type WindowClass C.ImGuiWindowClass
type WindowFlags C.ImGuiWindowFlags
type ChildFlags C.ImGuiChildFlags
type PlatformIO C.ImGuiPlatformIO
type Style C.ImGuiStyle
type FocusedFlags C.ImGuiFocusedFlags
