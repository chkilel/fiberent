clean:
		rm -rf target; \
		rm -f coverage.*

# Download verborse:
deps: clean
		go get -d -v ./...

# Dev with hot reload
docker-dev: deps
		docker-compose up

# Migrate Ent scheme to database
migrate:
	go run ./cmd/migration/main.go

# Start dev server if using docker container only for DB and adminer
# Go and Air must be installed locally and Air configured
# as per the docs https://github.com/cosmtrek/air
start:
	air

.PHONY: start migrate docker-dev
