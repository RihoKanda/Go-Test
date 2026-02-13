package model

import (
	"database/sql"
	"time"
)

// Enemy 敵データ構造体定義
type Enemy struct {
	EnemyID   int       `json:"enemy_id"`   //敵ID
	Name      string    `json:"name"`       //敵名
	Hp        int       `json:"hp"`         //体力
	Attack    int       `json:"attack"`     //攻撃力
	ExpReward int       `json:"exp_reward"` //経験値報酬
	CreatedAt time.Time `json:"created_at"` //作成日時
}

// GetEnemyByID IDで敵データを取得
func GetEnemyByID(db *sql.DB, enemyID int) (*Enemy, error) {
	enemy := &Enemy{}
	err := db.QueryRow(`
	    SELECT 
			enemy_id, 
			name, 
			hp, 
			attack, 
			exp_reward, 
			created_at
		FROM enemies WHERE enemy_id = ?
	`, enemyID).Scan(
		&enemy.EnemyID,
		&enemy.Name,
		&enemy.Hp,
		&enemy.Attack,
		&enemy.ExpReward,
		&enemy.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return enemy, nil
}

// GetAllEnemies 全ての敵データを取得
func GetAllEnemies(db *sql.DB) ([]*Enemy, error) {
	rows, err := db.Query(`
	    SELECT
			enemy_id,
			name,
			hp,
			attack,
			exp_reward,
			created_at
		FROM enemies ORDER BY enemy_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	enemies := []*Enemy{}
	for rows.Next() {
		enemy := &Enemy{}
		err := rows.Scan(
			&enemy.EnemyID,
			&enemy.Name,
			&enemy.Hp,
			&enemy.Attack,
			&enemy.ExpReward,
			&enemy.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		enemies = append(enemies, enemy)
	}
	return enemies, rows.Err()
}
