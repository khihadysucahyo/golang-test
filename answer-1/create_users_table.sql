CREATE TABLE IF NOT EXISTS users (
	id serial,
	username varchar(15),
	parent bigint
);
