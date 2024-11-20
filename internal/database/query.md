# query

```sql
SELECT 
    characters.name AS character_name, 
    animes.name AS anime_name
FROM 
    characters
JOIN 
    animes 
ON 
    characters.anime_id = animes.id;



CREATE TABLE IF NOT EXISTS animes(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS characters (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name text NOT NULL UNIQUE,
	anime_id INTEGER,
	image BLOB,
	FOREIGN KEY(anime_id) REFERENCES animes(id) ON DELETE CASCADE ON UPDATE CASCADE
);


DROP TABLE IF EXISTS animes;
DROP TABLE IF EXISTS characters;

```
