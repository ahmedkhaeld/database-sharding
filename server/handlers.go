package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Response struct {
	URL    string
	UrlID  string
	Server string
}

//GetHandler retrieves stored url based on the shortened-url
func GetHandler(w http.ResponseWriter, r *http.Request) {
	// extract the url-id from the Url request
	//https://localhost:8081/Ypj+u  url-id =Ypj+u
	urlID := chi.URLParam(r, "url-id")

	// to query the database, first we need to know which database to hit
	// identify which node
	srv, ok := hr.GetNode(urlID)
	if !ok {
		log.Println(err, ok)
		return
	}

	// query the database
	conn, err := sql.Open("pgx", srv)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to connect to db: %v\n", err))
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	query := `SELECT url FROM url_table WHERE url_id = $1`
	row := conn.QueryRow(query, urlID)
	var URL string
	err = row.Scan(&URL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Row fetched")

	res := Response{
		URL:    URL,
		UrlID:  urlID,
		Server: srv,
	}

	err = WriteJSON(w, http.StatusOK, res, "response")
	if err != nil {
		log.Println(err)
		return
	}

}

//PostHandler creates a new shortened-url(url-id) for a given url
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// get the url from the query
	// hash the url and get the hashed url
	//based on the hashed url take the first 5 chars
	// that hash will determine which server is going to hit
	// write to the database

	url := r.URL.Query().Get("url")

	h := sha256.Sum256([]byte(url))
	hash := base64.StdEncoding.EncodeToString(h[:])
	urlID := hash[0:5]

	srv, ok := hr.GetNode(urlID)
	if !ok {
		log.Println(err, ok)
		return
	}
	//open connection with node and then write to it
	conn, err := sql.Open("pgx", srv)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to connect to db: %v\n", err))
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	query := `INSERT INTO url_table (url, url_id) VALUES ($1,$2)`
	_, err = conn.Exec(query, url, urlID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("url_table has new row!")

	res := Response{
		URL:    url,
		UrlID:  urlID,
		Server: srv,
	}

	err = WriteJSON(w, http.StatusOK, res, "response")
	if err != nil {
		log.Println(err)
		return
	}
}
