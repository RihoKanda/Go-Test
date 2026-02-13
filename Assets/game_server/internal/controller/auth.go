package controller

import (
	"encoding/json"
	"net/http"

	"github.com/saf1o/go-test/internal/database"
	"github.com/saf1o/go-test/internal/model"
)

// LoginRequest ログインリクエスト
type LoginRequest struct {
	DeviceID string `json:"device_id"`
}

// LoginResponse ログインレスポンス
type LoginResponse struct {
	User      *model.User `json:"user"`
	IsNewUser bool        `json:"is_new_user"`
}

// HandleLogin ログイン・ユーザー作成
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.DeviceID == "" {
		SendError(w, http.StatusBadRequest, "device_id is required")
		return
	}

	// 既存ユーザーを検索
	user, err := model.GetUserByDeviceID(database.DB, req.DeviceID)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "Database error")
		return
	}

	isNewUser := false
	if user == nil {
		user, err = model.CreateUser(database.DB, req.DeviceID)
		if err != nil {
			SendError(w, http.StatusInternalServerError, "Failed to create user")
			return
		}
		isNewUser = true
	}

	SendSuccess(w, LoginResponse{
		User:      user,
		IsNewUser: isNewUser,
	})
}
