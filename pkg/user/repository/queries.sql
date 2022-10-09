-- name: insert-user
INSERT INTO maintenance.users (username, password, user_first_name, user_last_name, user_role)
VALUES (?, AES_ENCRYPT(?, 'secure_key'), ?, ?, ?);

-- name: delete-user
DELETE FROM maintenance.users WHERE username = ?;