package apiserver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq" // ...
	"github.com/yellowpuki/app-restapi/internal/app/store/sqlstore"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := NewServer(store, sessionStore)

	log.Print("starting apiserver...\n")
	return http.ListenAndServe(config.BindAddr, srv)
}

// newDB ...
func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
