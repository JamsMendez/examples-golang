package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratePSQL "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var cleanUp []func()
	testDB, cleanUp = newDB()

	code := m.Run()

	for _, fn := range cleanUp {
		fn()
	}

	os.Exit(code)
}

func newDB() (*sql.DB, []func()) {
	dns := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("test", "test"),
		Path:   "library_db",
	}

	q := dns.Query()
	q.Add("sslmode", "disable")
	dns.RawQuery = q.Encode()

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Colud not construct pool: %s", err)
	}

	pool.MaxWait = 10 * time.Second

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	pwd, ok := dns.User.Password()
	if !ok {
		log.Fatalf("Could not get user password to Postgres: %s", err)
	}

	// pull an image, creates a container based on it and runs it
	options := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15.2",
		Env: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", pwd),
			fmt.Sprintf("POSTGRES_USER=%s", dns.User.Username()),
			fmt.Sprintf("POSTGRES_DB=%s", dns.Path),
		},
	}

	resource, err := pool.RunWithOptions(options, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// Tell docker to hard kill the container in 120 seconds
	err = resource.Expire(120)
	if err != nil {
		log.Fatalf("Could not sets expire associated container: %s", err)
	}

	var cleanup []func()

	purgePool := func() {
		if err = pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge container: %s", err)
		}
	}

	cleanup = append(cleanup, purgePool)

	dns.Host = fmt.Sprintf("%s:5432", resource.Container.NetworkSettings.IPAddress)
	log.Println("Connecting to database on url: ", dns.String())

	var db *sql.DB
	db, err = sql.Open("postgres", dns.String())
	if err != nil {
		log.Fatalf("Could not open DB: %s", err)
	}

	closeConnDB := func() {
		if err = db.Close(); err != nil {
			log.Fatalf("Could not close DB: %s", err)
		}
	}

	cleanup = append(cleanup, closeConnDB)

	if err = pool.Retry(func() error {
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not ping DB: %s", err)
	}

	driver, err := migratePSQL.WithInstance(db, &migratePSQL.Config{})
	if err != nil {
		log.Fatalf("Could not migrate (1): %s", err)
	}

	var m *migrate.Migrate
	m, err = migrate.NewWithDatabaseInstance("file://./../db/migration", "postgres", driver)
	if err != nil {
		log.Fatalf("Could not migrate (2): %s", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Could not migrate (3): %s", err)
	}

	return db, cleanup
}
