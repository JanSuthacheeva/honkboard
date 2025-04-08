#!/usr/bin/env make

# ===========================================================================
# HELPERS
# ===========================================================================

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ===========================================================================
# DEVELOPMENT
# ===========================================================================

## db/migrations/up: apply all new migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Run migrations'
	goose ${GOOSE_DRIVER} ${GOOSE_DBSTRING} up -dir ${GOOSE_MIGRATION_DIR}

## db/migrations/down: remove migrations
.PHONY: db/migrations/down
db/migrations/down:
	goose down

## db/migrations/new: create a new migration file (sql)
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Create new goose sql migration'
	goose create ${name} sql

## db/seed: seeds the db
.PHONY: db/seed
db/seed:
	@echo 'Goose seeding DB..'
	goose ${GOOSE_DRIVER} ${GOOSE_DB_STRING} up -dir ./database/seeders -no-versioning
