INSERT INTO account_service.users (user_id, email, credentials)
VALUES ($1, $2, $3)
RETURNING user_id;