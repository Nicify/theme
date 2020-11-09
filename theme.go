package theme

import (
	"image/color"

	"github.com/AllenDang/giu/imgui"
)

type EffectMethod struct {
	Push func()
	Pop  func()
}

type Palette struct {
	Tint   color.RGBA
	Hover  color.RGBA
	Active color.RGBA
	BG     color.RGBA
}

func SetThemeDark(style *imgui.Style) {
	style.SetColor(imgui.StyleColorText, RGBAToVec4(color.RGBA{230, 230, 230, 255}))
	style.SetColor(imgui.StyleColorWindowBg, RGBAToVec4(color.RGBA{50, 50, 50, 250}))
	style.SetColor(imgui.StyleColorFrameBg, RGBAToVec4(color.RGBA{10, 10, 10, 240}))
	style.SetColor(imgui.StyleColorFrameBgHovered, RGBAToVec4(color.RGBA{20, 20, 20, 250}))
	style.SetColor(imgui.StyleColorFrameBgActive, RGBAToVec4(color.RGBA{25, 25, 25, 250}))
	style.SetColor(imgui.StyleColorButton, RGBAToVec4(color.RGBA{100, 100, 100, 255}))
	style.SetColor(imgui.StyleColorButtonHovered, RGBAToVec4(color.RGBA{120, 120, 120, 240}))
	style.SetColor(imgui.StyleColorButtonActive, RGBAToVec4(color.RGBA{80, 80, 80, 245}))
	style.SetColor(imgui.StyleColorTab, RGBAToVec4(color.RGBA{0, 0, 0, 0}))
	style.SetColor(imgui.StyleColorTabActive, RGBAToVec4(color.RGBA{18, 150, 219, 255}))
	style.SetColor(imgui.StyleColorTabHovered, RGBAToVec4(color.RGBA{18, 150, 219, 255}))
	style.SetColor(imgui.StyleColorScrollbarBg, RGBAToVec4(color.RGBA{27, 27, 27, 250}))
	style.SetColor(imgui.StyleColorScrollbarGrab, RGBAToVec4(color.RGBA{73, 73, 73, 255}))
	style.SetColor(imgui.StyleColorScrollbarGrabActive, RGBAToVec4(color.RGBA{164, 164, 164, 255}))
	style.SetColor(imgui.StyleColorScrollbarGrabHovered, RGBAToVec4(color.RGBA{140, 140, 140, 250}))
	style.SetColor(imgui.StyleColorSliderGrab, RGBAToVec4(color.RGBA{73, 73, 73, 255}))
	style.SetColor(imgui.StyleColorSliderGrabActive, RGBAToVec4(color.RGBA{164, 164, 164, 255}))
	style.SetColor(imgui.StyleColorCheckMark, RGBAToVec4(color.RGBA{18, 150, 219, 255}))
}

func SetThemeLight(style *imgui.Style) {
	style.SetColor(imgui.StyleColorText, RGBAToVec4(color.RGBA{60, 71, 89, 255}))
	style.SetColor(imgui.StyleColorWindowBg, RGBAToVec4(color.RGBA{240, 240, 240, 255}))
	style.SetColor(imgui.StyleColorChildBg, RGBAToVec4(color.RGBA{229, 229, 229, 255}))
	style.SetColor(imgui.StyleColorFrameBg, RGBAToVec4(color.RGBA{229, 229, 229, 255}))
	style.SetColor(imgui.StyleColorFrameBgHovered, RGBAToVec4(color.RGBA{210, 210, 210, 255}))
	style.SetColor(imgui.StyleColorFrameBgActive, RGBAToVec4(color.RGBA{215, 215, 215, 255}))
	style.SetColor(imgui.StyleColorButton, RGBAToVec4(color.RGBA{87, 190, 252, 255}))
	style.SetColor(imgui.StyleColorButtonHovered, RGBAToVec4(color.RGBA{80, 185, 248, 255}))
	style.SetColor(imgui.StyleColorButtonActive, RGBAToVec4(color.RGBA{77, 180, 242, 255}))
	style.SetColor(imgui.StyleColorTab, RGBAToVec4(color.RGBA{0, 0, 0, 0}))
	style.SetColor(imgui.StyleColorTabActive, RGBAToVec4(color.RGBA{87, 190, 252, 255}))
	style.SetColor(imgui.StyleColorTabHovered, RGBAToVec4(color.RGBA{87, 190, 252, 255}))
	style.SetColor(imgui.StyleColorScrollbarBg, RGBAToVec4(color.RGBA{255, 255, 255, 0}))
	style.SetColor(imgui.StyleColorScrollbarGrab, RGBAToVec4(color.RGBA{87, 190, 252, 230}))
	style.SetColor(imgui.StyleColorScrollbarGrabActive, RGBAToVec4(color.RGBA{87, 190, 252, 245}))
	style.SetColor(imgui.StyleColorScrollbarGrabHovered, RGBAToVec4(color.RGBA{87, 190, 252, 255}))
	style.SetColor(imgui.StyleColorSliderGrab, RGBAToVec4(color.RGBA{87, 190, 252, 255}))
	style.SetColor(imgui.StyleColorSliderGrabActive, RGBAToVec4(color.RGBA{77, 180, 242, 255}))
	style.SetColor(imgui.StyleColorCheckMark, RGBAToVec4(color.RGBA{87, 190, 252, 255}))
}

