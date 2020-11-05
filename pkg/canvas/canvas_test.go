package canvas

import (
	"reflect"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want *Canvas
	}{
		{"standard", args{10, 10}, NewCanvas(10, 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCanvas(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCanvas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanvas_WritePixel(t *testing.T) {
	type args struct {
		i     int
		j     int
		color Color
	}
	tests := []struct {
		name string
		c    Canvas
		args args
		want *Canvas
	}{
		{"standard", *NewCanvas(10, 10), args{3, 3, Color{0.3, 0.3, 0.3}}, NewCanvas(10, 10).WritePixel(3, 3, Color{0.3, 0.3, 0.3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.WritePixel(tt.args.i, tt.args.j, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Canvas.WritePixel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanvas_GetPixel(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		c    Canvas
		args args
		want *Color
	}{
		{"standard", *NewCanvas(10, 10).WritePixel(3, 3, Color{1, 1, 1}), args{3, 3}, &Color{1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetPixel(tt.args.i, tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Canvas.GetPixel() = %v, want %v", got, tt.want)
			}
		})
	}
}
