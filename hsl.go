package colorx

import (
	"image/color"
	"math"
)

// HSLA is an implementation of the HSV (Hue, Saturation and Value) color model. HSV is also known as HSB (Hue,
// Saturation, Brightness).
type HSLA struct {
	H float64 // Hue ∈ [0, 360)
	S float64 // Saturation ∈ [0, 1]
	L float64 // Lightness ∈ [0, 1]
	A float64 // Alpha ∈ [0, 1]
}

// HSLAModel can convert the color to the HSLA color model defined in this package.
var HSLAModel = color.ModelFunc(hslaModel)

func hslaModel(c color.Color) color.Color {
	if _, ok := c.(HSLA); ok {
		return c
	}

	r, g, b, a := c.RGBA()
	h, s, l, ha := RGBAToHSLA(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
	return HSLA{
		H: h,
		S: s,
		L: l,
		A: ha,
	}
}

// RGBAToHSLA converts RGBA to Hue, Saturation, Value and Alpha.
func RGBAToHSLA(r, g, b, a uint8) (float64, float64, float64, float64) {
	var hue, saturation, lightness, alpha float64

	// Convert R, G and B to floats.
	red := float64(r) / math.MaxUint8
	green := float64(g) / math.MaxUint8
	blue := float64(b) / math.MaxUint8
	alpha = float64(a) / math.MaxUint8

	// Get the most and least dominant colors.
	cMax := math.Max(red, math.Max(green, blue))
	cMin := math.Min(red, math.Min(green, blue))

	// Get color delta.
	delta := cMax - cMin

	// Value is the lightness of the dominant color.
	lightness = (cMax + cMin) / 2.0

	// Saturation is derived from the lightness, but it's zero if cMax is zero (saturation is initialized as zero).
	if delta != 0.0 {
		saturation = delta / (1.0 - math.Abs(2.0*lightness-1.0))
	}

	// Hue is derived from the dominant color.
	switch cMax {
	case cMin: // delta == 0
		hue = 0.0

	case red:
		hue = math.FMA(60.0, math.Mod((green-blue)/delta, 6), 360.0)

	case green:
		hue = math.FMA(60.0, (blue-red)/delta+2, 360.0)

	case blue:
		hue = math.FMA(60.0, (red-green)/delta+4, 360.0)
	}

	hue = math.Mod(hue, 360.0)

	return hue, saturation, lightness, alpha
}

// RGBA returns the alpha-premultiplied red, green, blue and alpha values for the color.
func (hsla HSLA) RGBA() (r, g, b, a uint32) {
	var rgba color.RGBA

	hsla.H = math.Mod(hsla.H+360.0, 360.0)

	c := (1 - math.Abs(2*hsla.L-1)) * hsla.S
	x := c * (1 - math.Abs(math.Mod(hsla.H/60.0, 2.0)-1))
	m := hsla.L - c/2

	rgba.A = uint8(hsla.A * math.MaxUint8)

	// sextant will be the sextant of the dominant color.
	sextant, _ := math.Modf(hsla.H / 60.0)

	switch int(sextant) {
	case 0:
		rgba.R = uint8(math.Floor((c + m) * math.MaxUint8))
		rgba.G = uint8(math.Floor((x + m) * math.MaxUint8))
		rgba.B = uint8(math.Floor(m * math.MaxUint8))

	case 1:
		rgba.R = uint8(math.Floor((x + m) * math.MaxUint8))
		rgba.G = uint8(math.Floor((c + m) * math.MaxUint8))
		rgba.B = uint8(math.Floor(m * math.MaxUint8))

	case 2:
		rgba.R = uint8(math.Floor(m * math.MaxUint8))
		rgba.G = uint8(math.Floor((c + m) * math.MaxUint8))
		rgba.B = uint8(math.Floor((x + m) * math.MaxUint8))

	case 3:
		rgba.R = uint8(math.Floor(m * math.MaxUint8))
		rgba.G = uint8(math.Floor((x + m) * math.MaxUint8))
		rgba.B = uint8(math.Floor((c + m) * math.MaxUint8))

	case 4:
		rgba.R = uint8(math.Floor((x + m) * math.MaxUint8))
		rgba.G = uint8(math.Floor(m * math.MaxUint8))
		rgba.B = uint8(math.Floor((c + m) * math.MaxUint8))

	default: // case 5
		rgba.R = uint8(math.Floor((c + m) * math.MaxUint8))
		rgba.G = uint8(math.Floor(m * math.MaxUint8))
		rgba.B = uint8(math.Floor((x + m) * math.MaxUint8))
	}

	return rgba.RGBA()
}
