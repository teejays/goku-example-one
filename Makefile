
# Color Control Sequences for easy printing
RESET=\033[0m
RED=\033[31;1m
GREEN=\033[32;1m
YELLOW=\033[33;1m
BLUE=\033[34;1m
MAGENTA=\033[35;1m
CYAN=\033[36;1m
WHITE=\033[37;1m

GO=go
GOKU=goku

CURRENT_DIR := $(shell pwd)
PATH_TO_APP=.
PATH_TO_MODELS=$(PATH_TO_APP)/models/example.json
PATH_TO_GEN=../goku/generator
DB_USER=${USER}

# Group commands: do more than one thing at once
all: clean goku-generate db-start-if-stopped create-dbs generate-db-migration migrate-db remove-empty-migration-files

# Setup on a new machine
setup: setup-yamltodb setup-postgres
	brew install yamllint && \
	brew install python-yq

setup-postgres:
	brew install postgres

setup-yamltodb: setup-postgres
	python3 -m pip install --upgrade pip && \
	brew install openssl && \
	export LDFLAGS="-L/usr/local/opt/openssl@1.1/lib" && \
	export CPPFLAGS="-I/usr/local/opt/openssl@1.1/include" && \
	pip install psycopg2 && \
	pip3 install Pyrseas

help:
	./bin/generator.bin --help

# Run Commands: 

goku-generate:
	@echo "$(YELLOW)Running Goku...$(RESET)"
		$(GOKU) \
		--generator-dir="$(PATH_TO_GEN)" \
		--models-dir="$(PATH_TO_APP)/models" \
		--models-json-file="$(PATH_TO_MODELS)" \
		--templates-dir="$(PATH_TO_GEN)/templates" \
		--app-root-dir="$(PATH_TO_APP)" \
		--golang-root-dir="github.com/teejays/goku-example-one" \
		--sql-yaml-schema=true --sql-yaml-schema-dir="$(PATH_TO_APP)/db/schema" \
		--golang-type-definitions=true --golang-type-definitions-dir="$(PATH_TO_APP)/backend/types" \
		--golang-meta-info=true --golang-meta-info-dir="$(PATH_TO_APP)/backend/meta" \
 		--golang-dal=true --golang-dal-dir="$(PATH_TO_APP)/backend/dal" \
		--golang-type-filters=true --golang-type-filters-dir="$(PATH_TO_APP)/backend/filters" \
		--golang-generics=true --golang-generics-dir="$(PATH_TO_APP)/backend" \
		--golang-methods=true --golang-methods-dir="$(PATH_TO_APP)/backend/methods" \
		--golang-http-handlers=true --golang-http-handlers-dir="$(PATH_TO_APP)/backend/httphandlers" \
		--type-script-types=true --type-script-types-dir="$(PATH_TO_APP)/frontend/admin/src" \
		--graphql-schema=true --graphql-schema-dir="$(PATH_TO_APP)/backend" \
		--graphql-resolver=true --graphql-resolver-dir="$(PATH_TO_APP)/backend" \

# Generate migration SQL scripts
CMD_RM_MIGRATION=rm -rf $(PATH_TO_APP)/db/migration/future/*
CMD_CREATE_MIGRATION_FOLDER_FUTURE=mkdir -p $(PATH_TO_APP)/db/migration/future/{}
CMD_CREATE_MIGRATION_FOLDER_PRESENT=mkdir -p $(PATH_TO_APP)/db/migration/present/{}
CMD_CREATE_MIGRATION_FOLDER_PAST=mkdir -p $(PATH_TO_APP)/db/migration/past/{}
CMD_GENERATE_DB_MIGRATION=yamltodb -r "$(PATH_TO_APP)/db/schema/{}" -c "$(PATH_TO_APP)/db/pyrseas-yamltodb.config.yaml" -m -o $(PATH_TO_APP)/db/migration/future/{}/db.{}.migration.sql {}
generate-db-migration: create-dbs
	@echo "$(YELLOW)Generating DB Migrations...$(RESET)"
	$(CMD_RM_MIGRATION) && \
	xargs -t -I{} $(CMD_CREATE_MIGRATION_FOLDER_FUTURE) <$(PATH_TO_APP)/db/schema/databases.generated.txt && \
	xargs -t -I{} $(CMD_CREATE_MIGRATION_FOLDER_PRESENT) <$(PATH_TO_APP)/db/schema/databases.generated.txt && \
	xargs -t -I{} $(CMD_CREATE_MIGRATION_FOLDER_PAST) <$(PATH_TO_APP)/db/schema/databases.generated.txt && \
	xargs -t -I{} $(CMD_GENERATE_DB_MIGRATION) <$(PATH_TO_APP)/db/schema/databases.generated.txt

# Run the migrations: Move the migration.sql file to 'present', run it, move it to 'past'
CMD_MOVE_MIGRATIONS_TO_PRESENT=mv $(PATH_TO_APP)/db/migration/future/{}/db.{}.migration.sql $(PATH_TO_APP)/db/migration/present/{}/db.{}.migration.sql
CMD_MOVE_MIGRATIONS_TO_PAST=mv $(PATH_TO_APP)/db/migration/present/{}/db.{}.migration.sql $(PATH_TO_APP)/db/migration/past/{}/db.{}.$$(date +%Y_%m_%d_%H%M%S).migration.sql
CMD_RUN_MIGRATIONS=psql -U ${DB_USER} --dbname={} --single-transaction --file=$(PATH_TO_APP)/db/migration/present/{}/db.{}.migration.sql
migrate-db:
	@echo "$(YELLOW)Running DB Migrations...$(RESET)"
	xargs -t -I{} $(CMD_MOVE_MIGRATIONS_TO_PRESENT) <$(PATH_TO_APP)/db/schema/databases.generated.txt && \
	xargs -t -I{} $(CMD_RUN_MIGRATIONS) <$(PATH_TO_APP)/db/schema/databases.generated.txt && \
	xargs -t -I{} $(CMD_MOVE_MIGRATIONS_TO_PAST) <$(PATH_TO_APP)/db/schema/databases.generated.txt

CMD_TO_DELETE_EMPTY_FILES_IN_DIR=find $(PATH_TO_APP)/db/migration/past/{} -size  0 -print -delete
remove-empty-migration-files:
	@echo "$(YELLOW)Removing any empty migration files...$(RESET)"
	xargs -t -I{} $(CMD_TO_DELETE_EMPTY_FILES_IN_DIR) <$(PATH_TO_APP)/db/schema/databases.generated.txt

# Migrations for the test-db: Copy schema of non-test TB to test-db
CMD_SYNC_TEST_DB=dbtoyaml {} | yamltodb --update {}_test
sync-test-dbs: create-test-dbs
	xargs -n 1 -I{} sh -c "$(CMD_SYNC_TEST_DB)" <$(PATH_TO_APP)/db/schema/databases.generated.txt

# Intent: Remove all generated & git ignored files.
clean:
	@echo "$(YELLOW)Removing all generated files...$(RESET)"
	rm -rf ./bin/* && \
	rm -rf $(PATH_TO_APP)/db/schema/* && \
	rm -rf $(PATH_TO_APP)/db/migration/future/* && \
	find . -type d -name goku.generated -prune -exec rm -rf {} \;

# Create Databases (so we can create migrations)
CMD_CREATE_DB=./scripts/db_create.sh {}
create-dbs:
	@echo "$(YELLOW)Creating databases (if needed)...$(RESET)"
	xargs -n 1 -I{} $(CMD_CREATE_DB) <$(PATH_TO_APP)/db/schema/databases.generated.txt DB_USER=${DB_USER}

# Create a test database for each service, named <service>_test
CMD_CREATE_TEST_DB=./scripts/db_create.sh {}_test
create-test-dbs:
	xargs -n 1 -I{} $(CMD_CREATE_TEST_DB) <$(PATH_TO_APP)/db/schema/databases.generated.txt DB_USER=${DB_USER}

# We need a postgres role in these databases
CMD_SETUP_DB_ROLES=psql -d {} -c "CREATE ROLE postgres WITH LOGIN SUPERUSER CREATEROLE CREATEDB REPLICATION;"
setup-db-roles:
	xargs -n 1 -I{} $(CMD_SETUP_DB_ROLES) <$(PATH_TO_APP)/db/schema/databases.generated.txt

# app-run
run-backend: db-start-if-stopped app-run-backend 
run-frontend: app-build-frontend-admin app-run-frontend-admin

app-run-backend:
	$(GO) run $(PATH_TO_APP)/backend/main.go
app-run-gateway:
	$(GO) run $(PATH_TO_APP)/backend/gateway/gateway.go

PATH_TO_FRONTEND_ADMIN=$(PATH_TO_APP)/frontend/admin

app-build-frontend-admin:
	yarn --cwd=$(PATH_TO_FRONTEND_ADMIN)

app-run-frontend-admin:
	yarn --cwd=$(PATH_TO_FRONTEND_ADMIN) start

# Reference Commands: not needed to run
db-start:
	pg_ctl -D /usr/local/var/postgres start
db-stop:
	pg_ctl -D /usr/local/var/postgres stop
db-start-if-stopped:
ifeq ($(shell pg_ctl -D /usr/local/var/postgres status > /dev/null; echo $$?),3)
	pg_ctl -D /usr/local/var/postgres start
endif
db-destroy:
	xargs -n 1 -I{} dropdb {} <$(PATH_TO_APP)/db/schema/databases.generated.txt

db-connect:
	psql --db postgres

CMD_DBTOYAML=dbtoyaml -c $(PATH_TO_APP)/db/pyrseas-dbtoyaml.config.yaml -r $(PATH_TO_APP)/db/schema/{} -m {}

dbtoyaml:
	xargs -t -I{} $(CMD_DBTOYAML) <$(PATH_TO_APP)/db/schema/databases.generated.txt