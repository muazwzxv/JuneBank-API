CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL,
	created_at timestamp DEFAULT NOW(),
	updated_at timestamp DEFAULT NOW(),
	deleted_at timestamp DEFAULT null
);

---- create above / drop below ----
DROP TABLE IF EXISTS users