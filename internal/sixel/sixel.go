package sixel

import (
	"bufio"
	"bytes"
	"image/png"

	gosixel "github.com/mattn/go-sixel"
	"github.com/steipete/gifgrep/gifdecode"
)

const (
	cellWidthPixels  = 8
	cellHeightPixels = 16
)

func SendFrame(out *bufio.Writer, frame gifdecode.Frame, cols, rows int) error {
	img, err := png.Decode(bytes.NewReader(frame.PNG))
	if err != nil {
		return err
	}
	enc := gosixel.NewEncoder(out)
	enc.Dither = true
	enc.Colors = 255
	if cols > 0 {
		enc.Width = cols * cellWidthPixels
	}
	if rows > 0 {
		enc.Height = rows * cellHeightPixels
	}
	return enc.Encode(img)
}
