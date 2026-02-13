using System;

namespace IdleGame.Models
{
    /// <summary>
    /// APIレスポンス基底クラス
    /// </summary>
    [Serializable]
    public class APIResponse<T>
    {
        public bool success;
        public T data;
        public string error;
    }

    /// <summary>
    /// ユーザーデータ
    /// </summary>
    [Serializable]
    public class User
    {
        public int user_id;
        public string device_id;
        public string user_name;
        public int level;
        public long exp;
        public int attack_up;
        public int speed_up;
        public int hp_regen_up;
        public int evolution_stage;
        public bool is_idle;
        public string idle_started_at;
        public string created_at;
        public string updated_at;
    }

    /// <summary>
    /// 敵データ
    /// </summary>
    [Serializable]
    public class Enemy
    {
        public int enemy_id;
        public string enemy_name;
        public int enemy_hp;
        public int exp_reward;
        public string created_at;
    }

    /// <summary>
    /// ログインレスポンス
    /// </summary>
    [Serializable]
    public class LoginResponse
    {
        public User user;
        public bool is_new_user;
    }

    /// <summary>
    /// ゲーム状態レスポンス
    /// </summary>
    [Serializable]
    public class GameStateResponse
    {
        public User user;
        public Enemy[] enemies;
    }

    /// <summary>
    /// 放置開始レスポンス
    /// </summary>
    [Serializable]
    public class IdleStartResponse
    {
        public User user;
        public string started_at;
    }

    /// <summary>
    /// 放置終了レスポンス
    /// </summary>
    [Serializable]
    public class IdleFinishResponse
    {
        public User user;
        public long exp_gained;
        public int idle_minutes;
    }

    /// <summary>
    /// レベルアップレスポンス
    /// </summary>
    [Serializable]
    public class LevelUpResponse
    {
        public User user;
        public bool leveled_up;
        public int new_level;
    }

    /// <summary>
    /// 能力強化レスポンス
    /// </summary>
    [Serializable]
    public class UpgradeResponse
    {
        public User user;
        public string upgrade_type;
    }

    /// <summary>
    /// 進化レスポンス
    /// </summary>
    [Serializable]
    public class EvolveResponse
    {
        public User user;
        public int evolution_stage;
    }
}