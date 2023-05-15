package dbrepo

import (
	"backend/internal/models"
	"database/sql"
	"time"
	"context"
)


type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3    //later used (creates)to ensure that the interaction with the db is over after 3 seconds
									// cancel everything

/* (m *postgresDB.. is a receiver), the function takes no parameters
and must return a pointer to sql.DB. 
*/		
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	/* The coalesce() function is logic right in sql, 
		parameters are (image, '') if the image field is not null use that , 
		otherwise use an empty string, which gets around the problem of null values
	*/
	query := `
		select 
			id, title, release_date, runtime, 
			mpaa_rating, description, coalesce(image, ''),
			created_at, updated_at
		from
			movies
		order by
			title	
			`
	rows, err := m.DB.QueryContext(ctx, query)
	if err !=  nil {
		return nil, err
	}
	defer rows.Close()   //closes the connection to the database

	var movies []*models.Movie
	//use rows to populate the var movies, must be in the same order as the query
	for rows.Next(){
		var movie models.Movie
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
		if err!= nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}


	
	return movies, nil   //if you get this far there is no errors, and movies can be returned

	//clients can disappear without warning, 
}