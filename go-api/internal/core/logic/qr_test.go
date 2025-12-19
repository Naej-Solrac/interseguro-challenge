package logic

import (
	"math"
	"testing"
)

func TestCalculateQR(t *testing.T) {
	tests := []struct {
		name    string
		input   [][]float64
		wantErr bool
	}{
		{
			name: "Matriz 3x2 válida",
			input: [][]float64{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			wantErr: false,
		},
		{
			name: "Matriz 2x2 válida",
			input: [][]float64{
				{12, -51},
				{6, 167},
			},
			wantErr: false,
		},
		{
			name:    "Matriz vacía debe fallar",
			input:   [][]float64{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculateQR(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateQR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verificar que Q es ortogonal: Q^T * Q ≈ I
				if !isOrthogonal(result.Q, 1e-10) {
					t.Errorf("Q no es ortogonal")
				}

				// Verificar que R es triangular superior
				if !isUpperTriangular(result.R, 1e-10) {
					t.Errorf("R no es triangular superior")
				}

				// Verificar que Q * R ≈ A
				reconstructed := multiplyMatrices(result.Q, result.R)
				if !matricesEqual(reconstructed, tt.input, 1e-10) {
					t.Errorf("Q * R no reconstruye la matriz original")
				}
			}
		})
	}
}

// isOrthogonal verifica si Q^T * Q = I
func isOrthogonal(Q [][]float64, tolerance float64) bool {
	n := len(Q[0])
	qtq := multiplyMatrices(transpose(Q), Q)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			expected := 0.0
			if i == j {
				expected = 1.0
			}
			if math.Abs(qtq[i][j]-expected) > tolerance {
				return false
			}
		}
	}
	return true
}

// isUpperTriangular verifica si R es triangular superior
func isUpperTriangular(R [][]float64, tolerance float64) bool {
	for i := 0; i < len(R); i++ {
		for j := 0; j < i && j < len(R[i]); j++ {
			if math.Abs(R[i][j]) > tolerance {
				return false
			}
		}
	}
	return true
}

// transpose calcula la transpuesta de una matriz
func transpose(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])
	result := make([][]float64, cols)
	for i := range result {
		result[i] = make([]float64, rows)
		for j := 0; j < rows; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}

// multiplyMatrices multiplica dos matrices
func multiplyMatrices(A, B [][]float64) [][]float64 {
	rowsA := len(A)
	colsA := len(A[0])
	colsB := len(B[0])

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return result
}

// matricesEqual compara dos matrices con tolerancia
func matricesEqual(A, B [][]float64, tolerance float64) bool {
	if len(A) != len(B) {
		return false
	}
	for i := range A {
		if len(A[i]) != len(B[i]) {
			return false
		}
		for j := range A[i] {
			if math.Abs(A[i][j]-B[i][j]) > tolerance {
				return false
			}
		}
	}
	return true
}
