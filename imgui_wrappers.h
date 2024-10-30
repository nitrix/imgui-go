#define CIMGUI_DEFINE_ENUMS_AND_STRUCTS 1
#include "dist/cimgui/cimgui.h"

void wrap_ImGuiTextBuffer_appendf(ImGuiTextBuffer* self, const char* fmt);
void wrap_igBulletText(const char* fmt);
void wrap_igDebugLog(const char* fmt);
int wrap_igImFormatString(char* buf, size_t buf_size, const char* fmt);
void wrap_igImFormatStringToTempBuffer(const char** out_buf, const char** out_buf_end, const char* fmt);
void wrap_igLabelText(const char* label, const char* fmt);
void wrap_igLogText(const char* fmt);
void wrap_igSetItemTooltip(const char* fmt);
void wrap_igSetTooltip(const char* fmt);
void wrap_igText(const char* fmt);
void wrap_igTextColored(const ImVec4 col, const char* fmt);
void wrap_igTextDisabled(const char* fmt);
void wrap_igTextWrapped(const char* fmt);
bool wrap_igTreeNodeEx_Ptr(const void* ptr_id, ImGuiTreeNodeFlags flags, const char* fmt);
bool wrap_igTreeNodeEx_StrStr(const char* str_id, ImGuiTreeNodeFlags flags, const char* fmt);
bool wrap_igTreeNode_Ptr(const void* ptr_id, const char* fmt);
bool wrap_igTreeNode_StrStr(const char* str_id, const char* fmt);
