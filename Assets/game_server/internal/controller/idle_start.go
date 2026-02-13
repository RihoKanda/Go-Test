package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/saf1o/go-test/internal/database"
	"github.com/saf1o/go-test/internal/model"
)

// IdleStartRequest 放置開始リクエスト
type IdleStartRequest struct {
	UserID int `json:"user_id"`
}

// IdleStartResponse 放置開始レスポンス
type IdleStartResponse struct {
	User      *model.User `json:"user"`
	StartedAt time.Time   `json:"started_at"`
}

// HandleIdleStart 放置開始
func HandleIdleStart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req IdleStartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// ユーザー取得
	var user model.User
	err := database.DB.QueryRow(`
		SELECT 
		    user_id, 
		    device_id,
			user_name,
		    level,
		    exp,
			attack_up,
			speed_up,
			hp_regen_up,
			evolution_stage,
			is_idle,
		    idle_started_at,
			created_at,
			updated_at
		FROM users WHERE user_id = ?
	`, req.UserID).Scan(
		&user.UserID,
		&user.DeviceID,
		&user.UserName,
		&user.Level,
		&user.Exp,
		&user.AttackUp,
		&user.SpeedUp,
		&user.HPRegenUp,
		//&user.EvolutionStage,
		&user.IsIdle,
		&user.IdleStartedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		SendError(w, http.StatusNotFound, "User not found")
		return
	}
	// 放置中の場合はエラー
	if user.IsIdle {
		SendError(w, http.StatusBadRequest, "Already in idle mode")
		return
	}

	// 放置開始
	now := time.Now()
	user.IsIdle = true
	user.IdleStartedAt = &now

	// ユーザー情報更新
	err = model.UpdateUser(database.DB, &user)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	SendSuccess(w, IdleStartResponse{
		User:      &user,
		StartedAt: now,
	})
}
