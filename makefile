.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./db/migrations ${name}
.PHONY: db/migrations/up
db/migrations/up: 
	@echo 'Running up migrations...'
	migrate -path ./db/migrations -database ${url} up
.PHONY: run
run:
	@echo "Loading environment variables from .env file"
	@set -o allexport; source ./load_env.sh; set +o allexport; \
	echo "Running Go application"; \
	go run ./cmd/main.go
.PHONY: tailwind
tailwind:
	tailwindcss -i ./internal/web/static/css/input.css -o ./internal/web/static/css/styles.css --watch
.PHONY:templ
templ:
	templ generate --watch
.PHONY: air
air:
	@echo "Loading environment variables from .env file"
	@set -o allexport; source ./load_env.sh; set +o allexport; \
	echo "Running air"; \
	air -c .air.toml
.PHONY: lint
lint:
	golangci-lint run
.PHONY: test
test:
	go test  ./... 
.PHONY: coverage
coverage:
	go test  -coverprofile=coverage.out ./... ;
	go tool cover -func=coverage.out
.PHONY: build
build:
	go build -o ./bin ./cmd