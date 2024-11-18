package models

import (
	"context"
)

type Anime struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (a *Anime) CreateAnime(anime Anime) (*Anime, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO animes (name)
		VALUES($1) RETURNING name
	`
	// erro
	_, err := db.ExecContext(
		ctx,
		query,
		anime.Name,
	)
	if err != nil {
		return nil, err
	}

	return &anime, nil

}

func (a *Anime) ListAnimes() ([]*Anime, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT
		* FROM animes`
	// erro
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var animes []*Anime
	for rows.Next() {
		var anime Anime
		err := rows.Scan(
			&anime.ID,
			&anime.Name,
		)
		if err != nil {
			return nil, err
		}
		animes = append(animes, &anime)
	}

	return animes, nil
}
