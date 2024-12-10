package utils

import (
	"encoding/json"
	"log"
)

func JsonDecoder[T any](msg []byte, target *T) {
	if err := json.Unmarshal(msg, target); err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}
}
