package test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
)

type ShutDownFunc func()

type Config struct {
	Schemas  []string
	TestData []string
	Schema   string
}

const (
	defaultPort                 = "5432"
	defaultSSLMode              = "disable"
	defaultSchema               = "public"
	defaultConnectionTimeout    = 50 * time.Second
	defaultConnectionPollPeriod = 500 * time.Millisecond
)

// MustConnectToPG connects to test docker postgres container
func MustConnectToPG(cfg *Config) (*sql.DB, ShutDownFunc) {
	return mustConnectToContainer("127.0.0.1", cfg)
}

func mustConnectToContainer(hostname string, cfg *Config) (*sql.DB, ShutDownFunc) {
	var binds map[string]string
	schema := defaultSchema
	if cfg != nil {
		const prefixTemplate = "/docker-entrypoint-initdb.d/%d-"
		binds = make(map[string]string, len(cfg.Schemas)+len(cfg.TestData))
		i := 0
		for _, shm := range cfg.Schemas {
			binds[shm] = setFilePrefix(fmt.Sprintf(prefixTemplate, i), shm)
			i++
		}
		for _, td := range cfg.TestData {
			binds[td] = setFilePrefix(fmt.Sprintf(prefixTemplate, i), td)
			i++
		}
		if cfg.Schema != "" {
			schema = cfg.Schema
		}
	}

	req := testcontainers.ContainerRequest{
		Image:        "postgres:11.6-alpine",
		ExposedPorts: []string{defaultPort + "/tcp"},
		NetworkMode:  testcontainers.Bridge,
		BindMounts:   binds,
		WaitingFor:   wait.ForLog("database system is ready to accept connections"),
	}

	ctx := context.Background()
	pgc, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to create container: %v\n", err)
	}

	port, err := pgc.MappedPort(ctx, defaultPort)
	if err != nil {
		log.Fatalf("Failed to map port %s: %v\n", defaultPort, err)
	}

	var db *sql.DB
	db, err = connectWithTimeout(defaultConnectionTimeout, hostname, port.Port(), "postgres", "postgres", "", defaultSSLMode, schema)
	if err != nil {
		mustTerminate(ctx, pgc)
		log.Fatalf("Failed to connect to container: %v\n", err)
	}

	return db, func() { mustTerminate(ctx, pgc) }
}

func setFilePrefix(prefix, filename string) string {
	_, file := filepath.Split(filename)
	return prefix + file
}

func mustTerminate(ctx context.Context, c testcontainers.Container) {
	if err := c.Terminate(ctx); err != nil {
		log.Fatalf("Failed to terminate container: %v\n", err)
	}
}

func connectWithTimeout(timeout time.Duration, host string, port string, name string, user string, pwd string, sslMode string, schema string) (*sql.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	dbCh := make(chan *sql.DB)
	defer close(dbCh)

	errCh := make(chan error)
	defer close(errCh)

	go connect(dbCh, errCh, host, port, name, user, pwd, sslMode, schema)

	select {
	case db := <-dbCh:
		return db, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func connect(dbCh chan *sql.DB, errCh chan error, host string, port string, name string, user string, pwd string, sslMode string, schema string) {
	connStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user='%s' password='%s' sslmode=%s search_path=%s",
		host, port, name, user, pwd, sslMode, schema,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		errCh <- err
		return
	}

	err = db.Ping()
	for err != nil {
		time.Sleep(defaultConnectionPollPeriod)
		err = db.Ping()
	}

	dbCh <- db
}
