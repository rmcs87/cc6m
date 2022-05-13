package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rmcs87/cc6m/pkg/models/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "Porta do Endereço")

	dsn := flag.String("dns", "RNNRfSNISQ:jMFmtb1Hby@tcp(remotemysql.com)/RNNRfSNISQ?parseTime=true", "MySql DSN")

	flag.Parse()

	infoLogger := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLogger.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLogger,
		infoLog:  infoLogger,
		snippets: &mysql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	infoLogger.Printf("Iniciando o Servidor na Porta %s\n", *addr)
	err = srv.ListenAndServe()
	errorLogger.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
