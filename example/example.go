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

	imgui.ImplGlfw_Init(window)
	defer imgui.ImplGlfw_Shutdown()

	imgui.ImplOpenGL3_Init(window)
	defer imgui.ImplOpenGL3_Shutdown()

	registerCallbacks(ctx, window)

	window.Centerize()
	window.Show()

	for !window.ShouldClose() {
		glfw.PollEvents()

		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		imgui.ImplGlfw_NewFrame()
		imgui.ImplOpenGL3_NewFrame()
		imgui.NewFrame()

		imgui.ShowDemoWindow()
		imgui.Render()

		imgui.ImplOpenGL3_RenderDrawData(imgui.GetDrawData())

		window.SwapBuffers()
	}
}

func registerCallbacks(ctx *imgui.Context, window *glfw.Window) {
	/*
		window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode glfw.Scancode, action glfw.Action, mods glfw.ModifierKey) {
			if ctx.IO.WantCaptureKeyboard {
				return
			}

			if key == glfw.KeyEscape && action == glfw.Press {
				w.SetShouldClose(true)
			}
		})

		window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
			if ctx.IO.WantCaptureMouse {
				return
			}

			if button == glfw.MouseButtonLeft && action == glfw.Press {
				x, y := w.GetCursorPos()
				fmt.Println("Clicked at:", x, y)
			}
		})

		window.SetScrollCallback(func(w *glfw.Window, xoff float64, yoff float64) {
			if ctx.IO.WantCaptureMouse {
				return
			}

			fmt.Println("Scrolled:", xoff, yoff)
		})

		window.SetCursorPosCallback(func(w *glfw.Window, x float64, y float64) {
			if ctx.IO.WantCaptureMouse {
				return
			}

			fmt.Println("Cursor at:", x, y)
		})
	*/

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	})
}
