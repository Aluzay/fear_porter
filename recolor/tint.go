package recolor

import (
	"image"
	"image/color"

	"github.com/swim-services/swim_porter/utils"
)

type Tint struct {
	color color.RGBA
}

func NewTint(color color.RGBA) *Tint {
	return &Tint{color: color}
}

func (t *Tint) SetColor(color color.RGBA) {
	t.color = color
}

func (t *Tint) RecolorImage(in image.Image, fileName string) (image.Image, error) {
	bounds := in.Bounds()
	dst := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := in.At(x, y)
			r, g, b, a := pixelColor.RGBA()
			a >>= 8
			// Skip fully transparent pixels
			if a == 0 {
				continue
			}
			r >>= 8
			g >>= 8
			b >>= 8
			r = r * uint32(t.color.R) / 255 // Actual red
			g = g * uint32(t.color.G) / 255 // Actual green
			b = b * uint32(t.color.B) / 255 // Actual blue
			dst.SetRGBA(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}
	return dst, nil
}

func (t *Tint) DefaultList() []string {
	return utils.DEFAULT_RECOLOR_LIST
}
