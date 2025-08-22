package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

type postgresDBRepo struct {
	db *sql.DB
}

const dbTimeout = time.Second * 3

func (m *postgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT
			id,
			title,
			release_date,
			run_time,
			mpaa_rating,
			description,
			coalesce(image, ''),
			created_at,
			updated_at
		FROM 
			movies
		ORDER BY
			 title
	`

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie
	for rows.Next() {
		movie := models.Movie{}
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	return movies, nil
}
