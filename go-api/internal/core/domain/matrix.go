package domain

// MatrixInput representa la matriz que entra
type MatrixInput struct {
	Data [][]float64 `json:"data"`
}

// QRFactorization representa la matriz que sale (QR)
type QRFactorization struct {
	Q [][]float64 `json:"Q"`
	R [][]float64 `json:"R"`
}

// FullResponse combina QR + Estadísticas de Node
type FullResponse struct {
	QR    QRFactorization `json:"qr_factorization"`
	Stats interface{}     `json:"statistics"`
}