-- name: insert-notification
INSERT INTO maintenance.notifications (user_id, task_id)
VALUES (?, ?)
ON DUPLICATE KEY UPDATE is_update = 1;

-- name: delete-notification
DELETE FROM maintenance.notifications WHERE id = ?;

-- name: clear-notifications
DELETE FROM maintenance.notifications WHERE user_id = ?;

-- name: list-notifications
SELECT
    id AS ID,
    user_id AS UserID,
    task_id AS TaskID,
    is_update AS IsUpdate,
    created_at AS CreatedAt
FROM maintenance.notifications WHERE user_id = ?;

