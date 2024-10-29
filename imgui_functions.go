package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS 1
// #include "dist/include/cimgui/cimgui.h"
// #include "imgui_wrappers.h"
// #include <stdbool.h>
import "C"
import "fmt"
import "unsafe"
import "github.com/go-gl/mathgl/mgl32"

func AcceptDragDropPayload(ty string, flags DragDropFlags) *Payload {
	a0 := stringPool.StoreCString(ty)
	a1 := (C.ImGuiDragDropFlags)(flags)
	call := C.igAcceptDragDropPayload(a0, a1)
	r := (*Payload)(call)
	return r
}

func ActivateItemByID(id ID) {
	a0 := (C.ImGuiID)(id)
	C.igActivateItemByID(a0)
}

func AddContextHook(context *Context, hook *ContextHook) ID {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(context))
	a1 := (*C.ImGuiContextHook)(unsafe.Pointer(hook))
	call := C.igAddContextHook(a0, a1)
	r := (ID)(call)
	return r
}

func AddSettingsHandler(handler *SettingsHandler) {
	a0 := (*C.ImGuiSettingsHandler)(unsafe.Pointer(handler))
	C.igAddSettingsHandler(a0)
}

func AlignTextToFramePadding() {
	C.igAlignTextToFramePadding()
}

func ArrowButton(strId string, dir Dir) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiDir)(dir)
	call := C.igArrowButton(a0, a1)
	r := (bool)(call)
	return r
}

func ArrowButtonEx(strId string, dir Dir, sizeArg mgl32.Vec2, flags ButtonFlags) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiDir)(dir)
	a2 := mglVec2ToImVec2(sizeArg)
	a3 := (C.ImGuiButtonFlags)(flags)
	call := C.igArrowButtonEx(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func Begin(name string, pOpen *bool, flags WindowFlags) bool {
	a0 := stringPool.StoreCString(name)
	a1 := (*C.bool)(unsafe.Pointer(pOpen))
	a2 := (C.ImGuiWindowFlags)(flags)
	call := C.igBegin(a0, a1, a2)
	r := (bool)(call)
	return r
}

func BeginBoxSelect(scopeRect Rect, window *Window, boxSelectId ID, msFlags MultiSelectFlags) bool {
	a0 := (C.ImRect)(scopeRect)
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a2 := (C.ImGuiID)(boxSelectId)
	a3 := (C.ImGuiMultiSelectFlags)(msFlags)
	call := C.igBeginBoxSelect(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func BeginChildEx(name string, id ID, sizeArg mgl32.Vec2, childFlags ChildFlags, windowFlags WindowFlags) bool {
	a0 := stringPool.StoreCString(name)
	a1 := (C.ImGuiID)(id)
	a2 := mglVec2ToImVec2(sizeArg)
	a3 := (C.ImGuiChildFlags)(childFlags)
	a4 := (C.ImGuiWindowFlags)(windowFlags)
	call := C.igBeginChildEx(a0, a1, a2, a3, a4)
	r := (bool)(call)
	return r
}

func BeginChild_ID(id ID, size mgl32.Vec2, childFlags ChildFlags, windowFlags WindowFlags) bool {
	a0 := (C.ImGuiID)(id)
	a1 := mglVec2ToImVec2(size)
	a2 := (C.ImGuiChildFlags)(childFlags)
	a3 := (C.ImGuiWindowFlags)(windowFlags)
	call := C.igBeginChild_ID(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func BeginChild_Str(strId string, size mgl32.Vec2, childFlags ChildFlags, windowFlags WindowFlags) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := mglVec2ToImVec2(size)
	a2 := (C.ImGuiChildFlags)(childFlags)
	a3 := (C.ImGuiWindowFlags)(windowFlags)
	call := C.igBeginChild_Str(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func BeginColumns(strId string, count int, flags OldColumnFlags) {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.int)(count)
	a2 := (C.ImGuiOldColumnFlags)(flags)
	C.igBeginColumns(a0, a1, a2)
}

func BeginCombo(label string, previewValue string, flags ComboFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(previewValue)
	a2 := (C.ImGuiComboFlags)(flags)
	call := C.igBeginCombo(a0, a1, a2)
	r := (bool)(call)
	return r
}

func BeginComboPopup(popupId ID, bb Rect, flags ComboFlags) bool {
	a0 := (C.ImGuiID)(popupId)
	a1 := (C.ImRect)(bb)
	a2 := (C.ImGuiComboFlags)(flags)
	call := C.igBeginComboPopup(a0, a1, a2)
	r := (bool)(call)
	return r
}

func BeginComboPreview() bool {
	call := C.igBeginComboPreview()
	r := (bool)(call)
	return r
}

func BeginDisabled(disabled bool) {
	a0 := (C.bool)(disabled)
	C.igBeginDisabled(a0)
}

func BeginDisabledOverrideReenable() {
	C.igBeginDisabledOverrideReenable()
}

func BeginDockableDragDropSource(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igBeginDockableDragDropSource(a0)
}

func BeginDockableDragDropTarget(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igBeginDockableDragDropTarget(a0)
}

func BeginDocked(window *Window, pOpen *bool) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (*C.bool)(unsafe.Pointer(pOpen))
	C.igBeginDocked(a0, a1)
}

func BeginDragDropSource(flags DragDropFlags) bool {
	a0 := (C.ImGuiDragDropFlags)(flags)
	call := C.igBeginDragDropSource(a0)
	r := (bool)(call)
	return r
}

func BeginDragDropTarget() bool {
	call := C.igBeginDragDropTarget()
	r := (bool)(call)
	return r
}

func BeginDragDropTargetCustom(bb Rect, id ID) bool {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	call := C.igBeginDragDropTargetCustom(a0, a1)
	r := (bool)(call)
	return r
}

func BeginErrorTooltip() bool {
	call := C.igBeginErrorTooltip()
	r := (bool)(call)
	return r
}

func BeginGroup() {
	C.igBeginGroup()
}

func BeginItemTooltip() bool {
	call := C.igBeginItemTooltip()
	r := (bool)(call)
	return r
}

func BeginListBox(label string, size mgl32.Vec2) bool {
	a0 := stringPool.StoreCString(label)
	a1 := mglVec2ToImVec2(size)
	call := C.igBeginListBox(a0, a1)
	r := (bool)(call)
	return r
}

func BeginMainMenuBar() bool {
	call := C.igBeginMainMenuBar()
	r := (bool)(call)
	return r
}

func BeginMenu(label string, enabled bool) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.bool)(enabled)
	call := C.igBeginMenu(a0, a1)
	r := (bool)(call)
	return r
}

func BeginMenuBar() bool {
	call := C.igBeginMenuBar()
	r := (bool)(call)
	return r
}

func BeginMenuEx(label string, icon string, enabled bool) bool {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(icon)
	a2 := (C.bool)(enabled)
	call := C.igBeginMenuEx(a0, a1, a2)
	r := (bool)(call)
	return r
}

func BeginMultiSelect(flags MultiSelectFlags, selectionSize int, itemsCount int) *MultiSelectIO {
	a0 := (C.ImGuiMultiSelectFlags)(flags)
	a1 := (C.int)(selectionSize)
	a2 := (C.int)(itemsCount)
	call := C.igBeginMultiSelect(a0, a1, a2)
	r := (*MultiSelectIO)(call)
	return r
}

func BeginPopup(strId string, flags WindowFlags) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiWindowFlags)(flags)
	call := C.igBeginPopup(a0, a1)
	r := (bool)(call)
	return r
}

func BeginPopupContextItem(strId string, popupFlags PopupFlags) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiPopupFlags)(popupFlags)
	call := C.igBeginPopupContextItem(a0, a1)
	r := (bool)(call)
	return r
}

func BeginPopupContextVoid(strId string, popupFlags PopupFlags) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiPopupFlags)(popupFlags)
	call := C.igBeginPopupContextVoid(a0, a1)
	r := (bool)(call)
	return r
}

func BeginPopupContextWindow(strId string, popupFlags PopupFlags) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiPopupFlags)(popupFlags)
	call := C.igBeginPopupContextWindow(a0, a1)
	r := (bool)(call)
	return r
}

func BeginPopupEx(id ID, extraWindowFlags WindowFlags) bool {
	a0 := (C.ImGuiID)(id)
	a1 := (C.ImGuiWindowFlags)(extraWindowFlags)
	call := C.igBeginPopupEx(a0, a1)
	r := (bool)(call)
	return r
}

func BeginPopupModal(name string, pOpen *bool, flags WindowFlags) bool {
	a0 := stringPool.StoreCString(name)
	a1 := (*C.bool)(unsafe.Pointer(pOpen))
	a2 := (C.ImGuiWindowFlags)(flags)
	call := C.igBeginPopupModal(a0, a1, a2)
	r := (bool)(call)
	return r
}

func BeginTabBar(strId string, flags TabBarFlags) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiTabBarFlags)(flags)
	call := C.igBeginTabBar(a0, a1)
	r := (bool)(call)
	return r
}

func BeginTabBarEx(tabBar *TabBar, bb Rect, flags TabBarFlags) bool {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (C.ImRect)(bb)
	a2 := (C.ImGuiTabBarFlags)(flags)
	call := C.igBeginTabBarEx(a0, a1, a2)
	r := (bool)(call)
	return r
}

func BeginTabItem(label string, pOpen *bool, flags TabItemFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.bool)(unsafe.Pointer(pOpen))
	a2 := (C.ImGuiTabItemFlags)(flags)
	call := C.igBeginTabItem(a0, a1, a2)
	r := (bool)(call)
	return r
}

func BeginTable(strId string, columns int, flags TableFlags, outerSize mgl32.Vec2, innerWidth float32) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.int)(columns)
	a2 := (C.ImGuiTableFlags)(flags)
	a3 := mglVec2ToImVec2(outerSize)
	a4 := (C.float)(innerWidth)
	call := C.igBeginTable(a0, a1, a2, a3, a4)
	r := (bool)(call)
	return r
}

func BeginTableEx(name string, id ID, columnsCount int, flags TableFlags, outerSize mgl32.Vec2, innerWidth float32) bool {
	a0 := stringPool.StoreCString(name)
	a1 := (C.ImGuiID)(id)
	a2 := (C.int)(columnsCount)
	a3 := (C.ImGuiTableFlags)(flags)
	a4 := mglVec2ToImVec2(outerSize)
	a5 := (C.float)(innerWidth)
	call := C.igBeginTableEx(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func BeginTooltip() bool {
	call := C.igBeginTooltip()
	r := (bool)(call)
	return r
}

func BeginTooltipEx(tooltipFlags TooltipFlags, extraWindowFlags WindowFlags) bool {
	a0 := (C.ImGuiTooltipFlags)(tooltipFlags)
	a1 := (C.ImGuiWindowFlags)(extraWindowFlags)
	call := C.igBeginTooltipEx(a0, a1)
	r := (bool)(call)
	return r
}

func BeginTooltipHidden() bool {
	call := C.igBeginTooltipHidden()
	r := (bool)(call)
	return r
}

func BeginViewportSideBar(name string, viewport *Viewport, dir Dir, size float32, windowFlags WindowFlags) bool {
	a0 := stringPool.StoreCString(name)
	a1 := (*C.ImGuiViewport)(unsafe.Pointer(viewport))
	a2 := (C.ImGuiDir)(dir)
	a3 := (C.float)(size)
	a4 := (C.ImGuiWindowFlags)(windowFlags)
	call := C.igBeginViewportSideBar(a0, a1, a2, a3, a4)
	r := (bool)(call)
	return r
}

func BringWindowToDisplayBack(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igBringWindowToDisplayBack(a0)
}

func BringWindowToDisplayBehind(window *Window, aboveWindow *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(aboveWindow))
	C.igBringWindowToDisplayBehind(a0, a1)
}

func BringWindowToDisplayFront(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igBringWindowToDisplayFront(a0)
}

func BringWindowToFocusFront(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igBringWindowToFocusFront(a0)
}

func Bullet() {
	C.igBullet()
}

func BulletText(vfmt string, vargs ...interface{}) {
	a0 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	C.wrap_igBulletText(a0)
}

func Button(label string, size mgl32.Vec2) bool {
	a0 := stringPool.StoreCString(label)
	a1 := mglVec2ToImVec2(size)
	call := C.igButton(a0, a1)
	r := (bool)(call)
	return r
}

func ButtonBehavior(bb Rect, id ID, outHovered *bool, outHeld *bool, flags ButtonFlags) bool {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	a2 := (*C.bool)(unsafe.Pointer(outHovered))
	a3 := (*C.bool)(unsafe.Pointer(outHeld))
	a4 := (C.ImGuiButtonFlags)(flags)
	call := C.igButtonBehavior(a0, a1, a2, a3, a4)
	r := (bool)(call)
	return r
}

func ButtonEx(label string, sizeArg mgl32.Vec2, flags ButtonFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := mglVec2ToImVec2(sizeArg)
	a2 := (C.ImGuiButtonFlags)(flags)
	call := C.igButtonEx(a0, a1, a2)
	r := (bool)(call)
	return r
}

func CalcItemSize(pOut *mgl32.Vec2, size mgl32.Vec2, defaultW float32, defaultH float32) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := mglVec2ToImVec2(size)
	a2 := (C.float)(defaultW)
	a3 := (C.float)(defaultH)
	C.igCalcItemSize(a0, a1, a2, a3)
}

func CalcItemWidth() float32 {
	call := C.igCalcItemWidth()
	r := (float32)(call)
	return r
}

func CalcRoundingFlagsForRectInRect(rIn Rect, rOuter Rect, threshold float32) DrawFlags {
	a0 := (C.ImRect)(rIn)
	a1 := (C.ImRect)(rOuter)
	a2 := (C.float)(threshold)
	call := C.igCalcRoundingFlagsForRectInRect(a0, a1, a2)
	r := (DrawFlags)(call)
	return r
}

func CalcTextSize(pOut *mgl32.Vec2, text string, textEnd string, hideTextAfterDoubleHash bool, wrapWidth float32) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := stringPool.StoreCString(text)
	a2 := stringPool.StoreCString(textEnd)
	a3 := (C.bool)(hideTextAfterDoubleHash)
	a4 := (C.float)(wrapWidth)
	C.igCalcTextSize(a0, a1, a2, a3, a4)
}

func CalcTypematicRepeatAmount(t0 float32, t1 float32, repeatDelay float32, repeatRate float32) int {
	a0 := (C.float)(t0)
	a1 := (C.float)(t1)
	a2 := (C.float)(repeatDelay)
	a3 := (C.float)(repeatRate)
	call := C.igCalcTypematicRepeatAmount(a0, a1, a2, a3)
	r := (int)(call)
	return r
}

func CalcWindowNextAutoFitSize(pOut *mgl32.Vec2, window *Window) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igCalcWindowNextAutoFitSize(a0, a1)
}

func CalcWrapWidthForPos(pos mgl32.Vec2, wrapPosX float32) float32 {
	a0 := mglVec2ToImVec2(pos)
	a1 := (C.float)(wrapPosX)
	call := C.igCalcWrapWidthForPos(a0, a1)
	r := (float32)(call)
	return r
}

func CallContextHooks(context *Context, ty ContextHookType) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(context))
	a1 := (C.ImGuiContextHookType)(ty)
	C.igCallContextHooks(a0, a1)
}

func Checkbox(label string, v *bool) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.bool)(unsafe.Pointer(v))
	call := C.igCheckbox(a0, a1)
	r := (bool)(call)
	return r
}

func CheckboxFlags_IntPtr(label string, flags *int, flagsValue int) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(flags))
	a2 := (C.int)(flagsValue)
	call := C.igCheckboxFlags_IntPtr(a0, a1, a2)
	r := (bool)(call)
	return r
}

func CheckboxFlags_S64Ptr(label string, flags *S64, flagsValue S64) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.ImS64)(unsafe.Pointer(flags))
	a2 := (C.ImS64)(flagsValue)
	call := C.igCheckboxFlags_S64Ptr(a0, a1, a2)
	r := (bool)(call)
	return r
}

func CheckboxFlags_U64Ptr(label string, flags *U64, flagsValue U64) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.ImU64)(unsafe.Pointer(flags))
	a2 := (C.ImU64)(flagsValue)
	call := C.igCheckboxFlags_U64Ptr(a0, a1, a2)
	r := (bool)(call)
	return r
}

func CheckboxFlags_UintPtr(label string, flags *uint, flagsValue uint) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.uint)(unsafe.Pointer(flags))
	a2 := (C.uint)(flagsValue)
	call := C.igCheckboxFlags_UintPtr(a0, a1, a2)
	r := (bool)(call)
	return r
}

func ClearActiveID() {
	C.igClearActiveID()
}

func ClearDragDrop() {
	C.igClearDragDrop()
}

func ClearIniSettings() {
	C.igClearIniSettings()
}

func ClearWindowSettings(name string) {
	a0 := stringPool.StoreCString(name)
	C.igClearWindowSettings(a0)
}

func CloseButton(id ID, pos mgl32.Vec2) bool {
	a0 := (C.ImGuiID)(id)
	a1 := mglVec2ToImVec2(pos)
	call := C.igCloseButton(a0, a1)
	r := (bool)(call)
	return r
}

func CloseCurrentPopup() {
	C.igCloseCurrentPopup()
}

func ClosePopupToLevel(remaining int, restoreFocusToWindowUnderPopup bool) {
	a0 := (C.int)(remaining)
	a1 := (C.bool)(restoreFocusToWindowUnderPopup)
	C.igClosePopupToLevel(a0, a1)
}

func ClosePopupsExceptModals() {
	C.igClosePopupsExceptModals()
}

func ClosePopupsOverWindow(refWindow *Window, restoreFocusToWindowUnderPopup bool) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(refWindow))
	a1 := (C.bool)(restoreFocusToWindowUnderPopup)
	C.igClosePopupsOverWindow(a0, a1)
}

func CollapseButton(id ID, pos mgl32.Vec2, dockNode *DockNode) bool {
	a0 := (C.ImGuiID)(id)
	a1 := mglVec2ToImVec2(pos)
	a2 := (*C.ImGuiDockNode)(unsafe.Pointer(dockNode))
	call := C.igCollapseButton(a0, a1, a2)
	r := (bool)(call)
	return r
}

func CollapsingHeader_BoolPtr(label string, pVisible *bool, flags TreeNodeFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.bool)(unsafe.Pointer(pVisible))
	a2 := (C.ImGuiTreeNodeFlags)(flags)
	call := C.igCollapsingHeader_BoolPtr(a0, a1, a2)
	r := (bool)(call)
	return r
}

func CollapsingHeader_TreeNodeFlags(label string, flags TreeNodeFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.ImGuiTreeNodeFlags)(flags)
	call := C.igCollapsingHeader_TreeNodeFlags(a0, a1)
	r := (bool)(call)
	return r
}

func ColorButton(descId string, col mgl32.Vec4, flags ColorEditFlags, size mgl32.Vec2) bool {
	a0 := stringPool.StoreCString(descId)
	a1 := mglVec4ToImVec4(col)
	a2 := (C.ImGuiColorEditFlags)(flags)
	a3 := mglVec2ToImVec2(size)
	call := C.igColorButton(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func ColorConvertFloat4ToU32(in mgl32.Vec4) U32 {
	a0 := mglVec4ToImVec4(in)
	call := C.igColorConvertFloat4ToU32(a0)
	r := (U32)(call)
	return r
}

func ColorConvertHSVtoRGB(h float32, s float32, v float32, outR *float32, outG *float32, outB *float32) {
	a0 := (C.float)(h)
	a1 := (C.float)(s)
	a2 := (C.float)(v)
	a3 := (*C.float)(unsafe.Pointer(outR))
	a4 := (*C.float)(unsafe.Pointer(outG))
	a5 := (*C.float)(unsafe.Pointer(outB))
	C.igColorConvertHSVtoRGB(a0, a1, a2, a3, a4, a5)
}

func ColorConvertRGBtoHSV(r float32, g float32, b float32, outH *float32, outS *float32, outV *float32) {
	a0 := (C.float)(r)
	a1 := (C.float)(g)
	a2 := (C.float)(b)
	a3 := (*C.float)(unsafe.Pointer(outH))
	a4 := (*C.float)(unsafe.Pointer(outS))
	a5 := (*C.float)(unsafe.Pointer(outV))
	C.igColorConvertRGBtoHSV(a0, a1, a2, a3, a4, a5)
}

func ColorConvertU32ToFloat4(pOut *mgl32.Vec4, in U32) {
	a0 := (*C.ImVec4)(unsafe.Pointer(&pOut[0]))
	a1 := (C.ImU32)(in)
	C.igColorConvertU32ToFloat4(a0, a1)
}

func ColorEdit3(label string, col *mgl32.Vec3, flags ColorEditFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(col))
	a2 := (C.ImGuiColorEditFlags)(flags)
	call := C.igColorEdit3(a0, a1, a2)
	r := (bool)(call)
	return r
}

func ColorEdit4(label string, col *mgl32.Vec4, flags ColorEditFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(&col[0]))
	a2 := (C.ImGuiColorEditFlags)(flags)
	call := C.igColorEdit4(a0, a1, a2)
	r := (bool)(call)
	return r
}

