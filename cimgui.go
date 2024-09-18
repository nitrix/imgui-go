package main

// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -DCIMGUI_USE_GLFW -DCIMGUI_USE_OPENGL3
// #cgo LDFLAGS: -Lbuild -Lbuild/glfw/src -Lbuild/glad -lcimgui -lglfw3 -lglad
// #cgo LDFLAGS: -lopengl32 -lgdi32 -lc++ -lc++abi
// #cgo CFLAGS: -I. -Iglfw/include
// #include "cimgui/cimgui.h"
// #include "cimgui/generator/output/cimgui_impl.h"
// #include "glfw/include/GLFW/glfw3.h"
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	C.glfwInit()
	defer C.glfwTerminate()

	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MAJOR, 3)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MINOR, 3)
	C.glfwWindowHint(C.GLFW_OPENGL_PROFILE, C.GLFW_OPENGL_CORE_PROFILE)
	C.glfwWindowHint(C.GLFW_OPENGL_FORWARD_COMPAT, C.GLFW_TRUE)
	C.glfwWindowHint(C.GLFW_VISIBLE, C.GLFW_FALSE)

	// title := C.CString("Demo")
	title := (*C.char)(unsafe.Pointer(unsafe.StringData("Demo\x00")))

	window, err := C.glfwCreateWindow(1280, 720, title, nil, nil)
	if err != nil {
		panic(err)
	}

	C.glfwMakeContextCurrent(window)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	C.glfwSwapInterval(1)

	primary := C.glfwGetPrimaryMonitor()
	if primary != nil {
		mode := C.glfwGetVideoMode(primary)
		if mode != nil {
			var x, y C.int
			C.glfwGetWindowSize(window, &x, &y)
			C.glfwSetWindowPos(window, (mode.width/2)-(x/2), (mode.height/2)-(y/2))
		}
	}

	C.glfwShowWindow(window)

	ctx := C.igCreateContext(nil)
	C.ImGui_ImplGlfw_InitForOpenGL(window, true)
	C.ImGui_ImplOpenGL3_Init(C.CString("#version 330 core"))

	for C.glfwWindowShouldClose(window) != C.GLFW_TRUE {
		C.glfwPollEvents()

		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		C.ImGui_ImplOpenGL3_NewFrame()
		C.ImGui_ImplGlfw_NewFrame()
		C.igNewFrame()

		C.igShowDemoWindow(nil)

		C.igRender()

		C.ImGui_ImplOpenGL3_RenderDrawData(C.igGetDrawData())

		C.glfwSwapBuffers(window)
	}

	C.ImGui_ImplOpenGL3_Shutdown()
	C.ImGui_ImplGlfw_Shutdown()
	C.igDestroyContext(ctx)
}
