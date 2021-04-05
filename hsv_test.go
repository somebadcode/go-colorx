package colorx

import (
	"image/color"
	"testing"

	"github.com/somebadcode/go-mathx"
)

func TestHSVAModel(t *testing.T) {
	type args struct {
		c color.Color
	}
	tests := []struct {
		name string
		args args
		want HSVA
	}{
		{
			name: "rgba",
			args: args{
				c: color.RGBA{R: 0x80, G: 0x80},
			},
			want: HSVA{H: 60.0, S: 1.0, V: 0.5},
		},
		{
			name: "hsva",
			args: args{
				c: HSVA{H: 60.0, S: 1.0, V: 0.5},
			},
			want: HSVA{H: 60.0, S: 1.0, V: 0.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := HSVAModel.Convert(tt.args.c).(HSVA)
			if !ok {
				t.Errorf("HSVAModel.Convert() got V = %T, want %T", got, tt.want)
			}

			if !mathx.EqualP(got.H, tt.want.H, 1e-2) {
				t.Errorf("HSVAModel.Convert() got H = %f, want %f", got.H, tt.want.H)
			}
			if !mathx.EqualP(got.S, tt.want.S, 1e-2) {
				t.Errorf("HSVAModel.Convert() got S = %f, want %f", got.S, tt.want.S)
			}
			if !mathx.EqualP(got.V, tt.want.V, 1e-2) {
				t.Errorf("HSVAModel.Convert() got V = %f, want %f", got.V, tt.want.V)
			}
		})
	}
}

func TestRGBToHSV(t *testing.T) {
	type args struct {
		r uint8
		g uint8
		b uint8
	}
	tests := []struct {
		name string
		args args
		want HSVA
	}{
		{
			name: "black",
			args: args{},
			want: HSVA{},
		},
		{
			name: "white",
			args: args{r: 0xFF, g: 0xFF, b: 0xFF},
			want: HSVA{V: 1.0},
		},
		{
			name: "red",
			args: args{r: 0xFF},
			want: HSVA{S: 1.0, V: 1.0},
		},
		{
			name: "lime",
			args: args{g: 0xFF},
			want: HSVA{H: 120.0, S: 1.0, V: 1.0},
		},
		{
			name: "blue",
			args: args{b: 0xFF},
			want: HSVA{H: 240.0, S: 1.0, V: 1.0},
		},
		{
			name: "yellow",
			args: args{r: 0xFF, g: 0xFF},
			want: HSVA{H: 60.0, S: 1.0, V: 1.0},
		},
		{
			name: "cyan",
			args: args{g: 0xFF, b: 0xFF},
			want: HSVA{H: 180.0, S: 1.0, V: 1.0},
		},
		{
			name: "magenta",
			args: args{r: 0xFF, b: 0xFF},
			want: HSVA{H: 300.0, S: 1.0, V: 1.0},
		},
		{
			name: "silver",
			args: args{r: 0xBF, g: 0xBF, b: 0xBF},
			want: HSVA{V: 0.75},
		},
		{
			name: "gray",
			args: args{r: 0x80, g: 0x80, b: 0x80},
			want: HSVA{V: 0.5},
		},
		{
			name: "maroon",
			args: args{r: 0x80},
			want: HSVA{S: 1.0, V: 0.5},
		},
		{
			name: "olive",
			args: args{r: 0x80, g: 0x80},
			want: HSVA{H: 60.0, S: 1.0, V: 0.5},
		},
		{
			name: "green",
			args: args{g: 0x80},
			want: HSVA{H: 120.0, S: 1.0, V: 0.5},
		},
		{
			name: "purple",
			args: args{r: 0x80, b: 0x80},
			want: HSVA{H: 300.0, S: 1.0, V: 0.5},
		},
		{
			name: "teal",
			args: args{g: 0x80, b: 0x80},
			want: HSVA{H: 180.0, S: 1.0, V: 0.5},
		},
		{
			name: "navy",
			args: args{b: 0x80},
			want: HSVA{H: 240.0, S: 1.0, V: 0.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RGBAToHSVA(tt.args.r, tt.args.g, tt.args.b, 0)
			if !mathx.EqualP(got.H, tt.want.H, 1e-2) {
				t.Errorf("RGBAToHSVA() got H = %f, want %f", got.H, tt.want.H)
			}
			if !mathx.EqualP(got.S, tt.want.S, 1e-2) {
				t.Errorf("RGBAToHSVA() got S = %f, want %f", got.S, tt.want.S)
			}
			if !mathx.EqualP(got.V, tt.want.V, 1e-2) {
				t.Errorf("RGBAToHSVA() got V = %f, want %f", got.V, tt.want.V)
			}
		})
	}
}

