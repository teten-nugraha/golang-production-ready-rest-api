.PHONY: run-dev run-prod swagger test

run-dev:
	export APP_ENV=development && go run cmd/main.go

run-prod:
	export APP_ENV=production && go run cmd/main.go

swagger:
	swag init -g cmd/main.go --output pkg/docs

test:
	go test ./...