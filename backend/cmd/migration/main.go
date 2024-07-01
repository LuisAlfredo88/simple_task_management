package main

import (
	"stm/initialize/db/sqlite"
)

func main() {
	db := sqlite.MigrationDBInstance()
	sqlite.DoMigration(db)
}