func ColorEditOptionsPopup(col *float32, flags ColorEditFlags) {
	a0 := (*C.float)(unsafe.Pointer(col))
	a1 := (C.ImGuiColorEditFlags)(flags)
	C.igColorEditOptionsPopup(a0, a1)
}

func ColorPicker3(label string, col *mgl32.Vec3, flags ColorEditFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(col))
	a2 := (C.ImGuiColorEditFlags)(flags)
	call := C.igColorPicker3(a0, a1, a2)
	r := (bool)(call)
	return r
}

func ColorPicker4(label string, col *mgl32.Vec4, flags ColorEditFlags, refCol *float32) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(&col[0]))
	a2 := (C.ImGuiColorEditFlags)(flags)
	a3 := (*C.float)(unsafe.Pointer(refCol))
	call := C.igColorPicker4(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func ColorPickerOptionsPopup(refCol *float32, flags ColorEditFlags) {
	a0 := (*C.float)(unsafe.Pointer(refCol))
	a1 := (C.ImGuiColorEditFlags)(flags)
	C.igColorPickerOptionsPopup(a0, a1)
}

func ColorTooltip(text string, col *float32, flags ColorEditFlags) {
	a0 := stringPool.StoreCString(text)
	a1 := (*C.float)(unsafe.Pointer(col))
	a2 := (C.ImGuiColorEditFlags)(flags)
	C.igColorTooltip(a0, a1, a2)
}

func Columns(count int, id string, borders bool) {
	a0 := (C.int)(count)
	a1 := stringPool.StoreCString(id)
	a2 := (C.bool)(borders)
	C.igColumns(a0, a1, a2)
}

func Combo_Str(label string, currentItem *int, itemsSeparatedByZeros string, popupMaxHeightInItems int) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(currentItem))
	a2 := stringPool.StoreCString(itemsSeparatedByZeros)
	a3 := (C.int)(popupMaxHeightInItems)
	call := C.igCombo_Str(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func Combo_Str_arr(label string, currentItem *int, items []string, itemsCount int, popupMaxHeightInItems int) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(currentItem))
	a2 := (**C.char)(unsafe.Pointer(&items[0]))
	a3 := (C.int)(itemsCount)
	a4 := (C.int)(popupMaxHeightInItems)
	call := C.igCombo_Str_arr(a0, a1, a2, a3, a4)
	r := (bool)(call)
	return r
}

func ConvertSingleModFlagToKey(key Key) Key {
	a0 := (C.ImGuiKey)(key)
	call := C.igConvertSingleModFlagToKey(a0)
	r := (Key)(call)
	return r
}

func CreateContext(sharedFontAtlas *FontAtlas) *Context {
	a0 := (*C.ImFontAtlas)(unsafe.Pointer(sharedFontAtlas))
	call := C.igCreateContext(a0)
	r := (*Context)(call)
	return r
}

func CreateNewWindowSettings(name string) *WindowSettings {
	a0 := stringPool.StoreCString(name)
	call := C.igCreateNewWindowSettings(a0)
	r := (*WindowSettings)(call)
	return r
}

func DataTypeApplyFromText(buf string, dataType DataType, pData unsafe.Pointer, format string, pDataWhenEmpty unsafe.Pointer) bool {
	a0 := stringPool.StoreCString(buf)
	a1 := (C.ImGuiDataType)(dataType)
	a2 := (unsafe.Pointer)(pData)
	a3 := stringPool.StoreCString(format)
	a4 := (unsafe.Pointer)(pDataWhenEmpty)
	call := C.igDataTypeApplyFromText(a0, a1, a2, a3, a4)
	r := (bool)(call)
	return r
}

func DataTypeApplyOp(dataType DataType, op int, output unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer) {
	a0 := (C.ImGuiDataType)(dataType)
	a1 := (C.int)(op)
	a2 := (unsafe.Pointer)(output)
	a3 := (unsafe.Pointer)(arg1)
	a4 := (unsafe.Pointer)(arg2)
	C.igDataTypeApplyOp(a0, a1, a2, a3, a4)
}

func DataTypeClamp(dataType DataType, pData unsafe.Pointer, pMin unsafe.Pointer, pMax unsafe.Pointer) bool {
	a0 := (C.ImGuiDataType)(dataType)
	a1 := (unsafe.Pointer)(pData)
	a2 := (unsafe.Pointer)(pMin)
	a3 := (unsafe.Pointer)(pMax)
	call := C.igDataTypeClamp(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func DataTypeCompare(dataType DataType, arg1 unsafe.Pointer, arg2 unsafe.Pointer) int {
	a0 := (C.ImGuiDataType)(dataType)
	a1 := (unsafe.Pointer)(arg1)
	a2 := (unsafe.Pointer)(arg2)
	call := C.igDataTypeCompare(a0, a1, a2)
	r := (int)(call)
	return r
}

func DataTypeFormatString(buf string, bufSize int, dataType DataType, pData unsafe.Pointer, format string) int {
	a0 := stringPool.StoreCString(buf)
	a1 := (C.int)(bufSize)
	a2 := (C.ImGuiDataType)(dataType)
	a3 := (unsafe.Pointer)(pData)
	a4 := stringPool.StoreCString(format)
	call := C.igDataTypeFormatString(a0, a1, a2, a3, a4)
	r := (int)(call)
	return r
}

func DataTypeGetInfo(dataType DataType) *DataTypeInfo {
	a0 := (C.ImGuiDataType)(dataType)
	call := C.igDataTypeGetInfo(a0)
	r := (*DataTypeInfo)(call)
	return r
}

func DataTypeIsZero(dataType DataType, pData unsafe.Pointer) bool {
	a0 := (C.ImGuiDataType)(dataType)
	a1 := (unsafe.Pointer)(pData)
	call := C.igDataTypeIsZero(a0, a1)
	r := (bool)(call)
	return r
}

func DestroyContext(ctx *Context) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	C.igDestroyContext(a0)
}

func DestroyPlatformWindow(viewport *ViewportP) {
	a0 := (*C.ImGuiViewportP)(unsafe.Pointer(viewport))
	C.igDestroyPlatformWindow(a0)
}

func DestroyPlatformWindows() {
	C.igDestroyPlatformWindows()
}

func DockBuilderAddNode(nodeId ID, flags DockNodeFlags) ID {
	a0 := (C.ImGuiID)(nodeId)
	a1 := (C.ImGuiDockNodeFlags)(flags)
	call := C.igDockBuilderAddNode(a0, a1)
	r := (ID)(call)
	return r
}

func DockBuilderCopyWindowSettings(srcName string, dstName string) {
	a0 := stringPool.StoreCString(srcName)
	a1 := stringPool.StoreCString(dstName)
	C.igDockBuilderCopyWindowSettings(a0, a1)
}

func DockBuilderDockWindow(windowName string, nodeId ID) {
	a0 := stringPool.StoreCString(windowName)
	a1 := (C.ImGuiID)(nodeId)
	C.igDockBuilderDockWindow(a0, a1)
}

func DockBuilderFinish(nodeId ID) {
	a0 := (C.ImGuiID)(nodeId)
	C.igDockBuilderFinish(a0)
}

func DockBuilderGetCentralNode(nodeId ID) *DockNode {
	a0 := (C.ImGuiID)(nodeId)
	call := C.igDockBuilderGetCentralNode(a0)
	r := (*DockNode)(call)
	return r
}

func DockBuilderGetNode(nodeId ID) *DockNode {
	a0 := (C.ImGuiID)(nodeId)
	call := C.igDockBuilderGetNode(a0)
	r := (*DockNode)(call)
	return r
}

func DockBuilderRemoveNode(nodeId ID) {
	a0 := (C.ImGuiID)(nodeId)
	C.igDockBuilderRemoveNode(a0)
}

func DockBuilderRemoveNodeChildNodes(nodeId ID) {
	a0 := (C.ImGuiID)(nodeId)
	C.igDockBuilderRemoveNodeChildNodes(a0)
}

func DockBuilderRemoveNodeDockedWindows(nodeId ID, clearSettingsRefs bool) {
	a0 := (C.ImGuiID)(nodeId)
	a1 := (C.bool)(clearSettingsRefs)
	C.igDockBuilderRemoveNodeDockedWindows(a0, a1)
}

func DockBuilderSetNodePos(nodeId ID, pos mgl32.Vec2) {
	a0 := (C.ImGuiID)(nodeId)
	a1 := mglVec2ToImVec2(pos)
	C.igDockBuilderSetNodePos(a0, a1)
}

func DockBuilderSetNodeSize(nodeId ID, size mgl32.Vec2) {
	a0 := (C.ImGuiID)(nodeId)
	a1 := mglVec2ToImVec2(size)
	C.igDockBuilderSetNodeSize(a0, a1)
}

func DockBuilderSplitNode(nodeId ID, splitDir Dir, sizeRatioForNodeAtDir float32, outIdAtDir *ID, outIdAtOppositeDir *ID) ID {
	a0 := (C.ImGuiID)(nodeId)
	a1 := (C.ImGuiDir)(splitDir)
	a2 := (C.float)(sizeRatioForNodeAtDir)
	a3 := (*C.ImGuiID)(unsafe.Pointer(outIdAtDir))
	a4 := (*C.ImGuiID)(unsafe.Pointer(outIdAtOppositeDir))
	call := C.igDockBuilderSplitNode(a0, a1, a2, a3, a4)
	r := (ID)(call)
	return r
}

func DockContextCalcDropPosForDocking(target *Window, targetNode *DockNode, payloadWindow *Window, payloadNode *DockNode, splitDir Dir, splitOuter bool, outPos *mgl32.Vec2) bool {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(target))
	a1 := (*C.ImGuiDockNode)(unsafe.Pointer(targetNode))
	a2 := (*C.ImGuiWindow)(unsafe.Pointer(payloadWindow))
	a3 := (*C.ImGuiDockNode)(unsafe.Pointer(payloadNode))
	a4 := (C.ImGuiDir)(splitDir)
	a5 := (C.bool)(splitOuter)
	a6 := (*C.ImVec2)(unsafe.Pointer(&outPos[0]))
	call := C.igDockContextCalcDropPosForDocking(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func DockContextClearNodes(ctx *Context, rootId ID, clearSettingsRefs bool) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	a1 := (C.ImGuiID)(rootId)
	a2 := (C.bool)(clearSettingsRefs)
	C.igDockContextClearNodes(a0, a1, a2)
}

func DockContextEndFrame(ctx *Context) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	C.igDockContextEndFrame(a0)
}

func DockContextFindNodeByID(ctx *Context, id ID) *DockNode {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	a1 := (C.ImGuiID)(id)
	call := C.igDockContextFindNodeByID(a0, a1)
	r := (*DockNode)(call)
	return r
}

func DockContextGenNodeID(ctx *Context) ID {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	call := C.igDockContextGenNodeID(a0)
	r := (ID)(call)
	return r
}

func DockContextInitialize(ctx *Context) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	C.igDockContextInitialize(a0)
}

func DockContextNewFrameUpdateDocking(ctx *Context) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	C.igDockContextNewFrameUpdateDocking(a0)
}

func DockContextNewFrameUpdateUndocking(ctx *Context) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	C.igDockContextNewFrameUpdateUndocking(a0)
}

func DockContextProcessUndockNode(ctx *Context, node *DockNode) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	a1 := (*C.ImGuiDockNode)(unsafe.Pointer(node))
	C.igDockContextProcessUndockNode(a0, a1)
}

func DockContextProcessUndockWindow(ctx *Context, window *Window, clearPersistentDockingRef bool) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a2 := (C.bool)(clearPersistentDockingRef)
	C.igDockContextProcessUndockWindow(a0, a1, a2)
}

func DockContextQueueDock(ctx *Context, target *Window, targetNode *DockNode, payload *Window, splitDir Dir, splitRatio float32, splitOuter bool) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(target))
	a2 := (*C.ImGuiDockNode)(unsafe.Pointer(targetNode))
	a3 := (*C.ImGuiWindow)(unsafe.Pointer(payload))
	a4 := (C.ImGuiDir)(splitDir)
	a5 := (C.float)(splitRatio)
	a6 := (C.bool)(splitOuter)
	C.igDockContextQueueDock(a0, a1, a2, a3, a4, a5, a6)
}

func DockContextQueueUndockNode(ctx *Context, node *DockNode) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	a1 := (*C.ImGuiDockNode)(unsafe.Pointer(node))
	C.igDockContextQueueUndockNode(a0, a1)
}

func DockContextQueueUndockWindow(ctx *Context, window *Window) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igDockContextQueueUndockWindow(a0, a1)
}

func DockContextRebuildNodes(ctx *Context) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	C.igDockContextRebuildNodes(a0)
}

func DockContextShutdown(ctx *Context) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	C.igDockContextShutdown(a0)
}

func DockNodeBeginAmendTabBar(node *DockNode) bool {
	a0 := (*C.ImGuiDockNode)(unsafe.Pointer(node))
	call := C.igDockNodeBeginAmendTabBar(a0)
	r := (bool)(call)
	return r
}

func DockNodeEndAmendTabBar() {
	C.igDockNodeEndAmendTabBar()
}

func DockNodeGetDepth(node *DockNode) int {
	a0 := (*C.ImGuiDockNode)(unsafe.Pointer(node))
	call := C.igDockNodeGetDepth(a0)
	r := (int)(call)
	return r
}

func DockNodeGetRootNode(node *DockNode) *DockNode {
	a0 := (*C.ImGuiDockNode)(unsafe.Pointer(node))
	call := C.igDockNodeGetRootNode(a0)
	r := (*DockNode)(call)
	return r
}

func DockNodeGetWindowMenuButtonId(node *DockNode) ID {
	a0 := (*C.ImGuiDockNode)(unsafe.Pointer(node))
	call := C.igDockNodeGetWindowMenuButtonId(a0)
	r := (ID)(call)
	return r
}

func DockNodeIsInHierarchyOf(node *DockNode, parent *DockNode) bool {
	a0 := (*C.ImGuiDockNode)(unsafe.Pointer(node))
	a1 := (*C.ImGuiDockNode)(unsafe.Pointer(parent))
	call := C.igDockNodeIsInHierarchyOf(a0, a1)
	r := (bool)(call)
	return r
}

func DockNodeWindowMenuHandler_Default(ctx *Context, node *DockNode, tabBar *TabBar) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	a1 := (*C.ImGuiDockNode)(unsafe.Pointer(node))
	a2 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	C.igDockNodeWindowMenuHandler_Default(a0, a1, a2)
}

func DockSpace(dockspaceId ID, size mgl32.Vec2, flags DockNodeFlags, windowClass *WindowClass) ID {
	a0 := (C.ImGuiID)(dockspaceId)
	a1 := mglVec2ToImVec2(size)
	a2 := (C.ImGuiDockNodeFlags)(flags)
	a3 := (*C.ImGuiWindowClass)(unsafe.Pointer(windowClass))
	call := C.igDockSpace(a0, a1, a2, a3)
	r := (ID)(call)
	return r
}

func DockSpaceOverViewport(dockspaceId ID, viewport *Viewport, flags DockNodeFlags, windowClass *WindowClass) ID {
	a0 := (C.ImGuiID)(dockspaceId)
	a1 := (*C.ImGuiViewport)(unsafe.Pointer(viewport))
	a2 := (C.ImGuiDockNodeFlags)(flags)
	a3 := (*C.ImGuiWindowClass)(unsafe.Pointer(windowClass))
	call := C.igDockSpaceOverViewport(a0, a1, a2, a3)
	r := (ID)(call)
	return r
}

func DragBehavior(id ID, dataType DataType, pV unsafe.Pointer, vSpeed float32, pMin unsafe.Pointer, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	a0 := (C.ImGuiID)(id)
	a1 := (C.ImGuiDataType)(dataType)
	a2 := (unsafe.Pointer)(pV)
	a3 := (C.float)(vSpeed)
	a4 := (unsafe.Pointer)(pMin)
	a5 := (unsafe.Pointer)(pMax)
	a6 := stringPool.StoreCString(format)
	a7 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragBehavior(a0, a1, a2, a3, a4, a5, a6, a7)
	r := (bool)(call)
	return r
}

