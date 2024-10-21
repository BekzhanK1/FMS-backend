build:
	go build -o bin/user-service cmd/app/main.go

test:
	go test -v ./...

run: build
	./bin/user-service

migration:
	migrate create -ext sql -dir migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

migrate-down-up:
	go run cmd/migrate/main.go down
	go run cmd/migrate/main.go up