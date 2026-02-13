package model

import (
	"database/sql"
	"time"
)

// ユーザーデータ構造体定義
type User struct {
	UserID         int        `json:"user_id"`         //ユーザーID
	DeviceID       string     `json:"device_id"`       //デバイスID
	UserName       string     `json:"user_name"`       //ユーザー名
	Level          int        `json:"level"`           //レベル
	Exp            int64      `json:"exp"`             //経験値
	AttackUp       int        `json:"attack_up"`       //攻撃力アップ
	SpeedUp        int        `json:"speed_up"`        //攻撃速度アップ
	HPRegenUp      int        `json:"hp_regen_up"`     //HP回復速度アップ
	EvolutionStage int        `json:"evolution_stage"` //進化段階
	IsIdle         bool       `json:"is_idle"`         //放置中フラグ
	IdleStartedAt  *time.Time `json:"idle_started_at"` //放置開始フラグ
	CreatedAt      time.Time  `json:"created_at"`      //作成日時
	UpdatedAt      time.Time  `json:"updated_at"`      //更新日時
}

// UPdateType　強化タイプ
type UpdateType string

const (
	UpdateAttack  UpdateType = "attack"
	UpdateSpeed   UpdateType = "speed"
	UpdateHPRegen UpdateType = "hp_regen"
)

// GetExpForNextLevel レベルアップに必要な経験値を計算
func GetExpForNextLevel(level int) int64 {
	// レベル１から指定レベルまでの累積必要経験値
	return int64(level * 100)
}

// CalculateUpgradeCount レベルに応じた最大強化回数を計算
func CalculateUpgradeCount(level int) int {
	count := 0
	for i := 1; i <= level; i++ {
		if i%10 == 0 || i%5 == 0 {
			count += 2
		} else {
			count += 1
		}
	}
	return count
}

// GetUserByDeviceID デバイスIDからユーザー情報を取得
func GetUserByDeviceID(db *sql.DB, deviceID string) (*User, error) {
	user := &User{}
	err := db.QueryRow(`
	    SELECT 
			user_id, 
			device_id, 
			user_name, 
			level, 
			exp, 
			attack_up, 
			speed_up, 
			hp_regen_up,
			is_idle, 
			idle_started_at, 
			created_at, 
			updated_at 
		FROM users WHERE device_id = ?
	`, deviceID).Scan(
		&user.UserID,
		&user.DeviceID,
		&user.UserName,
		&user.Level,
		&user.Exp,
		&user.AttackUp,
		&user.SpeedUp,
		&user.HPRegenUp,
		&user.IsIdle,
		&user.IdleStartedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser 新規ユーザーを作成
func CreateUser(db *sql.DB, deviceID string) (*User, error) {
	result, err := db.Exec(`
	    INSERT INTO users (device_id, user_name, level, exp,)
		VALUES (?, ?, 1, 0)
	`, deviceID, "Player")

	if err != nil {
		return nil, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetUserByDeviceID(db, deviceID)
}

// UpdateUser ユーザー情報を更新
func UpdateUser(db *sql.DB, user *User) error {
	_, err := db.Exec(`
	    UPDATE users SET
			user_name = ?, 
			level = ?,
			exp = ?,
			attack_up = ?,
			speed_up = ?,
			hp_regen_up = ?,
			is_idle = ?,
			idle_started_at = ?,
			updated_at = ?
		WHERE user_id = ?
	`,
		user.UserName,
		user.Level,
		user.Exp,
		user.AttackUp,
		user.SpeedUp,
		user.HPRegenUp,
		user.IsIdle,
		user.IdleStartedAt,
		user.UpdatedAt,
		user.UserID,
	)
	return err
}

// LevelUp ユーザーレベルアップ処理
func (user *User) LevelUp() bool {
	requiredExp := GetExpForNextLevel(user.Level)
	if user.Exp >= requiredExp {
		user.Level++
		user.Exp -= requiredExp
		return true
	}
	return false
}

// PowerUp ユーザー能力強化処理
func (user *User) PowerUp(updateType UpdateType) {
	switch updateType {
	case UpdateAttack:
		user.AttackUp++
	case UpdateSpeed:
		user.SpeedUp++
	case UpdateHPRegen:
		user.HPRegenUp++
	}
}

// GetCurrentUpgradeCount 現在の強化回数を取得
func (user *User) GetCurrentUpgradeCount() int {
	return user.AttackUp + user.SpeedUp + user.HPRegenUp
}

// CanUpgrade 能力強化可能か判定
func (user *User) CanUpgrade() bool {
	maxUpgrades := CalculateUpgradeCount(user.Level)
	currentUpgrade := user.GetCurrentUpgradeCount()
	return currentUpgrade < maxUpgrades
}
