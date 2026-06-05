package migration

import (
	"context"
	"embed"
	"io/fs"
	"sort"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed *.sql
var migrationsFS embed.FS

type Migrator struct {
	db *pgxpool.Pool
}

func NewMigrator(db *pgxpool.Pool) *Migrator {
	return &Migrator{db: db}
}

func (m *Migrator) Up(ctx context.Context) error {
	files, err := fs.Glob(migrationsFS, "*.sql")
	if err != nil {
		return err
	}

	sort.Strings(files)

	for _, file := range files {
		content, err := fs.ReadFile(migrationsFS, file)
		if err != nil {
			return err
		}

		if _, err := m.db.Exec(ctx, string(content)); err != nil {
			return err
		}
	}

	return nil
}
