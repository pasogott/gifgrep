package ansi

import (
	"bufio"
	"bytes"
	"image"
	"image/png"
	"io"
)

func SendFrame(w io.Writer, pngData []byte, cols, rows int) error {
	data, err := RenderFrame(pngData, cols, rows)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func RenderFrame(pngData []byte, cols, rows int) ([]byte, error) {
	if cols <= 0 || rows <= 0 {
		return nil, nil
	}
	img, err := png.Decode(bytes.NewReader(pngData))
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	if bounds.Dx() <= 0 || bounds.Dy() <= 0 {
		return nil, nil
	}

	var buf bytes.Buffer
	out := bufio.NewWriter(&buf)
	targetH := rows * 2
	for y := 0; y < rows; y++ {
		var lastFG, lastBG rgb
		haveFG := false
		haveBG := false
		for x := 0; x < cols; x++ {
			top := sample(img, bounds, x, y*2, cols, targetH)
			bottom := sample(img, bounds, x, y*2+1, cols, targetH)
			if !haveFG || top != lastFG {
				writeSGRColor(out, "38", top)
				lastFG = top
				haveFG = true
			}
			if !haveBG || bottom != lastBG {
				writeSGRColor(out, "48", bottom)
				lastBG = bottom
				haveBG = true
			}
			_, _ = out.WriteString("▀")
		}
		_, _ = out.WriteString("\x1b[0m")
		if y != rows-1 {
			_, _ = out.WriteString("\r\n")
		}
	}
	if err := out.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type rgb struct {
	r byte
	g byte
	b byte
}

func writeSGRColor(out *bufio.Writer, prefix string, c rgb) {
	_, _ = out.WriteString("\x1b[")
	_, _ = out.WriteString(prefix)
	_, _ = out.WriteString(";2;")
	writeByte(out, c.r)
	_, _ = out.WriteString(";")
	writeByte(out, c.g)
	_, _ = out.WriteString(";")
	writeByte(out, c.b)
	_, _ = out.WriteString("m")
}

func sample(img image.Image, bounds image.Rectangle, tx, ty, targetW, targetH int) rgb {
	x := bounds.Min.X + (tx*bounds.Dx()+targetW/2)/targetW
	y := bounds.Min.Y + (ty*bounds.Dy()+targetH/2)/targetH
	if x >= bounds.Max.X {
		x = bounds.Max.X - 1
	}
	if y >= bounds.Max.Y {
		y = bounds.Max.Y - 1
	}
	r, g, b, _ := img.At(x, y).RGBA()
	return rgb{byte(r >> 8), byte(g >> 8), byte(b >> 8)}
}

func writeByte(out *bufio.Writer, b byte) {
	if b >= 100 {
		_ = out.WriteByte('0' + b/100)
		b %= 100
		_ = out.WriteByte('0' + b/10)
		_ = out.WriteByte('0' + b%10)
		return
	}
	if b >= 10 {
		_ = out.WriteByte('0' + b/10)
		_ = out.WriteByte('0' + b%10)
		return
	}
	_ = out.WriteByte('0' + b)
}
