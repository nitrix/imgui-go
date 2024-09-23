package main

func doPlatformSpecific() {
	copyFile("build/libcimgui.a", "dist/linux/libcimgui.a")
	copyFile("build/thirdparty/glad/libglad.a", "dist/linux/libglad.a")
	copyFile("build/thirdparty/glfw/src/libglfw3.a", "dist/linux/libglfw3.a")
}
