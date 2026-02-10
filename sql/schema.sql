CREATE DATABASE IF NOT EXISTS idle_game;
USE idle_game;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    device_id VARCHAR(64) NOT NULL UNIQUE,
    level INT NOT NULL DEFAULT 1
);

CREATE TABLE idle_stats (
    user_id INT PRIMARY KEY,
    start_at DATETIME NOT NULL,
    last_active_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);