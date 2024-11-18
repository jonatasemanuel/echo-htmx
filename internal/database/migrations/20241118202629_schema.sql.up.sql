-- Add up migration script here
CREATE TABLE IF NOT EXISTS animes(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name text NOT NULL UNIQUE
);
