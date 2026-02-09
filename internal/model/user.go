package model

import "time"

type User struct {
	ID        int       `db:"id"`
	DeviceID  string    `db:"device_id"`
	UserName  string    `db:"user_name"`
	Level     int       `db:"level"`
	CreatedAt time.Time `db:"created_at"`
}