func DragFloat(label string, v *float32, vSpeed float32, vMin float32, vMax float32, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(v))
	a2 := (C.float)(vSpeed)
	a3 := (C.float)(vMin)
	a4 := (C.float)(vMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragFloat(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func DragFloat2(label string, v [2]float32, vSpeed float32, vMin float32, vMax float32, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(&v[0]))
	a2 := (C.float)(vSpeed)
	a3 := (C.float)(vMin)
	a4 := (C.float)(vMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragFloat2(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func DragFloat3(label string, v *mgl32.Vec3, vSpeed float32, vMin float32, vMax float32, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(v))
	a2 := (C.float)(vSpeed)
	a3 := (C.float)(vMin)
	a4 := (C.float)(vMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragFloat3(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func DragFloat4(label string, v *mgl32.Vec4, vSpeed float32, vMin float32, vMax float32, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(&v[0]))
	a2 := (C.float)(vSpeed)
	a3 := (C.float)(vMin)
	a4 := (C.float)(vMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragFloat4(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func DragFloatRange2(label string, vCurrentMin *float32, vCurrentMax *float32, vSpeed float32, vMin float32, vMax float32, format string, formatMax string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(vCurrentMin))
	a2 := (*C.float)(unsafe.Pointer(vCurrentMax))
	a3 := (C.float)(vSpeed)
	a4 := (C.float)(vMin)
	a5 := (C.float)(vMax)
	a6 := stringPool.StoreCString(format)
	a7 := stringPool.StoreCString(formatMax)
	a8 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragFloatRange2(a0, a1, a2, a3, a4, a5, a6, a7, a8)
	r := (bool)(call)
	return r
}

func DragInt(label string, v *int, vSpeed float32, vMin int, vMax int, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(v))
	a2 := (C.float)(vSpeed)
	a3 := (C.int)(vMin)
	a4 := (C.int)(vMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragInt(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func DragInt2(label string, v [2]int, vSpeed float32, vMin int, vMax int, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(&v[0]))
	a2 := (C.float)(vSpeed)
	a3 := (C.int)(vMin)
	a4 := (C.int)(vMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragInt2(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func DragInt3(label string, v [3]int, vSpeed float32, vMin int, vMax int, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(&v[0]))
	a2 := (C.float)(vSpeed)
	a3 := (C.int)(vMin)
	a4 := (C.int)(vMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragInt3(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func DragInt4(label string, v [4]int, vSpeed float32, vMin int, vMax int, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(&v[0]))
	a2 := (C.float)(vSpeed)
	a3 := (C.int)(vMin)
	a4 := (C.int)(vMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragInt4(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func DragIntRange2(label string, vCurrentMin *int, vCurrentMax *int, vSpeed float32, vMin int, vMax int, format string, formatMax string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(vCurrentMin))
	a2 := (*C.int)(unsafe.Pointer(vCurrentMax))
	a3 := (C.float)(vSpeed)
	a4 := (C.int)(vMin)
	a5 := (C.int)(vMax)
	a6 := stringPool.StoreCString(format)
	a7 := stringPool.StoreCString(formatMax)
	a8 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragIntRange2(a0, a1, a2, a3, a4, a5, a6, a7, a8)
	r := (bool)(call)
	return r
}

func DragScalar(label string, dataType DataType, pData unsafe.Pointer, vSpeed float32, pMin unsafe.Pointer, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.ImGuiDataType)(dataType)
	a2 := (unsafe.Pointer)(pData)
	a3 := (C.float)(vSpeed)
	a4 := (unsafe.Pointer)(pMin)
	a5 := (unsafe.Pointer)(pMax)
	a6 := stringPool.StoreCString(format)
	a7 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragScalar(a0, a1, a2, a3, a4, a5, a6, a7)
	r := (bool)(call)
	return r
}

func DragScalarN(label string, dataType DataType, pData unsafe.Pointer, components int, vSpeed float32, pMin unsafe.Pointer, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.ImGuiDataType)(dataType)
	a2 := (unsafe.Pointer)(pData)
	a3 := (C.int)(components)
	a4 := (C.float)(vSpeed)
	a5 := (unsafe.Pointer)(pMin)
	a6 := (unsafe.Pointer)(pMax)
	a7 := stringPool.StoreCString(format)
	a8 := (C.ImGuiSliderFlags)(flags)
	call := C.igDragScalarN(a0, a1, a2, a3, a4, a5, a6, a7, a8)
	r := (bool)(call)
	return r
}

func Dummy(size mgl32.Vec2) {
	a0 := mglVec2ToImVec2(size)
	C.igDummy(a0)
}

func End() {
	C.igEnd()
}

func EndBoxSelect(scopeRect Rect, msFlags MultiSelectFlags) {
	a0 := (C.ImRect)(scopeRect)
	a1 := (C.ImGuiMultiSelectFlags)(msFlags)
	C.igEndBoxSelect(a0, a1)
}

func EndChild() {
	C.igEndChild()
}

func EndColumns() {
	C.igEndColumns()
}

func EndCombo() {
	C.igEndCombo()
}

func EndComboPreview() {
	C.igEndComboPreview()
}

func EndDisabled() {
	C.igEndDisabled()
}

func EndDisabledOverrideReenable() {
	C.igEndDisabledOverrideReenable()
}

func EndDragDropSource() {
	C.igEndDragDropSource()
}

func EndDragDropTarget() {
	C.igEndDragDropTarget()
}

func EndErrorTooltip() {
	C.igEndErrorTooltip()
}

func EndFrame() {
	C.igEndFrame()
}

func EndGroup() {
	C.igEndGroup()
}

func EndListBox() {
	C.igEndListBox()
}

func EndMainMenuBar() {
	C.igEndMainMenuBar()
}

func EndMenu() {
	C.igEndMenu()
}

func EndMenuBar() {
	C.igEndMenuBar()
}

func EndMultiSelect() *MultiSelectIO {
	call := C.igEndMultiSelect()
	r := (*MultiSelectIO)(call)
	return r
}

func EndPopup() {
	C.igEndPopup()
}

func EndTabBar() {
	C.igEndTabBar()
}

func EndTabItem() {
	C.igEndTabItem()
}

func EndTable() {
	C.igEndTable()
}

func EndTooltip() {
	C.igEndTooltip()
}

func ErrorCheckEndFrameFinalizeErrorTooltip() {
	C.igErrorCheckEndFrameFinalizeErrorTooltip()
}

func ErrorCheckUsingSetCursorPosToExtendParentBoundaries() {
	C.igErrorCheckUsingSetCursorPosToExtendParentBoundaries()
}

func ErrorLog(msg string) bool {
	a0 := stringPool.StoreCString(msg)
	call := C.igErrorLog(a0)
	r := (bool)(call)
	return r
}

func ErrorRecoveryStoreState(stateOut *ErrorRecoveryState) {
	a0 := (*C.ImGuiErrorRecoveryState)(unsafe.Pointer(stateOut))
	C.igErrorRecoveryStoreState(a0)
}

func ErrorRecoveryTryToRecoverState(stateIn *ErrorRecoveryState) {
	a0 := (*C.ImGuiErrorRecoveryState)(unsafe.Pointer(stateIn))
	C.igErrorRecoveryTryToRecoverState(a0)
}

func ErrorRecoveryTryToRecoverWindowState(stateIn *ErrorRecoveryState) {
	a0 := (*C.ImGuiErrorRecoveryState)(unsafe.Pointer(stateIn))
	C.igErrorRecoveryTryToRecoverWindowState(a0)
}

func FindBestWindowPosForPopup(pOut *mgl32.Vec2, window *Window) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igFindBestWindowPosForPopup(a0, a1)
}

func FindBestWindowPosForPopupEx(pOut *mgl32.Vec2, refPos mgl32.Vec2, size mgl32.Vec2, lastDir *Dir, rOuter Rect, rAvoid Rect, policy PopupPositionPolicy) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := mglVec2ToImVec2(refPos)
	a2 := mglVec2ToImVec2(size)
	a3 := (*C.ImGuiDir)(unsafe.Pointer(lastDir))
	a4 := (C.ImRect)(rOuter)
	a5 := (C.ImRect)(rAvoid)
	a6 := (C.ImGuiPopupPositionPolicy)(policy)
	C.igFindBestWindowPosForPopupEx(a0, a1, a2, a3, a4, a5, a6)
}

func FindBlockingModal(window *Window) *Window {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	call := C.igFindBlockingModal(a0)
	r := (*Window)(call)
	return r
}

func FindBottomMostVisibleWindowWithinBeginStack(window *Window) *Window {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	call := C.igFindBottomMostVisibleWindowWithinBeginStack(a0)
	r := (*Window)(call)
	return r
}

func FindHoveredViewportFromPlatformWindowStack(mousePlatformPos mgl32.Vec2) *ViewportP {
	a0 := mglVec2ToImVec2(mousePlatformPos)
	call := C.igFindHoveredViewportFromPlatformWindowStack(a0)
	r := (*ViewportP)(call)
	return r
}

func FindHoveredWindowEx(pos mgl32.Vec2, findFirstAndInAnyViewport bool, outHoveredWindow **Window, outHoveredWindowUnderMovingWindow **Window) {
	a0 := mglVec2ToImVec2(pos)
	a1 := (C.bool)(findFirstAndInAnyViewport)
	a2 := (**C.ImGuiWindow)(unsafe.Pointer(outHoveredWindow))
	a3 := (**C.ImGuiWindow)(unsafe.Pointer(outHoveredWindowUnderMovingWindow))
	C.igFindHoveredWindowEx(a0, a1, a2, a3)
}

func FindOrCreateColumns(window *Window, id ID) *OldColumns {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImGuiID)(id)
	call := C.igFindOrCreateColumns(a0, a1)
	r := (*OldColumns)(call)
	return r
}

func FindRenderedTextEnd(text string, textEnd string) string {
	a0 := stringPool.StoreCString(text)
	a1 := stringPool.StoreCString(textEnd)
	call := C.igFindRenderedTextEnd(a0, a1)
	r := C.GoString(call)
	return r
}

func FindSettingsHandler(typeName string) *SettingsHandler {
	a0 := stringPool.StoreCString(typeName)
	call := C.igFindSettingsHandler(a0)
	r := (*SettingsHandler)(call)
	return r
}

func FindViewportByID(id ID) *Viewport {
	a0 := (C.ImGuiID)(id)
	call := C.igFindViewportByID(a0)
	r := (*Viewport)(call)
	return r
}

func FindViewportByPlatformHandle(platformHandle unsafe.Pointer) *Viewport {
	a0 := (unsafe.Pointer)(platformHandle)
	call := C.igFindViewportByPlatformHandle(a0)
	r := (*Viewport)(call)
	return r
}

func FindWindowByID(id ID) *Window {
	a0 := (C.ImGuiID)(id)
	call := C.igFindWindowByID(a0)
	r := (*Window)(call)
	return r
}

func FindWindowByName(name string) *Window {
	a0 := stringPool.StoreCString(name)
	call := C.igFindWindowByName(a0)
	r := (*Window)(call)
	return r
}

func FindWindowDisplayIndex(window *Window) int {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	call := C.igFindWindowDisplayIndex(a0)
	r := (int)(call)
	return r
}

func FindWindowSettingsByID(id ID) *WindowSettings {
	a0 := (C.ImGuiID)(id)
	call := C.igFindWindowSettingsByID(a0)
	r := (*WindowSettings)(call)
	return r
}

func FindWindowSettingsByWindow(window *Window) *WindowSettings {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	call := C.igFindWindowSettingsByWindow(a0)
	r := (*WindowSettings)(call)
	return r
}

func FixupKeyChord(keyChord KeyChord) KeyChord {
	a0 := (C.ImGuiKeyChord)(keyChord)
	call := C.igFixupKeyChord(a0)
	r := (KeyChord)(call)
	return r
}

func FocusItem() {
	C.igFocusItem()
}

func FocusTopMostWindowUnderOne(underThisWindow *Window, ignoreWindow *Window, filterViewport *Viewport, flags FocusRequestFlags) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(underThisWindow))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(ignoreWindow))
	a2 := (*C.ImGuiViewport)(unsafe.Pointer(filterViewport))
	a3 := (C.ImGuiFocusRequestFlags)(flags)
	C.igFocusTopMostWindowUnderOne(a0, a1, a2, a3)
}

func FocusWindow(window *Window, flags FocusRequestFlags) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImGuiFocusRequestFlags)(flags)
	C.igFocusWindow(a0, a1)
}

func GcAwakeTransientWindowBuffers(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igGcAwakeTransientWindowBuffers(a0)
}

func GcCompactTransientMiscBuffers() {
	C.igGcCompactTransientMiscBuffers()
}

func GcCompactTransientWindowBuffers(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igGcCompactTransientWindowBuffers(a0)
}

func GetActiveID() ID {
	call := C.igGetActiveID()
	r := (ID)(call)
	return r
}

func GetBackgroundDrawList(viewport *Viewport) *DrawList {
	a0 := (*C.ImGuiViewport)(unsafe.Pointer(viewport))
	call := C.igGetBackgroundDrawList(a0)
	r := (*DrawList)(call)
	return r
}

func GetBoxSelectState(id ID) *BoxSelectState {
	a0 := (C.ImGuiID)(id)
	call := C.igGetBoxSelectState(a0)
	r := (*BoxSelectState)(call)
	return r
}

func GetClipboardText() string {
	call := C.igGetClipboardText()
	r := C.GoString(call)
	return r
}

func GetColorU32_Col(idx Col, alphaMul float32) U32 {
	a0 := (C.ImGuiCol)(idx)
	a1 := (C.float)(alphaMul)
	call := C.igGetColorU32_Col(a0, a1)
	r := (U32)(call)
	return r
}

func GetColorU32_U32(col U32, alphaMul float32) U32 {
	a0 := (C.ImU32)(col)
	a1 := (C.float)(alphaMul)
	call := C.igGetColorU32_U32(a0, a1)
	r := (U32)(call)
	return r
}

func GetColorU32_Vec4(col mgl32.Vec4) U32 {
	a0 := mglVec4ToImVec4(col)
	call := C.igGetColorU32_Vec4(a0)
	r := (U32)(call)
	return r
}

func GetColumnIndex() int {
	call := C.igGetColumnIndex()
	r := (int)(call)
	return r
}

func GetColumnNormFromOffset(columns *OldColumns, offset float32) float32 {
	a0 := (*C.ImGuiOldColumns)(unsafe.Pointer(columns))
	a1 := (C.float)(offset)
	call := C.igGetColumnNormFromOffset(a0, a1)
	r := (float32)(call)
	return r
}

func GetColumnOffset(columnIndex int) float32 {
	a0 := (C.int)(columnIndex)
	call := C.igGetColumnOffset(a0)
	r := (float32)(call)
	return r
}

func GetColumnOffsetFromNorm(columns *OldColumns, offsetNorm float32) float32 {
	a0 := (*C.ImGuiOldColumns)(unsafe.Pointer(columns))
	a1 := (C.float)(offsetNorm)
	call := C.igGetColumnOffsetFromNorm(a0, a1)
	r := (float32)(call)
	return r
}

func GetColumnWidth(columnIndex int) float32 {
	a0 := (C.int)(columnIndex)
	call := C.igGetColumnWidth(a0)
	r := (float32)(call)
	return r
}

func GetColumnsCount() int {
	call := C.igGetColumnsCount()
	r := (int)(call)
	return r
}

func GetColumnsID(strId string, count int) ID {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.int)(count)
	call := C.igGetColumnsID(a0, a1)
	r := (ID)(call)
	return r
}

func GetContentRegionAvail(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetContentRegionAvail(a0)
}

func GetCurrentContext() *Context {
	call := C.igGetCurrentContext()
	r := (*Context)(call)
	return r
}

func GetCurrentFocusScope() ID {
	call := C.igGetCurrentFocusScope()
	r := (ID)(call)
	return r
}

func GetCurrentTabBar() *TabBar {
	call := C.igGetCurrentTabBar()
	r := (*TabBar)(call)
	return r
}

func GetCurrentTable() *Table {
	call := C.igGetCurrentTable()
	r := (*Table)(call)
	return r
}

func GetCurrentWindow() *Window {
	call := C.igGetCurrentWindow()
	r := (*Window)(call)
	return r
}

func GetCurrentWindowRead() *Window {
	call := C.igGetCurrentWindowRead()
	r := (*Window)(call)
	return r
}

func GetCursorPos(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetCursorPos(a0)
}

func GetCursorPosX() float32 {
	call := C.igGetCursorPosX()
	r := (float32)(call)
	return r
}

func GetCursorPosY() float32 {
	call := C.igGetCursorPosY()
	r := (float32)(call)
	return r
}

func GetCursorScreenPos(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetCursorScreenPos(a0)
}

func GetCursorStartPos(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetCursorStartPos(a0)
}

func GetDefaultFont() *Font {
	call := C.igGetDefaultFont()
	r := (*Font)(call)
	return r
}

func GetDragDropPayload() *Payload {
	call := C.igGetDragDropPayload()
	r := (*Payload)(call)
	return r
}

func GetDrawData() *DrawData {
	call := C.igGetDrawData()
	r := (*DrawData)(call)
	return r
}

func GetDrawListSharedData() *DrawListSharedData {
	call := C.igGetDrawListSharedData()
	r := (*DrawListSharedData)(call)
	return r
}

func GetFocusID() ID {
	call := C.igGetFocusID()
	r := (ID)(call)
	return r
}

func GetFont() *Font {
	call := C.igGetFont()
	r := (*Font)(call)
	return r
}

func GetFontSize() float32 {
	call := C.igGetFontSize()
	r := (float32)(call)
	return r
}

func GetFontTexUvWhitePixel(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetFontTexUvWhitePixel(a0)
}

func GetForegroundDrawList_ViewportPtr(viewport *Viewport) *DrawList {
	a0 := (*C.ImGuiViewport)(unsafe.Pointer(viewport))
	call := C.igGetForegroundDrawList_ViewportPtr(a0)
	r := (*DrawList)(call)
	return r
}

func GetForegroundDrawList_WindowPtr(window *Window) *DrawList {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	call := C.igGetForegroundDrawList_WindowPtr(a0)
	r := (*DrawList)(call)
	return r
}

func GetFrameCount() int {
	call := C.igGetFrameCount()
	r := (int)(call)
	return r
}

func GetFrameHeight() float32 {
	call := C.igGetFrameHeight()
	r := (float32)(call)
	return r
}

func GetFrameHeightWithSpacing() float32 {
	call := C.igGetFrameHeightWithSpacing()
	r := (float32)(call)
	return r
}

func GetHoveredID() ID {
	call := C.igGetHoveredID()
	r := (ID)(call)
	return r
}

func GetIDWithSeed_Int(n int, seed ID) ID {
	a0 := (C.int)(n)
	a1 := (C.ImGuiID)(seed)
	call := C.igGetIDWithSeed_Int(a0, a1)
	r := (ID)(call)
	return r
}

func GetIDWithSeed_Str(strIdBegin string, strIdEnd string, seed ID) ID {
	a0 := stringPool.StoreCString(strIdBegin)
	a1 := stringPool.StoreCString(strIdEnd)
	a2 := (C.ImGuiID)(seed)
	call := C.igGetIDWithSeed_Str(a0, a1, a2)
	r := (ID)(call)
	return r
}

func GetID_Int(intId int) ID {
	a0 := (C.int)(intId)
	call := C.igGetID_Int(a0)
	r := (ID)(call)
	return r
}

func GetID_Ptr(ptrId unsafe.Pointer) ID {
	a0 := (unsafe.Pointer)(ptrId)
	call := C.igGetID_Ptr(a0)
	r := (ID)(call)
	return r
}

func GetID_Str(strId string) ID {
	a0 := stringPool.StoreCString(strId)
	call := C.igGetID_Str(a0)
	r := (ID)(call)
	return r
}

func GetID_StrStr(strIdBegin string, strIdEnd string) ID {
	a0 := stringPool.StoreCString(strIdBegin)
	a1 := stringPool.StoreCString(strIdEnd)
	call := C.igGetID_StrStr(a0, a1)
	r := (ID)(call)
	return r
}

func GetIO() *IO {
	call := C.igGetIO()
	r := (*IO)(call)
	return r
}

func GetInputTextState(id ID) *InputTextState {
	a0 := (C.ImGuiID)(id)
	call := C.igGetInputTextState(a0)
	r := (*InputTextState)(call)
	return r
}

func GetItemFlags() ItemFlags {
	call := C.igGetItemFlags()
	r := (ItemFlags)(call)
	return r
}

func GetItemID() ID {
	call := C.igGetItemID()
	r := (ID)(call)
	return r
}

func GetItemRectMax(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetItemRectMax(a0)
}

func GetItemRectMin(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetItemRectMin(a0)
}

func GetItemRectSize(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetItemRectSize(a0)
}

func GetItemStatusFlags() ItemStatusFlags {
	call := C.igGetItemStatusFlags()
	r := (ItemStatusFlags)(call)
	return r
}

func GetKeyChordName(keyChord KeyChord) string {
	a0 := (C.ImGuiKeyChord)(keyChord)
	call := C.igGetKeyChordName(a0)
	r := C.GoString(call)
	return r
}

func GetKeyData_ContextPtr(ctx *Context, key Key) *KeyData {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	a1 := (C.ImGuiKey)(key)
	call := C.igGetKeyData_ContextPtr(a0, a1)
	r := (*KeyData)(call)
	return r
}

func GetKeyData_Key(key Key) *KeyData {
	a0 := (C.ImGuiKey)(key)
	call := C.igGetKeyData_Key(a0)
	r := (*KeyData)(call)
	return r
}

func GetKeyMagnitude2d(pOut *mgl32.Vec2, keyLeft Key, keyRight Key, keyUp Key, keyDown Key) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := (C.ImGuiKey)(keyLeft)
	a2 := (C.ImGuiKey)(keyRight)
	a3 := (C.ImGuiKey)(keyUp)
	a4 := (C.ImGuiKey)(keyDown)
	C.igGetKeyMagnitude2d(a0, a1, a2, a3, a4)
}

func GetKeyName(key Key) string {
	a0 := (C.ImGuiKey)(key)
	call := C.igGetKeyName(a0)
	r := C.GoString(call)
	return r
}

func GetKeyOwner(key Key) ID {
	a0 := (C.ImGuiKey)(key)
	call := C.igGetKeyOwner(a0)
	r := (ID)(call)
	return r
}

func GetKeyOwnerData(ctx *Context, key Key) *KeyOwnerData {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	a1 := (C.ImGuiKey)(key)
	call := C.igGetKeyOwnerData(a0, a1)
	r := (*KeyOwnerData)(call)
	return r
}

func GetKeyPressedAmount(key Key, repeatDelay float32, rate float32) int {
	a0 := (C.ImGuiKey)(key)
	a1 := (C.float)(repeatDelay)
	a2 := (C.float)(rate)
	call := C.igGetKeyPressedAmount(a0, a1, a2)
	r := (int)(call)
	return r
}

func GetMainViewport() *Viewport {
	call := C.igGetMainViewport()
	r := (*Viewport)(call)
	return r
}

func GetMouseClickedCount(button MouseButton) int {
	a0 := (C.ImGuiMouseButton)(button)
	call := C.igGetMouseClickedCount(a0)
	r := (int)(call)
	return r
}

func GetMouseCursor() MouseCursor {
	call := C.igGetMouseCursor()
	r := (MouseCursor)(call)
	return r
}

func GetMouseDragDelta(pOut *mgl32.Vec2, button MouseButton, lockThreshold float32) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := (C.ImGuiMouseButton)(button)
	a2 := (C.float)(lockThreshold)
	C.igGetMouseDragDelta(a0, a1, a2)
}

func GetMousePos(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetMousePos(a0)
}

func GetMousePosOnOpeningCurrentPopup(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetMousePosOnOpeningCurrentPopup(a0)
}

func GetMultiSelectState(id ID) *MultiSelectState {
	a0 := (C.ImGuiID)(id)
	call := C.igGetMultiSelectState(a0)
	r := (*MultiSelectState)(call)
	return r
}

func GetNavTweakPressedAmount(axis Axis) float32 {
	a0 := (C.ImGuiAxis)(axis)
	call := C.igGetNavTweakPressedAmount(a0)
	r := (float32)(call)
	return r
}

func GetPlatformIO() *PlatformIO {
	call := C.igGetPlatformIO()
	r := (*PlatformIO)(call)
	return r
}

func GetPopupAllowedExtentRect(pOut *Rect, window *Window) {
	a0 := (*C.ImRect)(unsafe.Pointer(pOut))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igGetPopupAllowedExtentRect(a0, a1)
}

func GetScrollMaxX() float32 {
	call := C.igGetScrollMaxX()
	r := (float32)(call)
	return r
}

func GetScrollMaxY() float32 {
	call := C.igGetScrollMaxY()
	r := (float32)(call)
	return r
}

func GetScrollX() float32 {
	call := C.igGetScrollX()
	r := (float32)(call)
	return r
}

func GetScrollY() float32 {
	call := C.igGetScrollY()
	r := (float32)(call)
	return r
}

func GetShortcutRoutingData(keyChord KeyChord) *KeyRoutingData {
	a0 := (C.ImGuiKeyChord)(keyChord)
	call := C.igGetShortcutRoutingData(a0)
	r := (*KeyRoutingData)(call)
	return r
}

func GetStateStorage() *Storage {
	call := C.igGetStateStorage()
	r := (*Storage)(call)
	return r
}

func GetStyle() *Style {
	call := C.igGetStyle()
	r := (*Style)(call)
	return r
}

func GetStyleColorName(idx Col) string {
	a0 := (C.ImGuiCol)(idx)
	call := C.igGetStyleColorName(a0)
	r := C.GoString(call)
	return r
}

func GetStyleColorVec4(idx Col) *mgl32.Vec4 {
	a0 := (C.ImGuiCol)(idx)
	call := C.igGetStyleColorVec4(a0)
	r := (*mgl32.Vec4)(unsafe.Pointer(call))
	return r
}

func GetStyleVarInfo(idx StyleVar) *DataVarInfo {
	a0 := (C.ImGuiStyleVar)(idx)
	call := C.igGetStyleVarInfo(a0)
	r := (*DataVarInfo)(call)
	return r
}

func GetTextLineHeight() float32 {
	call := C.igGetTextLineHeight()
	r := (float32)(call)
	return r
}

func GetTextLineHeightWithSpacing() float32 {
	call := C.igGetTextLineHeightWithSpacing()
	r := (float32)(call)
	return r
}

func GetTime() float64 {
	call := C.igGetTime()
	r := (float64)(call)
	return r
}

func GetTopMostAndVisiblePopupModal() *Window {
	call := C.igGetTopMostAndVisiblePopupModal()
	r := (*Window)(call)
	return r
}

func GetTopMostPopupModal() *Window {
	call := C.igGetTopMostPopupModal()
	r := (*Window)(call)
	return r
}

func GetTreeNodeToLabelSpacing() float32 {
	call := C.igGetTreeNodeToLabelSpacing()
	r := (float32)(call)
	return r
}

func GetTypematicRepeatRate(flags InputFlags, repeatDelay *float32, repeatRate *float32) {
	a0 := (C.ImGuiInputFlags)(flags)
	a1 := (*C.float)(unsafe.Pointer(repeatDelay))
	a2 := (*C.float)(unsafe.Pointer(repeatRate))
	C.igGetTypematicRepeatRate(a0, a1, a2)
}

func GetTypingSelectRequest(flags TypingSelectFlags) *TypingSelectRequest {
	a0 := (C.ImGuiTypingSelectFlags)(flags)
	call := C.igGetTypingSelectRequest(a0)
	r := (*TypingSelectRequest)(call)
	return r
}

func GetVersion() string {
	call := C.igGetVersion()
	r := C.GoString(call)
	return r
}

func GetViewportPlatformMonitor(viewport *Viewport) *PlatformMonitor {
	a0 := (*C.ImGuiViewport)(unsafe.Pointer(viewport))
	call := C.igGetViewportPlatformMonitor(a0)
	r := (*PlatformMonitor)(call)
	return r
}

func GetWindowAlwaysWantOwnTabBar(window *Window) bool {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	call := C.igGetWindowAlwaysWantOwnTabBar(a0)
	r := (bool)(call)
	return r
}

func GetWindowDockID() ID {
	call := C.igGetWindowDockID()
	r := (ID)(call)
	return r
}

func GetWindowDockNode() *DockNode {
	call := C.igGetWindowDockNode()
	r := (*DockNode)(call)
	return r
}

func GetWindowDpiScale() float32 {
	call := C.igGetWindowDpiScale()
	r := (float32)(call)
	return r
}

func GetWindowDrawList() *DrawList {
	call := C.igGetWindowDrawList()
	r := (*DrawList)(call)
	return r
}

func GetWindowHeight() float32 {
	call := C.igGetWindowHeight()
	r := (float32)(call)
	return r
}

func GetWindowPos(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetWindowPos(a0)
}

func GetWindowResizeBorderID(window *Window, dir Dir) ID {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImGuiDir)(dir)
	call := C.igGetWindowResizeBorderID(a0, a1)
	r := (ID)(call)
	return r
}

