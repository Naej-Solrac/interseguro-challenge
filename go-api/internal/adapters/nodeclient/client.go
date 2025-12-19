package nodeclient

import (
	"bytes"
	"challenge-go/internal/core/ports"
	"encoding/json"
	"fmt"
	"net/http"
)

type adapter struct {
	nodeURL string
}

// NewNodeAdapter crea el cliente
func NewNodeAdapter(url string) ports.NodeStatsPort {
	return &adapter{nodeURL: url}
}

// GetStats envía los datos a la API de Node.js
func (a *adapter) GetStats(Q, R [][]float64) (interface{}, error) {
	// Preparamos el JSON para enviar
	payload := map[string][][]float64{
		"Q": Q,
		"R": R,
	}
	jsonData, _ := json.Marshal(payload)

	// Hacemos el POST
	resp, err := http.Post(a.nodeURL+"/stats", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error conectando a Node: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("node respondió error: %d", resp.StatusCode)
	}

	// Leemos la respuesta de Node (las estadísticas)
	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}