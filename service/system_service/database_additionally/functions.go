package databaseadditionally

import "gorm.io/gorm"

func FunctionMigrations(db *gorm.DB) {
	// генерация реферального кода
	db.Exec(`
        CREATE OR REPLACE FUNCTION generate_referral_code() RETURNS TRIGGER AS $$
        DECLARE
            chars TEXT := 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
            result TEXT := '';
            i INT := 0;
        BEGIN
            FOR i IN 1..10 LOOP
                result := result || SUBSTRING(chars FROM (floor(random()*length(chars))+1)::INT FOR 1);
            END LOOP;
            NEW.referral_code := result;
            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;
    `)

	// Function: update_updated_at_column
	db.Exec(`
        CREATE OR REPLACE FUNCTION update_updated_at_column()
        RETURNS TRIGGER AS $$
        BEGIN
            NEW.updated_at = now();
            RETURN NEW;
        END;
        $$ language 'plpgsql';
    `)

	// Function: set_nickname
	db.Exec(`
        CREATE OR REPLACE FUNCTION set_nickname()
        RETURNS TRIGGER AS $$
        BEGIN
            IF NEW.nickname IS NULL OR NEW.nickname = '' THEN
                NEW.nickname := 'user_' || NEW.serial_id;
            END IF;
            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;
    `)

	// Function: set_nickname_admin
	db.Exec(`
        CREATE OR REPLACE FUNCTION set_nickname_admin()
        RETURNS TRIGGER AS $$
        BEGIN
            IF NEW.nickname IS NULL OR NEW.nickname = '' THEN
                NEW.nickname := 'admin_' || NEW.serial_id;
            END IF;
            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;
    `)

	// Триггер для генерации реферального кода
	db.Exec(`
        CREATE TRIGGER before_insert_users_referral_code
        BEFORE INSERT ON users
        FOR EACH ROW
        EXECUTE FUNCTION generate_referral_code();
    `)

	// Триггер для установки никнейма
	db.Exec(`
        CREATE TRIGGER before_insert_users_nickname
        BEFORE INSERT ON users
        FOR EACH ROW
        EXECUTE FUNCTION set_nickname();
    `)

	db.Exec("ALTER SEQUENCE users_serial_id_seq RESTART WITH 100000")
	db.Exec("ALTER SEQUENCE users_serial_id_seq START WITH 100000")
	db.Exec("ALTER SEQUENCE users_serial_id_seq MINVALUE 100000")
}
