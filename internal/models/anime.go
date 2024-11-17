package models

import (
	"context"
	"fmt"
)

type Anime struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (a *Anime) CreateAnime(anime Anime) (*Anime, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO anime (anime) VALUES (?) RETURNING id,name`
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
		id,
		name FROM anime`
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
	fmt.Println(animes)
	return animes, nil
}
