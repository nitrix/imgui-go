#define CIMGUI_DEFINE_ENUMS_AND_STRUCTS 1
#include "dist/include/cimgui/cimgui.h"

void wrap_ImGuiTextBuffer_appendf(ImGuiTextBuffer* self, const char* fmt) {
	ImGuiTextBuffer_appendf(self, fmt);
}

void wrap_igBulletText(const char* fmt) {
	igBulletText(fmt);
}

void wrap_igDebugLog(const char* fmt) {
	igDebugLog(fmt);
}

int wrap_igImFormatString(char* buf, size_t buf_size, const char* fmt) {
	return igImFormatString(buf, buf_size, fmt);
}

void wrap_igImFormatStringToTempBuffer(const char** out_buf, const char** out_buf_end, const char* fmt) {
	igImFormatStringToTempBuffer(out_buf, out_buf_end, fmt);
}

void wrap_igLabelText(const char* label, const char* fmt) {
	igLabelText(label, fmt);
}

void wrap_igLogText(const char* fmt) {
	igLogText(fmt);
}

void wrap_igSetItemTooltip(const char* fmt) {
	igSetItemTooltip(fmt);
}

void wrap_igSetTooltip(const char* fmt) {
	igSetTooltip(fmt);
}

void wrap_igText(const char* fmt) {
	igText(fmt);
}

void wrap_igTextColored(const ImVec4 col, const char* fmt) {
	igTextColored(col, fmt);
}

void wrap_igTextDisabled(const char* fmt) {
	igTextDisabled(fmt);
}

void wrap_igTextWrapped(const char* fmt) {
	igTextWrapped(fmt);
}

bool wrap_igTreeNodeEx_Ptr(const void* ptr_id, ImGuiTreeNodeFlags flags, const char* fmt) {
	return igTreeNodeEx_Ptr(ptr_id, flags, fmt);
}

bool wrap_igTreeNodeEx_StrStr(const char* str_id, ImGuiTreeNodeFlags flags, const char* fmt) {
	return igTreeNodeEx_StrStr(str_id, flags, fmt);
}

bool wrap_igTreeNode_Ptr(const void* ptr_id, const char* fmt) {
	return igTreeNode_Ptr(ptr_id, fmt);
}

bool wrap_igTreeNode_StrStr(const char* str_id, const char* fmt) {
	return igTreeNode_StrStr(str_id, fmt);
}

