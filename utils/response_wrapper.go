package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// Message function
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Response function
func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(data)
}

func TokenResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(data)
}

// LogError to log error handling
func LogError(function string, str string, err error) {
	log.Printf("[%v] %v : %v", function, str, err)
}
