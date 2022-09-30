DELETE FROM account_service.telegram_sessions
WHERE session_id=$1
RETURNING session_id, user_id;