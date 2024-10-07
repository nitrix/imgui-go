package imgui

// #include "cimgui/cimgui.h"
import "C"

type DrawData C.ImDrawData
type FontAtlas C.ImFontAtlas

type Context C.ImGuiContext
type ID C.ImGuiID
type WindowFlags C.ImGuiWindowFlags
type DockNodeFlags C.ImGuiDockNodeFlags
type WindowClass C.ImGuiWindowClass
type Style C.ImGuiStyle
type StyleMod C.ImGuiStyleMod
type StyleVar C.ImGuiStyleVar
type Col C.ImGuiCol
type Cond C.ImGuiCond
type Viewport C.ImGuiViewport
