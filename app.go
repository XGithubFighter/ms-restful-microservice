// app.go

package main

import (
    "fmt"
    "log"
    
    "database/sql"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) Run(addr string) { }

func (a *App) Initialize(user, password, dbname string) {
    connectionString :=
        fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

    var err error
    a.DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }

    a.Router = mux.NewRouter()
}