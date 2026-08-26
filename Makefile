include .env
export 

export PROJECT_ROOT=${shell cd}

env-up:
	docker compose up -d todoapp-postgres

env-down:
	docker compose down todoapp-postgres

env-port-forward:
	docker compose up -d port-forwarder
env-port-close:
	docker compose down port-forwarder

env-cleanup:
	@cmd /V:ON /C "set /p ans=\"Clear all volume files? [y/N]: \" && if /I \"!ans!\"==\"y\" (docker compose down todoapp-postgres port-forwarder && if exist out\pgdata (rmdir /s /q out\pgdata) && echo Files deleted\!) else (echo Cleanup canceled.)"

migrate-create:
	@cmd /C "if not defined seq (echo Error: Missing seq parameter && exit 1) else (docker compose run --rm todoapp-postgres-migrate create -ext sql -dir /migrations %seq%)"

migrate-up:
	@$(MAKE) migrate-action action=up

migrate-down:
	@$(MAKE) migrate-action action=down

migrate-action:
	@cmd /C "if not defined action (echo Error: Missing action parameter. Usage: make migrate-action action=up && exit 1) else (docker compose run --rm todoapp-postgres-migrate -path /migrations -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable %action%)"


todoapp-run:
	go mod tidy
	@cmd /C "set LOGGER_FOLDER=$(PROJECT_ROOT)\out\logs&& set POSTGRES_HOST=localhost&& go run ./cmd/todoapp/main.go"
