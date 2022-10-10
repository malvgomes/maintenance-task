-- name: insert-user
INSERT INTO maintenance.users (username, password, user_first_name, user_last_name, user_role)
VALUES (?, AES_ENCRYPT(?, 'secure_key'), ?, ?, ?);

-- name: delete-user
DELETE FROM maintenance.users WHERE id = ?;

-- name: get-user
SELECT
    id AS ID,
    username AS Username,
    user_first_name AS FirstName,
    user_last_name AS LastName,
    user_role AS UserRole,
    created_at AS CreatedAt,
    updated_at AS UpdatedAt
FROM maintenance.users WHERE username = ? AND password = AES_ENCRYPT(?, 'secure_key');

-- name: get-user-by-id
SELECT
    id AS ID,
    username AS Username,
    user_first_name AS FirstName,
    user_last_name AS LastName,
    user_role AS UserRole,
    created_at AS CreatedAt,
    updated_at AS UpdatedAt
FROM maintenance.users WHERE id = ?;

-- name: get-users-by-role
SELECT
    id AS ID,
    username AS Username,
    user_first_name AS FirstName,
    user_last_name AS LastName,
    user_role AS UserRole,
    created_at AS CreatedAt,
    updated_at AS UpdatedAt
FROM maintenance.users WHERE user_role = ?;