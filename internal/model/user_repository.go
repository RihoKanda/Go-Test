package model

// DB操作関連
import (
	"database/sql"
	//"errors"
	//"time"
)

// DeviceIDからユーザーを取得する
func FindUserByDeviceID(deviceID string) (*User, error) {
	row := DB.QueryRow(`
		SELECT
			user_id,
			device_id,
			user_name,
			level,
			exp,
			attack_up,
			speed_up,
			is_idle,
			idle_started_at,
			created_at,
			updated_at
		FROM users
		WHERE device_id = ?
	`, deviceID)

	var user User
	err := row.Scan(
		&user.UserID,
		&user.DeviceID,
		&user.UserName,
		&user.Level,
		&user.Exp,
		&user.AttackUp,
		&user.SpeedUp,
		&user.IsIdle,
		&user.IdleStartedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ユーザーを作成する(新規ログイン)
func CreateUser(deviceID string) (*User, error) {
	result, err := DB.Exec(`
		INSERT INTO users (device_id, level, exp)
		VALUES (?, 1, 0)
	`, deviceID)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return FindUserByID(int(id))
}

// user_idからユーザーを取得する(内部)
func FindUserByID(UserID int) (*User, error) {
	row := DB.QueryRow(`
		SELECT
			user_id,
			device_id,
			user_name,
			level,
			exp,
			attack_up,
			speed_up,
			is_idle,
			idle_started_at,
			created_at,
			updated_at
		FROM users
		WHERE user_id = ?
	`, UserID)

	var user User
	err := row.Scan(
		&user.UserID,
		&user.DeviceID,
		&user.UserName,
		&user.Level,
		&user.Exp,
		&user.AttackUp,
		&user.SpeedUp,
		&user.IsIdle,
		&user.IdleStartedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 放置開始
func StartIdle(userID int) error {
	_, err := DB.Exec(`
		UPDATE users
		SET
			is_idle = TRUE,
			idle_started_at = NOW()
		WHERE user_id = ?
	`, userID)
	return err
}

// 放置終了+経験値付与
func EndIdle(userID int, gainedExp uint64) error {
	_, err := DB.Exec(`
		UPDATE users
		SET
			is_idle = FALSE,
			idle_started_at = NULL,
			exp = exp + ?
		WHERE user_id = ?
	`, gainedExp, userID)
	return err
}
