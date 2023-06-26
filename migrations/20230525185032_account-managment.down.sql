-- migrate -path migrations -database "postgres://root:root@localhost/db?sslmode=disable" down

DROP TABLE IF EXISTS reservations;

DROP TABLE IF EXISTS accounts;

DROP TABLE IF EXISTS users;


