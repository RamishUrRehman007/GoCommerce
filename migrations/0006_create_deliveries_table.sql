BEGIN;

DO
$do$
DECLARE
    current_migration_number integer := 6;
BEGIN
    IF NOT EXISTS (SELECT migration_number FROM migrations WHERE migration_number = current_migration_number) THEN
        CREATE TABLE deliveries (
            id SERIAL PRIMARY KEY,
            order_item_id INT NOT NULL REFERENCES order_items(id),
            delivered_quantity INT NOT NULL
        );

        INSERT INTO migrations(migration_number) VALUES (current_migration_number);
    ELSE
        RAISE NOTICE 'Already ran migration %.', current_migration_number;
    END IF;
END
$do$;

COMMIT;
