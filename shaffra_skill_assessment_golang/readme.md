# Fix the Issues on the buggy_project

## Error 1: Did not handle database error

The code ignored the error that might results while while initializing the db connection

 - how how I handled it
   receive the error as potential return item from the following code:

   `
   db, err = sql.Open("postgres", "user=postgres dbname=test sslmode=disable")
   if err != nil {
      log.Fatal(err)
   }
   `
if error exists, exit the program given that the program role is to work with the database if a connection cannot be established then it has no business running.

## Error 2: Using sync.WaitGroup to wait for goroutines inside HTTP handlers is problematic.

This approach is not suitable for handling HTTP requests as it does not manage the lifecycle of HTTP requests correctly and can lead to issues such as resource leaks or incorrect output.

Although in this case scenario go routines are not needed since it can handle traffic all the same.

All the same ,I did use of go routines and lock mechanism using mutex to ensure consistency while interacting with the database especially while dealing with operation that affect the database.

## Error 3: Ingoring database possible query error

They once again ignore the error that might results while quering the database

here
`
rows, _ := db.Query("SELECT name FROM users")
`

Always handle the error! always - do not ignore any error
and here is the solution

`
_, err := db.Exec("INSERT INTO users (name) VALUES ('" + username + "')")
if err != nil {
			http.Error(w, "Error querying database", http.StatusInternalServerError)
			return
		}
`
## Error 4: SQL Injection Vulnerability

this code :
`
_, err := db.Exec("INSERT INTO users (name) VALUES ('" + username + "')")

`
This code is vulnerable to SQL injection attacks because it concatenates user input directly into the SQL query.

Solution: 
Use parameterized queries to prevent SQL injection.
`
_, err := db.Exec("INSERT INTO users (name) VALUES ($1)", username)
`


## Error 5: Ignoring the scan error

They ignored the scan error here is the solution

here:
`
rows.Scan(&name)
`
and here is the solution.

`
if err := rows.Scan(&name); err != nil {
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
   ,
## Error 6: Simulating Long Database Operations

Problem: time.Sleep in the createUser function simulates a long operation, but it's not a realistic or advisable way to handle such cases. It can lead to performance issues and unresponsive servers.

so  I removed it in my code

## summary
This project major error revolved around ignoring errors and dealing with go routines the wrong way.

