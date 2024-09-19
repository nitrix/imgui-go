package main

func main() {
	copyFile("build/libcimgui.a", "dist/windows/libcimgui.a")
	copyFile("build/thirdparty/glad/libglad.a", "dist/windows/libglad.a")
	copyFile("build/thirdparty/glfw/src/libglfw3.a", "dist/windows/libglfw3.a")

	copyFile("thirdparty/glfw/include/GLFW/glfw3.h", "dist/include/GLFW/glfw3.h")
	copyFile("thirdparty/glfw/include/GLFW/glfw3native.h", "dist/include/GLFW/glfw3native.h")
	copyFile("thirdparty/glad/include/glad.h", "dist/include/glad/glad.h")
	copyFile("thirdparty/cimgui/cimgui.h", "dist/include/cimgui/cimgui.h")
	copyFile("thirdparty/cimgui/generator/output/cimgui_impl.h", "dist/include/cimgui/cimgui_impl.h")
}
