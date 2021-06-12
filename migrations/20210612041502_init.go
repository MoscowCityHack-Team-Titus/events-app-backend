package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

const migrationUp = `
	CREATE TABLE users (
		id serial NOT NULL,
		username VARCHAR(255) NOT NULL,
		additional VARCHAR(255)
	);
`

const migrationDown = `
	DROP TABLE users;
`

func init() {
	goose.AddMigration(upInit, downInit)
}

func upInit(tx *sql.Tx) error {
	_, err := tx.Exec(migrationUp)
	if err != nil {
		return err
	}

	return nil
}

func downInit(tx *sql.Tx) error {
	_, err := tx.Exec(migrationDown)
	if err != nil {
		return err
	}

	return nil
}
