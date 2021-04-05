// The HSVA color space implementation uses 64-bit floating point mathematics.
package colorx

import (
	"image/color"
	"math"

	"github.com/somebadcode/go-mathx"
)

// HSVA is an implementation of the HSV (Hue, Saturation and Value) color model. HSV is also known as HSB (Hue,
// Saturation, Brightness).
type HSVA struct {
	H float64 // Hue ∈ [0, 360)
	S float64 // Saturation ∈ [0, 1]
	V float64 // Value/Brightness ∈ [0, 1]
	A float64 // Alpha ∈ [0, 1]
}

var HSVAModel = color.ModelFunc(hsvaModel)

func hsvaModel(c color.Color) color.Color {
	if _, ok := c.(HSVA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBAToHSVA(
		uint8(r>>8),
		uint8(g>>8),
		uint8(b>>8),
		uint8(a>>8),
	)
}

// RGBAToHSVA converts RGBA to HSVA.
func RGBAToHSVA(r, g, b, a uint8) HSVA {
	var hsva HSVA
	var cMin, cMax float64
	var red, green, blue, alpha float64
	var delta float64

	// Convert R, G and B to floats.
	red = float64(r) / math.MaxUint8
	green = float64(g) / math.MaxUint8
	blue = float64(b) / math.MaxUint8
	alpha = float64(a) / math.MaxUint8

	// Get the most and least dominant colors.
	cMax = math.Max(red, math.Max(green, blue))
	cMin = math.Min(red, math.Min(green, blue))

	// Value is derived from the dominant color.
	hsva.V = cMax

	// Alpha requires no further transformation.
	hsva.A = alpha

	// Get color delta.
	delta = cMax - cMin

	// Saturation is derived from the delta but it's zero if cMax is zero (S is initialized as zero).
	if cMax != 0.0 {
		hsva.S = delta / cMax
	}

	// Hue is derived from the most dominant color.
	switch cMax {
	case cMin: // delta == 0
		hsva.H = 0.0

	case red:
		hsva.H = math.FMA(60.0, math.Mod((green-blue)/delta, 6), 360.0)

	case green:
		hsva.H = math.FMA(60.0, (blue-red)/delta+2, 360.0)

	case blue:
		hsva.H = math.FMA(60.0, (red-green)/delta+4, 360.0)
	}

	hsva.H = math.Mod(hsva.H, 360.0)

	return hsva
}

// RGBA returns the alpha-premultiplied red, green, blue and alpha values for the color.
func (hsva HSVA) RGBA() (r, g, b, a uint32) {
	var rgba color.RGBA
	var angle float64
	var f, i, p, q, t float64

	rgba.A = uint8(hsva.A * math.MaxUint8)

	if mathx.Equal(hsva.S, 0.0) {
		rgba.R = uint8(hsva.V * math.MaxUint8)
		rgba.G = uint8(hsva.V * math.MaxUint8)
		rgba.B = uint8(hsva.V * math.MaxUint8)
		return rgba.RGBA()
	}

	angle = math.Mod(hsva.H+360.0, 360.0)

	i, f = math.Modf(angle / 60.0)

	p = hsva.V * (1.0 - hsva.S)
	q = hsva.V * (1.0 - (hsva.S * f))
	t = hsva.V * (1.0 - (hsva.S * (1.0 - f)))

	switch i {
	case 0:
		rgba.R = uint8(math.Floor(hsva.V * math.MaxUint8))
		rgba.G = uint8(math.Floor(t * math.MaxUint8))
		rgba.B = uint8(math.Floor(p * math.MaxUint8))

	case 1:
		rgba.R = uint8(math.Floor(q * math.MaxUint8))
		rgba.G = uint8(math.Floor(hsva.V * math.MaxUint8))
		rgba.B = uint8(math.Floor(p * math.MaxUint8))

	case 2:
		rgba.R = uint8(math.Floor(p * math.MaxUint8))
		rgba.G = uint8(math.Floor(hsva.V * math.MaxUint8))
		rgba.B = uint8(math.Floor(t * math.MaxUint8))

	case 3:
		rgba.R = uint8(math.Floor(p * math.MaxUint8))
		rgba.G = uint8(math.Floor(q * math.MaxUint8))
		rgba.B = uint8(math.Floor(hsva.V * math.MaxUint8))

	case 4:
		rgba.R = uint8(math.Floor(t * math.MaxUint8))
		rgba.G = uint8(math.Floor(p * math.MaxUint8))
		rgba.B = uint8(math.Floor(hsva.V * math.MaxUint8))

	default: // 5
		rgba.R = uint8(math.Floor(hsva.V * math.MaxUint8))
		rgba.G = uint8(math.Floor(p * math.MaxUint8))
		rgba.B = uint8(math.Floor(q * math.MaxUint8))
	}

	return rgba.RGBA()
}
