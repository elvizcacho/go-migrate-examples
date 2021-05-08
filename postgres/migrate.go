package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
)

func up(m *migrate.Migrate)  {
	fmt.Println("Migrate up")
	err := m.Up()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Migrate up was successful")
	return
}

func down(m *migrate.Migrate)  {
	err := m.Down()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Migrate down was successful")
	return
}

func create(fileName string)  {
	path := fmt.Sprintf("./migrations/%s.sql", fileName)
	ioutil.WriteFile(path, []byte{}, 0644)
}

func main() {
	err := godotenv.Load("/Users/juanvizcaino/Projects/juapp/juapp-git-service/.env")
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	fmt.Println(dbURL)
	m, err := migrate.New(
		"file:///Users/juanvizcaino/Projects/juapp/juapp-git-service/db/migrations",
		dbURL,
	)

	if err != nil {
		panic(err)
		return
	}



	isDown := flag.Bool("down", false, "migrate-down")
	isUp := flag.Bool("up", false, "migrate-up")
	fileName := flag.String("create", "", "migrate-create")
	flag.Parse()

	fmt.Println(*isDown, *isUp)

	if *isDown {
		down(m)
	}

	if *isUp {
		up(m)
	}

	if *fileName != "" {
		create(fileName)
	}

	return
}



