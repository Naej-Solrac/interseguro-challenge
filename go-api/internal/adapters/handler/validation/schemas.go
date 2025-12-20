package validation

// LoginRequest representa el payload del endpoint /login
type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

// MatrixRequest representa el payload del endpoint /process
type MatrixRequest struct {
	Data [][]float64 `json:"data" validate:"required,min=1,dive,min=1"`
}
