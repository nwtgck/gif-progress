package gif_progress

import (
	"image"
	"image/color"
	"image/gif"
)

func AddProgressBar(inOutGif *gif.GIF, barTop bool, barHeight int, barColor color.RGBA) {
	// NOTE: inOutGif is changed destructively
	width := inOutGif.Config.Width
	height := inOutGif.Config.Height
	imageLen := len(inOutGif.Image)
	// Image size
	imageSize := image.Rect(0, 0, width, height)
	// Previous frame
	previousImage := inOutGif.Image[0].SubImage(imageSize)
	firstPalette := inOutGif.Image[0].Palette
	for i, paletted := range inOutGif.Image {
		// Create new empty paletted
		newPaletted := image.NewPaletted(imageSize, firstPalette)
		// Copy previous frame to the new paletted
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				newPaletted.Set(x, y, previousImage.At(x, y))
			}
		}
		// Copy whole image to the new paletted
		rect := paletted.Rect
		for x := rect.Min.X; x < rect.Max.X; x++ {
			for y := rect.Min.Y; y < rect.Max.Y; y++ {
				newPaletted.Set(x, y, paletted.At(x, y))
			}
		}
		// Save as previous
		previousImage = newPaletted.SubImage(imageSize)

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
