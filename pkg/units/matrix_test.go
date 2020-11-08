package units

import (
	"reflect"
	"testing"
)

func TestDefaultMatrix(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{"standard", args{10, 10}, DefaultMatrix(10, 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultMatrix(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrix(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			"standard",
			args{[][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			}},
			NewMatrix([][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16}},
			)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Dot(t *testing.T) {
	type args struct {
		m1 *Matrix
	}
	tests := []struct {
		name    string
		m       Matrix
		args    args
		want    *Matrix
		wantErr bool
	}{
		{
			"standard",
			*NewMatrix([][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2}}),
			args{
				NewMatrix([][]float64{
					{-2, 1, 2, 3},
					{3, 2, 1, -1},
					{4, 3, 6, 5},
					{1, 2, 7, 8},
				})},
			NewMatrix([][]float64{
				{20, 22, 50, 48},
				{44, 54, 114, 108},
				{40, 58, 110, 102},
				{16, 26, 46, 42},
			}),
			false,
		},
		{
			"Different matrix size",
			*NewMatrix([][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2}}),
			args{
				NewMatrix([][]float64{
					{-2, 1, 2, 3},
					{3, 2, 1, -1},
					{4, 3, 6, 5},
				})},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Dot(tt.args.m1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Dot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copyMatrix(t *testing.T) {
	type args struct {
		m *Matrix
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			"standard",
			args{
				NewMatrix([][]float64{
					{-2, 1, 2, 3},
					{3, 2, 1, -1},
					{4, 3, 6, 5},
					{1, 2, 7, 8},
				})},
			NewMatrix([][]float64{
				{-2, 1, 2, 3},
				{3, 2, 1, -1},
				{4, 3, 6, 5},
				{1, 2, 7, 8},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyMatrix(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentityMatrix(t *testing.T) {
	tests := []struct {
		name string
		want *Matrix
	}{
		{
			"standard",
			NewMatrix([][]float64{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IdentityMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdentityMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Transpose(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want *Matrix
	}{
		{
			"standard",
			*NewMatrix([][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			}),
			NewMatrix([][]float64{
				{1, 5, 9, 13},
				{2, 6, 10, 14},
				{3, 7, 11, 15},
				{4, 8, 12, 16},
			}),
		},
		{
			"uneven matrix",
			*NewMatrix([][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			}),
			NewMatrix([][]float64{
				{1, 5, 9},
				{2, 6, 10},
				{3, 7, 11},
				{4, 8, 12},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Submatrix(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want *Matrix
	}{
		{
			"3 x 3",
			*NewMatrix([][]float64{
				{1, 5, 0},
				{-3, 2, 7},
				{0, 6, -3},
			}),
			args{0, 2},
			NewMatrix([][]float64{
				{-3, 2},
				{0, 6},
			}),
		},
		{
			"4 x 4",
			*NewMatrix([][]float64{
				{-6, 1, 1, 6},
				{-8, 5, 8, 6},
				{-1, 0, 8, 2},
				{-7, 1, -1, 1},
			}),
			args{2, 1},
			NewMatrix([][]float64{
				{-6, 1, 6},
				{-8, 8, 6},
				{-7, -1, 1},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Submatrix(tt.args.i, tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Submatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Minor(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want float64
	}{
		{
			"3 x 3",
			*NewMatrix([][]float64{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			}),
			args{0, 0},
			-12,
		},
		{
			"3 x 3",
			*NewMatrix([][]float64{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			}),
			args{1, 0},
			25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Minor(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Matrix.Minor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Cofactor(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want float64
	}{
		{
			"3 x 3",
			*NewMatrix([][]float64{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			}),
			args{0, 0},
			56,
		},
		{
			"3 x 3",
			*NewMatrix([][]float64{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			}),
			args{0, 0},
			-12,
		},
		{
			"3 x 3",
			*NewMatrix([][]float64{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			}),
			args{1, 0},
			-25,
		},
		{
			"4 x 4",
			*NewMatrix([][]float64{
				{-2, -8, 3, 5},
				{-3, 1, 7, 3},
				{1, 2, -9, 6},
				{-6, 7, 7, -9},
			}),
			args{0, 0},
			690,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Cofactor(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Matrix.Cofactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Determinant(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want float64
	}{
		{
			"2 x 2",
			*NewMatrix([][]float64{{1, 5}, {-3, 2}}),
			17,
		},
		{
			"3 x 3",
			*NewMatrix([][]float64{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			}),
			-196,
		},
		{
			"4 x 4",
			*NewMatrix([][]float64{
				{-5, 2, 6, -8},
				{1, -5, 1, 8},
				{7, 7, -6, -7},
				{1, -3, 7, 4},
			}),
			532,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Determinant(); got != tt.want {
				t.Errorf("Matrix.Determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_TupleMultiply(t *testing.T) {
	type args struct {
		t *Tuple
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want *Tuple
	}{
		{"standard", *TranslationMatrix(5, -3, 2), args{NewPoint(-3, 4, 5)}, NewPoint(2, 1, 7)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.TupleMultiply(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.TupleMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Invert(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want *Matrix
	}{
		{
			"standard",
			*NewMatrix([][]float64{
				{8, -5, 9, 2},
				{7, 5, 6, 1},
				{-6, 0, 9, 6},
				{-3, 0, -9, -4},
			}),
			NewMatrix([][]float64{
				{-0.153846, -0.153846, -0.282051, -0.538462},
				{-0.076923, 0.123077, 0.025641, 0.030769},
				{0.358974, 0.358974, 0.435897, 0.923077},
				{-0.692308, -0.692308, -0.769231, -1.923077},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Invert(); !got.Equal(tt.want) {
				t.Errorf("Matrix.Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslationMatrix(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TranslationMatrix(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TranslationMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
