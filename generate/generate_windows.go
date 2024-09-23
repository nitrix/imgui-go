package main

func doPlatformSpecific() {
	copyFile("build/libcimgui.a", "dist/windows/libcimgui.a")
	copyFile("build/thirdparty/glad/libglad.a", "dist/windows/libglad.a")
	copyFile("build/thirdparty/glfw/src/libglfw3.a", "dist/windows/libglfw3.a")
}
