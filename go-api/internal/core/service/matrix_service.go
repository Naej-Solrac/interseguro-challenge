package service

import (
	"challenge-go/internal/core/domain"
	"challenge-go/internal/core/logic"
	"challenge-go/internal/core/ports"
)

// service implementa la interfaz MatrixService
type service struct {
	nodePort ports.NodeStatsPort // Dependencia: Alguien que hable con Node
}

// NewService es el constructor
func NewService(n ports.NodeStatsPort) ports.MatrixService {
	return &service{
		nodePort: n,
	}
}

// Process es la función principal que orquesta todo
func (s *service) Process(input domain.MatrixInput) (domain.FullResponse, error) {
	// 1. Usamos la lógica pura para calcular QR
	qrResult, err := logic.CalculateQR(input.Data)
	if err != nil {
		return domain.FullResponse{}, err
	}

	// 2. Usamos el puerto para pedir estadísticas a Node
	stats, err := s.nodePort.GetStats(qrResult.Q, qrResult.R)
	if err != nil {
		return domain.FullResponse{}, err
	}

	// 3. Devolvemos todo junto
	return domain.FullResponse{
		QR:    qrResult,
		Stats: stats,
	}, nil
}