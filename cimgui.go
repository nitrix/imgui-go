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
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	C.glfwInit()
	defer C.glfwTerminate()

	var major, minor, rev C.int
	C.glfwGetVersion(&major, &minor, &rev)
	fmt.Println("GLFW Version", major, minor, rev)

	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MAJOR, 3)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MINOR, 3)
	C.glfwWindowHint(C.GLFW_OPENGL_PROFILE, C.GLFW_OPENGL_CORE_PROFILE)
	C.glfwWindowHint(C.GLFW_OPENGL_FORWARD_COMPAT, C.GLFW_TRUE)
	C.glfwWindowHint(C.GLFW_VISIBLE, C.GLFW_FALSE)

	window, err := C.glfwCreateWindow(1280, 720, C.CString("Demo"), nil, nil)
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
			C.glfwGetWindowPos(window, &x, &y)
			C.glfwSetWindowPos(window, x+(mode.width-1280)/2, y+(mode.height-720)/2)
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

// #if IMGUI_HAS_DOCK
// ctx->IO.ConfigFlags |= ImGuiConfigFlags_DockingEnable;
// ctx->IO.ConfigFlags |= ImGuiConfigFlags_ViewportsEnable;
// ctx->IO.ConfigDockingWithShift = true;
// #endif

// igDockSpaceOverViewport(0, NULL, ImGuiDockNodeFlags_PassthruCentralNode, NULL);

// #if IMGUI_HAS_DOCK
// if (ctx->IO.ConfigFlags & ImGuiConfigFlags_ViewportsEnable) {
// 	GLFWwindow *previous = glfwGetCurrentContext();
// 	igUpdatePlatformWindows();
// 	igRenderPlatformWindowsDefault(NULL, NULL);
// 	glfwMakeContextCurrent(previous);
// }
// #endif
