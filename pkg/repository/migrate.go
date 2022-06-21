package repository

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"
	// need for migrate
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// need for migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getDSN(cfg Config) string {
	const format = "postgres://%s:%s@%s:%s/%s?sslmode=%s"

	dsn := fmt.Sprintf(
		format,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	return dsn
}

func InitMigrate(cfg Config) {
	dsn := getDSN(cfg)

	if dsn == "" {
		logrus.Fatal("migrate: environment variable not declared")
	}

	m, err := migrate.New("file://migrations", dsn)

	if err != nil {
		logrus.Fatal(err)
	}

	if err := m.Up(); err != nil {
		logrus.Debug(err)
	}

	m.Close()
}
