-include app.env
export $(shell test -f app.env && sed 's/=.*//' app.env)

prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest
	@echo "Installing goose"
	go install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "Installing wire"
	go install github.com/google/wire/cmd/wire@latest
	@echo "Copy app.env.example to app.env if app.env does not exists"
	test -f app.env || cp -a app.env.example app.env
	@echo "Install Husky"
	cd ../ && go install github.com/go-courier/husky/cmd/husky@latest && husky init

dockerconf:
	docker-compose --env-file ./app.env config

start:
	docker-compose --env-file ./app.env up -d

stop:
	docker-compose stop

destroy:
	docker-compose down -v

clean:
	[ -d dist ] && rm dist/*

di:
	@wire

# Usage: make seq=${seq} migration
gen_migration:
	$(info Generating migration sequence $(seq) to migration)
	goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" create $(seq) sql

migup:
	goose -allow-missing -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migup1:
	goose -allow-missing -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up-by-one

migdown1:
	goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

migreset:
	goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" reset

migstat:
	goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" status

migver:
	goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" version

migfix:
	goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" fix

execdb:
	docker-compose exec db psql -U postgres $(DB_NAME)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

build: di
	go build -o dist/app app.go wire_gen.go

rundev: di
	export APP_ENV=local && go run app.go wire_gen.go

.PHONY: prepare dockerconf start stop destroy clean di gen_migration migup migup1 migdown1 migreset migstat migver migfix execdb sqlc test mock build
