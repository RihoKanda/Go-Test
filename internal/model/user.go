package model

import "time"

type User struct {
	UserID      int        `json:"user_id"`       //ユーザーID
	DeviceID    string     `json:"device_id"`     //デバイスID
	UserName    string     `json:"user_name"`     //ユーザー名
	Level       int        `json:"level"`         //レベル
	Exp         int64      `json:"exp"`           //経験値
	AttackUp    int        `json:"attack_up"`     //攻撃力アップ
	SpeedUp     int        `json:"speed_up"`      //攻撃速度アップ
	IsIdle      bool       `json:"is_idle"`       //放置中フラグ
	IdleStartAt *time.Time `json:"idle_start_at"` //放置開始フラグ
	CreatedAt   time.Time  `json:"created_at"`    //作成日時
	UpdatedAt   time.Time  `json:"updated_at"`    //更新日時
}
