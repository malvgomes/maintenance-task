CREATE DATABASE IF NOT EXISTS maintenance CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE TABLE IF NOT EXISTS maintenance.users (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(32) NOT NULL,
    password BLOB NOT NULL,
    user_first_name VARCHAR(50) NOT NULL,
    user_last_name VARCHAR(50) NULL DEFAULT NULL,
    user_role ENUM('MANAGER', 'TECHNICIAN') NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE (username),
    INDEX (user_role)
);

CREATE TABLE IF NOT EXISTS maintenance.tasks (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    summary BLOB NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (id),
    INDEX (user_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS maintenance.notifications (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    task_id INT NOT NULL,
    is_update BOOLEAN NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (user_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE
);

INSERT INTO maintenance.users (id, username,password,user_first_name,user_last_name,user_role)
VALUES
    (1, 'john_123', AES_ENCRYPT('password_john_123', 'secure_key'), 'John', '123', 'MANAGER'),
    (2, 'bob_456', AES_ENCRYPT('password_bob_456', 'secure_key'), 'Bob', '456', 'TECHNICIAN'),
    (3, 'clayton_789', AES_ENCRYPT('password_clayton_789', 'secure_key'), 'Clayton', '789', 'TECHNICIAN');

INSERT INTO maintenance.tasks (id, user_id, summary)
VALUES
    (1, 2, AES_ENCRYPT('Bob removed the trash from the break room', 'secure_key')),
    (2, 2, AES_ENCRYPT('Bob finished washing the garage', 'secure_key')),
    (3, 3, AES_ENCRYPT('Clayton washed the reception windows', 'secure_key'));

INSERT INTO maintenance.notifications (user_id, task_id)
VALUES
    (2, 1),
    (2, 2),
    (3, 3);

