package backgroundjob

import (
	"log"
	databaseadditionally "system_service/database_additionally"
	"time"
)

func (m *BackgroundJobModule) StartDatabaseSync() {
	log.Println("📆 StartDatabaseSysns: start", time.Now().UTC())
	grmDB := m.ipc.GormDatabase

	err := databaseadditionally.EnumMigrations(grmDB)
	if err != nil {
		return
	}

	err = databaseadditionally.MigrateTables(grmDB)
	if err != nil {
		return
	}

	databaseadditionally.FunctionMigrations(grmDB)
}
