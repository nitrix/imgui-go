package main

/*
#cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -DCIMGUI_USE_GLFW -DCIMGUI_USE_OPENGL3
#cgo CPPFLAGS: -Ithirdparty/cimgui/imgui -Ithirdparty/glfw/include
#include "thirdparty/cimgui/cimgui.h"
#include "thirdparty/cimgui/generator/output/cimgui_impl.h"
#include "thirdparty/glfw/include/GLFW/glfw3.h"
//#cgo CPPFLAGS: -DIMGUI_IMPL_API=\"extern\\\t\\\"C\\\"\" -DIMGUI_IMPL_OPENGL_LOADER_GLAD

// GLFW below

// Windows Build Tags
// ----------------
// GLFW Options:
#cgo windows CFLAGS: -D_GLFW_WIN32 -Ithirdparty/glfw/deps/mingw

// Linker Options:
#cgo windows LDFLAGS: -lgdi32

#cgo !gles2,windows LDFLAGS: -lopengl32
#cgo gles2,windows LDFLAGS: -lGLESv2

// Darwin Build Tags
// ----------------
// GLFW Options:
#cgo darwin CFLAGS: -D_GLFW_COCOA -Wno-deprecated-declarations

// Linker Options:
#cgo darwin LDFLAGS: -framework Cocoa -framework IOKit -framework CoreVideo

#cgo !gles2,darwin LDFLAGS: -framework OpenGL
#cgo gles2,darwin LDFLAGS: -lGLESv2

// Linux Build Tags
// ----------------
// GLFW Options:
#cgo linux,!wayland CFLAGS: -D_GLFW_X11
#cgo linux,wayland CFLAGS: -D_GLFW_WAYLAND -D_GNU_SOURCE

// Linker Options:
#cgo linux,!gles1,!gles2,!gles3,!vulkan LDFLAGS: -lGL
#cgo linux,gles1 LDFLAGS: -lGLESv1
#cgo linux,gles2 LDFLAGS: -lGLESv2
#cgo linux,gles3 LDFLAGS: -lGLESv3
#cgo linux,vulkan LDFLAGS: -lvulkan
#cgo linux,!wayland LDFLAGS: -lX11 -lXrandr -lXxf86vm -lXi -lXcursor -lm -lXinerama -ldl -lrt
#cgo linux,wayland LDFLAGS: -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon -lm -ldl -lrt

// BSD Build Tags
// ----------------
// GLFW Options:
#cgo freebsd,!wayland netbsd,!wayland openbsd pkg-config: x11 xau xcb xdmcp
#cgo freebsd,wayland netbsd,wayland pkg-config: wayland-client wayland-cursor wayland-egl epoll-shim
#cgo freebsd netbsd openbsd CFLAGS: -D_GLFW_HAS_DLOPEN
#cgo freebsd,!wayland netbsd,!wayland openbsd CFLAGS: -D_GLFW_X11 -D_GLFW_HAS_GLXGETPROCADDRESSARB
#cgo freebsd,wayland netbsd,wayland CFLAGS: -D_GLFW_WAYLAND

// Linker Options:
#cgo freebsd netbsd openbsd LDFLAGS: -lm
*/
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

	ctx.IO.IniFilename = nil

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
