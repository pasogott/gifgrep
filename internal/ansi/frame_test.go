package ansi

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"strings"
	"testing"
)

func TestSendFrameTruecolorHalfBlocks(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 2))
	img.Set(0, 0, color.RGBA{R: 255, A: 255})
	img.Set(0, 1, color.RGBA{G: 255, A: 255})
	var pngBuf bytes.Buffer
	if err := png.Encode(&pngBuf, img); err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	if err := SendFrame(&out, pngBuf.Bytes(), 1, 1); err != nil {
		t.Fatal(err)
	}
	text := out.String()
	if !strings.Contains(text, "\x1b[38;2;255;0;0m") {
		t.Fatalf("missing foreground truecolor: %q", text)
	}
	if !strings.Contains(text, "\x1b[48;2;0;255;0m") {
		t.Fatalf("missing background truecolor: %q", text)
	}
	if !strings.Contains(text, "▀") {
		t.Fatalf("missing half block: %q", text)
	}
}

func TestRenderFrameReusesColorsAcrossRuns(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 4, 2))
	for y := 0; y < 2; y++ {
		for x := 0; x < 4; x++ {
			img.Set(x, y, color.RGBA{R: 10, G: 20, B: 30, A: 255})
		}
	}
	var pngBuf bytes.Buffer
	if err := png.Encode(&pngBuf, img); err != nil {
		t.Fatal(err)
	}

	data, err := RenderFrame(pngBuf.Bytes(), 4, 1)
	if err != nil {
		t.Fatal(err)
	}
	text := string(data)
	if got := strings.Count(text, "\x1b[38;2;10;20;30m"); got != 1 {
		t.Fatalf("foreground color emitted %d times: %q", got, text)
	}
	if got := strings.Count(text, "\x1b[48;2;10;20;30m"); got != 1 {
		t.Fatalf("background color emitted %d times: %q", got, text)
	}
}