func GetWindowResizeCornerID(window *Window, n int) ID {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.int)(n)
	call := C.igGetWindowResizeCornerID(a0, a1)
	r := (ID)(call)
	return r
}

func GetWindowScrollbarID(window *Window, axis Axis) ID {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImGuiAxis)(axis)
	call := C.igGetWindowScrollbarID(a0, a1)
	r := (ID)(call)
	return r
}

func GetWindowScrollbarRect(pOut *Rect, window *Window, axis Axis) {
	a0 := (*C.ImRect)(unsafe.Pointer(pOut))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a2 := (C.ImGuiAxis)(axis)
	C.igGetWindowScrollbarRect(a0, a1, a2)
}

func GetWindowSize(pOut *mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	C.igGetWindowSize(a0)
}

func GetWindowViewport() *Viewport {
	call := C.igGetWindowViewport()
	r := (*Viewport)(call)
	return r
}

func GetWindowWidth() float32 {
	call := C.igGetWindowWidth()
	r := (float32)(call)
	return r
}

func Indent(indentW float32) {
	a0 := (C.float)(indentW)
	C.igIndent(a0)
}

func Initialize() {
	C.igInitialize()
}

func InputDouble(label string, v *float64, step float64, stepFast float64, format string, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.double)(unsafe.Pointer(v))
	a2 := (C.double)(step)
	a3 := (C.double)(stepFast)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputDouble(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func InputFloat(label string, v *float32, step float32, stepFast float32, format string, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(v))
	a2 := (C.float)(step)
	a3 := (C.float)(stepFast)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputFloat(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func InputFloat2(label string, v [2]float32, format string, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(&v[0]))
	a2 := stringPool.StoreCString(format)
	a3 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputFloat2(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func InputFloat3(label string, v *mgl32.Vec3, format string, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(v))
	a2 := stringPool.StoreCString(format)
	a3 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputFloat3(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func InputFloat4(label string, v *mgl32.Vec4, format string, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(&v[0]))
	a2 := stringPool.StoreCString(format)
	a3 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputFloat4(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func InputInt(label string, v *int, step int, stepFast int, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(v))
	a2 := (C.int)(step)
	a3 := (C.int)(stepFast)
	a4 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputInt(a0, a1, a2, a3, a4)
	r := (bool)(call)
	return r
}

func InputInt2(label string, v [2]int, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(&v[0]))
	a2 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputInt2(a0, a1, a2)
	r := (bool)(call)
	return r
}

func InputInt3(label string, v [3]int, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(&v[0]))
	a2 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputInt3(a0, a1, a2)
	r := (bool)(call)
	return r
}

func InputInt4(label string, v [4]int, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(&v[0]))
	a2 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputInt4(a0, a1, a2)
	r := (bool)(call)
	return r
}

func InputScalar(label string, dataType DataType, pData unsafe.Pointer, pStep unsafe.Pointer, pStepFast unsafe.Pointer, format string, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.ImGuiDataType)(dataType)
	a2 := (unsafe.Pointer)(pData)
	a3 := (unsafe.Pointer)(pStep)
	a4 := (unsafe.Pointer)(pStepFast)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputScalar(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func InputScalarN(label string, dataType DataType, pData unsafe.Pointer, components int, pStep unsafe.Pointer, pStepFast unsafe.Pointer, format string, flags InputTextFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.ImGuiDataType)(dataType)
	a2 := (unsafe.Pointer)(pData)
	a3 := (C.int)(components)
	a4 := (unsafe.Pointer)(pStep)
	a5 := (unsafe.Pointer)(pStepFast)
	a6 := stringPool.StoreCString(format)
	a7 := (C.ImGuiInputTextFlags)(flags)
	call := C.igInputScalarN(a0, a1, a2, a3, a4, a5, a6, a7)
	r := (bool)(call)
	return r
}

func InputText(label string, buf string, bufSize uint, flags InputTextFlags, callback InputTextCallback, userData unsafe.Pointer) bool {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(buf)
	a2 := (C.size_t)(bufSize)
	a3 := (C.ImGuiInputTextFlags)(flags)
	a4 := (C.ImGuiInputTextCallback)(callback)
	a5 := (unsafe.Pointer)(userData)
	call := C.igInputText(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func InputTextDeactivateHook(id ID) {
	a0 := (C.ImGuiID)(id)
	C.igInputTextDeactivateHook(a0)
}

func InputTextEx(label string, hint string, buf string, bufSize int, sizeArg mgl32.Vec2, flags InputTextFlags, callback InputTextCallback, userData unsafe.Pointer) bool {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(hint)
	a2 := stringPool.StoreCString(buf)
	a3 := (C.int)(bufSize)
	a4 := mglVec2ToImVec2(sizeArg)
	a5 := (C.ImGuiInputTextFlags)(flags)
	a6 := (C.ImGuiInputTextCallback)(callback)
	a7 := (unsafe.Pointer)(userData)
	call := C.igInputTextEx(a0, a1, a2, a3, a4, a5, a6, a7)
	r := (bool)(call)
	return r
}

func InputTextMultiline(label string, buf string, bufSize uint, size mgl32.Vec2, flags InputTextFlags, callback InputTextCallback, userData unsafe.Pointer) bool {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(buf)
	a2 := (C.size_t)(bufSize)
	a3 := mglVec2ToImVec2(size)
	a4 := (C.ImGuiInputTextFlags)(flags)
	a5 := (C.ImGuiInputTextCallback)(callback)
	a6 := (unsafe.Pointer)(userData)
	call := C.igInputTextMultiline(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func InputTextWithHint(label string, hint string, buf string, bufSize uint, flags InputTextFlags, callback InputTextCallback, userData unsafe.Pointer) bool {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(hint)
	a2 := stringPool.StoreCString(buf)
	a3 := (C.size_t)(bufSize)
	a4 := (C.ImGuiInputTextFlags)(flags)
	a5 := (C.ImGuiInputTextCallback)(callback)
	a6 := (unsafe.Pointer)(userData)
	call := C.igInputTextWithHint(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func InvisibleButton(strId string, size mgl32.Vec2, flags ButtonFlags) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := mglVec2ToImVec2(size)
	a2 := (C.ImGuiButtonFlags)(flags)
	call := C.igInvisibleButton(a0, a1, a2)
	r := (bool)(call)
	return r
}

func IsActiveIdUsingNavDir(dir Dir) bool {
	a0 := (C.ImGuiDir)(dir)
	call := C.igIsActiveIdUsingNavDir(a0)
	r := (bool)(call)
	return r
}

func IsAliasKey(key Key) bool {
	a0 := (C.ImGuiKey)(key)
	call := C.igIsAliasKey(a0)
	r := (bool)(call)
	return r
}

func IsAnyItemActive() bool {
	call := C.igIsAnyItemActive()
	r := (bool)(call)
	return r
}

func IsAnyItemFocused() bool {
	call := C.igIsAnyItemFocused()
	r := (bool)(call)
	return r
}

func IsAnyItemHovered() bool {
	call := C.igIsAnyItemHovered()
	r := (bool)(call)
	return r
}

func IsAnyMouseDown() bool {
	call := C.igIsAnyMouseDown()
	r := (bool)(call)
	return r
}

func IsClippedEx(bb Rect, id ID) bool {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	call := C.igIsClippedEx(a0, a1)
	r := (bool)(call)
	return r
}

func IsDragDropActive() bool {
	call := C.igIsDragDropActive()
	r := (bool)(call)
	return r
}

func IsDragDropPayloadBeingAccepted() bool {
	call := C.igIsDragDropPayloadBeingAccepted()
	r := (bool)(call)
	return r
}

func IsGamepadKey(key Key) bool {
	a0 := (C.ImGuiKey)(key)
	call := C.igIsGamepadKey(a0)
	r := (bool)(call)
	return r
}

func IsItemActivated() bool {
	call := C.igIsItemActivated()
	r := (bool)(call)
	return r
}

func IsItemActive() bool {
	call := C.igIsItemActive()
	r := (bool)(call)
	return r
}

func IsItemClicked(mouseButton MouseButton) bool {
	a0 := (C.ImGuiMouseButton)(mouseButton)
	call := C.igIsItemClicked(a0)
	r := (bool)(call)
	return r
}

func IsItemDeactivated() bool {
	call := C.igIsItemDeactivated()
	r := (bool)(call)
	return r
}

func IsItemDeactivatedAfterEdit() bool {
	call := C.igIsItemDeactivatedAfterEdit()
	r := (bool)(call)
	return r
}

func IsItemEdited() bool {
	call := C.igIsItemEdited()
	r := (bool)(call)
	return r
}

func IsItemFocused() bool {
	call := C.igIsItemFocused()
	r := (bool)(call)
	return r
}

func IsItemHovered(flags HoveredFlags) bool {
	a0 := (C.ImGuiHoveredFlags)(flags)
	call := C.igIsItemHovered(a0)
	r := (bool)(call)
	return r
}

func IsItemToggledOpen() bool {
	call := C.igIsItemToggledOpen()
	r := (bool)(call)
	return r
}

func IsItemToggledSelection() bool {
	call := C.igIsItemToggledSelection()
	r := (bool)(call)
	return r
}

func IsItemVisible() bool {
	call := C.igIsItemVisible()
	r := (bool)(call)
	return r
}

func IsKeyChordPressed_InputFlags(keyChord KeyChord, flags InputFlags, ownerId ID) bool {
	a0 := (C.ImGuiKeyChord)(keyChord)
	a1 := (C.ImGuiInputFlags)(flags)
	a2 := (C.ImGuiID)(ownerId)
	call := C.igIsKeyChordPressed_InputFlags(a0, a1, a2)
	r := (bool)(call)
	return r
}

func IsKeyChordPressed_Nil(keyChord KeyChord) bool {
	a0 := (C.ImGuiKeyChord)(keyChord)
	call := C.igIsKeyChordPressed_Nil(a0)
	r := (bool)(call)
	return r
}

func IsKeyDown_ID(key Key, ownerId ID) bool {
	a0 := (C.ImGuiKey)(key)
	a1 := (C.ImGuiID)(ownerId)
	call := C.igIsKeyDown_ID(a0, a1)
	r := (bool)(call)
	return r
}

func IsKeyDown_Nil(key Key) bool {
	a0 := (C.ImGuiKey)(key)
	call := C.igIsKeyDown_Nil(a0)
	r := (bool)(call)
	return r
}

func IsKeyPressed_Bool(key Key, repeat bool) bool {
	a0 := (C.ImGuiKey)(key)
	a1 := (C.bool)(repeat)
	call := C.igIsKeyPressed_Bool(a0, a1)
	r := (bool)(call)
	return r
}

func IsKeyPressed_InputFlags(key Key, flags InputFlags, ownerId ID) bool {
	a0 := (C.ImGuiKey)(key)
	a1 := (C.ImGuiInputFlags)(flags)
	a2 := (C.ImGuiID)(ownerId)
	call := C.igIsKeyPressed_InputFlags(a0, a1, a2)
	r := (bool)(call)
	return r
}

func IsKeyReleased_ID(key Key, ownerId ID) bool {
	a0 := (C.ImGuiKey)(key)
	a1 := (C.ImGuiID)(ownerId)
	call := C.igIsKeyReleased_ID(a0, a1)
	r := (bool)(call)
	return r
}

func IsKeyReleased_Nil(key Key) bool {
	a0 := (C.ImGuiKey)(key)
	call := C.igIsKeyReleased_Nil(a0)
	r := (bool)(call)
	return r
}

func IsKeyboardKey(key Key) bool {
	a0 := (C.ImGuiKey)(key)
	call := C.igIsKeyboardKey(a0)
	r := (bool)(call)
	return r
}

func IsLRModKey(key Key) bool {
	a0 := (C.ImGuiKey)(key)
	call := C.igIsLRModKey(a0)
	r := (bool)(call)
	return r
}

func IsLegacyKey(key Key) bool {
	a0 := (C.ImGuiKey)(key)
	call := C.igIsLegacyKey(a0)
	r := (bool)(call)
	return r
}

func IsMouseClicked_Bool(button MouseButton, repeat bool) bool {
	a0 := (C.ImGuiMouseButton)(button)
	a1 := (C.bool)(repeat)
	call := C.igIsMouseClicked_Bool(a0, a1)
	r := (bool)(call)
	return r
}

func IsMouseClicked_InputFlags(button MouseButton, flags InputFlags, ownerId ID) bool {
	a0 := (C.ImGuiMouseButton)(button)
	a1 := (C.ImGuiInputFlags)(flags)
	a2 := (C.ImGuiID)(ownerId)
	call := C.igIsMouseClicked_InputFlags(a0, a1, a2)
	r := (bool)(call)
	return r
}

func IsMouseDoubleClicked_ID(button MouseButton, ownerId ID) bool {
	a0 := (C.ImGuiMouseButton)(button)
	a1 := (C.ImGuiID)(ownerId)
	call := C.igIsMouseDoubleClicked_ID(a0, a1)
	r := (bool)(call)
	return r
}

func IsMouseDoubleClicked_Nil(button MouseButton) bool {
	a0 := (C.ImGuiMouseButton)(button)
	call := C.igIsMouseDoubleClicked_Nil(a0)
	r := (bool)(call)
	return r
}

func IsMouseDown_ID(button MouseButton, ownerId ID) bool {
	a0 := (C.ImGuiMouseButton)(button)
	a1 := (C.ImGuiID)(ownerId)
	call := C.igIsMouseDown_ID(a0, a1)
	r := (bool)(call)
	return r
}

func IsMouseDown_Nil(button MouseButton) bool {
	a0 := (C.ImGuiMouseButton)(button)
	call := C.igIsMouseDown_Nil(a0)
	r := (bool)(call)
	return r
}

func IsMouseDragPastThreshold(button MouseButton, lockThreshold float32) bool {
	a0 := (C.ImGuiMouseButton)(button)
	a1 := (C.float)(lockThreshold)
	call := C.igIsMouseDragPastThreshold(a0, a1)
	r := (bool)(call)
	return r
}

func IsMouseDragging(button MouseButton, lockThreshold float32) bool {
	a0 := (C.ImGuiMouseButton)(button)
	a1 := (C.float)(lockThreshold)
	call := C.igIsMouseDragging(a0, a1)
	r := (bool)(call)
	return r
}

func IsMouseHoveringRect(rMin mgl32.Vec2, rMax mgl32.Vec2, clip bool) bool {
	a0 := mglVec2ToImVec2(rMin)
	a1 := mglVec2ToImVec2(rMax)
	a2 := (C.bool)(clip)
	call := C.igIsMouseHoveringRect(a0, a1, a2)
	r := (bool)(call)
	return r
}

func IsMouseKey(key Key) bool {
	a0 := (C.ImGuiKey)(key)
	call := C.igIsMouseKey(a0)
	r := (bool)(call)
	return r
}

func IsMousePosValid(mousePos *mgl32.Vec2) bool {
	a0 := (*C.ImVec2)(unsafe.Pointer(&mousePos[0]))
	call := C.igIsMousePosValid(a0)
	r := (bool)(call)
	return r
}

func IsMouseReleased_ID(button MouseButton, ownerId ID) bool {
	a0 := (C.ImGuiMouseButton)(button)
	a1 := (C.ImGuiID)(ownerId)
	call := C.igIsMouseReleased_ID(a0, a1)
	r := (bool)(call)
	return r
}

func IsMouseReleased_Nil(button MouseButton) bool {
	a0 := (C.ImGuiMouseButton)(button)
	call := C.igIsMouseReleased_Nil(a0)
	r := (bool)(call)
	return r
}

func IsNamedKey(key Key) bool {
	a0 := (C.ImGuiKey)(key)
	call := C.igIsNamedKey(a0)
	r := (bool)(call)
	return r
}

func IsNamedKeyOrMod(key Key) bool {
	a0 := (C.ImGuiKey)(key)
	call := C.igIsNamedKeyOrMod(a0)
	r := (bool)(call)
	return r
}

func IsPopupOpen_ID(id ID, popupFlags PopupFlags) bool {
	a0 := (C.ImGuiID)(id)
	a1 := (C.ImGuiPopupFlags)(popupFlags)
	call := C.igIsPopupOpen_ID(a0, a1)
	r := (bool)(call)
	return r
}

func IsPopupOpen_Str(strId string, flags PopupFlags) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiPopupFlags)(flags)
	call := C.igIsPopupOpen_Str(a0, a1)
	r := (bool)(call)
	return r
}

func IsRectVisible_Nil(size mgl32.Vec2) bool {
	a0 := mglVec2ToImVec2(size)
	call := C.igIsRectVisible_Nil(a0)
	r := (bool)(call)
	return r
}

func IsRectVisible_Vec2(rectMin mgl32.Vec2, rectMax mgl32.Vec2) bool {
	a0 := mglVec2ToImVec2(rectMin)
	a1 := mglVec2ToImVec2(rectMax)
	call := C.igIsRectVisible_Vec2(a0, a1)
	r := (bool)(call)
	return r
}

func IsWindowAbove(potentialAbove *Window, potentialBelow *Window) bool {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(potentialAbove))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(potentialBelow))
	call := C.igIsWindowAbove(a0, a1)
	r := (bool)(call)
	return r
}

func IsWindowAppearing() bool {
	call := C.igIsWindowAppearing()
	r := (bool)(call)
	return r
}

func IsWindowChildOf(window *Window, potentialParent *Window, popupHierarchy bool, dockHierarchy bool) bool {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(potentialParent))
	a2 := (C.bool)(popupHierarchy)
	a3 := (C.bool)(dockHierarchy)
	call := C.igIsWindowChildOf(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func IsWindowCollapsed() bool {
	call := C.igIsWindowCollapsed()
	r := (bool)(call)
	return r
}

func IsWindowContentHoverable(window *Window, flags HoveredFlags) bool {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImGuiHoveredFlags)(flags)
	call := C.igIsWindowContentHoverable(a0, a1)
	r := (bool)(call)
	return r
}

func IsWindowDocked() bool {
	call := C.igIsWindowDocked()
	r := (bool)(call)
	return r
}

func IsWindowFocused(flags FocusedFlags) bool {
	a0 := (C.ImGuiFocusedFlags)(flags)
	call := C.igIsWindowFocused(a0)
	r := (bool)(call)
	return r
}

func IsWindowHovered(flags HoveredFlags) bool {
	a0 := (C.ImGuiHoveredFlags)(flags)
	call := C.igIsWindowHovered(a0)
	r := (bool)(call)
	return r
}

func IsWindowNavFocusable(window *Window) bool {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	call := C.igIsWindowNavFocusable(a0)
	r := (bool)(call)
	return r
}

func IsWindowWithinBeginStackOf(window *Window, potentialParent *Window) bool {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(potentialParent))
	call := C.igIsWindowWithinBeginStackOf(a0, a1)
	r := (bool)(call)
	return r
}

func ItemAdd(bb Rect, id ID, navBb *Rect, extraFlags ItemFlags) bool {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	a2 := (*C.ImRect)(unsafe.Pointer(navBb))
	a3 := (C.ImGuiItemFlags)(extraFlags)
	call := C.igItemAdd(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func ItemHoverable(bb Rect, id ID, itemFlags ItemFlags) bool {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	a2 := (C.ImGuiItemFlags)(itemFlags)
	call := C.igItemHoverable(a0, a1, a2)
	r := (bool)(call)
	return r
}

func ItemSize_Rect(bb Rect, textBaselineY float32) {
	a0 := (C.ImRect)(bb)
	a1 := (C.float)(textBaselineY)
	C.igItemSize_Rect(a0, a1)
}

func ItemSize_Vec2(size mgl32.Vec2, textBaselineY float32) {
	a0 := mglVec2ToImVec2(size)
	a1 := (C.float)(textBaselineY)
	C.igItemSize_Vec2(a0, a1)
}

func KeepAliveID(id ID) {
	a0 := (C.ImGuiID)(id)
	C.igKeepAliveID(a0)
}

func LabelText(label string, vfmt string, vargs ...interface{}) {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	C.wrap_igLabelText(a0, a1)
}

func ListBox_Str_arr(label string, currentItem *int, items []string, itemsCount int, heightInItems int) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(currentItem))
	a2 := (**C.char)(unsafe.Pointer(&items[0]))
	a3 := (C.int)(itemsCount)
	a4 := (C.int)(heightInItems)
	call := C.igListBox_Str_arr(a0, a1, a2, a3, a4)
	r := (bool)(call)
	return r
}

func LoadIniSettingsFromDisk(iniFilename string) {
	a0 := stringPool.StoreCString(iniFilename)
	C.igLoadIniSettingsFromDisk(a0)
}

func LoadIniSettingsFromMemory(iniData string, iniSize uint) {
	a0 := stringPool.StoreCString(iniData)
	a1 := (C.size_t)(iniSize)
	C.igLoadIniSettingsFromMemory(a0, a1)
}

func LocalizeGetMsg(key LocKey) string {
	a0 := (C.ImGuiLocKey)(key)
	call := C.igLocalizeGetMsg(a0)
	r := C.GoString(call)
	return r
}

func LocalizeRegisterEntries(entries *LocEntry, count int) {
	a0 := (*C.ImGuiLocEntry)(unsafe.Pointer(entries))
	a1 := (C.int)(count)
	C.igLocalizeRegisterEntries(a0, a1)
}

func LogBegin(ty LogType, autoOpenDepth int) {
	a0 := (C.ImGuiLogType)(ty)
	a1 := (C.int)(autoOpenDepth)
	C.igLogBegin(a0, a1)
}

func LogButtons() {
	C.igLogButtons()
}

func LogFinish() {
	C.igLogFinish()
}

func LogRenderedText(refPos *mgl32.Vec2, text string, textEnd string) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&refPos[0]))
	a1 := stringPool.StoreCString(text)
	a2 := stringPool.StoreCString(textEnd)
	C.igLogRenderedText(a0, a1, a2)
}

