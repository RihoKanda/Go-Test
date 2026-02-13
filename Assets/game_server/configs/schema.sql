--CREATE DATABASE IF NOT EXISTS idle_game;
--USE idle_game;

-- ユーザーテーブル
CREATE TABLE IF NOT EXISTS users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    device_id CHAR(36) UNIQUE NOT NULL,
    user_name VARCHAR(50) NOT NULL DEFAULT 'Player',
    level INT NOT NULL DEFAULT 1,
    exp BIGINT NOT NULL DEFAULT 0,
    attack_up INT NOT NULL DEFAULT 0,
    speed_up INT NOT NULL DEFAULT 0,
    hp_regen_up INT NOT NULL DEFAULT 0,
    evolution_stage INT NOT NULL DEFAULT 0,
    is_idle BOOLEAN NOT NULL DEFAULT FALSE,
    idle_started_at DATETIME,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_device_id(device_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 敵データテーブル
CREATE TABLE IF NOT EXISTS enemies (
    enemy_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    hp INT NOT NULL,
    attack INT NOT NULL,
    exp_reward INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 初期敵データ挿入
INSERT INTO enemies (name, hp, attack, exp_reward) VALUES
('Slime', 100, 10, 20),
('Goblin', 300, 30, 50),
('Orc', 700, 70, 100),
('Dragon', 5000, 200, 1000);