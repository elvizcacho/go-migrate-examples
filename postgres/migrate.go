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

func getMigrate() (*migrate.Migrate, error)  {
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	fmt.Println(dbURL)
	return migrate.New(
		"file:///Users/juanvizcaino/migrations/postgres/db/migrations",
		dbURL,
	)
}

func up() error {
	m, err := getMigrate()

	if err != nil {
		return err
	}

	fmt.Println("Migrate up")
	err = m.Up()

	if err != nil {
		return err
	}

	fmt.Println("Migrate up was successful")
	return nil
}

func down() error  {
	m, err := getMigrate()

	err = m.Down()

	if err != nil {
		return err
	}

	fmt.Println("Migrate down was successful")
	return nil
}

func getVersion() (string, error)  {
	m, err := getMigrate()
	if err != nil {
		return "", err
	}

	currentVersion, _, err := m.Version()

	newVersion := fmt.Sprintf("%06d", currentVersion + 1)

	return newVersion, err
}

func create(fileName string, migrationsPath string) error {
	newVersion, err := getVersion()
	upPath := fmt.Sprintf("%s/%s_%s.up.sql", migrationsPath, newVersion, fileName)
	downPath := fmt.Sprintf("%s/%s_%s.down.sql", migrationsPath, newVersion, fileName)
	err = ioutil.WriteFile(upPath, []byte{}, 0644)
	err = ioutil.WriteFile(downPath, []byte{}, 0644)
	if err != nil {
		return err
	}

	return nil
}



func main() {
	err := godotenv.Load("./.env")

	if err != nil {
		fmt.Println(err)
		return
	}

	isDown := flag.Bool("down", false, "migrate-down")
	isUp := flag.Bool("up", false, "migrate-up")
	fileName := flag.String("create", "", "migrate-create")
	migrationsPath := flag.String("migrationsPath", "", "migrate-create")
	flag.Parse()

	if *fileName != "" && *migrationsPath != "" {
		err := create(*fileName, *migrationsPath)
		if err != nil {
			fmt.Println(err)
			return 
		}
		return
	}

	if *isDown {
		err := down()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	if *isUp {
		err := up()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	return
}



