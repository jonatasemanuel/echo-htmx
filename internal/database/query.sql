-- name: GetAnime :one
SELECT * FROM animes
WHERE id = ? LIMIT 1;

-- name: ListAnimes :many
SELECT * FROM animes
ORDER BY name;
