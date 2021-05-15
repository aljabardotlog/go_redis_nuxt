BEGIN;

    CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password text,
    status INT(3),
    created_by VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(20) NOT NULL,
    updated_at TIMESTAMP,
    deleted_by VARCHAR(20) NOT NULL,
    deleted_at TIMESTAMP,
    is_deleted BOOLEAN
    );

COMMIT;