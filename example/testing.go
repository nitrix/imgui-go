package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/nitrix/imgui-go"
)

var b bool

func testing() {
	imgui.Begin("Testing", nil, imgui.WindowFlags_None)
	imgui.Text("Hello %d!", 42)
	imgui.Checkbox("Check me", &b)
	if imgui.Button("Something", mgl32.Vec2{100, 32}) {
		fmt.Println("Clicked!")
	}
	imgui.End()
}
