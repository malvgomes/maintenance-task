-- name: insert-notification
INSERT INTO maintenance.notifications (user_id, task_id)
VALUES (?, ?)
ON DUPLICATE KEY UPDATE is_update = 1;
