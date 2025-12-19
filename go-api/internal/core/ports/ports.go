package ports

import "challenge-go/internal/core/domain"

// MatrixService define qué hace nuestro servicio (entrada)
type MatrixService interface {
	Process(input domain.MatrixInput) (domain.FullResponse, error)
}

// NodeStatsPort define cómo llamamos a Node (salida)
type NodeStatsPort interface {
	GetStats(Q, R [][]float64) (interface{}, error)
}