package main

/** reason for the blank identifier: include but dont explicitly use
*/
import(
	"database/sql"
	"log"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string)  (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()    //checking to see if database can be pinged else not connected
	if err!= nil{
		return nil, err
	}

	return db, nil
}

/* Creating a second function with a (receiver of app that point to the application)
	called connectToDB, which takes no parameters and returns a pointer
	to sql.db and potentially an error. This function seems to do exactly what 
	the function above does, however it has the receiver of app.  If the code 
	reached to the log line then it has connected succesfully to the database.

	Connection and err are populated the openDB(app.DSN) is called.  
	It then will returns the connection and nil since there is no error
*/

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgres!")
	return connection, nil
}



/*
	These functions are separateed because in the future, 
	if you are using multiple databases, a function for each like 
	openDB, for another like openmongo and then pass it to connectto db
*/



