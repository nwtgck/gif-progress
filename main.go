
package main

import (
	"fmt"
	"image/color"
	"image/gif"
	"os"
)

// (base: http://tech.nitoyon.com/ja/blog/2016/01/07/go-animated-gif-gen/)
// TODO: Handle error
func main() {
	// TODO: Hard code name
	f, _ := os.Open("rgb.gif")
	defer f.Close()
	inGif, _ := gif.DecodeAll(f)
	width := inGif.Config.Width
	height := inGif.Config.Height
	image_len := len(inGif.Image)
	bar_top := false
	// TODO: Hard code
	bar_height := 5
	for i, paletted := range inGif.Image {
		w := int(float32(width) * ((float32(i)+1)/float32(image_len)))
		fmt.Println(w)
		for x := 0; x < w; x++ {
			for h := 0; h < bar_height; h++ {
				var y = h
				if !bar_top {
					y = height - h
				}
				// TODO: Hard code bar color
				paletted.Set(x, y, color.RGBA{204, 204, 204, 255})
			}
		}
	}
	// TODO: Hard code name
	f2, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	_ = gif.EncodeAll(f2, inGif)
}
