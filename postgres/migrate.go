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
	"path/filepath"
)

type Migrate struct {
	migrationsPath string
	migrate *migrate.Migrate
}

func (m *Migrate) loadMigrate() error  {
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	fmt.Println(dbURL)

	absMigrationPath, err := filepath.Abs(m.migrationsPath)

	if err != nil{
		return err
	}

	m.migrate, err = migrate.New(
		fmt.Sprintf("file://%s", absMigrationPath),
		dbURL,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *Migrate) up() error {
	fmt.Println("Migrate up")
	err := m.migrate.Steps(1)

	if err != nil {
		return err
	}

	fmt.Println("Migrate up was successful")
	return nil
}

func (m *Migrate) upAll() error {
	fmt.Println("Migrate up all")
	err := m.migrate.Up()

	if err != nil {
		return err
	}

	fmt.Println("Migrate up all was successful")
	return nil
}

func (m *Migrate) down() error  {
	fmt.Println("Migrate down")
	err := m.migrate.Steps(-1)

	if err != nil {
		return err
	}

	fmt.Println("Migrate down was successful")
	return nil
}

func (m *Migrate) downAll() error  {
	fmt.Println("Migrate down all")
	err := m.migrate.Down()

	if err != nil {
		return err
	}

	fmt.Println("Migrate down all was successful")
	return nil
}

func (m *Migrate) getVersion() (string, error)  {
	currentVersion, _, err := m.migrate.Version()

	newVersion := fmt.Sprintf("%06d", currentVersion + 1)

	return newVersion, err
}

func (m *Migrate) create(fileName string) error {
	fmt.Println("Create migration")

	newVersion, err := m.getVersion()
	fmt.Println("newVersion", newVersion)
	upPath := fmt.Sprintf("%s/%s_%s.up.sql", m.migrationsPath, newVersion, fileName)
	downPath := fmt.Sprintf("%s/%s_%s.down.sql", m.migrationsPath, newVersion, fileName)
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

	var m Migrate

	isDown := flag.Bool("down", false, "migrate-down")
	isUp := flag.Bool("up", false, "migrate-up")
	isUpAll := flag.Bool("upAll", false, "migrate-up-all")
	isDownAll := flag.Bool("downAll", false, "migrate-down-all")
	fileName := flag.String("create", "", "migrate-create")
	migrationsPath := flag.String("migrationsPath", "", "migrate-create")
	flag.Parse()

	m.migrationsPath = *migrationsPath
	err = m.loadMigrate()
	if err != nil {
		fmt.Println(err)
		return
	}

	if *fileName != "" && *migrationsPath != "" {
		err := m.create(*fileName)
		if err != nil {
			fmt.Println(err)
			return 
		}
		return
	}

	if *isDown && *migrationsPath != "" {
		err := m.down()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	if *isUp && *migrationsPath != "" {
		err := m.up()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	if *isUpAll && *migrationsPath != "" {
		err := m.upAll()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	if *isDownAll && *migrationsPath != "" {
		err := m.downAll()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	fmt.Println("Invalid arguments")
	return
}



