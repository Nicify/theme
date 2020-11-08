package theme

import (
	"image/color"

	"github.com/AllenDang/giu/imgui"
)

func RGBAToVec4(rgba color.RGBA) imgui.Vec4 {
	return imgui.Vec4{
		X: float32(rgba.R) / 255,
		Y: float32(rgba.G) / 255,
		Z: float32(rgba.B) / 255,
		W: float32(rgba.A) / 255,
	}
}
