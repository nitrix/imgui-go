package imgui

//go:generate go run ./generator

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -Idist/include
// #cgo linux LDFLAGS: -lm
// #cgo darwin LDFLAGS: -framework CoreFoundation -framework OpenGL -framework Cocoa -framework IOKit -framework QuartzCore
// #include "dist/cimgui/cimgui.h"
// #include <stdlib.h>
import "C"
import (
	"github.com/go-gl/mathgl/mgl32"
	_ "github.com/nitrix/imgui-go/dist/cimgui"
)

// These are from enums.
type MouseSource C.ImGuiMouseSource
type Axis C.ImGuiAxis
type ContextHookType C.ImGuiContextHookType
type PopupPositionPolicy C.ImGuiPopupPositionPolicy
type InputEventType C.ImGuiInputEventType
type Key C.ImGuiKey
type WindowDockStyleCol C.ImGuiWindowDockStyleCol
type NavLayer C.ImGuiNavLayer
type SortDirection C.ImGuiSortDirection
type PlotType C.ImGuiPlotType
type LogType C.ImGuiLogType
type SelectionRequestType C.ImGuiSelectionRequestType
type InputSource C.ImGuiInputSource
type DockNodeState C.ImGuiDockNodeState
type LocKey C.ImGuiLocKey
type Dir C.ImGuiDir

var stringPool StringPool

func NewFrame() {
	stringPool.Reset()
	C.igNewFrame()
}

func imVec2ToMglVec2(v C.ImVec2) mgl32.Vec2 {
	return mgl32.Vec2{float32(v.x), float32(v.y)}
}

func imVec4ToMglVec4(v C.ImVec4) mgl32.Vec4 {
	return mgl32.Vec4{float32(v.x), float32(v.y), float32(v.z), float32(v.w)}
}

func mglVec2ToImVec2(v mgl32.Vec2) C.ImVec2 {
	return C.ImVec2{x: C.float(v.X()), y: C.float(v.Y())}
}

func mglVec4ToImVec4(v mgl32.Vec4) C.ImVec4 {
	return C.ImVec4{x: C.float(v.X()), y: C.float(v.Y()), z: C.float(v.Z()), w: C.float(v.W())}
}
