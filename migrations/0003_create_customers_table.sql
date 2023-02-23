BEGIN;

DO
$do$
DECLARE
    current_migration_number integer := 3;
BEGIN
    IF NOT EXISTS (SELECT migration_number FROM migrations WHERE migration_number = current_migration_number) THEN
        CREATE EXTENSION IF NOT EXISTS citext;
        CREATE TABLE customers (
            user_id CITEXT PRIMARY KEY,
            login CITEXT NOT NULL,
            password TEXT NOT NULL,
            name CITEXT NOT NULL,
            company_id INT NOT NULL REFERENCES customer_companies(id),
            credit_cards TEXT NOT NULL
        );

        INSERT INTO migrations(migration_number) VALUES (current_migration_number);
    ELSE
        RAISE NOTICE 'Already ran migration %.', current_migration_number;
    END IF;
END
$do$;

COMMIT;
