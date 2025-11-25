MIGRATE = migrate
MIGRATIONS_DIR = db/migrations
DB_URL = "postgres://admin:123456@localhost:5432/family_db?sslmode=disable&timezone=Asia%2FJakarta"



## Create new migration
create-migration:
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) $(name)

## Run migrations
migrate-up:
	$(MIGRATE) -database $(DB_URL) -path $(MIGRATIONS_DIR) up

## Rollback 1 step
migrate-down:
	$(MIGRATE) -database $(DB_URL) -path $(MIGRATIONS_DIR) down 1

## Drop all migrations
migrate-drop:
	$(MIGRATE) -database $(DB_URL) -path $(MIGRATIONS_DIR) drop -f
