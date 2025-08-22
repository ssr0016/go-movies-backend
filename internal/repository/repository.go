package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
	GetuserByEmail(email string) (*models.User, error)
	GetuserByID(id int) (*models.User, error)
}
