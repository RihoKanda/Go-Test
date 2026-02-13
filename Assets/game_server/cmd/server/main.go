package main

import (
	"log"
	"net/http"
	"os"

	"github.com/saf1o/go-test/internal/controller"
	"github.com/saf1o/go-test/internal/database"
)

// ルーティング設定
func main() {
	// データベース接続設定
	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     3306,
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", ""),
		Database: getEnv("DB_NAME", "idle_game"),
	}

	// データベース初期化
	if err := database.Initialize(dbConfig); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// ルーティング設定
	mux := http.NewServeMux()

	// 認証
	mux.HandleFunc("/api/auth/login", corsMiddleware(controller.LoginHandler)) // ログイン処理

	// ゲーム状態取得
	mux.HandleFunc("/api/game/state", corsMiddleware(controller.HandleGameState))

	// 放置システム
	mux.HandleFunc("/api/idle/start", corsMiddleware(controller.HandleIdleStart))   // 放置開始
	mux.HandleFunc("/api/idle/finish", corsMiddleware(controller.HandleIdleFinish)) // 放置終了

	// ユーザー操作
	mux.HandleFunc("/api/user/levelup", corsMiddleware(controller.HandleLevelUp)) // レベルアップ
	mux.HandleFunc("/api/user/upgrade", corsMiddleware(controller.HandleUpgrade)) // 能力振り分け
	//mux.HandleFunc("/api/user/evolve", corsMiddleware(controller.HandleEvolve))   // 進化

	// ヘルスチェック
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// サーバー起動
	port := getEnv("PORT", "8080")
	log.Printf("Server starting at port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// conrsMiddleware CORS対応のミドルウェア
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// getEnv 環境変数を取得 存在しない場合=デフォルト値を返す
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// 	model.InitDB("idle_game.db")

// 	http.HandleFunc("/api/auth/login", controller.LoginHandler)     // ログイン処理
// 	http.HandleFunc("/api/game/state", controller.GameStateHandler) // ゲーム状態取得

// 	http.HandleFunc("/api/idle/start", controller.StartIdleHandler)   // 放置開始
// 	http.HandleFunc("/api/idle/finish", controller.FinishIdleHandler) // 放置終了

// 	http.HandleFunc("/api/user/levelup", controller.LevelUpHandler) // レベルアップ
// 	http.HandleFunc("/api/user/powerup", controller.PowerUpHandler) // 能力振り分け
// 	http.HandleFunc("/api/user/evolve", controller.EnvolveHandler)  // 進化

// 	log.Println("Server started at :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
