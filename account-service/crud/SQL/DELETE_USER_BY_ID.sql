DELETE FROM account_service.users
WHERE user_id=$1
RETURNING user_id, email, credentials;