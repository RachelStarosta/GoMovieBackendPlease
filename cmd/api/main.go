package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN string
	Domain string
	DB repository.DatabaseRepo
}
	//DB *sql.DB   // this is a pointer (sql.db)to a pool of database connections right in the application, this is temporary
				 // change this to the repository pattern, which makes things easier to test and change
					

// entry point for the application, only one main function every time
func main() {
	/**
	*  how to connect to the database, where is the database repo, what is the jwt seq
	 */
	var app application

	//read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to the database
	conn, err := app.connectToDB();
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}   //we want to close these when we are done with them
	//Change: defer conn.Close()   to  //keyword defer means execute code after this keyword just when the function that contains the keyword ends aka app.DB
	defer app.DB.Connection().Close()
	
	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	//start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
