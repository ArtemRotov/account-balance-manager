-- migrate -path migrations -database "postgres://root:root@localhost/db?sslmode=disable" up

CREATE TABLE users (
	id 			SERIAL PRIMARY KEY,
	username 	VARCHAR(255) not null unique,
	password	VARCHAR(255) not null,
	created_at  TIMESTAMP    not null default now()
);