func UseLayoutFlat() EffectMethod {
	return EffectMethod{
		func() {
			imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0)
			imgui.PushStyleVarFloat(imgui.StyleVarFrameBorderSize, 0)
			imgui.PushStyleVarFloat(imgui.StyleVarChildBorderSize, 1)
			imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 0)
			imgui.PushStyleVarFloat(imgui.StyleVarChildRounding, 0)
			imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 0)
			imgui.PushStyleVarFloat(imgui.StyleVarTabRounding, 0)
			imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{X: 8, Y: 6})
		},
		func() { imgui.PopStyleVarV(8) },
	}
}

func UseStyleButtonDark() EffectMethod {
	return EffectMethod{
		func() {
			imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 2)
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 0, Y: 0})
			imgui.PushStyleColor(imgui.StyleColorButton, imgui.Vec4{X: 0.125, Y: 0.125, Z: 0.125, W: 1})
			imgui.PushStyleColor(imgui.StyleColorButtonHovered, imgui.Vec4{X: 0.125, Y: 0.125, Z: 0.125, W: 0.90})
			imgui.PushStyleColor(imgui.StyleColorButtonActive, imgui.Vec4{X: 0.125, Y: 0.125, Z: 0.125, W: 0.95})
		},
		func() {
			imgui.PopStyleVarV(2)
			imgui.PopStyleColorV(3)
		},
	}
}

func UseStyleButtonGhost() EffectMethod {
	return EffectMethod{
		func() {
			imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 2)
			imgui.PushStyleColor(imgui.StyleColorButton, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
			imgui.PushStyleColor(imgui.StyleColorButtonHovered, imgui.Vec4{X: 0.90, Y: 0.90, Z: 0.90, W: 240})
			imgui.PushStyleColor(imgui.StyleColorButtonActive, imgui.Vec4{X: 0.85, Y: 0.85, Z: 0.85, W: 250})
		},
		func() {
			imgui.PopStyleVarV(1)
			imgui.PopStyleColorV(3)
		},
	}
}

func UseStyleButtonLight() EffectMethod {
	return EffectMethod{
		func() {
			imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 0)
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 0, Y: 0})
			imgui.PushStyleColor(imgui.StyleColorText, RGBAToVec4(color.RGBA{0, 0, 0, 255}))
			imgui.PushStyleColor(imgui.StyleColorButton, RGBAToVec4(color.RGBA{229, 229, 229, 255}))
			imgui.PushStyleColor(imgui.StyleColorButtonHovered, RGBAToVec4(color.RGBA{210, 210, 210, 255}))
			imgui.PushStyleColor(imgui.StyleColorButtonActive, RGBAToVec4(color.RGBA{215, 215, 215, 255}))
		},
		func() {
			imgui.PopStyleVarV(2)
			imgui.PopStyleColorV(4)
		},
	}
}

func UseFont(font imgui.Font) EffectMethod {
	return EffectMethod{
		func() {
			imgui.PushFont(font)
		},
		func() {
			imgui.PopFont()
		},
	}
}
