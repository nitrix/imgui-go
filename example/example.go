package main

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/nitrix/imgui-go/glfw"
	"github.com/nitrix/imgui-go/imgui"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.False)

	window, err := glfw.CreateWindow(1280, 720, "Demo", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	glfw.SwapInterval(1)

	ctx, err := imgui.CreateContext()
	if err != nil {
		panic(err)
	}
	defer ctx.Destroy()

	ctx.IO.IniFilename = nil
	ctx.IO.ConfigFlags |= imgui.ConfigFlags_ViewportsEnable
	ctx.IO.ConfigFlags |= imgui.ConfigFlags_DockingEnable
	ctx.IO.ConfigDockingWithShift = true

	imgui.ImplGlfw_Init(window)
	defer imgui.ImplGlfw_Shutdown()

	imgui.ImplOpenGL3_Init(window)
	defer imgui.ImplOpenGL3_Shutdown()

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	})

	window.Centerize()
	window.Show()

	for !window.ShouldClose() {
		glfw.PollEvents()

		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		imgui.ImplGlfw_NewFrame()
		imgui.ImplOpenGL3_NewFrame()
		imgui.NewFrame()

		imgui.DockSpaceOverViewport(0, nil, imgui.DockNodeFlags_PassthruCentralNode, nil)

		imgui.ShowDemoWindow()
		imgui.Render()

		if ctx.IO.ConfigFlags&imgui.ConfigFlags_ViewportsEnable > 0 {
			previous := glfw.GetCurrentContext()
			imgui.UpdatePlatformWindows()
			imgui.RenderPlatformWindowsDefault(nil, nil)
			previous.MakeContextCurrent()
		}

		imgui.ImplOpenGL3_RenderDrawData(imgui.GetDrawData())

		window.SwapBuffers()
	}
}