func LogSetNextTextDecoration(prefix string, suffix string) {
	a0 := stringPool.StoreCString(prefix)
	a1 := stringPool.StoreCString(suffix)
	C.igLogSetNextTextDecoration(a0, a1)
}

func LogText(vfmt string, vargs ...interface{}) {
	a0 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	C.wrap_igLogText(a0)
}

func LogToBuffer(autoOpenDepth int) {
	a0 := (C.int)(autoOpenDepth)
	C.igLogToBuffer(a0)
}

func LogToClipboard(autoOpenDepth int) {
	a0 := (C.int)(autoOpenDepth)
	C.igLogToClipboard(a0)
}

func LogToFile(autoOpenDepth int, filename string) {
	a0 := (C.int)(autoOpenDepth)
	a1 := stringPool.StoreCString(filename)
	C.igLogToFile(a0, a1)
}

func LogToTTY(autoOpenDepth int) {
	a0 := (C.int)(autoOpenDepth)
	C.igLogToTTY(a0)
}

func MarkIniSettingsDirty_Nil() {
	C.igMarkIniSettingsDirty_Nil()
}

func MarkIniSettingsDirty_WindowPtr(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igMarkIniSettingsDirty_WindowPtr(a0)
}

func MarkItemEdited(id ID) {
	a0 := (C.ImGuiID)(id)
	C.igMarkItemEdited(a0)
}

func MemAlloc(size uint) unsafe.Pointer {
	a0 := (C.size_t)(size)
	call := C.igMemAlloc(a0)
	r := (unsafe.Pointer)(call)
	return r
}

func MemFree(ptr unsafe.Pointer) {
	a0 := (unsafe.Pointer)(ptr)
	C.igMemFree(a0)
}

func MenuItemEx(label string, icon string, shortcut string, selected bool, enabled bool) bool {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(icon)
	a2 := stringPool.StoreCString(shortcut)
	a3 := (C.bool)(selected)
	a4 := (C.bool)(enabled)
	call := C.igMenuItemEx(a0, a1, a2, a3, a4)
	r := (bool)(call)
	return r
}

func MenuItem_Bool(label string, shortcut string, selected bool, enabled bool) bool {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(shortcut)
	a2 := (C.bool)(selected)
	a3 := (C.bool)(enabled)
	call := C.igMenuItem_Bool(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func MenuItem_BoolPtr(label string, shortcut string, pSelected *bool, enabled bool) bool {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(shortcut)
	a2 := (*C.bool)(unsafe.Pointer(pSelected))
	a3 := (C.bool)(enabled)
	call := C.igMenuItem_BoolPtr(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func MouseButtonToKey(button MouseButton) Key {
	a0 := (C.ImGuiMouseButton)(button)
	call := C.igMouseButtonToKey(a0)
	r := (Key)(call)
	return r
}

func MultiSelectAddSetAll(ms *MultiSelectTempData, selected bool) {
	a0 := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(ms))
	a1 := (C.bool)(selected)
	C.igMultiSelectAddSetAll(a0, a1)
}

func MultiSelectAddSetRange(ms *MultiSelectTempData, selected bool, rangeDir int, firstItem SelectionUserData, lastItem SelectionUserData) {
	a0 := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(ms))
	a1 := (C.bool)(selected)
	a2 := (C.int)(rangeDir)
	a3 := (C.ImGuiSelectionUserData)(firstItem)
	a4 := (C.ImGuiSelectionUserData)(lastItem)
	C.igMultiSelectAddSetRange(a0, a1, a2, a3, a4)
}

func MultiSelectItemFooter(id ID, pSelected *bool, pPressed *bool) {
	a0 := (C.ImGuiID)(id)
	a1 := (*C.bool)(unsafe.Pointer(pSelected))
	a2 := (*C.bool)(unsafe.Pointer(pPressed))
	C.igMultiSelectItemFooter(a0, a1, a2)
}

func MultiSelectItemHeader(id ID, pSelected *bool, pButtonFlags *ButtonFlags) {
	a0 := (C.ImGuiID)(id)
	a1 := (*C.bool)(unsafe.Pointer(pSelected))
	a2 := (*C.ImGuiButtonFlags)(unsafe.Pointer(pButtonFlags))
	C.igMultiSelectItemHeader(a0, a1, a2)
}

func NavClearPreferredPosForAxis(axis Axis) {
	a0 := (C.ImGuiAxis)(axis)
	C.igNavClearPreferredPosForAxis(a0)
}

func NavHighlightActivated(id ID) {
	a0 := (C.ImGuiID)(id)
	C.igNavHighlightActivated(a0)
}

func NavInitRequestApplyResult() {
	C.igNavInitRequestApplyResult()
}

func NavInitWindow(window *Window, forceReinit bool) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.bool)(forceReinit)
	C.igNavInitWindow(a0, a1)
}

func NavMoveRequestApplyResult() {
	C.igNavMoveRequestApplyResult()
}

func NavMoveRequestButNoResultYet() bool {
	call := C.igNavMoveRequestButNoResultYet()
	r := (bool)(call)
	return r
}

func NavMoveRequestCancel() {
	C.igNavMoveRequestCancel()
}

func NavMoveRequestForward(moveDir Dir, clipDir Dir, moveFlags NavMoveFlags, scrollFlags ScrollFlags) {
	a0 := (C.ImGuiDir)(moveDir)
	a1 := (C.ImGuiDir)(clipDir)
	a2 := (C.ImGuiNavMoveFlags)(moveFlags)
	a3 := (C.ImGuiScrollFlags)(scrollFlags)
	C.igNavMoveRequestForward(a0, a1, a2, a3)
}

func NavMoveRequestResolveWithLastItem(result *NavItemData) {
	a0 := (*C.ImGuiNavItemData)(unsafe.Pointer(result))
	C.igNavMoveRequestResolveWithLastItem(a0)
}

func NavMoveRequestResolveWithPastTreeNode(result *NavItemData, treeNodeData *TreeNodeStackData) {
	a0 := (*C.ImGuiNavItemData)(unsafe.Pointer(result))
	a1 := (*C.ImGuiTreeNodeStackData)(unsafe.Pointer(treeNodeData))
	C.igNavMoveRequestResolveWithPastTreeNode(a0, a1)
}

func NavMoveRequestSubmit(moveDir Dir, clipDir Dir, moveFlags NavMoveFlags, scrollFlags ScrollFlags) {
	a0 := (C.ImGuiDir)(moveDir)
	a1 := (C.ImGuiDir)(clipDir)
	a2 := (C.ImGuiNavMoveFlags)(moveFlags)
	a3 := (C.ImGuiScrollFlags)(scrollFlags)
	C.igNavMoveRequestSubmit(a0, a1, a2, a3)
}

func NavMoveRequestTryWrapping(window *Window, moveFlags NavMoveFlags) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImGuiNavMoveFlags)(moveFlags)
	C.igNavMoveRequestTryWrapping(a0, a1)
}

func NavUpdateCurrentWindowIsScrollPushableX() {
	C.igNavUpdateCurrentWindowIsScrollPushableX()
}

func NewLine() {
	C.igNewLine()
}

func NextColumn() {
	C.igNextColumn()
}

func OpenPopupEx(id ID, popupFlags PopupFlags) {
	a0 := (C.ImGuiID)(id)
	a1 := (C.ImGuiPopupFlags)(popupFlags)
	C.igOpenPopupEx(a0, a1)
}

func OpenPopupOnItemClick(strId string, popupFlags PopupFlags) {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiPopupFlags)(popupFlags)
	C.igOpenPopupOnItemClick(a0, a1)
}

func OpenPopup_ID(id ID, popupFlags PopupFlags) {
	a0 := (C.ImGuiID)(id)
	a1 := (C.ImGuiPopupFlags)(popupFlags)
	C.igOpenPopup_ID(a0, a1)
}

func OpenPopup_Str(strId string, popupFlags PopupFlags) {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiPopupFlags)(popupFlags)
	C.igOpenPopup_Str(a0, a1)
}

func PlotHistogram_FloatPtr(label string, values *float32, valuesCount int, valuesOffset int, overlayText string, scaleMin float32, scaleMax float32, graphSize mgl32.Vec2, stride int) {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(values))
	a2 := (C.int)(valuesCount)
	a3 := (C.int)(valuesOffset)
	a4 := stringPool.StoreCString(overlayText)
	a5 := (C.float)(scaleMin)
	a6 := (C.float)(scaleMax)
	a7 := mglVec2ToImVec2(graphSize)
	a8 := (C.int)(stride)
	C.igPlotHistogram_FloatPtr(a0, a1, a2, a3, a4, a5, a6, a7, a8)
}

func PlotLines_FloatPtr(label string, values *float32, valuesCount int, valuesOffset int, overlayText string, scaleMin float32, scaleMax float32, graphSize mgl32.Vec2, stride int) {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(values))
	a2 := (C.int)(valuesCount)
	a3 := (C.int)(valuesOffset)
	a4 := stringPool.StoreCString(overlayText)
	a5 := (C.float)(scaleMin)
	a6 := (C.float)(scaleMax)
	a7 := mglVec2ToImVec2(graphSize)
	a8 := (C.int)(stride)
	C.igPlotLines_FloatPtr(a0, a1, a2, a3, a4, a5, a6, a7, a8)
}

func PopClipRect() {
	C.igPopClipRect()
}

func PopColumnsBackground() {
	C.igPopColumnsBackground()
}

func PopFocusScope() {
	C.igPopFocusScope()
}

func PopFont() {
	C.igPopFont()
}

func PopID() {
	C.igPopID()
}

func PopItemFlag() {
	C.igPopItemFlag()
}

func PopItemWidth() {
	C.igPopItemWidth()
}

func PopStyleColor(count int) {
	a0 := (C.int)(count)
	C.igPopStyleColor(a0)
}

func PopStyleVar(count int) {
	a0 := (C.int)(count)
	C.igPopStyleVar(a0)
}

func PopTextWrapPos() {
	C.igPopTextWrapPos()
}

func ProgressBar(fraction float32, sizeArg mgl32.Vec2, overlay string) {
	a0 := (C.float)(fraction)
	a1 := mglVec2ToImVec2(sizeArg)
	a2 := stringPool.StoreCString(overlay)
	C.igProgressBar(a0, a1, a2)
}

func PushClipRect(clipRectMin mgl32.Vec2, clipRectMax mgl32.Vec2, intersectWithCurrentClipRect bool) {
	a0 := mglVec2ToImVec2(clipRectMin)
	a1 := mglVec2ToImVec2(clipRectMax)
	a2 := (C.bool)(intersectWithCurrentClipRect)
	C.igPushClipRect(a0, a1, a2)
}

func PushColumnClipRect(columnIndex int) {
	a0 := (C.int)(columnIndex)
	C.igPushColumnClipRect(a0)
}

func PushColumnsBackground() {
	C.igPushColumnsBackground()
}

func PushFocusScope(id ID) {
	a0 := (C.ImGuiID)(id)
	C.igPushFocusScope(a0)
}

func PushFont(font *Font) {
	a0 := (*C.ImFont)(unsafe.Pointer(font))
	C.igPushFont(a0)
}

func PushID_Int(intId int) {
	a0 := (C.int)(intId)
	C.igPushID_Int(a0)
}

func PushID_Ptr(ptrId unsafe.Pointer) {
	a0 := (unsafe.Pointer)(ptrId)
	C.igPushID_Ptr(a0)
}

func PushID_Str(strId string) {
	a0 := stringPool.StoreCString(strId)
	C.igPushID_Str(a0)
}

func PushID_StrStr(strIdBegin string, strIdEnd string) {
	a0 := stringPool.StoreCString(strIdBegin)
	a1 := stringPool.StoreCString(strIdEnd)
	C.igPushID_StrStr(a0, a1)
}

func PushItemFlag(option ItemFlags, enabled bool) {
	a0 := (C.ImGuiItemFlags)(option)
	a1 := (C.bool)(enabled)
	C.igPushItemFlag(a0, a1)
}

func PushItemWidth(itemWidth float32) {
	a0 := (C.float)(itemWidth)
	C.igPushItemWidth(a0)
}

func PushMultiItemsWidths(components int, widthFull float32) {
	a0 := (C.int)(components)
	a1 := (C.float)(widthFull)
	C.igPushMultiItemsWidths(a0, a1)
}

func PushOverrideID(id ID) {
	a0 := (C.ImGuiID)(id)
	C.igPushOverrideID(a0)
}

func PushStyleColor_U32(idx Col, col U32) {
	a0 := (C.ImGuiCol)(idx)
	a1 := (C.ImU32)(col)
	C.igPushStyleColor_U32(a0, a1)
}

func PushStyleColor_Vec4(idx Col, col mgl32.Vec4) {
	a0 := (C.ImGuiCol)(idx)
	a1 := mglVec4ToImVec4(col)
	C.igPushStyleColor_Vec4(a0, a1)
}

func PushStyleVarX(idx StyleVar, valX float32) {
	a0 := (C.ImGuiStyleVar)(idx)
	a1 := (C.float)(valX)
	C.igPushStyleVarX(a0, a1)
}

func PushStyleVarY(idx StyleVar, valY float32) {
	a0 := (C.ImGuiStyleVar)(idx)
	a1 := (C.float)(valY)
	C.igPushStyleVarY(a0, a1)
}

func PushStyleVar_Float(idx StyleVar, val float32) {
	a0 := (C.ImGuiStyleVar)(idx)
	a1 := (C.float)(val)
	C.igPushStyleVar_Float(a0, a1)
}

func PushStyleVar_Vec2(idx StyleVar, val mgl32.Vec2) {
	a0 := (C.ImGuiStyleVar)(idx)
	a1 := mglVec2ToImVec2(val)
	C.igPushStyleVar_Vec2(a0, a1)
}

func PushTextWrapPos(wrapLocalPosX float32) {
	a0 := (C.float)(wrapLocalPosX)
	C.igPushTextWrapPos(a0)
}

func RadioButton_Bool(label string, active bool) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.bool)(active)
	call := C.igRadioButton_Bool(a0, a1)
	r := (bool)(call)
	return r
}

func RadioButton_IntPtr(label string, v *int, vButton int) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(v))
	a2 := (C.int)(vButton)
	call := C.igRadioButton_IntPtr(a0, a1, a2)
	r := (bool)(call)
	return r
}

func RemoveContextHook(context *Context, hookToRemove ID) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(context))
	a1 := (C.ImGuiID)(hookToRemove)
	C.igRemoveContextHook(a0, a1)
}

func RemoveSettingsHandler(typeName string) {
	a0 := stringPool.StoreCString(typeName)
	C.igRemoveSettingsHandler(a0)
}

func Render() {
	C.igRender()
}

func RenderArrow(drawList *DrawList, pos mgl32.Vec2, col U32, dir Dir, scale float32) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := mglVec2ToImVec2(pos)
	a2 := (C.ImU32)(col)
	a3 := (C.ImGuiDir)(dir)
	a4 := (C.float)(scale)
	C.igRenderArrow(a0, a1, a2, a3, a4)
}

func RenderArrowDockMenu(drawList *DrawList, pMin mgl32.Vec2, sz float32, col U32) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := mglVec2ToImVec2(pMin)
	a2 := (C.float)(sz)
	a3 := (C.ImU32)(col)
	C.igRenderArrowDockMenu(a0, a1, a2, a3)
}

func RenderArrowPointingAt(drawList *DrawList, pos mgl32.Vec2, halfSz mgl32.Vec2, direction Dir, col U32) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := mglVec2ToImVec2(pos)
	a2 := mglVec2ToImVec2(halfSz)
	a3 := (C.ImGuiDir)(direction)
	a4 := (C.ImU32)(col)
	C.igRenderArrowPointingAt(a0, a1, a2, a3, a4)
}

func RenderBullet(drawList *DrawList, pos mgl32.Vec2, col U32) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := mglVec2ToImVec2(pos)
	a2 := (C.ImU32)(col)
	C.igRenderBullet(a0, a1, a2)
}

func RenderCheckMark(drawList *DrawList, pos mgl32.Vec2, col U32, sz float32) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := mglVec2ToImVec2(pos)
	a2 := (C.ImU32)(col)
	a3 := (C.float)(sz)
	C.igRenderCheckMark(a0, a1, a2, a3)
}

func RenderColorRectWithAlphaCheckerboard(drawList *DrawList, pMin mgl32.Vec2, pMax mgl32.Vec2, fillCol U32, gridStep float32, gridOff mgl32.Vec2, rounding float32, flags DrawFlags) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := mglVec2ToImVec2(pMin)
	a2 := mglVec2ToImVec2(pMax)
	a3 := (C.ImU32)(fillCol)
	a4 := (C.float)(gridStep)
	a5 := mglVec2ToImVec2(gridOff)
	a6 := (C.float)(rounding)
	a7 := (C.ImDrawFlags)(flags)
	C.igRenderColorRectWithAlphaCheckerboard(a0, a1, a2, a3, a4, a5, a6, a7)
}

func RenderDragDropTargetRect(bb Rect, itemClipRect Rect) {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImRect)(itemClipRect)
	C.igRenderDragDropTargetRect(a0, a1)
}

func RenderFrame(pMin mgl32.Vec2, pMax mgl32.Vec2, fillCol U32, borders bool, rounding float32) {
	a0 := mglVec2ToImVec2(pMin)
	a1 := mglVec2ToImVec2(pMax)
	a2 := (C.ImU32)(fillCol)
	a3 := (C.bool)(borders)
	a4 := (C.float)(rounding)
	C.igRenderFrame(a0, a1, a2, a3, a4)
}

