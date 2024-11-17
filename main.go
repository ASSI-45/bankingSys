package main

import (
    "database/sql"
    "log"
    "fmt"
    _ "github.com/lib/pq"
)

func main()  {
  // connect to db
  db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/bankingSys?sslmode=disable")
  //logs error
  if err != nil {
      log.Fatal(err)
  }
  // checks connection
  err = db.Ping()

  if err != nil {
      log.Fatal(err)
  }

  fmt.Println("Successfully connected to PostgreSQL!")


  // closes the connection
  defer db.Close()

}

