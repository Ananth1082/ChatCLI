package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func init() {
	db, err := sql.Open("sqlite", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            age INTEGER NOT NULL
        );
    `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO users(name, age) VALUES (?, ?)", "Alice", 30)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("User: %d, Name: %s, Age: %d", id, name, age)
	}
}
