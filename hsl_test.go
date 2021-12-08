package colorx

import (
	"image/color"
	"testing"

	"github.com/somebadcode/go-colorx/v2/internal/mathx"
)

func TestHSLAModel(t *testing.T) {
	type args struct {
		c color.Color
	}
	tests := []struct {
		name string
		args args
		want HSLA
	}{
		{
			name: "rgba",
			args: args{
				c: color.RGBA{
					R: 128,
					G: 128,
					B: 128,
					A: 128,
				},
			},
			want: HSLA{
				L: 0.5,
				A: 0.5,
			},
		},
		{
			name: "hsla",
			args: args{
				c: HSLA{
					L: 0.5,
					A: 0.5,
				},
			},
			want: HSLA{
				L: 0.5,
				A: 0.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := HSLAModel.Convert(tt.args.c).(HSLA)
			if !ok {
				t.Errorf("HSLAModel.Convert() got L = %T, want %T", got, tt.want)
			}

			if !mathx.EqualP(got.H, tt.want.H, 1e-2) {
				t.Errorf("HSLAModel.Convert() got H = %f, want %f", got.H, tt.want.H)
			}
			if !mathx.EqualP(got.S, tt.want.S, 1e-2) {
				t.Errorf("HSLAModel.Convert() got S = %f, want %f", got.S, tt.want.S)
			}
			if !mathx.EqualP(got.L, tt.want.L, 1e-2) {
				t.Errorf("HSLAModel.Convert() got L = %f, want %f", got.L, tt.want.L)
			}
		})
	}
}

func TestRGBToHSLA(t *testing.T) {
	type args struct {
		r uint8
		g uint8
		b uint8
		a uint8
	}
	tests := []struct {
		name  string
		args  args
		wantH float64
		wantS float64
		wantL float64
		wantA float64
	}{
		{
			name:  "black",
			args:  args{0, 0, 0, 255},
			wantA: 1.0,
		},
		{
			name:  "white",
			args:  args{255, 255, 255, 255},
			wantL: 1.0,
			wantA: 1.0,
		},
		{
			name:  "red",
			args:  args{255, 0, 0, 255},
			wantS: 1.0,
			wantL: 0.5,
			wantA: 1.0,
		},
		{
			name:  "lime",
			args:  args{0, 255, 0, 255},
			wantH: 120.0,
			wantS: 1.0,
			wantL: 0.5,
			wantA: 1.0,
		},
		{
			name:  "blue",
			args:  args{0, 0, 255, 255},
			wantH: 240.0,
			wantS: 1.0,
			wantL: 0.5,
			wantA: 1.0,
		},
		{
			name:  "yellow",
			args:  args{255, 255, 0, 255},
			wantH: 60.0,
			wantS: 1.0,
			wantL: 0.5,
			wantA: 1.0,
		},
		{
			name:  "cyan",
			args:  args{0, 255, 255, 255},
			wantH: 180.0,
			wantS: 1.0,
			wantL: 0.5,
			wantA: 1.0,
		},
		{
			name:  "magenta",
			args:  args{255, 0, 255, 255},
			wantH: 300.0,
			wantS: 1.0,
			wantL: 0.5,
			wantA: 1.0,
		},
		{
			name:  "silver",
			args:  args{191, 191, 191, 255},
			wantL: 0.75,
			wantA: 1.0,
		},
		{
			name:  "gray",
			args:  args{128, 128, 128, 255},
			wantL: 0.50,
			wantA: 1.0,
		},
		{
			name:  "maroon",
			args:  args{128, 0, 0, 255},
			wantS: 1.0,
			wantL: 0.25,
			wantA: 1.0,
		},
		{
			name:  "olive",
			args:  args{128, 128, 0, 255},
			wantH: 60.0,
			wantS: 1.0,
			wantL: 0.25,
			wantA: 1.0,
		},
		{
			name:  "green",
			args:  args{0, 128, 0, 255},
			wantH: 120.0,
			wantS: 1.0,
			wantL: 0.25,
			wantA: 1.0,
		},
		{
			name:  "purple",
			args:  args{128, 0, 128, 255},
			wantH: 300.0,
			wantS: 1.0,
			wantL: 0.25,
			wantA: 1.0,
		},
		{
			name:  "teal",
			args:  args{0, 128, 128, 255},
			wantH: 180.0,
			wantS: 1.0,
			wantL: 0.25,
			wantA: 1.0,
		},
		{
			name:  "navy",
			args:  args{0, 0, 128, 255},
			wantH: 240.0,
			wantS: 1.0,
			wantL: 0.25,
			wantA: 1.0,
		},
		{
			name:  "navy_alpha",
			args:  args{0, 0, 128, 128},
			wantH: 240.0,
			wantS: 1.0,
			wantL: 0.25,
			wantA: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotH, gotS, gotV, gotA := RGBAToHSLA(tt.args.r, tt.args.g, tt.args.b, tt.args.a)
			if !mathx.EqualP(gotH, tt.wantH, 1e-2) {
				t.Errorf("RGBAToHSLA() got H = %f, want %f", gotH, tt.wantH)
			}
			if !mathx.EqualP(gotS, tt.wantS, 1e-2) {
				t.Errorf("RGBAToHSLA() got S = %f, want %f", gotS, tt.wantS)
			}
			if !mathx.EqualP(gotV, tt.wantL, 1e-2) {
				t.Errorf("RGBAToHSLA() got L = %f, want %f", gotV, tt.wantL)
			}
			if !mathx.EqualP(gotA, tt.wantA, 1e-2) {
				t.Errorf("RGBAToHSLA() got A = %f, want %f", gotA, tt.wantA)
			}
		})
	}
}

