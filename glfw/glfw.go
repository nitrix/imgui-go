package glfw

// #cgo CFLAGS: -I../dist/include
// #cgo windows LDFLAGS: -L../dist/windows
// #cgo windows LDFLAGS: -lopengl32 -lgdi32
// #cgo linux LDFLAGS: -L../dist/linux
// #cgo LDFLAGS: -lglfw3
// #include "GLFW/glfw3.h"
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	ContextVersionMajor = C.GLFW_CONTEXT_VERSION_MAJOR
	ContextVersionMinor = C.GLFW_CONTEXT_VERSION_MINOR
	OpenGLProfile       = C.GLFW_OPENGL_PROFILE
	OpenGLCoreProfile   = C.GLFW_OPENGL_CORE_PROFILE
	OpenGLForwardCompat = C.GLFW_OPENGL_FORWARD_COMPAT
	Visible             = C.GLFW_VISIBLE
	True                = C.GLFW_TRUE
	False               = C.GLFW_FALSE
)

type Monitor C.GLFWmonitor
type Window C.GLFWwindow
type VideoMode struct {
	Width  int
	Height int
}

func Init() error {
	if C.glfwInit() != C.GLFW_TRUE {
		return fmt.Errorf("initialization failed")
	}
	return nil
}

func Terminate() {
	C.glfwTerminate()
}

func WindowHint(hint int, value int) {
	C.glfwWindowHint(C.int(hint), C.int(value))
}

func CreateWindow(width, height int, title string, monitor *Monitor, share *Window) (*Window, error) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	window := C.glfwCreateWindow(C.int(width), C.int(height), cTitle, (*C.GLFWmonitor)(monitor), (*C.GLFWwindow)(share))
	if window == nil {
		return nil, fmt.Errorf("failed to create window")
	}
	return (*Window)(window), nil
}

func SwapInterval(n int) {
	C.glfwSwapInterval(C.int(n))
}

func MakeContextCurrent(w *Window) {
	C.glfwMakeContextCurrent(w.Handle())
}

func PollEvents() {
	C.glfwPollEvents()
}

func GetPrimaryMonitor() *Monitor {
	return (*Monitor)(C.glfwGetPrimaryMonitor())
}

func (w *Window) Handle() *C.GLFWwindow {
	return (*C.GLFWwindow)(w)
}

func (w *Window) Show() {
	C.glfwShowWindow(w.Handle())
}

func (w *Window) ShouldClose() bool {
	return C.glfwWindowShouldClose(w.Handle()) == C.GLFW_TRUE
}

func (w *Window) SwapBuffers() {
	C.glfwSwapBuffers(w.Handle())
}

func (w *Window) Destroy() {
	C.glfwDestroyWindow(w.Handle())
}

func (w *Window) GetSize() (int, int) {
	var width, height C.int
	C.glfwGetWindowSize(w.Handle(), &width, &height)
	return int(width), int(height)
}

func (w *Window) SetPos(x, y int) {
	C.glfwSetWindowPos(w.Handle(), C.int(x), C.int(y))
}

func (w *Window) Centerize() {
	primary := GetPrimaryMonitor()
	if primary != nil {
		mode := primary.GetVideoMode()
		if mode != nil {
			x, y := w.GetSize()
			w.SetPos((mode.Width/2)-(x/2), (mode.Height/2)-(y/2))
		}
	}
}

func (m *Monitor) GetVideoMode() *VideoMode {
	videoMode := C.glfwGetVideoMode((*C.GLFWmonitor)(m))
	return &VideoMode{
		Width:  int(videoMode.width),
		Height: int(videoMode.height),
	}
}