func RenderFrameBorder(pMin mgl32.Vec2, pMax mgl32.Vec2, rounding float32) {
	a0 := mglVec2ToImVec2(pMin)
	a1 := mglVec2ToImVec2(pMax)
	a2 := (C.float)(rounding)
	C.igRenderFrameBorder(a0, a1, a2)
}

func RenderMouseCursor(pos mgl32.Vec2, scale float32, mouseCursor MouseCursor, colFill U32, colBorder U32, colShadow U32) {
	a0 := mglVec2ToImVec2(pos)
	a1 := (C.float)(scale)
	a2 := (C.ImGuiMouseCursor)(mouseCursor)
	a3 := (C.ImU32)(colFill)
	a4 := (C.ImU32)(colBorder)
	a5 := (C.ImU32)(colShadow)
	C.igRenderMouseCursor(a0, a1, a2, a3, a4, a5)
}

func RenderNavCursor(bb Rect, id ID, flags NavRenderCursorFlags) {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	a2 := (C.ImGuiNavRenderCursorFlags)(flags)
	C.igRenderNavCursor(a0, a1, a2)
}

func RenderPlatformWindowsDefault(platformRenderArg unsafe.Pointer, rendererRenderArg unsafe.Pointer) {
	a0 := (unsafe.Pointer)(platformRenderArg)
	a1 := (unsafe.Pointer)(rendererRenderArg)
	C.igRenderPlatformWindowsDefault(a0, a1)
}

func RenderRectFilledRangeH(drawList *DrawList, rect Rect, col U32, xStartNorm float32, xEndNorm float32, rounding float32) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := (C.ImRect)(rect)
	a2 := (C.ImU32)(col)
	a3 := (C.float)(xStartNorm)
	a4 := (C.float)(xEndNorm)
	a5 := (C.float)(rounding)
	C.igRenderRectFilledRangeH(a0, a1, a2, a3, a4, a5)
}

func RenderRectFilledWithHole(drawList *DrawList, outer Rect, inner Rect, col U32, rounding float32) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := (C.ImRect)(outer)
	a2 := (C.ImRect)(inner)
	a3 := (C.ImU32)(col)
	a4 := (C.float)(rounding)
	C.igRenderRectFilledWithHole(a0, a1, a2, a3, a4)
}

func RenderText(pos mgl32.Vec2, text string, textEnd string, hideTextAfterHash bool) {
	a0 := mglVec2ToImVec2(pos)
	a1 := stringPool.StoreCString(text)
	a2 := stringPool.StoreCString(textEnd)
	a3 := (C.bool)(hideTextAfterHash)
	C.igRenderText(a0, a1, a2, a3)
}

func RenderTextClipped(posMin mgl32.Vec2, posMax mgl32.Vec2, text string, textEnd string, textSizeIfKnown *mgl32.Vec2, align mgl32.Vec2, clipRect *Rect) {
	a0 := mglVec2ToImVec2(posMin)
	a1 := mglVec2ToImVec2(posMax)
	a2 := stringPool.StoreCString(text)
	a3 := stringPool.StoreCString(textEnd)
	a4 := (*C.ImVec2)(unsafe.Pointer(&textSizeIfKnown[0]))
	a5 := mglVec2ToImVec2(align)
	a6 := (*C.ImRect)(unsafe.Pointer(clipRect))
	C.igRenderTextClipped(a0, a1, a2, a3, a4, a5, a6)
}

func RenderTextClippedEx(drawList *DrawList, posMin mgl32.Vec2, posMax mgl32.Vec2, text string, textEnd string, textSizeIfKnown *mgl32.Vec2, align mgl32.Vec2, clipRect *Rect) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := mglVec2ToImVec2(posMin)
	a2 := mglVec2ToImVec2(posMax)
	a3 := stringPool.StoreCString(text)
	a4 := stringPool.StoreCString(textEnd)
	a5 := (*C.ImVec2)(unsafe.Pointer(&textSizeIfKnown[0]))
	a6 := mglVec2ToImVec2(align)
	a7 := (*C.ImRect)(unsafe.Pointer(clipRect))
	C.igRenderTextClippedEx(a0, a1, a2, a3, a4, a5, a6, a7)
}

func RenderTextEllipsis(drawList *DrawList, posMin mgl32.Vec2, posMax mgl32.Vec2, clipMaxX float32, ellipsisMaxX float32, text string, textEnd string, textSizeIfKnown *mgl32.Vec2) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := mglVec2ToImVec2(posMin)
	a2 := mglVec2ToImVec2(posMax)
	a3 := (C.float)(clipMaxX)
	a4 := (C.float)(ellipsisMaxX)
	a5 := stringPool.StoreCString(text)
	a6 := stringPool.StoreCString(textEnd)
	a7 := (*C.ImVec2)(unsafe.Pointer(&textSizeIfKnown[0]))
	C.igRenderTextEllipsis(a0, a1, a2, a3, a4, a5, a6, a7)
}

func RenderTextWrapped(pos mgl32.Vec2, text string, textEnd string, wrapWidth float32) {
	a0 := mglVec2ToImVec2(pos)
	a1 := stringPool.StoreCString(text)
	a2 := stringPool.StoreCString(textEnd)
	a3 := (C.float)(wrapWidth)
	C.igRenderTextWrapped(a0, a1, a2, a3)
}

func ResetMouseDragDelta(button MouseButton) {
	a0 := (C.ImGuiMouseButton)(button)
	C.igResetMouseDragDelta(a0)
}

func SameLine(offsetFromStartX float32, spacing float32) {
	a0 := (C.float)(offsetFromStartX)
	a1 := (C.float)(spacing)
	C.igSameLine(a0, a1)
}

func SaveIniSettingsToDisk(iniFilename string) {
	a0 := stringPool.StoreCString(iniFilename)
	C.igSaveIniSettingsToDisk(a0)
}

func SaveIniSettingsToMemory(outIniSize *uint) string {
	a0 := (*C.size_t)(unsafe.Pointer(outIniSize))
	call := C.igSaveIniSettingsToMemory(a0)
	r := C.GoString(call)
	return r
}

func ScaleWindowsInViewport(viewport *ViewportP, scale float32) {
	a0 := (*C.ImGuiViewportP)(unsafe.Pointer(viewport))
	a1 := (C.float)(scale)
	C.igScaleWindowsInViewport(a0, a1)
}

func ScrollToBringRectIntoView(window *Window, rect Rect) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImRect)(rect)
	C.igScrollToBringRectIntoView(a0, a1)
}

func ScrollToItem(flags ScrollFlags) {
	a0 := (C.ImGuiScrollFlags)(flags)
	C.igScrollToItem(a0)
}

func ScrollToRect(window *Window, rect Rect, flags ScrollFlags) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImRect)(rect)
	a2 := (C.ImGuiScrollFlags)(flags)
	C.igScrollToRect(a0, a1, a2)
}

func ScrollToRectEx(pOut *mgl32.Vec2, window *Window, rect Rect, flags ScrollFlags) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a2 := (C.ImRect)(rect)
	a3 := (C.ImGuiScrollFlags)(flags)
	C.igScrollToRectEx(a0, a1, a2, a3)
}

func Scrollbar(axis Axis) {
	a0 := (C.ImGuiAxis)(axis)
	C.igScrollbar(a0)
}

func ScrollbarEx(bb Rect, id ID, axis Axis, pScrollV *S64, availV S64, contentsV S64, flags DrawFlags) bool {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	a2 := (C.ImGuiAxis)(axis)
	a3 := (*C.ImS64)(unsafe.Pointer(pScrollV))
	a4 := (C.ImS64)(availV)
	a5 := (C.ImS64)(contentsV)
	a6 := (C.ImDrawFlags)(flags)
	call := C.igScrollbarEx(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func Selectable_Bool(label string, selected bool, flags SelectableFlags, size mgl32.Vec2) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.bool)(selected)
	a2 := (C.ImGuiSelectableFlags)(flags)
	a3 := mglVec2ToImVec2(size)
	call := C.igSelectable_Bool(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func Selectable_BoolPtr(label string, pSelected *bool, flags SelectableFlags, size mgl32.Vec2) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.bool)(unsafe.Pointer(pSelected))
	a2 := (C.ImGuiSelectableFlags)(flags)
	a3 := mglVec2ToImVec2(size)
	call := C.igSelectable_BoolPtr(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func Separator() {
	C.igSeparator()
}

func SeparatorEx(flags SeparatorFlags, thickness float32) {
	a0 := (C.ImGuiSeparatorFlags)(flags)
	a1 := (C.float)(thickness)
	C.igSeparatorEx(a0, a1)
}

func SeparatorText(label string) {
	a0 := stringPool.StoreCString(label)
	C.igSeparatorText(a0)
}

func SeparatorTextEx(id ID, label string, labelEnd string, extraWidth float32) {
	a0 := (C.ImGuiID)(id)
	a1 := stringPool.StoreCString(label)
	a2 := stringPool.StoreCString(labelEnd)
	a3 := (C.float)(extraWidth)
	C.igSeparatorTextEx(a0, a1, a2, a3)
}

func SetActiveID(id ID, window *Window) {
	a0 := (C.ImGuiID)(id)
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igSetActiveID(a0, a1)
}

func SetActiveIdUsingAllKeyboardKeys() {
	C.igSetActiveIdUsingAllKeyboardKeys()
}

func SetAllocatorFunctions(allocFunc MemAllocFunc, freeFunc MemFreeFunc, userData unsafe.Pointer) {
	a0 := (C.ImGuiMemAllocFunc)(allocFunc)
	a1 := (C.ImGuiMemFreeFunc)(freeFunc)
	a2 := (unsafe.Pointer)(userData)
	C.igSetAllocatorFunctions(a0, a1, a2)
}

func SetClipboardText(text string) {
	a0 := stringPool.StoreCString(text)
	C.igSetClipboardText(a0)
}

func SetColorEditOptions(flags ColorEditFlags) {
	a0 := (C.ImGuiColorEditFlags)(flags)
	C.igSetColorEditOptions(a0)
}

func SetColumnOffset(columnIndex int, offsetX float32) {
	a0 := (C.int)(columnIndex)
	a1 := (C.float)(offsetX)
	C.igSetColumnOffset(a0, a1)
}

func SetColumnWidth(columnIndex int, width float32) {
	a0 := (C.int)(columnIndex)
	a1 := (C.float)(width)
	C.igSetColumnWidth(a0, a1)
}

func SetCurrentContext(ctx *Context) {
	a0 := (*C.ImGuiContext)(unsafe.Pointer(ctx))
	C.igSetCurrentContext(a0)
}

func SetCurrentFont(font *Font) {
	a0 := (*C.ImFont)(unsafe.Pointer(font))
	C.igSetCurrentFont(a0)
}

func SetCurrentViewport(window *Window, viewport *ViewportP) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (*C.ImGuiViewportP)(unsafe.Pointer(viewport))
	C.igSetCurrentViewport(a0, a1)
}

func SetCursorPos(localPos mgl32.Vec2) {
	a0 := mglVec2ToImVec2(localPos)
	C.igSetCursorPos(a0)
}

func SetCursorPosX(localX float32) {
	a0 := (C.float)(localX)
	C.igSetCursorPosX(a0)
}

func SetCursorPosY(localY float32) {
	a0 := (C.float)(localY)
	C.igSetCursorPosY(a0)
}

func SetCursorScreenPos(pos mgl32.Vec2) {
	a0 := mglVec2ToImVec2(pos)
	C.igSetCursorScreenPos(a0)
}

func SetDragDropPayload(ty string, data unsafe.Pointer, sz uint, cond Cond) bool {
	a0 := stringPool.StoreCString(ty)
	a1 := (unsafe.Pointer)(data)
	a2 := (C.size_t)(sz)
	a3 := (C.ImGuiCond)(cond)
	call := C.igSetDragDropPayload(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func SetFocusID(id ID, window *Window) {
	a0 := (C.ImGuiID)(id)
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igSetFocusID(a0, a1)
}

func SetHoveredID(id ID) {
	a0 := (C.ImGuiID)(id)
	C.igSetHoveredID(a0)
}

func SetItemDefaultFocus() {
	C.igSetItemDefaultFocus()
}

func SetItemKeyOwner_InputFlags(key Key, flags InputFlags) {
	a0 := (C.ImGuiKey)(key)
	a1 := (C.ImGuiInputFlags)(flags)
	C.igSetItemKeyOwner_InputFlags(a0, a1)
}

func SetItemKeyOwner_Nil(key Key) {
	a0 := (C.ImGuiKey)(key)
	C.igSetItemKeyOwner_Nil(a0)
}

func SetItemTooltip(vfmt string, vargs ...interface{}) {
	a0 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	C.wrap_igSetItemTooltip(a0)
}

func SetKeyOwner(key Key, ownerId ID, flags InputFlags) {
	a0 := (C.ImGuiKey)(key)
	a1 := (C.ImGuiID)(ownerId)
	a2 := (C.ImGuiInputFlags)(flags)
	C.igSetKeyOwner(a0, a1, a2)
}

func SetKeyOwnersForKeyChord(key KeyChord, ownerId ID, flags InputFlags) {
	a0 := (C.ImGuiKeyChord)(key)
	a1 := (C.ImGuiID)(ownerId)
	a2 := (C.ImGuiInputFlags)(flags)
	C.igSetKeyOwnersForKeyChord(a0, a1, a2)
}

func SetKeyboardFocusHere(offset int) {
	a0 := (C.int)(offset)
	C.igSetKeyboardFocusHere(a0)
}

func SetLastItemData(itemId ID, inFlags ItemFlags, statusFlags ItemStatusFlags, itemRect Rect) {
	a0 := (C.ImGuiID)(itemId)
	a1 := (C.ImGuiItemFlags)(inFlags)
	a2 := (C.ImGuiItemStatusFlags)(statusFlags)
	a3 := (C.ImRect)(itemRect)
	C.igSetLastItemData(a0, a1, a2, a3)
}

func SetMouseCursor(cursorType MouseCursor) {
	a0 := (C.ImGuiMouseCursor)(cursorType)
	C.igSetMouseCursor(a0)
}

func SetNavCursorVisible(visible bool) {
	a0 := (C.bool)(visible)
	C.igSetNavCursorVisible(a0)
}

func SetNavCursorVisibleAfterMove() {
	C.igSetNavCursorVisibleAfterMove()
}

func SetNavFocusScope(focusScopeId ID) {
	a0 := (C.ImGuiID)(focusScopeId)
	C.igSetNavFocusScope(a0)
}

func SetNavID(id ID, navLayer NavLayer, focusScopeId ID, rectRel Rect) {
	a0 := (C.ImGuiID)(id)
	a1 := (C.ImGuiNavLayer)(navLayer)
	a2 := (C.ImGuiID)(focusScopeId)
	a3 := (C.ImRect)(rectRel)
	C.igSetNavID(a0, a1, a2, a3)
}

func SetNavWindow(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igSetNavWindow(a0)
}

func SetNextFrameWantCaptureKeyboard(wantCaptureKeyboard bool) {
	a0 := (C.bool)(wantCaptureKeyboard)
	C.igSetNextFrameWantCaptureKeyboard(a0)
}

func SetNextFrameWantCaptureMouse(wantCaptureMouse bool) {
	a0 := (C.bool)(wantCaptureMouse)
	C.igSetNextFrameWantCaptureMouse(a0)
}

func SetNextItemAllowOverlap() {
	C.igSetNextItemAllowOverlap()
}

func SetNextItemOpen(isOpen bool, cond Cond) {
	a0 := (C.bool)(isOpen)
	a1 := (C.ImGuiCond)(cond)
	C.igSetNextItemOpen(a0, a1)
}

func SetNextItemRefVal(dataType DataType, pData unsafe.Pointer) {
	a0 := (C.ImGuiDataType)(dataType)
	a1 := (unsafe.Pointer)(pData)
	C.igSetNextItemRefVal(a0, a1)
}

func SetNextItemSelectionUserData(selectionUserData SelectionUserData) {
	a0 := (C.ImGuiSelectionUserData)(selectionUserData)
	C.igSetNextItemSelectionUserData(a0)
}

func SetNextItemShortcut(keyChord KeyChord, flags InputFlags) {
	a0 := (C.ImGuiKeyChord)(keyChord)
	a1 := (C.ImGuiInputFlags)(flags)
	C.igSetNextItemShortcut(a0, a1)
}

func SetNextItemStorageID(storageId ID) {
	a0 := (C.ImGuiID)(storageId)
	C.igSetNextItemStorageID(a0)
}

func SetNextItemWidth(itemWidth float32) {
	a0 := (C.float)(itemWidth)
	C.igSetNextItemWidth(a0)
}

func SetNextWindowBgAlpha(alpha float32) {
	a0 := (C.float)(alpha)
	C.igSetNextWindowBgAlpha(a0)
}

func SetNextWindowClass(windowClass *WindowClass) {
	a0 := (*C.ImGuiWindowClass)(unsafe.Pointer(windowClass))
	C.igSetNextWindowClass(a0)
}

func SetNextWindowCollapsed(collapsed bool, cond Cond) {
	a0 := (C.bool)(collapsed)
	a1 := (C.ImGuiCond)(cond)
	C.igSetNextWindowCollapsed(a0, a1)
}

func SetNextWindowContentSize(size mgl32.Vec2) {
	a0 := mglVec2ToImVec2(size)
	C.igSetNextWindowContentSize(a0)
}

func SetNextWindowDockID(dockId ID, cond Cond) {
	a0 := (C.ImGuiID)(dockId)
	a1 := (C.ImGuiCond)(cond)
	C.igSetNextWindowDockID(a0, a1)
}

func SetNextWindowFocus() {
	C.igSetNextWindowFocus()
}

func SetNextWindowPos(pos mgl32.Vec2, cond Cond, pivot mgl32.Vec2) {
	a0 := mglVec2ToImVec2(pos)
	a1 := (C.ImGuiCond)(cond)
	a2 := mglVec2ToImVec2(pivot)
	C.igSetNextWindowPos(a0, a1, a2)
}

func SetNextWindowRefreshPolicy(flags WindowRefreshFlags) {
	a0 := (C.ImGuiWindowRefreshFlags)(flags)
	C.igSetNextWindowRefreshPolicy(a0)
}

func SetNextWindowScroll(scroll mgl32.Vec2) {
	a0 := mglVec2ToImVec2(scroll)
	C.igSetNextWindowScroll(a0)
}

func SetNextWindowSize(size mgl32.Vec2, cond Cond) {
	a0 := mglVec2ToImVec2(size)
	a1 := (C.ImGuiCond)(cond)
	C.igSetNextWindowSize(a0, a1)
}

func SetNextWindowSizeConstraints(sizeMin mgl32.Vec2, sizeMax mgl32.Vec2, customCallback SizeCallback, customCallbackData unsafe.Pointer) {
	a0 := mglVec2ToImVec2(sizeMin)
	a1 := mglVec2ToImVec2(sizeMax)
	a2 := (C.ImGuiSizeCallback)(customCallback)
	a3 := (unsafe.Pointer)(customCallbackData)
	C.igSetNextWindowSizeConstraints(a0, a1, a2, a3)
}

func SetNextWindowViewport(viewportId ID) {
	a0 := (C.ImGuiID)(viewportId)
	C.igSetNextWindowViewport(a0)
}

func SetScrollFromPosX_Float(localX float32, centerXRatio float32) {
	a0 := (C.float)(localX)
	a1 := (C.float)(centerXRatio)
	C.igSetScrollFromPosX_Float(a0, a1)
}

func SetScrollFromPosX_WindowPtr(window *Window, localX float32, centerXRatio float32) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.float)(localX)
	a2 := (C.float)(centerXRatio)
	C.igSetScrollFromPosX_WindowPtr(a0, a1, a2)
}

func SetScrollFromPosY_Float(localY float32, centerYRatio float32) {
	a0 := (C.float)(localY)
	a1 := (C.float)(centerYRatio)
	C.igSetScrollFromPosY_Float(a0, a1)
}

func SetScrollFromPosY_WindowPtr(window *Window, localY float32, centerYRatio float32) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.float)(localY)
	a2 := (C.float)(centerYRatio)
	C.igSetScrollFromPosY_WindowPtr(a0, a1, a2)
}

func SetScrollHereX(centerXRatio float32) {
	a0 := (C.float)(centerXRatio)
	C.igSetScrollHereX(a0)
}

func SetScrollHereY(centerYRatio float32) {
	a0 := (C.float)(centerYRatio)
	C.igSetScrollHereY(a0)
}

func SetScrollX_Float(scrollX float32) {
	a0 := (C.float)(scrollX)
	C.igSetScrollX_Float(a0)
}

func SetScrollX_WindowPtr(window *Window, scrollX float32) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.float)(scrollX)
	C.igSetScrollX_WindowPtr(a0, a1)
}

func SetScrollY_Float(scrollY float32) {
	a0 := (C.float)(scrollY)
	C.igSetScrollY_Float(a0)
}

func SetScrollY_WindowPtr(window *Window, scrollY float32) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.float)(scrollY)
	C.igSetScrollY_WindowPtr(a0, a1)
}

