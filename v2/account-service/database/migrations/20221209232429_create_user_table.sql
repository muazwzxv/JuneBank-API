-- +goose Up
-- +goose StatementBegin
CREATE TABLE user(
	id int NOT NULL,
	owner VARCHAR(60) NOT NULL DEFAULT "",
	balance DECIMAL NOT NULL DEFAULT 0,
	currency varchar(10) NOT NULL DEFAULT "",
	created_at  TIMESTAMP NOT NULL,
	PRIMARY KEY (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user
-- +goose StatementEnd
