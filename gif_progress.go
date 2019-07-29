package gif_progress

import (
	"image/color"
	"image/gif"
)

func AddProgressBar(inOutGif *gif.GIF, barTop bool, barHeight int) {
	// NOTE: inOutGif is changed destructively
	width := inOutGif.Config.Width
	height := inOutGif.Config.Height
	image_len := len(inOutGif.Image)
	for i, paletted := range inOutGif.Image {
		w := int(float32(width) * ((float32(i)+1)/float32(image_len)))
		for x := 0; x < w; x++ {
			for h := 0; h < barHeight; h++ {
				var y = h
				if !barTop {
					y = height - h
				}
				// TODO: Hard code bar color
				paletted.Set(x, y, color.RGBA{204, 204, 204, 255})
			}
		}
	}
}
