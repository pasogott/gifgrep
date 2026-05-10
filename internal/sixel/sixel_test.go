package sixel

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"image/png"
	"strings"
	"testing"

	"github.com/steipete/gifgrep/gifdecode"
)

func TestSendFrameEmitsSixelDCS(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{R: 255, A: 255})
	img.Set(1, 0, color.RGBA{G: 255, A: 255})
	img.Set(0, 1, color.RGBA{B: 255, A: 255})
	img.Set(1, 1, color.RGBA{R: 255, G: 255, A: 255})

	var pngBuf bytes.Buffer
	if err := png.Encode(&pngBuf, img); err != nil {
		t.Fatalf("encode png: %v", err)
	}

	var outBuf bytes.Buffer
	out := bufio.NewWriter(&outBuf)
	if err := SendFrame(out, gifdecode.Frame{PNG: pngBuf.Bytes()}, 4, 2); err != nil {
		t.Fatalf("send frame: %v", err)
	}
	_ = out.Flush()

	got := outBuf.String()
	if !strings.Contains(got, "\x1bP") || !strings.Contains(got, "q") || !strings.Contains(got, "\x1b\\") {
		t.Fatalf("expected sixel DCS output, got %q", got)
	}
}
