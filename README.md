# colorx

This package implements the following color models:

### HSV/HSL - Hue, Saturation, Value
HSV is implemented through the concrete type `HSVA`.

### CSS - Cascading Style Sheets
A variation of the RGBA color model where the alpha/opacity is stored as a floating point number between 0.0 and 1.0.
This allows you to work with the colors using the `colors` package and convert it to this special color model that can
give you strings that specify the color according to CSS specifications.

Examples:
- `rgb(255,191,128)`
- `rgba(255,191,128,0.75)`
- `#FFBF80`
- `#FFBF80BF`
