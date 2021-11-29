# colorx - extra color models
![Lint & Test](https://github.com/somebadcode/go-colorx/actions/workflows/main.yml/badge.svg)

This package contains extra color models to complement the models in `image/color`. By using a different color model
than the default `color.RGBA` model, you can specify or modify the color in a different way.

**Import:** `github.com/somebadcode/go-colorx/v2`

### HSV/HSB - Hue, Saturation, Value/Brightness
HSV is implemented through the concrete type `HSVA`. The HSV color model was designed to more closely reflect how
humans perceive colors. You can use HSV to generate gradients or a color palette for graphs that are easier to perceive
for colorblind users. You can also use this color model to make images monochrome or have more saturated colors.

### CSS - Cascading Style Sheets
A variation of the RGBA color model where the alpha/opacity is stored as a floating point number between 0.0 and 1.0.
This allows you to work with the colors using the `image/colors` package and convert it to this special color model that
can give you strings that specify the color according to CSS specifications. The `CSS` color model implements the
stringer interface.

Examples:
- `rgb(255,191,128)`
- `rgba(255,191,128,0.75)`
- `#FFBF80`
- `#FFBF80BF`
