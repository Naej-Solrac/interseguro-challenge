package logic

import (
	"challenge-go/internal/core/domain"
	"errors"
	"math"
)

func CalculateQR(matrix [][]float64) (domain.QRFactorization, error) {
	if len(matrix) == 0 {
		return domain.QRFactorization{}, errors.New("la matriz no puede estar vacía")
	}

	m := len(matrix)
	n := len(matrix[0])

	// Inicializar Q y R
	Q := make([][]float64, m)
	for i := range Q {
		Q[i] = make([]float64, n)
	}
	R := make([][]float64, n)
	for i := range R {
		R[i] = make([]float64, n)
	}

	// Copia temporal transpuesta
	a := make([][]float64, n)
	for j := 0; j < n; j++ {
		a[j] = make([]float64, m)
		for i := 0; i < m; i++ {
			a[j][i] = matrix[i][j]
		}
	}

	// Gram-Schmidt
	for k := 0; k < n; k++ {
		norm := 0.0
		for i := 0; i < m; i++ {
			norm += a[k][i] * a[k][i]
		}
		R[k][k] = math.Sqrt(norm)

		for i := 0; i < m; i++ {
			if R[k][k] != 0 {
				Q[i][k] = a[k][i] / R[k][k]
			} else {
				Q[i][k] = 0
			}
		}

		for j := k + 1; j < n; j++ {
			dotProduct := 0.0
			for i := 0; i < m; i++ {
				dotProduct += Q[i][k] * a[j][i]
			}
			R[k][j] = dotProduct

			for i := 0; i < m; i++ {
				a[j][i] = a[j][i] - R[k][j]*Q[i][k]
			}
		}
	}

	return domain.QRFactorization{Q: Q, R: R}, nil
}