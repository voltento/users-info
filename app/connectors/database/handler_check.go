package database

import (
	"context"
	"github.com/go-pg/pg"
)

// According to https://github.com/go-pg/pg/wiki/FAQ#how-to-check-connection-health
func (d *dataBase) Check(ctx context.Context) error {
	var n int
	_, err := d.db.WithContext(ctx).QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}