func TestHSV_RGBA(t *testing.T) {
	type fields struct {
		H float64
		S float64
		V float64
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
				V: 1.0,
			},
			wantR: 0xFFFF,
			wantG: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "red",
			fields: fields{
				S: 1.0,
				V: 1.0,
			},
			wantR: 0xFFFF,
		},
		{
			name: "lime",
			fields: fields{
				H: 120.0,
				S: 1.0,
				V: 1.0,
			},
			wantG: 0xFFFF,
		},
		{
			name: "blue",
			fields: fields{
				H: 240.0,
				S: 1.0,
				V: 1.0,
			},
			wantB: 0xFFFF,
		},
		{
			name: "yellow",
			fields: fields{
				H: 60.0,
				S: 1.0,
				V: 1.0,
			},
			wantR: 0xFFFF,
			wantG: 0xFFFF,
		},
		{
			name: "cyan",
			fields: fields{
				H: 180.0,
				S: 1.0,
				V: 1.0,
			},
			wantG: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "cyan_negative",
			fields: fields{
				H: -180.0,
				S: 1.0,
				V: 1.0,
			},
			wantG: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "magenta",
			fields: fields{
				H: 300.0,
				S: 1.0,
				V: 1.0,
			},
			wantR: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "magenta_negative",
			fields: fields{
				H: -60.0,
				S: 1.0,
				V: 1.0,
			},
			wantR: 0xFFFF,
			wantB: 0xFFFF,
		},
		{
			name: "silver",
			fields: fields{
				V: 0.75,
			},
			wantR: 0xBFBF,
			wantG: 0xBFBF,
			wantB: 0xBFBF,
		},
		{
			name: "gray",
			fields: fields{
				V: 0.5,
			},
			wantR: 0x7F7F,
			wantG: 0x7F7F,
			wantB: 0x7F7F,
		},
		{
			name: "maroon",
			fields: fields{
				S: 1.0,
				V: 0.5,
			},
			wantR: 0x7F7F,
		},
		{
			name: "olive",
			fields: fields{
				H: 60.0,
				S: 1.0,
				V: 0.5,
			},
			wantR: 0x7F7F,
			wantG: 0x7F7F,
		},
		{
			name: "green",
			fields: fields{
				H: 120.0,
				S: 1.0,
				V: 0.5,
			},
			wantG: 0x7F7F,
		},
		{
			name: "purple",
			fields: fields{
				H: 300.0,
				S: 1.0,
				V: 0.5,
			},
			wantR: 0x7F7F,
			wantB: 0x7F7F,
		},
		{
			name: "teal",
			fields: fields{
				H: 180.0,
				S: 1.0,
				V: 0.5,
			},
			wantG: 0x7F7F,
			wantB: 0x7F7F,
		},
		{
			name: "navy",
			fields: fields{
				H: 240.0,
				S: 1.0,
				V: 0.5,
			},
			wantB: 0x7F7F,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsv := HSVA{
				H: tt.fields.H,
				S: tt.fields.S,
				V: tt.fields.V,
			}
			gotR, gotG, gotB, gotA := hsv.RGBA()
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

func BenchmarkRGBToHSV(b *testing.B) {
	var red uint8 = 0xBF
	var green uint8 = 0x0F
	var blue uint8 = 0x7F

	for n := 0; n < b.N; n++ {
		RGBAToHSVA(red, green, blue, 0)
	}
}

func BenchmarkHSV_RGBA(b *testing.B) {
	hsv := HSVA{
		H: 180,
		S: 0.5,
		V: 0.5,
	}

	for n := 0; n < b.N; n++ {
		hsv.RGBA()
	}
}
