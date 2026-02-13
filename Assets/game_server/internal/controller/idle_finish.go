package controller

import (
	"encoding/json"
	"net/http"
	"time"

	model "github.com/saf1o/go-test/internal/database"
	_ "github.com/saf1o/go-test/internal/model"
)

// IdleFinishRequest 放置終了リクエスト
type IdleFinishRequest struct {
	UserID int `json:"user_id"`
}

// IdleFinishResponse 放置終了レスポンス
type IdleFinishResponse struct {
	User        *model.User `json:"user"`
	ExpGained   int64       `json:"exp_gained"`
	Idleminutes int         `json:"idle_minutes"`
}

// HandleIdleFinish 放置終了・報酬確定
func HandleIdleFinish(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req IdleFinishRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// ユーザー情報を取得
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

	// 放置中でない場合はエラー
	if !user.IsIdle || user.IdleStartedAt == nil {
		SendError(w, http.StatusBadRequest, "Not in idle mode")
		return
	}

	// 放置時間を計算
	now := time.Now()
	idleDuration := now.Sub(*user.IdleStartedAt)
	idleMinutes := int(idleDuration.Minutes())

	// 最大5時間制限
	if idleMinutes > 300 {
		idleMinutes = 300
	}

	// 経験値を計算:　与ダメージ　×　放置時間（分）
	// 与ダメージ　＝　攻撃力　×　（1 + 強化回数 × 0.1）
	baseAttack := 10.0 // 基本攻撃力
	attackMultiplier := 1.0 + float64(user.AttackUp)*0.1
	damage := baseAttack * attackMultiplier
	expGained := int64(damage) * int64(idleMinutes)
	user.Exp += expGained

	// 放置状態解除
	user.IsIdle = false
	user.IdleStartedAt = nil

	err = model.UpdateUser(database.DB, &user)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	SendSuccess(w, IdleFinishResponse{
		User:        &user,
		ExpGained:   expGained,
		Idleminutes: idleMinutes,
	})
}
