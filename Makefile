include .env

GOPACKAGES=$$(go list ./... )

# ensure makes sure that all dependencies will be downloaded
ensure:
	go get $(GOPACKAGES)

# tidy runs go modules tidy tool
tidy:
	@go mod tidy

# build the whole application using Go locally
build: swag
	@go build -v --ldflags="-s" ./cmd/api

# build and run service
run:
	@go run cmd/api/main.go

# run unit tests
test:
	go test -coverprofile coverage.out -failfast -short $(GOPACKAGES)
	go tool cover -func coverage.out | grep total

# run lint
lint: #fix
	golangci-lint run $(GOPACKAGES)

cover-clear:
	rm -f cover.out

coverage: 
	go test -tags="all" -covermode="count" -coverprofile="cover.out" $(GOPACKAGES)

# =========== Swagger =============
check/swag:
	which swag || go get -u github.com/swaggo/swag/cmd/swag

swag: check/swag
	swag init -g cmd/api/main.go

# =========== Working with Docker =============
docker/build:
	docker build -t transaction ./

docker/run:
	docker-compose up -d

docker/down:
	docker-compose down

docker/migrate/up:
	docker run --rm --network=host -v $(PWD)/migrations:/migrations migrate/migrate -path=/migrations -database "postgres://${DATABASE_USER}:${DATABASE_PWD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_BASE}?sslmode=disable" up

docker/migrate/down:
	docker run --rm --network=host -v $(PWD)/migrations:/migrations migrate/migrate -path=/migrations -database "postgres://${DATABASE_USER}:${DATABASE_PWD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_BASE}?sslmode=disable" down -all
