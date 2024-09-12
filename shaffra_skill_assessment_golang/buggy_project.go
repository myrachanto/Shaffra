package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var mu sync.Mutex

func main() {
	var err error
	// ensure you handle error from trying to connect to the db
	db, err = sql.Open("postgres", "user=postgres dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/create", createUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	go func() {
		rows, err := db.Query("SELECT name FROM users")
		if err != nil {
			http.Error(w, "Error querying database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var names []string
		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				http.Error(w, "Error scanning rows", http.StatusInternalServerError)
				return
			}
			names = append(names, name)

		}
		for _, name := range names {
			fmt.Fprintf(w, "User: %s\n", name)
		}
	}()
}

func createUser(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	go func() {
		username := r.URL.Query().Get("name")
		// Use parameterized queries to prevent SQL injection
		_, err := db.Exec("INSERT INTO users (name) VALUES ($1)", username)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User %s created successfully", username)
	}()
}