func TestHSLA_RGBA(t *testing.T) {
	type fields struct {
		H float64
		S float64
		L float64
		A float64
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
				L: 1.0,
			},
			wantR: 0xFFFF,
			wantG: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "red",
			fields: fields{
				S: 1.0,
				L: 0.5,
			},
			wantR: 0xFFFF,
		},
		{
			name: "lime",
			fields: fields{
				H: 120.0,
				S: 1.0,
				L: 0.5,
			},
			wantG: 0xFFFF,
		},
		{
			name: "blue",
			fields: fields{
				H: 240.0,
				S: 1.0,
				L: 0.5,
			},
			wantB: 0xFFFF,
		},
		{
			name: "yellow",
			fields: fields{
				H: 60.0,
				S: 1.0,
				L: 0.5,
			},
			wantR: 0xFFFF,
			wantG: 0xFFFF,
		},
		{
			name: "cyan",
			fields: fields{
				H: 180.0,
				S: 1.0,
				L: 0.5,
			},
			wantG: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "cyan_negative",
			fields: fields{
				H: -180.0,
				S: 1.0,
				L: 0.5,
			},
			wantG: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "magenta",
			fields: fields{
				H: 300.0,
				S: 1.0,
				L: 0.5,
			},
			wantR: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "magenta_negative",
			fields: fields{
				H: -60.0,
				S: 1.0,
				L: 0.5,
			},
			wantR: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "silver",
			fields: fields{
				L: 0.75,
			},
			wantR: 0xBFBF,
			wantG: 0xBFBF,
			wantB: 0xBFBF,
		},
		{
			name: "gray",
			fields: fields{
				L: 0.5,
			},
			wantR: 0x7F7F,
			wantG: 0x7F7F,
			wantB: 0x7F7F,
		},
		{
			name: "maroon",
			fields: fields{
				S: 1.0,
				L: 0.25,
			},
			wantR: 0x7F7F,
		},
		{
			name: "olive",
			fields: fields{
				H: 60.0,
				S: 1.0,
				L: 0.25,
			},
			wantR: 0x7F7F,
			wantG: 0x7F7F,
		},
		{
			name: "green",
			fields: fields{
				H: 120.0,
				S: 1.0,
				L: 0.25,
			},
			wantG: 0x7F7F,
		},
		{
			name: "purple",
			fields: fields{
				H: 300.0,
				S: 1.0,
				L: 0.25,
			},
			wantR: 0x7F7F,
			wantB: 0x7F7F,
		},
		{
			name: "teal",
			fields: fields{
				H: 180.0,
				S: 1.0,
				L: 0.25,
			},
			wantG: 0x7F7F,
			wantB: 0x7F7F,
		},
		{
			name: "navy",
			fields: fields{
				H: 240.0,
				S: 1.0,
				L: 0.25,
			},
			wantB: 0x7F7F,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsla := HSLA{
				H: tt.fields.H,
				S: tt.fields.S,
				L: tt.fields.L,
			}
			gotR, gotG, gotB, gotA := hsla.RGBA()
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

func BenchmarkRGBToHSLA(b *testing.B) {
	var red uint8 = 0xBF
	var green uint8 = 0x0F
	var blue uint8 = 0x7F

	for n := 0; n < b.N; n++ {
		RGBAToHSLA(red, green, blue, 0)
	}
}

func BenchmarkHSL_RGBA(b *testing.B) {
	hsl := HSLA{
		H: 180,
		S: 0.5,
		L: 0.5,
	}

	for n := 0; n < b.N; n++ {
		hsl.RGBA()
	}
}
