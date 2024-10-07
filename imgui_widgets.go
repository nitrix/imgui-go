package imgui

// #include "cimgui/cimgui.h"
// void wrap_igText(const char *text);
import "C"

// func Checkbox(label string, v *bool) bool {
// 	s := stringPool.StoreCString(label)
// 	return (bool)(C.igCheckbox(s, (*C.bool)(v)))
// }

// func Button(label string, v2 mgl32.Vec2) bool {
// 	s := stringPool.StoreCString(label)
// 	return (bool)(C.igButton(s, C.ImVec2{C.float(v2.X()), C.float(v2.Y())}))
// }
