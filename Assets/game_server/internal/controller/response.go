package controller

import (
	"encoding/json"
	"net/http"
)

// APIResponse 共通APIレスポンス構造体
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SendJSON JSONレスポンスを送信
func SendJSON(w http.ResponseWriter, statusCode int, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// SendSuccess 成功レスポンスを送信
func SendSuccess(w http.ResponseWriter, data interface{}) {
	SendJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
	})
}

// SendError エラーレスポンスを送信
func SendError(w http.ResponseWriter, statusCode int, message string) {
	SendJSON(w, statusCode, APIResponse{
		Success: false,
		Error:   message,
	})
}
