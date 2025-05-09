package main

import (
	"crypto/tls"
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
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jansuthacheeva/honkboard/internal/mailer"
	"github.com/jansuthacheeva/honkboard/internal/models"
	"github.com/joho/godotenv"
)

type application struct {
	cfg             config
	logger          *slog.Logger
	todos           models.TodoModelInterface
	users           *models.UserModel
	validationCodes *models.ValidationCodeModel
	sessionManager  *scs.SessionManager
	templateCache   map[string]*template.Template
	formDecoder     *form.Decoder
	mailer          *mailer.Mailer
}

type config struct {
	addr string
	dsn  string
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config

	flag.StringVar(&cfg.dsn, "dsn", os.Getenv("DB_STRING"), "MySQL data source name")
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "b8b37b30a398d8", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "0f69e0fd4663f7", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Honkboard <no-reply@honkboard.com>", "SMTP sender")

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

	formDecoder := form.NewDecoder()

	mailer, err := mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	app := application{
		cfg:             cfg,
		logger:          logger,
		todos:           &models.TodoModel{DB: db},
		users:           &models.UserModel{DB: db},
		validationCodes: &models.ValidationCodeModel{DB: db},
		sessionManager:  sessionManager,
		templateCache:   templateCache,
		formDecoder:     formDecoder,
		mailer:          mailer,
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         cfg.addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.logger.Info("Starting server", slog.String("addr", srv.Addr))

	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")

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
