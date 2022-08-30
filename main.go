package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fridgigator/backend/globals"
	"github.com/Fridgigator/backend/login"
	_ "github.com/lib/pq"
)

func main() {
	psqlconn := fmt.Sprintf("user=postgres password=%s host=db.cldkdlstnxjwpxyqnkwb.supabase.co port=5432 dbname=postgres", os.Getenv("POSTGRES_PASSWD"))

	// open database
	var err error
	globals.DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer globals.DB.Close()

	globals.StateMap = make(map[string]chan string, 1)

	port := os.Getenv("PORT")

	if port == "" {
		port = "1234"
	}
	fmt.Println("Ready")
	mux := http.NewServeMux()
	mux.HandleFunc("/login/get-state", login.GetOauth2StateString)
	mux.HandleFunc("/login/github-response", login.GithubResponse)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