func SetShortcutRouting(keyChord KeyChord, flags InputFlags, ownerId ID) bool {
	a0 := (C.ImGuiKeyChord)(keyChord)
	a1 := (C.ImGuiInputFlags)(flags)
	a2 := (C.ImGuiID)(ownerId)
	call := C.igSetShortcutRouting(a0, a1, a2)
	r := (bool)(call)
	return r
}

func SetStateStorage(storage *Storage) {
	a0 := (*C.ImGuiStorage)(unsafe.Pointer(storage))
	C.igSetStateStorage(a0)
}

func SetTabItemClosed(tabOrDockedWindowLabel string) {
	a0 := stringPool.StoreCString(tabOrDockedWindowLabel)
	C.igSetTabItemClosed(a0)
}

func SetTooltip(vfmt string, vargs ...interface{}) {
	a0 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	C.wrap_igSetTooltip(a0)
}

func SetWindowClipRectBeforeSetChannel(window *Window, clipRect Rect) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImRect)(clipRect)
	C.igSetWindowClipRectBeforeSetChannel(a0, a1)
}

func SetWindowCollapsed_Bool(collapsed bool, cond Cond) {
	a0 := (C.bool)(collapsed)
	a1 := (C.ImGuiCond)(cond)
	C.igSetWindowCollapsed_Bool(a0, a1)
}

func SetWindowCollapsed_Str(name string, collapsed bool, cond Cond) {
	a0 := stringPool.StoreCString(name)
	a1 := (C.bool)(collapsed)
	a2 := (C.ImGuiCond)(cond)
	C.igSetWindowCollapsed_Str(a0, a1, a2)
}

func SetWindowCollapsed_WindowPtr(window *Window, collapsed bool, cond Cond) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.bool)(collapsed)
	a2 := (C.ImGuiCond)(cond)
	C.igSetWindowCollapsed_WindowPtr(a0, a1, a2)
}

func SetWindowDock(window *Window, dockId ID, cond Cond) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImGuiID)(dockId)
	a2 := (C.ImGuiCond)(cond)
	C.igSetWindowDock(a0, a1, a2)
}

func SetWindowFocus_Nil() {
	C.igSetWindowFocus_Nil()
}

func SetWindowFocus_Str(name string) {
	a0 := stringPool.StoreCString(name)
	C.igSetWindowFocus_Str(a0)
}

func SetWindowFontScale(scale float32) {
	a0 := (C.float)(scale)
	C.igSetWindowFontScale(a0)
}

func SetWindowHiddenAndSkipItemsForCurrentFrame(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igSetWindowHiddenAndSkipItemsForCurrentFrame(a0)
}

func SetWindowHitTestHole(window *Window, pos mgl32.Vec2, size mgl32.Vec2) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := mglVec2ToImVec2(pos)
	a2 := mglVec2ToImVec2(size)
	C.igSetWindowHitTestHole(a0, a1, a2)
}

func SetWindowParentWindowForFocusRoute(window *Window, parentWindow *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(parentWindow))
	C.igSetWindowParentWindowForFocusRoute(a0, a1)
}

func SetWindowPos_Str(name string, pos mgl32.Vec2, cond Cond) {
	a0 := stringPool.StoreCString(name)
	a1 := mglVec2ToImVec2(pos)
	a2 := (C.ImGuiCond)(cond)
	C.igSetWindowPos_Str(a0, a1, a2)
}

func SetWindowPos_Vec2(pos mgl32.Vec2, cond Cond) {
	a0 := mglVec2ToImVec2(pos)
	a1 := (C.ImGuiCond)(cond)
	C.igSetWindowPos_Vec2(a0, a1)
}

func SetWindowPos_WindowPtr(window *Window, pos mgl32.Vec2, cond Cond) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := mglVec2ToImVec2(pos)
	a2 := (C.ImGuiCond)(cond)
	C.igSetWindowPos_WindowPtr(a0, a1, a2)
}

func SetWindowSize_Str(name string, size mgl32.Vec2, cond Cond) {
	a0 := stringPool.StoreCString(name)
	a1 := mglVec2ToImVec2(size)
	a2 := (C.ImGuiCond)(cond)
	C.igSetWindowSize_Str(a0, a1, a2)
}

func SetWindowSize_Vec2(size mgl32.Vec2, cond Cond) {
	a0 := mglVec2ToImVec2(size)
	a1 := (C.ImGuiCond)(cond)
	C.igSetWindowSize_Vec2(a0, a1)
}

func SetWindowSize_WindowPtr(window *Window, size mgl32.Vec2, cond Cond) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := mglVec2ToImVec2(size)
	a2 := (C.ImGuiCond)(cond)
	C.igSetWindowSize_WindowPtr(a0, a1, a2)
}

func SetWindowViewport(window *Window, viewport *ViewportP) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (*C.ImGuiViewportP)(unsafe.Pointer(viewport))
	C.igSetWindowViewport(a0, a1)
}

func ShadeVertsLinearColorGradientKeepAlpha(drawList *DrawList, vertStartIdx int, vertEndIdx int, gradientP0 mgl32.Vec2, gradientP1 mgl32.Vec2, col0 U32, col1 U32) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := (C.int)(vertStartIdx)
	a2 := (C.int)(vertEndIdx)
	a3 := mglVec2ToImVec2(gradientP0)
	a4 := mglVec2ToImVec2(gradientP1)
	a5 := (C.ImU32)(col0)
	a6 := (C.ImU32)(col1)
	C.igShadeVertsLinearColorGradientKeepAlpha(a0, a1, a2, a3, a4, a5, a6)
}

func ShadeVertsLinearUV(drawList *DrawList, vertStartIdx int, vertEndIdx int, a mgl32.Vec2, b mgl32.Vec2, uvA mgl32.Vec2, uvB mgl32.Vec2, clamp bool) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := (C.int)(vertStartIdx)
	a2 := (C.int)(vertEndIdx)
	a3 := mglVec2ToImVec2(a)
	a4 := mglVec2ToImVec2(b)
	a5 := mglVec2ToImVec2(uvA)
	a6 := mglVec2ToImVec2(uvB)
	a7 := (C.bool)(clamp)
	C.igShadeVertsLinearUV(a0, a1, a2, a3, a4, a5, a6, a7)
}

func ShadeVertsTransformPos(drawList *DrawList, vertStartIdx int, vertEndIdx int, pivotIn mgl32.Vec2, cosA float32, sinA float32, pivotOut mgl32.Vec2) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := (C.int)(vertStartIdx)
	a2 := (C.int)(vertEndIdx)
	a3 := mglVec2ToImVec2(pivotIn)
	a4 := (C.float)(cosA)
	a5 := (C.float)(sinA)
	a6 := mglVec2ToImVec2(pivotOut)
	C.igShadeVertsTransformPos(a0, a1, a2, a3, a4, a5, a6)
}

func Shortcut_ID(keyChord KeyChord, flags InputFlags, ownerId ID) bool {
	a0 := (C.ImGuiKeyChord)(keyChord)
	a1 := (C.ImGuiInputFlags)(flags)
	a2 := (C.ImGuiID)(ownerId)
	call := C.igShortcut_ID(a0, a1, a2)
	r := (bool)(call)
	return r
}

func Shortcut_Nil(keyChord KeyChord, flags InputFlags) bool {
	a0 := (C.ImGuiKeyChord)(keyChord)
	a1 := (C.ImGuiInputFlags)(flags)
	call := C.igShortcut_Nil(a0, a1)
	r := (bool)(call)
	return r
}

func ShowAboutWindow(pOpen *bool) {
	a0 := (*C.bool)(unsafe.Pointer(pOpen))
	C.igShowAboutWindow(a0)
}

func ShowDebugLogWindow(pOpen *bool) {
	a0 := (*C.bool)(unsafe.Pointer(pOpen))
	C.igShowDebugLogWindow(a0)
}

func ShowDemoWindow(pOpen *bool) {
	a0 := (*C.bool)(unsafe.Pointer(pOpen))
	C.igShowDemoWindow(a0)
}

func ShowFontAtlas(atlas *FontAtlas) {
	a0 := (*C.ImFontAtlas)(unsafe.Pointer(atlas))
	C.igShowFontAtlas(a0)
}

func ShowFontSelector(label string) {
	a0 := stringPool.StoreCString(label)
	C.igShowFontSelector(a0)
}

func ShowIDStackToolWindow(pOpen *bool) {
	a0 := (*C.bool)(unsafe.Pointer(pOpen))
	C.igShowIDStackToolWindow(a0)
}

func ShowMetricsWindow(pOpen *bool) {
	a0 := (*C.bool)(unsafe.Pointer(pOpen))
	C.igShowMetricsWindow(a0)
}

func ShowStyleEditor(ref *Style) {
	a0 := (*C.ImGuiStyle)(unsafe.Pointer(ref))
	C.igShowStyleEditor(a0)
}

func ShowStyleSelector(label string) bool {
	a0 := stringPool.StoreCString(label)
	call := C.igShowStyleSelector(a0)
	r := (bool)(call)
	return r
}

func ShowUserGuide() {
	C.igShowUserGuide()
}

func ShrinkWidths(items *ShrinkWidthItem, count int, widthExcess float32) {
	a0 := (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(items))
	a1 := (C.int)(count)
	a2 := (C.float)(widthExcess)
	C.igShrinkWidths(a0, a1, a2)
}

func Shutdown() {
	C.igShutdown()
}

