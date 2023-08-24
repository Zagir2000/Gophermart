package main

import (
	"flag"
	"os"
)

type FlagVar struct {
	runAddr       string
	databaseURI   string
	migrationsDir string
}

func NewFlagVarStruct() *FlagVar {
	return &FlagVar{}
}
func (f *FlagVar) parseFlags() error {
	//postgresql://postgres:docker@localhost:5432/asd?sslmode=disable
	// как аргумент -a со значением :8080 по умолчанию
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.StringVar(&f.runAddr, "a", "localhost:8081", "address and port to run server")
	flag.StringVar(&f.databaseURI, "d", "postgres://postgres:123456@localhost/gofermart?sslmode=disable", "database connection address")
	flag.StringVar(&f.migrationsDir, "m", "migrations", "migrations to db")
	flag.Parse()

	if envRunAddr, ok := os.LookupEnv("RUN_ADDRESS"); ok {
		f.runAddr = envRunAddr
	}

	if envDatabaseURI, ok := os.LookupEnv("DATABASE_URI"); ok {
		f.databaseURI = envDatabaseURI
	}

	if envMigrationsDir, ok := os.LookupEnv("MIGRATIONS_DIR"); ok {
		f.migrationsDir = envMigrationsDir
	}
	return nil
}
