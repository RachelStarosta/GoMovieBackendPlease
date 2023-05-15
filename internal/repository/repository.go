package repository

import (
	"backend/internal/models"
	"database/sql"
)

/* we are defining a type that is an interface so that everything this type
must include the following functions, (to satisfy this interface)
so far postgres_dbrepo is this one tyoe because it implements the function
*/
type DatabaseRepo interface {   //pretty much everything in go is an interface
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}