-- migrate -path migrations -database "postgres://root:root@localhost/db?sslmode=disable" up

CREATE TABLE users (
	id 			SERIAL PRIMARY KEY,
	username 	VARCHAR(255) NOT NULL UNIQUE,
	password	VARCHAR(255) NOT NULL,
	created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE accounts (
	id 			SERIAL PRIMARY KEY,
	user_id 	INT NOT NULL, --REFERENCES users (id),
	balance 	INT NOT NULL DEFAULT 0,
	created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
	FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE reservations (
	id 			SERIAL PRIMARY KEY,
	account_id 	INT NOT NULL,
	service_id 	INT NOT NULL,
	order_id 	INT NOT NULL,
	amount 		INT NOT NULL,
	created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
	FOREIGN KEY (account_id) REFERENCES accounts (id)
);