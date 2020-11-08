package canvas

import (
	"reflect"
	"testing"
)

func TestColor_Add(t *testing.T) {
	type args struct {
		c1 *Color
	}
	tests := []struct {
		name string
		c    Color
		args args
		want *Color
	}{
		{"standard", Color{0.9, 0.6, 0.75}, args{&Color{0.7, 0.1, 0.25}}, &Color{1.6, 0.7, 1.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Add(tt.args.c1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Substract(t *testing.T) {
	type args struct {
		c1 *Color
	}
	tests := []struct {
		name string
		c    Color
		args args
		want *Color
	}{
		{"standard", Color{0.9, 0.6, 0.75}, args{&Color{0.5, 0.1, 0.25}}, &Color{.4, .5, .5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Substract(tt.args.c1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.Substract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Scale(t *testing.T) {
	type args struct {
		s float64
	}
	tests := []struct {
		name string
		c    Color
		args args
		want *Color
	}{
		{"standard", Color{0.9, 0.6, 0.75}, args{2}, &Color{1.8, 1.2, 1.5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Scale(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Multiply(t *testing.T) {
	type args struct {
		c1 *Color
	}
	tests := []struct {
		name string
		c    Color
		args args
		want *Color
	}{
		{"standard", Color{1, 0.2, 5.}, args{&Color{0.9, 1, 0.1}}, &Color{0.9, 0.2, 0.5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Multiply(tt.args.c1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Clamp(t *testing.T) {
	tests := []struct {
		name string
		c    Color
		want *Color
	}{
		{"standard", Color{1.5, -0.3, 5}, &Color{1., 0, 1.}},
		{"Test values greater than 1", Color{1.5, 3.3, 5}, &Color{1., 1., 1.}},
		{"Test values less than 1", Color{-1, -0.1, -3.}, &Color{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Clamp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
