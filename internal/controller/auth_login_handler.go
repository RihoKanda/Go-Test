package controller

import (
	"encoding/json"
	"net/http"

	"github.com/saf1o/go-test/internal/model"
)

type LoginRequest struct {
	DeviceID string `json:"device_id"`
}

type LoginResponse struct {
	UserID int   `json:"user_id"`
	Level  int   `json:"level"`
	Exp    int64 `json:"exp"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	user, err := model.FindUserByDeviceID(req.DeviceID)
	if err != nil {
		user, err = model.CreateUser(req.DeviceID)
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
	}

	res := LoginResponse{
		UserID: user.UserID,
		Level:  user.Level,
		Exp:    user.Exp,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
