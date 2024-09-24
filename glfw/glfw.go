package glfw

// #cgo CFLAGS: -I../dist/include
// #cgo windows LDFLAGS: -L../dist/windows
// #cgo windows LDFLAGS: -lopengl32 -lgdi32
// #cgo linux LDFLAGS: -L../dist/linux
// #cgo LDFLAGS: -lglfw3
// #include "GLFW/glfw3.h"
// #include <stdlib.h>
// void glfwSetCursorPosCallback_fix(GLFWwindow *window);
// void glfwSetWindowSizeCallback_fix(GLFWwindow *window);
// void glfwSetFramebufferSizeCallback_fix(GLFWwindow *window);
import "C"

import (
	"fmt"
	"unsafe"
)

type Monitor C.GLFWmonitor

type Window struct {
	handle                  *C.GLFWwindow
	userPtr                 unsafe.Pointer
	cursorPosCallback       func(ww *Window, x float64, y float64)
	sizeCallback            func(ww *Window, x int, y int)
	framebuffersizeCallback func(ww *Window, x int, y int)
}

type VideoMode struct {
	Width  int
	Height int
}

//export goCursorPosCallback
func goCursorPosCallback(w *C.GLFWwindow, xpos, ypos C.double) {
	real := (*Window)(C.glfwGetWindowUserPointer(w))
	real.cursorPosCallback(real, float64(xpos), float64(ypos))
}

//export goSizeCallback
func goSizeCallback(w *C.GLFWwindow, width, height C.int) {
	real := (*Window)(C.glfwGetWindowUserPointer(w))
	real.sizeCallback(real, int(width), int(height))
}

//export goFramebufferSizeCallback
func goFramebufferSizeCallback(w *C.GLFWwindow, width, height C.int) {
	real := (*Window)(C.glfwGetWindowUserPointer(w))
	real.framebuffersizeCallback(real, int(width), int(height))
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
	shareHandle := (*C.GLFWwindow)(nil)
	if share != nil {
		shareHandle = share.Handle()
	}

	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	handle := C.glfwCreateWindow(C.int(width), C.int(height), cTitle, (*C.GLFWmonitor)(monitor), shareHandle)
	if handle == nil {
		return nil, fmt.Errorf("failed to create window")
	}
	window := &Window{
		handle:                  handle,
		cursorPosCallback:       func(w *Window, x float64, y float64) {},
		sizeCallback:            func(w *Window, x int, y int) {},
		framebuffersizeCallback: func(w *Window, x int, y int) {},
	}

	C.glfwSetWindowUserPointer(handle, unsafe.Pointer(window))
	C.glfwSetCursorPosCallback_fix(handle)
	C.glfwSetWindowSizeCallback_fix(handle)
	C.glfwSetFramebufferSizeCallback_fix(handle)

	return window, nil
}

func SwapInterval(n int) {
	C.glfwSwapInterval(C.int(n))
}

func PollEvents() {
	C.glfwPollEvents()
}

func DetachCurrentContext() {
	C.glfwMakeContextCurrent(nil)
}

func GetPrimaryMonitor() *Monitor {
	return (*Monitor)(C.glfwGetPrimaryMonitor())
}

func WaitEvents() {
	C.glfwWaitEvents()
}

func PostEmptyEvent() {
	C.glfwPostEmptyEvent()
}

func (w *Window) Handle() *C.GLFWwindow {
	return (*C.GLFWwindow)(w.handle)
}

func (w *Window) MakeContextCurrent() {
	C.glfwMakeContextCurrent(w.Handle())
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

func (w *Window) GetCursorPos() (float64, float64) {
	var x, y C.double
	C.glfwGetCursorPos(w.Handle(), &x, &y)
	return float64(x), float64(y)
}

func (w *Window) SetCursorPosCallback(f func(ww *Window, x float64, y float64)) {
	w.cursorPosCallback = f
}

func (w *Window) SetSizeCallback(f func(ww *Window, width int, height int)) {
	w.sizeCallback = f
}

func (w *Window) SetFramebufferSizeCallback(f func(ww *Window, width int, height int)) {
	w.framebuffersizeCallback = f
}

func (w *Window) GetFramebufferSize() (int, int) {
	var width, height C.int
	C.glfwGetFramebufferSize(w.Handle(), &width, &height)
	return int(width), int(height)
}

func (w *Window) SetShouldClose(b bool) {
	if b {
		C.glfwSetWindowShouldClose(w.Handle(), C.GLFW_TRUE)
	} else {
		C.glfwSetWindowShouldClose(w.Handle(), C.GLFW_FALSE)
	}
}

func (w *Window) GetSize() (int, int) {
	var width, height C.int
	C.glfwGetWindowSize(w.Handle(), &width, &height)
	return int(width), int(height)
}

func (w *Window) SetPos(x, y int) {
	C.glfwSetWindowPos(w.Handle(), C.int(x), C.int(y))
}

func (w *Window) SetWindowUserPointer(ptr unsafe.Pointer) {
	w.userPtr = ptr
}

func (w *Window) GetWindowUserPointer() unsafe.Pointer {
	return w.userPtr
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
