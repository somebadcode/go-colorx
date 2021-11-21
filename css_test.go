package colorx

import (
	"image/color"
	"reflect"
	"strings"
	"testing"
)

func TestCSSModel(t *testing.T) {
	type args struct {
		c color.Color
	}
	tests := []struct {
		name string
		args args
		want CSS
	}{
		{
			name: "rgba",
			args: args{
				c: color.RGBA{R: 0x80, G: 0x80},
			},
			want: CSS{R: 0x80, G: 0x80},
		},
		{
			name: "css",
			args: args{
				c: CSS{R: 0x80, G: 0x80, B: 0x80},
			},
			want: CSS{R: 0x80, G: 0x80, B: 0x80},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := CSSModel.Convert(tt.args.c).(CSS)
			if !ok {
				t.Errorf("CSSModel.Convert() got = %T, want %T", got, tt.want)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CSSModel.Convert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSS_HexString(t *testing.T) {
	type fields struct {
		r uint8
		g uint8
		b uint8
		a float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "black",
			fields: fields{a: 1.0},
			want:   "#000000",
		},
		{
			name:   "white",
			fields: fields{r: 0xFF, g: 0xFF, b: 0xFF, a: 1.0},
			want:   "#FFFFFF",
		},
		{
			name:   "red",
			fields: fields{r: 0xFF, a: 1.0},
			want:   "#FF0000",
		},
		{
			name:   "lime",
			fields: fields{g: 0xFF, a: 1.0},
			want:   "#00FF00",
		},
		{
			name:   "blue",
			fields: fields{b: 0xFF, a: 1.0},
			want:   "#0000FF",
		},
		{
			name:   "yellow",
			fields: fields{r: 0xFF, g: 0xFF, a: 1.0},
			want:   "#FFFF00",
		},
		{
			name:   "cyan",
			fields: fields{g: 0xFF, b: 0xFF, a: 1.0},
			want:   "#00FFFF",
		},
		{
			name:   "magenta",
			fields: fields{r: 0xFF, b: 0xFF, a: 1.0},
			want:   "#FF00FF",
		},
		{
			name:   "silver",
			fields: fields{r: 0xBF, g: 0xBF, b: 0xBF, a: 1.0},
			want:   "#BFBFBF",
		},
		{
			name:   "gray",
			fields: fields{r: 0x80, g: 0x80, b: 0x80, a: 1.0},
			want:   "#808080",
		},
		{
			name:   "maroon",
			fields: fields{r: 0x80, a: 1.0},
			want:   "#800000",
		},
		{
			name:   "olive",
			fields: fields{r: 0x80, g: 0x80, a: 1.0},
			want:   "#808000",
		},
		{
			name:   "green",
			fields: fields{g: 0x80, a: 1.0},
			want:   "#008000",
		},
		{
			name:   "purple",
			fields: fields{r: 0x80, b: 0x80, a: 1.0},
			want:   "#800080",
		},
		{
			name:   "teal",
			fields: fields{g: 0x80, b: 0x80, a: 1.0},
			want:   "#008080",
		},
		{
			name:   "navy",
			fields: fields{b: 0x80, a: 1.0},
			want:   "#000080",
		},
		{
			name:   "navy_alpha",
			fields: fields{b: 0x80, a: 0.5},
			want:   "#00008080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CSS{
				R:       tt.fields.r,
				G:       tt.fields.g,
				B:       tt.fields.b,
				Opacity: tt.fields.a,
			}
			if got := c.HexString(); !strings.EqualFold(got, tt.want) {
				t.Errorf("HexString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSS_RGBA(t *testing.T) {
	type fields struct {
		r uint8
		g uint8
		b uint8
		a float64
	}
	tests := []struct {
		name   string
		fields fields
		wantR  uint32
		wantG  uint32
		wantB  uint32
		wantA  uint32
	}{
		{
			name:   "black",
			fields: fields{},
		},
		{
			name: "white",
			fields: fields{
				r: 0xFF,
				g: 0xFF,
				b: 0xFF,
			},
			wantR: 0xFFFF,
			wantG: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "red",
			fields: fields{
				r: 0xFF,
			},
			wantR: 0xFFFF,
		},
		{
			name: "lime",
			fields: fields{
				g: 0xFF,
			},
			wantG: 0xFFFF,
		},
		{
			name: "blue",
			fields: fields{
				b: 0xFF,
			},
			wantB: 0xFFFF,
		},
		{
			name: "yellow",
			fields: fields{
				r: 0xFF,
				g: 0xFF,
			},
			wantR: 0xFFFF,
			wantG: 0xFFFF,
		},
		{
			name: "cyan",
			fields: fields{
				g: 0xFF,
				b: 0xFF,
			},
			wantG: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "magenta",
			fields: fields{
				r: 0xFF,
				b: 0xFF,
			},
			wantR: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "silver",
			fields: fields{
				r: 0xBF,
				g: 0xBF,
				b: 0xBF,
			},
			wantR: 0xBFBF,
			wantG: 0xBFBF,
			wantB: 0xBFBF,
		},
		{
			name: "gray",
			fields: fields{
				r: 0x7F,
				g: 0x7F,
				b: 0x7F,
			},
			wantR: 0x7F7F,
			wantG: 0x7F7F,
			wantB: 0x7F7F,
		},
		{
			name: "maroon",
			fields: fields{
				r: 0x7F,
			},
			wantR: 0x7F7F,
		},
		{
			name: "olive",
			fields: fields{
				r: 0x7F,
				g: 0x7F,
			},
			wantR: 0x7F7F,
			wantG: 0x7F7F,
		},
		{
			name: "green",
			fields: fields{
				g: 0x7F,
			},
			wantG: 0x7F7F,
		},
		{
			name: "purple",
			fields: fields{
				r: 0x7F,
				b: 0x7F,
			},
			wantR: 0x7F7F,
			wantB: 0x7F7F,
		},
		{
			name: "teal",
			fields: fields{
				g: 0x7F,
				b: 0x7F,
			},
			wantG: 0x7F7F,
			wantB: 0x7F7F,
		},
		{
			name: "navy",
			fields: fields{
				b: 0x7F,
			},
			wantB: 0x7F7F,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CSS{
				R:       tt.fields.r,
				G:       tt.fields.g,
				B:       tt.fields.b,
				Opacity: tt.fields.a,
			}
			gotR, gotG, gotB, gotA := c.RGBA()
			if gotR != tt.wantR {
				t.Errorf("RGBA() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("RGBA() gotG = %v, want %v", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("RGBA() gotB = %v, want %v", gotB, tt.wantB)
			}
			if gotA != tt.wantA {
				t.Errorf("RGBA() gotA = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func TestCSS_String(t *testing.T) {
	type fields struct {
		r uint8
		g uint8
		b uint8
		a float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "black",
			fields: fields{a: 1.0},
			want:   "rgb(0,0,0)",
		},
		{
			name:   "white",
			fields: fields{r: 0xFF, g: 0xFF, b: 0xFF, a: 1.0},
			want:   "rgb(255,255,255)",
		},
		{
			name:   "red",
			fields: fields{r: 0xFF, a: 1.0},
			want:   "rgb(255,0,0)",
		},
		{
			name:   "lime",
			fields: fields{g: 0xFF, a: 1.0},
			want:   "rgb(0,255,0)",
		},
		{
			name:   "blue",
			fields: fields{b: 0xFF, a: 1.0},
			want:   "rgb(0,0,255)",
		},
		{
			name:   "yellow",
			fields: fields{r: 0xFF, g: 0xFF, a: 1.0},
			want:   "rgb(255,255,0)",
		},
		{
			name:   "cyan",
			fields: fields{g: 0xFF, b: 0xFF, a: 1.0},
			want:   "rgb(0,255,255)",
		},
		{
			name:   "magenta",
			fields: fields{r: 0xFF, b: 0xFF, a: 1.0},
			want:   "rgb(255,0,255)",
		},
		{
			name:   "silver",
			fields: fields{r: 0xBF, g: 0xBF, b: 0xBF, a: 1.0},
			want:   "rgb(191,191,191)",
		},
		{
			name:   "gray",
			fields: fields{r: 0x80, g: 0x80, b: 0x80, a: 1.0},
			want:   "rgb(128,128,128)",
		},
		{
			name:   "maroon",
			fields: fields{r: 0x80, a: 1.0},
			want:   "rgb(128,0,0)",
		},
		{
			name:   "olive",
			fields: fields{r: 0x80, g: 0x80, a: 1.0},
			want:   "rgb(128,128,0)",
		},
		{
			name:   "green",
			fields: fields{g: 0x80, a: 1.0},
			want:   "rgb(0,128,0)",
		},
		{
			name:   "purple",
			fields: fields{r: 0x80, b: 0x80, a: 1.0},
			want:   "rgb(128,0,128)",
		},
		{
			name:   "teal",
			fields: fields{g: 0x80, b: 0x80, a: 1.0},
			want:   "rgb(0,128,128)",
		},
		{
			name:   "navy",
			fields: fields{b: 0x80, a: 1.0},
			want:   "rgb(0,0,128)",
		},
		{
			name:   "navy_alpha",
			fields: fields{b: 0x80, a: 0.5},
			want:   "rgba(0,0,128,0.5)",
		},
		{
			name:   "navy_alpha",
			fields: fields{b: 0x80, a: 0.543},
			want:   "rgba(0,0,128,0.54)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CSS{
				R:       tt.fields.r,
				G:       tt.fields.g,
				B:       tt.fields.b,
				Opacity: tt.fields.a,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRGBAToCSS(t *testing.T) {
	type args struct {
		r uint8
		g uint8
		b uint8
		a uint8
	}
	tests := []struct {
		name        string
		args        args
		wantR       uint8
		wantG       uint8
		wantB       uint8
		wantOpacity float64
	}{
		{
			args: args{
				r: 0xFF,
				g: 0xBF,
				b: 0x80,
				a: 0x80,
			},
			wantR:       0xFF,
			wantG:       0xBF,
			wantB:       0x80,
			wantOpacity: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB, gotOpacity := RGBAToCSS(tt.args.r, tt.args.g, tt.args.b, tt.args.a)
			if gotR != tt.wantR {
				t.Errorf("RGBAToCSS() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("RGBAToCSS() gotG = %v, want %v", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("RGBAToCSS() gotB = %v, want %v", gotB, tt.wantB)
			}
			if gotOpacity != tt.wantOpacity {
				t.Errorf("RGBAToCSS() gotOpacity = %v, want %v", gotOpacity, tt.wantOpacity)
			}
		})
	}
}

func BenchmarkCSS_HexString(b *testing.B) {
	c := CSS{
		R:       0xFF,
		G:       0xBF,
		B:       0x80,
		Opacity: 0.5,
	}

	for n := 0; n < b.N; n++ {
		_ = c.HexString()
	}
}
