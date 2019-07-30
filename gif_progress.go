package gif_progress

import (
	"image"
	"image/color"
	"image/gif"
)

func clonePix(pix []uint8) []uint8 {
	p := make([]uint8, len(pix))
	copy(p, pix)
	return p
}

func AddProgressBar(inOutGif *gif.GIF, barTop bool, barHeight int, barColor color.RGBA) {
	// NOTE: inOutGif is changed destructively
	width := inOutGif.Config.Width
	height := inOutGif.Config.Height
	imageLen := len(inOutGif.Image)
	// Image size
	imageSize := image.Rect(0, 0, width, height)
	// Previous pixels
	previousPix := inOutGif.Image[0].Pix
	for i, paletted := range inOutGif.Image {
		// Create new empty paletted
		newPaletted := image.NewPaletted(imageSize, paletted.Palette)
		// Copy previous frame to the new paletted
		newPaletted.Pix = previousPix
		// Copy diff to the new paletted
		rect := paletted.Rect
		for x := rect.Min.X; x < rect.Max.X; x++ {
			for y := rect.Min.Y; y < rect.Max.Y; y++ {
				newPaletted.Set(x, y, paletted.At(x, y))
			}
		}
		// Save as previous
		previousPix = clonePix(newPaletted.Pix)

		// Attach progress bar
		w := int(float32(width) * ((float32(i)+1)/float32(imageLen)))
		for x := 0; x < w; x++ {
			for h := 0; h < barHeight; h++ {
				var y = h
				if !barTop {
					y = height - h
				}
				newPaletted.Set(x, y, barColor)
			}
		}

		// Update image to new paletted
		inOutGif.Image[i] = newPaletted
	}
}
