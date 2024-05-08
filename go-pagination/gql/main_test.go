package gql

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratePSQL "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var cleanUp []func()
	testDB, cleanUp = newDB()

	code := m.Run()

	slices.Reverse(cleanUp)
	for _, fn := range cleanUp {
		fn()
	}

	os.Exit(code)
}

func newDB() (*sql.DB, []func()) {
	ctx := context.Background()

	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.2"),
		postgres.WithDatabase("library_db"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		testcontainers.WithWaitStrategy(
			wait.
				ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		log.Fatalf("failed not start up container to postgres %v", err)
	}

	connURI, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("failed to get connection string to postgres %v", err)
	}

	var cleanup []func()

	purgeContainer := func() {
		if err = container.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}

		log.Println("terminate container postgres success")
	}

	cleanup = append(cleanup, purgeContainer)

	var db *sql.DB
	db, err = sql.Open("postgres", connURI)
	if err != nil {
		log.Fatalf("Could not open DB: %s", err)
	}

	closeConnDB := func() {
		if err = db.Close(); err != nil {
			log.Fatalf("Could not close DB: %s", err)
			return
		}

		log.Println("close connection db success")
	}

	cleanup = append(cleanup, closeConnDB)

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
