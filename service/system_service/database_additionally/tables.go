package databaseadditionally

import (
	"cnpc_backend/core/typescore"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// MigrateTables - функция для миграции моделей и проверки существования таблиц
func MigrateTables(db *gorm.DB) error {
	println(" Initialize db migration...")
	// Жестко прописанные имена таблиц
	modelsToMigrate := map[interface{}]string{
		&typescore.UsersProviderControl{}: "users",
		&typescore.APIAccess{}:            "api_access",
		&typescore.LoginActivities{}:      "login_activities",
		&typescore.UserSystemAlerts{}:     "user_alerts",
		&typescore.BlackListIP{}:          "blacklist_ip",
		&typescore.UsersSubscriptions{}:   "user_subscriptions",
		&typescore.Subscription{}:         "subscriptions",
	}
	for model, tableName := range modelsToMigrate {
		if err := MigrateModel(db, model, tableName); err != nil {
			log.Println(" Migration failed for model: ", model, err)
			return err
		}
	}
	err := MigrateKeys(db)
	if err != nil {
		return err
	}
	return nil
}

func MigrateModel(db *gorm.DB, model interface{}, tableName string) error {
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(model)
	if err != nil {
		println(fmt.Sprintf(" Model parsing error: %s", err))
		return err
	}
	stmt.Schema.Table = tableName // Задание имени таблицы
	if !db.Migrator().HasTable(model) {
		println(fmt.Sprintf("⚠️ Table %s doesn't exist. Execute migration...", tableName))
		if err := db.Table(tableName).AutoMigrate(model); err != nil {
			return err
		}
	} else {
		for _, field := range stmt.Schema.Fields {
			if field.IgnoreMigration {
				continue // Пропустить поля, отмеченные для игнорирования
			}
			if !db.Migrator().HasColumn(model, field.DBName) {
				println(fmt.Sprintf("⚠️ Column %s doesn't exist in table %s. Execute migration...", field.DBName, tableName))
				if err := db.Table(tableName).AutoMigrate(model); err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}

func MigrateKeys(db *gorm.DB) error {
	//db.Exec("ALTER TABLE users ADD CONSTRAINT fk_language_code FOREIGN KEY (language) REFERENCES languages(code_639_1) ON DELETE RESTRICT ON UPDATE RESTRICT")
	db.Exec("ALTER TABLE user_alerts ADD CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(system_id) ON DELETE RESTRICT ON UPDATE RESTRICT")
	return nil
}
