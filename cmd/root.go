package cmd

import (
	gif_progress "github.com/nwtgck/gif-progress"
	"github.com/spf13/cobra"
	"image/gif"
	"os"
)

var barTop bool
var barHeight int


func init() {
	cobra.OnInitialize()
	RootCmd.Flags().BoolVar(&barTop,  "bar-top", false, "Bar is on top")
	RootCmd.Flags().IntVar(&barHeight,  "bar-height", 5, "Bar height")
}

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "gif-progress",
	Long:  "Attache progress bar to animated Gif",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Hard code stdin
		inOutGif, err := gif.DecodeAll(os.Stdin)
		if err != nil {
			return err
		}
		// Add progress bar to gif
		gif_progress.AddProgressBar(inOutGif, barTop, barHeight)
		// Write gif
		// TODO: Hard code stdout
		err = gif.EncodeAll(os.Stdout, inOutGif)
		return err
	},
}
