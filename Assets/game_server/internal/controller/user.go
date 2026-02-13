package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/saf1o/go-test/internal/database"
	"github.com/saf1o/go-test/internal/model"
)

// LevelUpRequest レベルアップリクエスト
type LevelUpRequest struct {
	UserID int `json:"user_id"`
}

// LevelUpResponse レベルアップレスポンス
type LevelUpResponse struct {
	User      *model.User `json:"user"`
	LeveledUp bool        `json:"leveled_up"`
	NewLevel  int         `json:"new_level,omitempty"`
}

// UpgradeRequest 能力強化リクエスト
type UpgradeRequest struct {
	UserID      int    `json:"user_id"`
	UpgradeType string `json:"upgrade_type"`
}

// UpgradeResponse 能力強化レスポンス
type UpgradeResponse struct {
	User        *model.User `json:"user"`
	UpgradeType string      `json:"upgrade_type"`
}

// EvolveRequest 進化リクエスト
type EvolveRequest struct {
	UserID int `json:"user_id"`
}

// EvolveResponse 進化レスポンス
type EvolveResponse struct {
	User           *model.User `json:"user"`
	EvolutionStage int         `json:"evolution_stage"`
}

// HandleLevelUp レベルアップ処理
func HandleLevelUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req LevelUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// ユーザー取得
	var user model.User
	err := databaase.DB.QueryRow(`
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
		&user.EvolutionStage,
		&user.IsIdle,
		&user.IdleStartedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		SendError(w, http.StatusNotFound, "User not found")
		return
	}

	// レベルアップ可能かチェック
	leveledUp := user.LevelUp()
	newLevel := user.Level

	if leveledUp {
		err = model.UpdateUser(database.DB, &user)
		if err != nil {
			SendError(w, http.StatusInternalServerError, "Failed to update user")
			return
		}
	}

	SendSuccess(w, LevelUpResponse{
		User:      &user,
		LeveledUp: leveledUp,
		NewLevel:  newLevel,
	})
}

// HandleUpgrade 能力強化処理
func HandleUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req UpgradeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// UpgradeTypeの検証
	var upgradeType model.UpgradeType
	switch req.UpgradeType {
	case "attack":
		upgradeType = model.UpgradeTypeAttack
	case "speed":
		upgradeType = model.UpgradeTypeSpeed
	case "hp_regen":
		upgradeType = model.UpgradeTypeHPRegen
	default:
		SendError(w, http.StatusBadRequest, "Invalid upgrade_type")
		return
	}

	// ユーザー取得
	var user model.User
	err := databaase.DB.QueryRow(`
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
		&user.EvolutionStage,
		&user.IsIdle,
		&user.IdleStartedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		SendError(w, http.StatusNotFound, "User not found")
		return
	}

	// 進化条件チェック
	requiredLevel := (user.EvolutionStage + 1) * 10
	if user.Level < requiredLevel {
		SendError(w, http.StatusBadRequest, "Level too low for evolution")
		return
	}

	// 進化処理
	user.EvolutionStage++

	err = model.UpdateUser(database.DB, &user)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	SendSuccess(w, EvolveResponse{
		User:           &user,
		EvolutionStage: user.EvolutionStage,
	})
}
