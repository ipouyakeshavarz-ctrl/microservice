package migrator

import (
	"database/sql"
	"fmt"
	"productapp/internal/repository/mysql"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	dbConfig   mysql.Config
	migrations *migrate.FileMigrationSource
	tableNames string
}

func New(dbConfig mysql.Config) Migrator {
	migrations := &migrate.FileMigrationSource{

		Dir: "./migrations",
	}

	return Migrator{dialect: "mysql",
		dbConfig:   dbConfig,
		migrations: migrations,
		tableNames: "productapp_migrations"}
}

func (m Migrator) Up() {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	migrate.SetTable(m.tableNames)

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't execute migration: %v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	migrate.SetTable(m.tableNames)

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't execute migration: %v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)

}
