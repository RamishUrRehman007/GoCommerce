BEGIN;

DO
$do$
DECLARE
    current_migration_number integer := 2;
BEGIN
    IF NOT EXISTS (SELECT migration_number FROM migrations WHERE migration_number = current_migration_number) THEN
        CREATE EXTENSION IF NOT EXISTS citext;
        CREATE TABLE customer_companies (
            id SERIAL PRIMARY KEY,
            company_name TEXT NOT NULL
        );

        INSERT INTO migrations(migration_number) VALUES (current_migration_number);
    ELSE
        RAISE NOTICE 'Already ran migration %.', current_migration_number;
    END IF;
END
$do$;

COMMIT;
