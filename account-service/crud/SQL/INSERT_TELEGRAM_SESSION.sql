INSERT INTO account_service.telegram_sessions (session_id, user_id)
VALUES ($1, $2)
RETURNING session_id, user_id;