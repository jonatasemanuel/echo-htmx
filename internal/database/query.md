# sql

```sql

-- Querys
SELECT 
    characters.name AS character_name, 
    animes.name AS anime_name
FROM 
    characters
JOIN 
    animes 
ON 
    characters.anime_id = animes.id;

-- MIGRATIONS
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

-- VALUES
INSERT INTO animes (name) VALUES 
    ('Bleach'),
    ('Black Clover'),
    ('Attack on Titan'),
    ('Yugi-Yo'),
    ('My Hero Academia'),
    ('Dragon Ball Z'),
    ('Demon Slayer'),
    ('One Punch Man'),
    ('Hunter x Hunter'),
    ('Death Note'),
    ('Fullmetal Alchemist'),
    ('Sword Art Online'),
    ('Tokyo Ghoul'),
    ('Fairy Tail');

INSERT INTO characters (name, anime_id, image)
VALUES
    ('Ace', 1, 'https://i.ibb.co/ZXchsTH/image.png'),
    ('Magna', 4, 'https://i.ibb.co/cFWnD5P/image.png'),
    ('Gin', 12, 'https://i.ibb.co/GJdMbCv/image.png'),
    ('Killer Bee', 2, 'https://i.ibb.co/fYtC0jf/image.png'),
    ('Sato', 8, 'https://i.ibb.co/TLKD3xF/image.png'),
    ('Air', 11, 'https://i.ibb.co/2K8mRY1/image.png'),
    ('Buchwald', 6, 'https://i.ibb.co/4jXZYNv/image.png'),
    ('Bojack', 9, 'https://i.ibb.co/vZ9tsPg/image.png'),
    ('Walter', 4, 'https://i.ibb.co/7CQhn9z/image.png'),
    ('Mikami', 13, 'https://i.ibb.co/6yY465G/image.png'),
    ('Asada', 15, 'https://i.ibb.co/WtNsq8D/image.png'),
    ('Bulla', 9, 'https://i.ibb.co/xC5ypP1/image.png'),
    ('Jose Porla', 17, 'https://i.ibb.co/61k2Nxg/image.png'),
    ('Marlene', 6, 'https://i.ibb.co/7SrncFR/image.png'),
    ('Kotetsu', 10, 'https://i.ibb.co/WBXxGBQ/image.png');

```
