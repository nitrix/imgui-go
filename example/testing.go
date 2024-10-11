package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/nitrix/imgui-go"
)

var dummyBool bool
var dummyInt int
var dummyVec3 = mgl32.Vec3{0.5, 0.5, 0.5}

func testing() {
	imgui.Begin("Testing", nil, imgui.WindowFlags_None)
	imgui.Text("Hello %d!", 42)
	imgui.Checkbox("Check me", &dummyBool)

	if imgui.Button("Something", mgl32.Vec2{100, 32}) {
		fmt.Println("Clicked!")
	}

	imgui.SliderInt("Height", &dummyInt, 0, 100, "%d", imgui.SliderFlags_None)

	imgui.LabelText("foo##foo", "Hover me")
	if imgui.IsItemHovered(imgui.HoveredFlags_ForTooltip) {
		imgui.BeginTooltip()
		imgui.Text("I'm a tooltip")
		imgui.EndTooltip()
	}

	if imgui.ColorPicker3("Color", &dummyVec3, imgui.ColorEditFlags_None) {
		fmt.Println(dummyVec3)
	}

	imgui.End()
}
