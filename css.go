package colorx

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"math"
	"strconv"
)

// CSS is a model for easy use of colors in Cascading Style Sheets.
type CSS struct {
	R       uint8
	G       uint8
	B       uint8
	Opacity float64
}

var (
	CSSModel = color.ModelFunc(cssModel)
)

func cssModel(c color.Color) color.Color {
	if _, ok := c.(CSS); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return CSS{
		R:       uint8(r >> 8),
		G:       uint8(g >> 8),
		B:       uint8(b >> 8),
		Opacity: float64(a>>8) / math.MaxUint8,
	}
}

func (c CSS) RGBA() (r, g, b, a uint32) {
	return color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: uint8(math.Min(c.SanitizedOpacity(), 1.0) * math.MaxUint8),
	}.RGBA()
}

func (c CSS) SanitizedOpacity() float64 {
	return math.Min(math.Abs(c.Opacity), 1.0)
}

// String returns the color in its CSS string format, either "rgb" or "rgba".
func (c *CSS) String() string {
	if c.SanitizedOpacity() < 1.0 {
		return fmt.Sprintf("rgba(%d,%d,%d,%s)", c.R, c.G, c.B, slimFloatString(c.SanitizedOpacity(), 2))
	}
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

// HexString returns the color in the hexadecimal format used by CSS.
func (c *CSS) HexString() string {
	b := make([]byte, 0, 4)
	b = append(b, c.R, c.G, c.B)

	if c.SanitizedOpacity() < 1.0 {
		b = append(b, opacityUint8(c.SanitizedOpacity()))
	}

	return "#" + hex.EncodeToString(b)
}

// RGBAToCSS converts Red, Green, Blue and Alpha to Red, Green, Blue and Opacity.
func RGBAToCSS(r, g, b, a uint8) (uint8, uint8, uint8, float64) {
	return r, g, b, math.Floor((float64(a)/float64(math.MaxUint8))*100.0) / 100.0
}

func slimFloatString(f float64, e int) string {
	if f*math.Pow10(e) == math.Floor(f*math.Pow10(e)) {
		return strconv.FormatFloat(f, 'f', -1, 64)
	}
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func opacityUint8(f float64) uint8 {
	return uint8(math.Round(f * float64(math.MaxUint8)))
}
