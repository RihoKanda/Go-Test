package controller

import (
	"net/http"
	"strconv"

	"github.com/saf1o/go-test/internal/database"
	"github.com/saf1o/go-test/internal/model"
)

// GameStateResponse ゲーム状態レスポンス
type GameStateResponse struct {
	User   *model.User   `json:"user"`
	Enmies []model.Enemy `json:"enemies"`
}

// HandleGameState ゲーム状態取得
func HandleGameState(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// クエリパラメータからuser_idを取得
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		SendError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		SendError(w, http.StatusBadRequest, "invalid user_id")
		return
	}

	// ユーザー情報取得
	var user model.User
	err = database.DB.QueryRow(`
		SELECT 
		    user_id, 
			device_id,
			user_name,
			level, 
			exp,
			attack_up,
			speed_up,
			hp_regen_up,
			evolition_stage,
			is_idle,
			idle_started_at,
			created_at,
			updated_at 
		FROM users WHERE user_id = ?
	`, userID).Scan(
		&user.UserID,
		&user.DeviceID,
		&user.UserName,
		&user.Level,
		&user.Exp,
		&user.AttackUp,
		&user.SpeedUp,
		&user.HPRegenUp,
		&user.EvolitionStage,
		&user.IsIdle,
		&user.IdleStartedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		SendError(w, http.StatusNotFound, "User not found")
		return
	}

	// 敵一覧取得
	enemies, err := model.GetAllEnemies(database.DB)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "Failed to retrieve enemies")
		return
	}

	SendSuccess(w, GameStateResponse{
		User:   &user,
		Enmies: enemies,
	})
}
