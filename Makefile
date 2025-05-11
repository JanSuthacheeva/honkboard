#!/usr/bin/env make
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

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

## web/run: Runs the project
.PHONY: web/run
web/run:
	docker-compose up

## db/mysql: Connect to the mysql shell inside docker
.PHONY: db/mysql
db/mysql:
	docker exec -it honkboard-honkboard_mysql-1 bash -c "mysql -u honkboard -p"

## db/migrations/up: apply all new migrations
.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Run migrations'
	goose ${GOOSE_DB_STRING} up -dir ${GOOSE_MIGRATION_DIR}

## db/migrations/down: remove migrations
.PHONY: db/migrations/down
db/migrations/down:
	goose down

## db/migrations/new $name: create a new migration file (sql)
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Create new goose sql migration'
	goose create ${name} sql -dir ${GOOSE_MIGRATION_DIR}

## db/seed: seeds the db
.PHONY: db/seed
db/seed:
	@echo 'Goose seeding DB..'
	goose ${GOOSE_DRIVER} ${GOOSE_DB_STRING} up -dir ./database/seeders -no-versioning

# ===========================================================================
# QUALITY
# ===========================================================================

## code/format: formats all go files
.PHONY: code/format
code/format:
	go fmt ./...
