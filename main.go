package main

import (
	"github.com/fogleman/gg"
)

func main() {
	const S = 1024

	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./azuki.ttf", 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("フォントを描画します", S/2, S/2, 0.5, 0.5)
	dc.SavePNG("out.png")
}