func SliderAngle(label string, vRad *float32, vDegreesMin float32, vDegreesMax float32, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(vRad))
	a2 := (C.float)(vDegreesMin)
	a3 := (C.float)(vDegreesMax)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderAngle(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func SliderBehavior(bb Rect, id ID, dataType DataType, pV unsafe.Pointer, pMin unsafe.Pointer, pMax unsafe.Pointer, format string, flags SliderFlags, outGrabBb *Rect) bool {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	a2 := (C.ImGuiDataType)(dataType)
	a3 := (unsafe.Pointer)(pV)
	a4 := (unsafe.Pointer)(pMin)
	a5 := (unsafe.Pointer)(pMax)
	a6 := stringPool.StoreCString(format)
	a7 := (C.ImGuiSliderFlags)(flags)
	a8 := (*C.ImRect)(unsafe.Pointer(outGrabBb))
	call := C.igSliderBehavior(a0, a1, a2, a3, a4, a5, a6, a7, a8)
	r := (bool)(call)
	return r
}

func SliderFloat(label string, v *float32, vMin float32, vMax float32, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(v))
	a2 := (C.float)(vMin)
	a3 := (C.float)(vMax)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderFloat(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func SliderFloat2(label string, v [2]float32, vMin float32, vMax float32, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(&v[0]))
	a2 := (C.float)(vMin)
	a3 := (C.float)(vMax)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderFloat2(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func SliderFloat3(label string, v *mgl32.Vec3, vMin float32, vMax float32, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(v))
	a2 := (C.float)(vMin)
	a3 := (C.float)(vMax)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderFloat3(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func SliderFloat4(label string, v *mgl32.Vec4, vMin float32, vMax float32, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.float)(unsafe.Pointer(&v[0]))
	a2 := (C.float)(vMin)
	a3 := (C.float)(vMax)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderFloat4(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func SliderInt(label string, v *int, vMin int, vMax int, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(v))
	a2 := (C.int)(vMin)
	a3 := (C.int)(vMax)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderInt(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func SliderInt2(label string, v [2]int, vMin int, vMax int, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(&v[0]))
	a2 := (C.int)(vMin)
	a3 := (C.int)(vMax)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderInt2(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func SliderInt3(label string, v [3]int, vMin int, vMax int, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(&v[0]))
	a2 := (C.int)(vMin)
	a3 := (C.int)(vMax)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderInt3(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func SliderInt4(label string, v [4]int, vMin int, vMax int, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (*C.int)(unsafe.Pointer(&v[0]))
	a2 := (C.int)(vMin)
	a3 := (C.int)(vMax)
	a4 := stringPool.StoreCString(format)
	a5 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderInt4(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func SliderScalar(label string, dataType DataType, pData unsafe.Pointer, pMin unsafe.Pointer, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.ImGuiDataType)(dataType)
	a2 := (unsafe.Pointer)(pData)
	a3 := (unsafe.Pointer)(pMin)
	a4 := (unsafe.Pointer)(pMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderScalar(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func SliderScalarN(label string, dataType DataType, pData unsafe.Pointer, components int, pMin unsafe.Pointer, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.ImGuiDataType)(dataType)
	a2 := (unsafe.Pointer)(pData)
	a3 := (C.int)(components)
	a4 := (unsafe.Pointer)(pMin)
	a5 := (unsafe.Pointer)(pMax)
	a6 := stringPool.StoreCString(format)
	a7 := (C.ImGuiSliderFlags)(flags)
	call := C.igSliderScalarN(a0, a1, a2, a3, a4, a5, a6, a7)
	r := (bool)(call)
	return r
}

func SmallButton(label string) bool {
	a0 := stringPool.StoreCString(label)
	call := C.igSmallButton(a0)
	r := (bool)(call)
	return r
}

func Spacing() {
	C.igSpacing()
}

func SplitterBehavior(bb Rect, id ID, axis Axis, size1 *float32, size2 *float32, minSize1 float32, minSize2 float32, hoverExtend float32, hoverVisibilityDelay float32, bgCol U32) bool {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	a2 := (C.ImGuiAxis)(axis)
	a3 := (*C.float)(unsafe.Pointer(size1))
	a4 := (*C.float)(unsafe.Pointer(size2))
	a5 := (C.float)(minSize1)
	a6 := (C.float)(minSize2)
	a7 := (C.float)(hoverExtend)
	a8 := (C.float)(hoverVisibilityDelay)
	a9 := (C.ImU32)(bgCol)
	call := C.igSplitterBehavior(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	r := (bool)(call)
	return r
}

func StartMouseMovingWindow(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igStartMouseMovingWindow(a0)
}

func StartMouseMovingWindowOrNode(window *Window, node *DockNode, undock bool) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (*C.ImGuiDockNode)(unsafe.Pointer(node))
	a2 := (C.bool)(undock)
	C.igStartMouseMovingWindowOrNode(a0, a1, a2)
}

func StyleColorsClassic(dst *Style) {
	a0 := (*C.ImGuiStyle)(unsafe.Pointer(dst))
	C.igStyleColorsClassic(a0)
}

func StyleColorsDark(dst *Style) {
	a0 := (*C.ImGuiStyle)(unsafe.Pointer(dst))
	C.igStyleColorsDark(a0)
}

func StyleColorsLight(dst *Style) {
	a0 := (*C.ImGuiStyle)(unsafe.Pointer(dst))
	C.igStyleColorsLight(a0)
}

func TabBarAddTab(tabBar *TabBar, tabFlags TabItemFlags, window *Window) {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (C.ImGuiTabItemFlags)(tabFlags)
	a2 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igTabBarAddTab(a0, a1, a2)
}

func TabBarCloseTab(tabBar *TabBar, tab *TabItem) {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (*C.ImGuiTabItem)(unsafe.Pointer(tab))
	C.igTabBarCloseTab(a0, a1)
}

func TabBarFindMostRecentlySelectedTabForActiveWindow(tabBar *TabBar) *TabItem {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	call := C.igTabBarFindMostRecentlySelectedTabForActiveWindow(a0)
	r := (*TabItem)(call)
	return r
}

func TabBarFindTabByID(tabBar *TabBar, tabId ID) *TabItem {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (C.ImGuiID)(tabId)
	call := C.igTabBarFindTabByID(a0, a1)
	r := (*TabItem)(call)
	return r
}

func TabBarFindTabByOrder(tabBar *TabBar, order int) *TabItem {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (C.int)(order)
	call := C.igTabBarFindTabByOrder(a0, a1)
	r := (*TabItem)(call)
	return r
}

func TabBarGetCurrentTab(tabBar *TabBar) *TabItem {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	call := C.igTabBarGetCurrentTab(a0)
	r := (*TabItem)(call)
	return r
}

func TabBarGetTabName(tabBar *TabBar, tab *TabItem) string {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (*C.ImGuiTabItem)(unsafe.Pointer(tab))
	call := C.igTabBarGetTabName(a0, a1)
	r := C.GoString(call)
	return r
}

func TabBarGetTabOrder(tabBar *TabBar, tab *TabItem) int {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (*C.ImGuiTabItem)(unsafe.Pointer(tab))
	call := C.igTabBarGetTabOrder(a0, a1)
	r := (int)(call)
	return r
}

func TabBarProcessReorder(tabBar *TabBar) bool {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	call := C.igTabBarProcessReorder(a0)
	r := (bool)(call)
	return r
}

func TabBarQueueFocus_Str(tabBar *TabBar, tabName string) {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := stringPool.StoreCString(tabName)
	C.igTabBarQueueFocus_Str(a0, a1)
}

func TabBarQueueFocus_TabItemPtr(tabBar *TabBar, tab *TabItem) {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (*C.ImGuiTabItem)(unsafe.Pointer(tab))
	C.igTabBarQueueFocus_TabItemPtr(a0, a1)
}

func TabBarQueueReorder(tabBar *TabBar, tab *TabItem, offset int) {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (*C.ImGuiTabItem)(unsafe.Pointer(tab))
	a2 := (C.int)(offset)
	C.igTabBarQueueReorder(a0, a1, a2)
}

func TabBarQueueReorderFromMousePos(tabBar *TabBar, tab *TabItem, mousePos mgl32.Vec2) {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (*C.ImGuiTabItem)(unsafe.Pointer(tab))
	a2 := mglVec2ToImVec2(mousePos)
	C.igTabBarQueueReorderFromMousePos(a0, a1, a2)
}

func TabBarRemoveTab(tabBar *TabBar, tabId ID) {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := (C.ImGuiID)(tabId)
	C.igTabBarRemoveTab(a0, a1)
}

func TabItemBackground(drawList *DrawList, bb Rect, flags TabItemFlags, col U32) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := (C.ImRect)(bb)
	a2 := (C.ImGuiTabItemFlags)(flags)
	a3 := (C.ImU32)(col)
	C.igTabItemBackground(a0, a1, a2, a3)
}

func TabItemButton(label string, flags TabItemFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.ImGuiTabItemFlags)(flags)
	call := C.igTabItemButton(a0, a1)
	r := (bool)(call)
	return r
}

func TabItemCalcSize_Str(pOut *mgl32.Vec2, label string, hasCloseButtonOrUnsavedMarker bool) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := stringPool.StoreCString(label)
	a2 := (C.bool)(hasCloseButtonOrUnsavedMarker)
	C.igTabItemCalcSize_Str(a0, a1, a2)
}

func TabItemCalcSize_WindowPtr(pOut *mgl32.Vec2, window *Window) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igTabItemCalcSize_WindowPtr(a0, a1)
}

func TabItemEx(tabBar *TabBar, label string, pOpen *bool, flags TabItemFlags, dockedWindow *Window) bool {
	a0 := (*C.ImGuiTabBar)(unsafe.Pointer(tabBar))
	a1 := stringPool.StoreCString(label)
	a2 := (*C.bool)(unsafe.Pointer(pOpen))
	a3 := (C.ImGuiTabItemFlags)(flags)
	a4 := (*C.ImGuiWindow)(unsafe.Pointer(dockedWindow))
	call := C.igTabItemEx(a0, a1, a2, a3, a4)
	r := (bool)(call)
	return r
}

func TabItemLabelAndCloseButton(drawList *DrawList, bb Rect, flags TabItemFlags, framePadding mgl32.Vec2, label string, tabId ID, closeButtonId ID, isContentsVisible bool, outJustClosed *bool, outTextClipped *bool) {
	a0 := (*C.ImDrawList)(unsafe.Pointer(drawList))
	a1 := (C.ImRect)(bb)
	a2 := (C.ImGuiTabItemFlags)(flags)
	a3 := mglVec2ToImVec2(framePadding)
	a4 := stringPool.StoreCString(label)
	a5 := (C.ImGuiID)(tabId)
	a6 := (C.ImGuiID)(closeButtonId)
	a7 := (C.bool)(isContentsVisible)
	a8 := (*C.bool)(unsafe.Pointer(outJustClosed))
	a9 := (*C.bool)(unsafe.Pointer(outTextClipped))
	C.igTabItemLabelAndCloseButton(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func TableAngledHeadersRow() {
	C.igTableAngledHeadersRow()
}

func TableAngledHeadersRowEx(rowId ID, angle float32, maxLabelWidth float32, data *TableHeaderData, dataCount int) {
	a0 := (C.ImGuiID)(rowId)
	a1 := (C.float)(angle)
	a2 := (C.float)(maxLabelWidth)
	a3 := (*C.ImGuiTableHeaderData)(unsafe.Pointer(data))
	a4 := (C.int)(dataCount)
	C.igTableAngledHeadersRowEx(a0, a1, a2, a3, a4)
}

func TableBeginApplyRequests(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableBeginApplyRequests(a0)
}

func TableBeginCell(table *Table, columnN int) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (C.int)(columnN)
	C.igTableBeginCell(a0, a1)
}

func TableBeginContextMenuPopup(table *Table) bool {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	call := C.igTableBeginContextMenuPopup(a0)
	r := (bool)(call)
	return r
}

func TableBeginInitMemory(table *Table, columnsCount int) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (C.int)(columnsCount)
	C.igTableBeginInitMemory(a0, a1)
}

func TableBeginRow(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableBeginRow(a0)
}

func TableCalcMaxColumnWidth(table *Table, columnN int) float32 {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (C.int)(columnN)
	call := C.igTableCalcMaxColumnWidth(a0, a1)
	r := (float32)(call)
	return r
}

func TableDrawBorders(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableDrawBorders(a0)
}

func TableDrawDefaultContextMenu(table *Table, flagsForSectionToDisplay TableFlags) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (C.ImGuiTableFlags)(flagsForSectionToDisplay)
	C.igTableDrawDefaultContextMenu(a0, a1)
}

func TableEndCell(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableEndCell(a0)
}

func TableEndRow(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableEndRow(a0)
}

func TableFindByID(id ID) *Table {
	a0 := (C.ImGuiID)(id)
	call := C.igTableFindByID(a0)
	r := (*Table)(call)
	return r
}

func TableFixColumnSortDirection(table *Table, column *TableColumn) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (*C.ImGuiTableColumn)(unsafe.Pointer(column))
	C.igTableFixColumnSortDirection(a0, a1)
}

func TableGcCompactSettings() {
	C.igTableGcCompactSettings()
}

func TableGcCompactTransientBuffers_TablePtr(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableGcCompactTransientBuffers_TablePtr(a0)
}

func TableGcCompactTransientBuffers_TableTempDataPtr(table *TableTempData) {
	a0 := (*C.ImGuiTableTempData)(unsafe.Pointer(table))
	C.igTableGcCompactTransientBuffers_TableTempDataPtr(a0)
}

func TableGetBoundSettings(table *Table) *TableSettings {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	call := C.igTableGetBoundSettings(a0)
	r := (*TableSettings)(call)
	return r
}

func TableGetCellBgRect(pOut *Rect, table *Table, columnN int) {
	a0 := (*C.ImRect)(unsafe.Pointer(pOut))
	a1 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a2 := (C.int)(columnN)
	C.igTableGetCellBgRect(a0, a1, a2)
}

func TableGetColumnCount() int {
	call := C.igTableGetColumnCount()
	r := (int)(call)
	return r
}

func TableGetColumnFlags(columnN int) TableColumnFlags {
	a0 := (C.int)(columnN)
	call := C.igTableGetColumnFlags(a0)
	r := (TableColumnFlags)(call)
	return r
}

func TableGetColumnIndex() int {
	call := C.igTableGetColumnIndex()
	r := (int)(call)
	return r
}

func TableGetColumnName_Int(columnN int) string {
	a0 := (C.int)(columnN)
	call := C.igTableGetColumnName_Int(a0)
	r := C.GoString(call)
	return r
}

func TableGetColumnName_TablePtr(table *Table, columnN int) string {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (C.int)(columnN)
	call := C.igTableGetColumnName_TablePtr(a0, a1)
	r := C.GoString(call)
	return r
}

func TableGetColumnNextSortDirection(column *TableColumn) SortDirection {
	a0 := (*C.ImGuiTableColumn)(unsafe.Pointer(column))
	call := C.igTableGetColumnNextSortDirection(a0)
	r := (SortDirection)(call)
	return r
}

func TableGetColumnResizeID(table *Table, columnN int, instanceNo int) ID {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (C.int)(columnN)
	a2 := (C.int)(instanceNo)
	call := C.igTableGetColumnResizeID(a0, a1, a2)
	r := (ID)(call)
	return r
}

func TableGetColumnWidthAuto(table *Table, column *TableColumn) float32 {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (*C.ImGuiTableColumn)(unsafe.Pointer(column))
	call := C.igTableGetColumnWidthAuto(a0, a1)
	r := (float32)(call)
	return r
}

func TableGetHeaderAngledMaxLabelWidth() float32 {
	call := C.igTableGetHeaderAngledMaxLabelWidth()
	r := (float32)(call)
	return r
}

func TableGetHeaderRowHeight() float32 {
	call := C.igTableGetHeaderRowHeight()
	r := (float32)(call)
	return r
}

func TableGetHoveredColumn() int {
	call := C.igTableGetHoveredColumn()
	r := (int)(call)
	return r
}

func TableGetHoveredRow() int {
	call := C.igTableGetHoveredRow()
	r := (int)(call)
	return r
}

func TableGetInstanceData(table *Table, instanceNo int) *TableInstanceData {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (C.int)(instanceNo)
	call := C.igTableGetInstanceData(a0, a1)
	r := (*TableInstanceData)(call)
	return r
}

func TableGetInstanceID(table *Table, instanceNo int) ID {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (C.int)(instanceNo)
	call := C.igTableGetInstanceID(a0, a1)
	r := (ID)(call)
	return r
}

func TableGetRowIndex() int {
	call := C.igTableGetRowIndex()
	r := (int)(call)
	return r
}

func TableGetSortSpecs() *TableSortSpecs {
	call := C.igTableGetSortSpecs()
	r := (*TableSortSpecs)(call)
	return r
}

func TableHeader(label string) {
	a0 := stringPool.StoreCString(label)
	C.igTableHeader(a0)
}

func TableHeadersRow() {
	C.igTableHeadersRow()
}

func TableLoadSettings(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableLoadSettings(a0)
}

func TableMergeDrawChannels(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableMergeDrawChannels(a0)
}

func TableNextColumn() bool {
	call := C.igTableNextColumn()
	r := (bool)(call)
	return r
}

func TableNextRow(rowFlags TableRowFlags, minRowHeight float32) {
	a0 := (C.ImGuiTableRowFlags)(rowFlags)
	a1 := (C.float)(minRowHeight)
	C.igTableNextRow(a0, a1)
}

func TableOpenContextMenu(columnN int) {
	a0 := (C.int)(columnN)
	C.igTableOpenContextMenu(a0)
}

func TablePopBackgroundChannel() {
	C.igTablePopBackgroundChannel()
}

func TablePushBackgroundChannel() {
	C.igTablePushBackgroundChannel()
}

func TableRemove(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableRemove(a0)
}

func TableResetSettings(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableResetSettings(a0)
}

func TableSaveSettings(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableSaveSettings(a0)
}

func TableSetBgColor(target TableBgTarget, color U32, columnN int) {
	a0 := (C.ImGuiTableBgTarget)(target)
	a1 := (C.ImU32)(color)
	a2 := (C.int)(columnN)
	C.igTableSetBgColor(a0, a1, a2)
}

func TableSetColumnEnabled(columnN int, v bool) {
	a0 := (C.int)(columnN)
	a1 := (C.bool)(v)
	C.igTableSetColumnEnabled(a0, a1)
}

func TableSetColumnIndex(columnN int) bool {
	a0 := (C.int)(columnN)
	call := C.igTableSetColumnIndex(a0)
	r := (bool)(call)
	return r
}

func TableSetColumnSortDirection(columnN int, sortDirection SortDirection, appendToSortSpecs bool) {
	a0 := (C.int)(columnN)
	a1 := (C.ImGuiSortDirection)(sortDirection)
	a2 := (C.bool)(appendToSortSpecs)
	C.igTableSetColumnSortDirection(a0, a1, a2)
}

func TableSetColumnWidth(columnN int, width float32) {
	a0 := (C.int)(columnN)
	a1 := (C.float)(width)
	C.igTableSetColumnWidth(a0, a1)
}

func TableSetColumnWidthAutoAll(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableSetColumnWidthAutoAll(a0)
}

func TableSetColumnWidthAutoSingle(table *Table, columnN int) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	a1 := (C.int)(columnN)
	C.igTableSetColumnWidthAutoSingle(a0, a1)
}

func TableSettingsAddSettingsHandler() {
	C.igTableSettingsAddSettingsHandler()
}

func TableSettingsCreate(id ID, columnsCount int) *TableSettings {
	a0 := (C.ImGuiID)(id)
	a1 := (C.int)(columnsCount)
	call := C.igTableSettingsCreate(a0, a1)
	r := (*TableSettings)(call)
	return r
}

func TableSettingsFindByID(id ID) *TableSettings {
	a0 := (C.ImGuiID)(id)
	call := C.igTableSettingsFindByID(a0)
	r := (*TableSettings)(call)
	return r
}

func TableSetupColumn(label string, flags TableColumnFlags, initWidthOrWeight float32, userId ID) {
	a0 := stringPool.StoreCString(label)
	a1 := (C.ImGuiTableColumnFlags)(flags)
	a2 := (C.float)(initWidthOrWeight)
	a3 := (C.ImGuiID)(userId)
	C.igTableSetupColumn(a0, a1, a2, a3)
}

func TableSetupDrawChannels(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableSetupDrawChannels(a0)
}

func TableSetupScrollFreeze(cols int, rows int) {
	a0 := (C.int)(cols)
	a1 := (C.int)(rows)
	C.igTableSetupScrollFreeze(a0, a1)
}

func TableSortSpecsBuild(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableSortSpecsBuild(a0)
}

func TableSortSpecsSanitize(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableSortSpecsSanitize(a0)
}

func TableUpdateBorders(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableUpdateBorders(a0)
}

func TableUpdateColumnsWeightFromWidth(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableUpdateColumnsWeightFromWidth(a0)
}

func TableUpdateLayout(table *Table) {
	a0 := (*C.ImGuiTable)(unsafe.Pointer(table))
	C.igTableUpdateLayout(a0)
}

func TeleportMousePos(pos mgl32.Vec2) {
	a0 := mglVec2ToImVec2(pos)
	C.igTeleportMousePos(a0)
}

func TempInputIsActive(id ID) bool {
	a0 := (C.ImGuiID)(id)
	call := C.igTempInputIsActive(a0)
	r := (bool)(call)
	return r
}

func TempInputScalar(bb Rect, id ID, label string, dataType DataType, pData unsafe.Pointer, format string, pClampMin unsafe.Pointer, pClampMax unsafe.Pointer) bool {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	a2 := stringPool.StoreCString(label)
	a3 := (C.ImGuiDataType)(dataType)
	a4 := (unsafe.Pointer)(pData)
	a5 := stringPool.StoreCString(format)
	a6 := (unsafe.Pointer)(pClampMin)
	a7 := (unsafe.Pointer)(pClampMax)
	call := C.igTempInputScalar(a0, a1, a2, a3, a4, a5, a6, a7)
	r := (bool)(call)
	return r
}

func TempInputText(bb Rect, id ID, label string, buf string, bufSize int, flags InputTextFlags) bool {
	a0 := (C.ImRect)(bb)
	a1 := (C.ImGuiID)(id)
	a2 := stringPool.StoreCString(label)
	a3 := stringPool.StoreCString(buf)
	a4 := (C.int)(bufSize)
	a5 := (C.ImGuiInputTextFlags)(flags)
	call := C.igTempInputText(a0, a1, a2, a3, a4, a5)
	r := (bool)(call)
	return r
}

func TestKeyOwner(key Key, ownerId ID) bool {
	a0 := (C.ImGuiKey)(key)
	a1 := (C.ImGuiID)(ownerId)
	call := C.igTestKeyOwner(a0, a1)
	r := (bool)(call)
	return r
}

func TestShortcutRouting(keyChord KeyChord, ownerId ID) bool {
	a0 := (C.ImGuiKeyChord)(keyChord)
	a1 := (C.ImGuiID)(ownerId)
	call := C.igTestShortcutRouting(a0, a1)
	r := (bool)(call)
	return r
}

func Text(vfmt string, vargs ...interface{}) {
	a0 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	C.wrap_igText(a0)
}

func TextColored(col mgl32.Vec4, vfmt string, vargs ...interface{}) {
	a0 := mglVec4ToImVec4(col)
	a1 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	C.wrap_igTextColored(a0, a1)
}

func TextDisabled(vfmt string, vargs ...interface{}) {
	a0 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	C.wrap_igTextDisabled(a0)
}

func TextEx(text string, textEnd string, flags TextFlags) {
	a0 := stringPool.StoreCString(text)
	a1 := stringPool.StoreCString(textEnd)
	a2 := (C.ImGuiTextFlags)(flags)
	C.igTextEx(a0, a1, a2)
}

func TextLink(label string) bool {
	a0 := stringPool.StoreCString(label)
	call := C.igTextLink(a0)
	r := (bool)(call)
	return r
}

func TextLinkOpenURL(label string, url string) {
	a0 := stringPool.StoreCString(label)
	a1 := stringPool.StoreCString(url)
	C.igTextLinkOpenURL(a0, a1)
}

func TextUnformatted(text string, textEnd string) {
	a0 := stringPool.StoreCString(text)
	a1 := stringPool.StoreCString(textEnd)
	C.igTextUnformatted(a0, a1)
}

func TextWrapped(vfmt string, vargs ...interface{}) {
	a0 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	C.wrap_igTextWrapped(a0)
}

func TranslateWindowsInViewport(viewport *ViewportP, oldPos mgl32.Vec2, newPos mgl32.Vec2, oldSize mgl32.Vec2, newSize mgl32.Vec2) {
	a0 := (*C.ImGuiViewportP)(unsafe.Pointer(viewport))
	a1 := mglVec2ToImVec2(oldPos)
	a2 := mglVec2ToImVec2(newPos)
	a3 := mglVec2ToImVec2(oldSize)
	a4 := mglVec2ToImVec2(newSize)
	C.igTranslateWindowsInViewport(a0, a1, a2, a3, a4)
}

func TreeNodeBehavior(id ID, flags TreeNodeFlags, label string, labelEnd string) bool {
	a0 := (C.ImGuiID)(id)
	a1 := (C.ImGuiTreeNodeFlags)(flags)
	a2 := stringPool.StoreCString(label)
	a3 := stringPool.StoreCString(labelEnd)
	call := C.igTreeNodeBehavior(a0, a1, a2, a3)
	r := (bool)(call)
	return r
}

func TreeNodeEx_Ptr(ptrId unsafe.Pointer, flags TreeNodeFlags, vfmt string, vargs ...interface{}) bool {
	a0 := (unsafe.Pointer)(ptrId)
	a1 := (C.ImGuiTreeNodeFlags)(flags)
	a2 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	call := C.wrap_igTreeNodeEx_Ptr(a0, a1, a2)
	r := (bool)(call)
	return r
}

func TreeNodeEx_Str(label string, flags TreeNodeFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := (C.ImGuiTreeNodeFlags)(flags)
	call := C.igTreeNodeEx_Str(a0, a1)
	r := (bool)(call)
	return r
}

func TreeNodeEx_StrStr(strId string, flags TreeNodeFlags, vfmt string, vargs ...interface{}) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := (C.ImGuiTreeNodeFlags)(flags)
	a2 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	call := C.wrap_igTreeNodeEx_StrStr(a0, a1, a2)
	r := (bool)(call)
	return r
}

func TreeNodeGetOpen(storageId ID) bool {
	a0 := (C.ImGuiID)(storageId)
	call := C.igTreeNodeGetOpen(a0)
	r := (bool)(call)
	return r
}

func TreeNodeSetOpen(storageId ID, open bool) {
	a0 := (C.ImGuiID)(storageId)
	a1 := (C.bool)(open)
	C.igTreeNodeSetOpen(a0, a1)
}

func TreeNodeUpdateNextOpen(storageId ID, flags TreeNodeFlags) bool {
	a0 := (C.ImGuiID)(storageId)
	a1 := (C.ImGuiTreeNodeFlags)(flags)
	call := C.igTreeNodeUpdateNextOpen(a0, a1)
	r := (bool)(call)
	return r
}

func TreeNode_Ptr(ptrId unsafe.Pointer, vfmt string, vargs ...interface{}) bool {
	a0 := (unsafe.Pointer)(ptrId)
	a1 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	call := C.wrap_igTreeNode_Ptr(a0, a1)
	r := (bool)(call)
	return r
}

func TreeNode_Str(label string) bool {
	a0 := stringPool.StoreCString(label)
	call := C.igTreeNode_Str(a0)
	r := (bool)(call)
	return r
}

func TreeNode_StrStr(strId string, vfmt string, vargs ...interface{}) bool {
	a0 := stringPool.StoreCString(strId)
	a1 := stringPool.StoreCString(fmt.Sprintf(vfmt, vargs...))
	call := C.wrap_igTreeNode_StrStr(a0, a1)
	r := (bool)(call)
	return r
}

func TreePop() {
	C.igTreePop()
}

func TreePushOverrideID(id ID) {
	a0 := (C.ImGuiID)(id)
	C.igTreePushOverrideID(a0)
}

func TreePush_Ptr(ptrId unsafe.Pointer) {
	a0 := (unsafe.Pointer)(ptrId)
	C.igTreePush_Ptr(a0)
}

func TreePush_Str(strId string) {
	a0 := stringPool.StoreCString(strId)
	C.igTreePush_Str(a0)
}

func Unindent(indentW float32) {
	a0 := (C.float)(indentW)
	C.igUnindent(a0)
}

func UpdateHoveredWindowAndCaptureFlags() {
	C.igUpdateHoveredWindowAndCaptureFlags()
}

func UpdateInputEvents(trickleFastInputs bool) {
	a0 := (C.bool)(trickleFastInputs)
	C.igUpdateInputEvents(a0)
}

func UpdateMouseMovingWindowEndFrame() {
	C.igUpdateMouseMovingWindowEndFrame()
}

func UpdateMouseMovingWindowNewFrame() {
	C.igUpdateMouseMovingWindowNewFrame()
}

func UpdatePlatformWindows() {
	C.igUpdatePlatformWindows()
}

func UpdateWindowParentAndRootLinks(window *Window, flags WindowFlags, parentWindow *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a1 := (C.ImGuiWindowFlags)(flags)
	a2 := (*C.ImGuiWindow)(unsafe.Pointer(parentWindow))
	C.igUpdateWindowParentAndRootLinks(a0, a1, a2)
}

func UpdateWindowSkipRefresh(window *Window) {
	a0 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	C.igUpdateWindowSkipRefresh(a0)
}

func VSliderFloat(label string, size mgl32.Vec2, v *float32, vMin float32, vMax float32, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := mglVec2ToImVec2(size)
	a2 := (*C.float)(unsafe.Pointer(v))
	a3 := (C.float)(vMin)
	a4 := (C.float)(vMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igVSliderFloat(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func VSliderInt(label string, size mgl32.Vec2, v *int, vMin int, vMax int, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := mglVec2ToImVec2(size)
	a2 := (*C.int)(unsafe.Pointer(v))
	a3 := (C.int)(vMin)
	a4 := (C.int)(vMax)
	a5 := stringPool.StoreCString(format)
	a6 := (C.ImGuiSliderFlags)(flags)
	call := C.igVSliderInt(a0, a1, a2, a3, a4, a5, a6)
	r := (bool)(call)
	return r
}

func VSliderScalar(label string, size mgl32.Vec2, dataType DataType, pData unsafe.Pointer, pMin unsafe.Pointer, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	a0 := stringPool.StoreCString(label)
	a1 := mglVec2ToImVec2(size)
	a2 := (C.ImGuiDataType)(dataType)
	a3 := (unsafe.Pointer)(pData)
	a4 := (unsafe.Pointer)(pMin)
	a5 := (unsafe.Pointer)(pMax)
	a6 := stringPool.StoreCString(format)
	a7 := (C.ImGuiSliderFlags)(flags)
	call := C.igVSliderScalar(a0, a1, a2, a3, a4, a5, a6, a7)
	r := (bool)(call)
	return r
}

func Value_Bool(prefix string, b bool) {
	a0 := stringPool.StoreCString(prefix)
	a1 := (C.bool)(b)
	C.igValue_Bool(a0, a1)
}

func Value_Float(prefix string, v float32, floatFormat string) {
	a0 := stringPool.StoreCString(prefix)
	a1 := (C.float)(v)
	a2 := stringPool.StoreCString(floatFormat)
	C.igValue_Float(a0, a1, a2)
}

func Value_Int(prefix string, v int) {
	a0 := stringPool.StoreCString(prefix)
	a1 := (C.int)(v)
	C.igValue_Int(a0, a1)
}

func Value_Uint(prefix string, v uint) {
	a0 := stringPool.StoreCString(prefix)
	a1 := (C.uint)(v)
	C.igValue_Uint(a0, a1)
}

func WindowPosAbsToRel(pOut *mgl32.Vec2, window *Window, p mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a2 := mglVec2ToImVec2(p)
	C.igWindowPosAbsToRel(a0, a1, a2)
}

func WindowPosRelToAbs(pOut *mgl32.Vec2, window *Window, p mgl32.Vec2) {
	a0 := (*C.ImVec2)(unsafe.Pointer(&pOut[0]))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a2 := mglVec2ToImVec2(p)
	C.igWindowPosRelToAbs(a0, a1, a2)
}

func WindowRectAbsToRel(pOut *Rect, window *Window, r Rect) {
	a0 := (*C.ImRect)(unsafe.Pointer(pOut))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a2 := (C.ImRect)(r)
	C.igWindowRectAbsToRel(a0, a1, a2)
}

func WindowRectRelToAbs(pOut *Rect, window *Window, r Rect) {
	a0 := (*C.ImRect)(unsafe.Pointer(pOut))
	a1 := (*C.ImGuiWindow)(unsafe.Pointer(window))
	a2 := (C.ImRect)(r)
	C.igWindowRectRelToAbs(a0, a1, a2)
}

