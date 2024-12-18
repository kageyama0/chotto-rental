-- 1ユーザーあたりの有効セッション数を制限するための古いセッションの無効化
WITH RankedSessions AS (
    SELECT id,
           ROW_NUMBER() OVER (
               PARTITION BY user_id
               ORDER BY last_accessed_at DESC
           ) as rn
    FROM sessions
    WHERE user_id = ? AND is_valid = true
)
UPDATE sessions
SET is_valid = false
WHERE id IN (
    SELECT id
    FROM RankedSessions
    WHERE rn > 5  -- 最大5セッションまで保持
);
