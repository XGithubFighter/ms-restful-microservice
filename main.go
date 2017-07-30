// main.go

package main

import (
	"fmt"
	"os"
)

func main() {
    a := App{}
    fmt.Printf("Pre - Initializing ...\n")

    os.Setenv("APP_DB_USERNAME", "postgres")
    os.Setenv("APP_DB_PASSWORD", "Anomitra")
    os.Setenv("APP_DB_NAME", "restdb")

    a.Initialize(
        os.Getenv("APP_DB_USERNAME"),
        os.Getenv("APP_DB_PASSWORD"),
        os.Getenv("APP_DB_NAME"))

    fmt.Printf("Post - Initializing ...\n")

    a.Run(":7777")
}