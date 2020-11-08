package units

import (
	"math"
	"reflect"
	"testing"
)

func TestNewPoint(t *testing.T) {
	type args struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name string
		args args
		want *Tuple
	}{
		{"standard", args{4.3, -4.2, 3.1}, &Tuple{4.3, -4.2, 3.1, 1.}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPoint(tt.args.X, tt.args.Y, tt.args.Z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVector(t *testing.T) {
	type args struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name string
		args args
		want *Tuple
	}{
		{"standard", args{4.3, -4.2, 3.1}, &Tuple{4.3, -4.2, 3.1, 0.}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVector(tt.args.X, tt.args.Y, tt.args.Z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_IsPoint(t *testing.T) {
	tests := []struct {
		name string
		t    Tuple
		want bool
	}{
		{"standard", Tuple{4.3, -4.2, 3.1, 1.}, true},
		{"standard", Tuple{4.3, -4.2, 3.1, 0.}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.IsPoint(); got != tt.want {
				t.Errorf("Tuple.IsPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_IsVector(t *testing.T) {
	tests := []struct {
		name string
		t    Tuple
		want bool
	}{
		{"standard", Tuple{4.3, -4.2, 3.1, 0.}, true},
		{"standard", Tuple{4.3, -4.2, 3.1, 1.}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.IsVector(); got != tt.want {
				t.Errorf("Tuple.IsVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Add(t *testing.T) {
	type args struct {
		t1 *Tuple
	}
	tests := []struct {
		name string
		t    Tuple
		args args
		want *Tuple
	}{
		{"standard", Tuple{3, -2, 5, 1}, args{&Tuple{-2, 3, 1, 0}}, &Tuple{1, 1, 6, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Add(tt.args.t1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuple.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Substract(t *testing.T) {
	type args struct {
		t1 *Tuple
	}
	tests := []struct {
		name string
		t    Tuple
		args args
		want *Tuple
	}{
		{"standard", Tuple{3, -2, 5, 1}, args{&Tuple{-2, 3, 1, 0}}, &Tuple{5, -5, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Substract(tt.args.t1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuple.Substract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Negate(t *testing.T) {
	tests := []struct {
		name string
		t    Tuple
		want *Tuple
	}{
		{"standard", Tuple{3, -2, 5, 1}, &Tuple{-3, 2, -5, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Negate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuple.Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Multiply(t *testing.T) {
	type args struct {
		s float64
	}
	tests := []struct {
		name string
		t    Tuple
		args args
		want *Tuple
	}{
		{"standard", Tuple{1, -2, 3, -4}, args{3.5}, &Tuple{3.5, -7, 10.5, -14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Multiply(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuple.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Divide(t *testing.T) {
	type args struct {
		s float64
	}
	tests := []struct {
		name string
		t    Tuple
		args args
		want *Tuple
	}{
		{"standard", Tuple{1, -2, 3, -4}, args{2}, &Tuple{.5, -1, 1.5, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Divide(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuple.Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Magnitude(t *testing.T) {
	tests := []struct {
		name string
		t    Tuple
		want float64
	}{
		{"(1, 0, 0)", Tuple{1, 0, 0, 0}, 1},
		{"(0, 1, 0)", Tuple{0, 1, 0, 0}, 1},
		{"(0, 0, 1)", Tuple{0, 0, 1, 0}, 1},
		{"(1, 2, 3)", Tuple{1, 2, 3, 0}, math.Sqrt(14)},
		{"(-1, -2, -3)", Tuple{-1, -2, -3, 0}, math.Sqrt(14)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Magnitude(); got != tt.want {
				t.Errorf("Tuple.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Normalize(t *testing.T) {
	tests := []struct {
		name string
		t    Tuple
		want *Tuple
	}{
		{"(4, 0, 0)", Tuple{4, 0, 0, 0}, &Tuple{1, 0, 0, 0}},
		{"(1, 2, 3)", Tuple{1, 2, 3, 0}, &Tuple{1 / math.Sqrt(14), 2 / math.Sqrt(14), 3 / math.Sqrt(14), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Normalize(); !TupleEqual(got, tt.want) {
				t.Errorf("Tuple.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Dot(t *testing.T) {
	type args struct {
		t1 *Tuple
	}
	tests := []struct {
		name string
		t    Tuple
		args args
		want float64
	}{
		{"standard", *NewVector(1, 2, 3), args{NewVector(2, 3, 4)}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Dot(tt.args.t1); !FloatEqual(got, tt.want) {
				t.Errorf("Tuple.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Cross(t *testing.T) {
	type args struct {
		t1 *Tuple
	}
	tests := []struct {
		name string
		t    Tuple
		args args
		want *Tuple
	}{
		{"standard", *NewVector(1, 2, 3), args{NewVector(2, 3, 4)}, NewVector(-1, 2, -1)},
		{"inverse", *NewVector(2, 3, 4), args{NewVector(1, 2, 3)}, NewVector(1, -2, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Cross(tt.args.t1); !TupleEqual(got, tt.want) {
				t.Errorf("Tuple.Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_MatrixMultiply(t *testing.T) {
	type args struct {
		m *Matrix
	}
	tests := []struct {
		name string
		t    Tuple
		args args
		want *Tuple
	}{
		{"standard", *NewPoint(-3, 4, 5), args{TranslationMatrix(5, -3, 2)}, NewPoint(2, 1, 7)},
		{"inverse", *NewPoint(-3, 4, 5), args{TranslationMatrix(5, -3, 2).Invert()}, NewPoint(-8, 7, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.MatrixMultiply(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuple.MatrixMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Translate(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		t    Tuple
		args args
		want *Tuple
	}{
		{"standard", *NewPoint(-3, 4, 5), args{5, -3, 2}, NewPoint(2, 1, 7)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Translate(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuple.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
