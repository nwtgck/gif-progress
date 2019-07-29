package cmd

import (
	gif_progress "github.com/nwtgck/gif-progress"
	"github.com/spf13/cobra"
	"gopkg.in/go-playground/colors.v1"
	"image/color"
	"image/gif"
	"os"
)

var barTop bool
var barHeight int
var barColorString string


func init() {
	cobra.OnInitialize()
	RootCmd.Flags().BoolVar(&barTop,  "bar-top", false, "Bar is on top")
	RootCmd.Flags().IntVar(&barHeight,  "bar-height", 5, "Bar height")
	RootCmd.Flags().StringVar(&barColorString,  "bar-color", "#ccc", "Bar color")
}

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "gif-progress",
	Long:  "Attach progress bar to animated GIF",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Parse color
		hex, err := colors.ParseHEX(barColorString)
		if err != nil {
			return err
		}
		c := hex.ToRGB()
		barColor := color.RGBA{c.R, c.G, c.B, 255}
		// TODO: Hard code stdin
		inOutGif, err := gif.DecodeAll(os.Stdin)
		if err != nil {
			return err
		}
		// Add progress bar to gif
		gif_progress.AddProgressBar(inOutGif, barTop, barHeight, barColor)
		// Write gif
		// TODO: Hard code stdout
		err = gif.EncodeAll(os.Stdout, inOutGif)
		return err
	},
}
