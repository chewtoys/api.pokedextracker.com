package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		if _, err := db.Exec(`
			CREATE TABLE games (
				id TEXT PRIMARY KEY,
				name TEXT NOT NULL,
				game_family_id TEXT NOT NULL,
				"order" INTEGER NOT NULL
			)
		`); err != nil {
			return err
		}
		if _, err := db.Exec("ALTER TABLE games ADD CONSTRAINT games_game_family_id_foreign FOREIGN KEY (game_family_id) REFERENCES game_families (id)"); err != nil {
			return err
		}
		if _, err := db.Exec("CREATE INDEX games_game_family_id_index ON games (game_family_id)"); err != nil {
			return err
		}
		_, err := db.Exec(`ALTER TABLE games ADD CONSTRAINT games_order_unique UNIQUE ("order")`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE games")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20180817141015_create_games_table", up, down, opts)
}