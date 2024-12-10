package utils

import (
	"encoding/json"
	"log"
	"os"
)

var utilsLogger = log.New(os.Stdout, "UTILS: ", log.Ldate|log.Ltime|log.Lshortfile)

func JsonDecoder[T any](msg []byte, target *T) {
	if err := json.Unmarshal(msg, target); err != nil {
		utilsLogger.Printf("Error translating JSON: %v", err)
	}
}

func JsonEncoder[T any](body T) []byte {
	encodedJson, err := json.Marshal(body)
	if err != nil {
		utilsLogger.Fatalf("Error creating JSON: %v", err)
	}
	return encodedJson
}
