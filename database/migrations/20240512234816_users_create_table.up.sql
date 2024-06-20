CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(200) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/WA-API" -path database/migrations up
-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/gin-gorm" -path database/migrations down
-- migrate create -ext sql -dir database/migrations  comments_create_table