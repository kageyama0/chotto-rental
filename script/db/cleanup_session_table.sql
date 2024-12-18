-- 期限切れまたは無効なセッションの削除（バッチ処理用）
DELETE FROM sessions
WHERE expires_at < CURRENT_TIMESTAMP
   OR is_valid = false;
