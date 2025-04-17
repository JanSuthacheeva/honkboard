package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jansuthacheeva/honkboard/internal/models"
	"github.com/joho/godotenv"
)

type application struct {
	cfg            config
	logger         *slog.Logger
	todos          *models.TodoModel
	sessionManager *scs.SessionManager
	templateCache  map[string]*template.Template
}

type config struct {
	addr string
	dsn  string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config

	flag.StringVar(&cfg.dsn, "dsn", os.Getenv("DB_STRING"), "MySQL data source name")
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := openDB(cfg.dsn)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer db.Close()

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := application{
		cfg:            cfg,
		logger:         logger,
		todos:          &models.TodoModel{DB: db},
		sessionManager: sessionManager,
		templateCache:  templateCache,
	}

	app.logger.Info("Starting server", slog.String("addr", cfg.addr))

	err = http.ListenAndServe(cfg.addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
