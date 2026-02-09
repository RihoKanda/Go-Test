package model

import (
	//"github.com/naoina/goam"
	"database/sql"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

func FindUserByDeviceID(deviceID string) (*User, error) {
	row := DB.QueryRow(
		`SELECT id, device_id, user_name, level, created_at
		 FROM users
		WHERE device_id = ?`,
		deviceID,
	)

	var user User

	err := row.Scan(
		&user.ID,
		&user.DeviceID,
		&user.UserName,
		&user.Level,
		&user.CreatedAt,
	)

	// err := goam.
	// 	New(DB).
	// 	Select("*").
	// 	From("users").
	// 	Where("device_id = ?", deviceID).
	// 	Scan(&user)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(deviceID string) (*User, error) {
	result, err := DB.Exec(
		`INSERT INTO users (device_id, user_name, level)
		 VALUES (?, ?, ?)`,
		deviceID,
		"NoName",
		1,
	)
	// user := &User{
	// 	DeviceID: deviceID,
	// 	UserName: "NoName",
	// 	Level:    1,
	// }

	// _, err := goam.
	// 	New(DB).
	// 	Insert("users").
	// 	Columns("device_id", "user_name", "level").
	// 	Values(user.DeviceID, user.UserName, user.Level).
	// 	Exec()

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return FindUserByID(int(id))
}

func FindUserByID(id int) (*User, error) {
	row := DB.QueryRow(
		`SELECT id, device_id, user_name, level, created_at
		 FROM users
		 WHERE id = ?`,
		id,
	)

	var user User
	err := row.Scan(
		&user.ID,
		&user.DeviceID,
		&user.UserName,
		&user.Level,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
