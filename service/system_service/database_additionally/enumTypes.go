package databaseadditionally

import "gorm.io/gorm"

func EnumMigrations(db *gorm.DB) error {
	// Создание типов ENUM
	err := CreateEnumType(db, "user_status", []string{"active", "banned"}, "Статус пользователя")
	if err != nil {
		return err
	}
	err = CreateEnumType(db, "gender", []string{"male", "female"}, "Пол")
	if err != nil {
		return err
	}
	err = CreateEnumType(db, "sub_status", []string{"active", "expired"}, "Статус подписки")
	if err != nil {
		return err
	}
	err = CreateEnumType(db, "payment_status", []string{"pending", "completed", "failed"}, "Статус платежа")
	if err != nil {
		return err
	}

	return nil
}
