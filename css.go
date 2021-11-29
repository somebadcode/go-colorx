package colorx

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"math"
	"strconv"
)

// CSS an implementation of the color model used in Cascading Style Sheets.
type CSS struct {
	R       uint8   // Red
	G       uint8   // Green
	B       uint8   // Blue
	Opacity float64 // Opacity ∈ [0.0, 1.0]
}

// CSSModel can convert the color to the CSS color model defined in this package.
var CSSModel = color.ModelFunc(cssModel)

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

// RGBA returns the alpha-premultiplied red, green, blue and alpha values for the color.
func (c CSS) RGBA() (r, g, b, a uint32) {
	return color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: uint8(math.Min(c.SanitizedOpacity(), 1.0) * math.MaxUint8),
	}.RGBA()
}

// SanitizedOpacity returns the absolute value of opacity in the range [0.0, 1.0].
func (c CSS) SanitizedOpacity() float64 {
	return math.Min(math.Abs(c.Opacity), 1.0)
}

// String returns the color in its CSS string format, either "rgb" or "rgba".
func (c CSS) String() string {
	if c.SanitizedOpacity() < 1.0 {
		return fmt.Sprintf("rgba(%d,%d,%d,%s)", c.R, c.G, c.B, slimFloatString(c.SanitizedOpacity(), 2))
	}
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

// HexString returns the color in the hexadecimal format used by CSS.
func (c CSS) HexString() string {
	b := make([]byte, 0, 4)
	b = append(b, c.R, c.G, c.B)

	if c.SanitizedOpacity() < 1.0 {
		b = append(b, opacityUint8(c.SanitizedOpacity()))
	}

	return "#" + hex.EncodeToString(b)
}

// RGBAToCSS converts red, green, blue and alpha to red, green, blue and opacity.
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
