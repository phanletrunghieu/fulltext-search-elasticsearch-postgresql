.PHONY: dep local-db build run

dep:
	dep ensure -v

local-db:
	@docker-compose -p fulltext-search down
	@docker-compose -p fulltext-search up -d
	@echo "Waiting for database connection..."
	@while ! docker exec fulltext-search_postgres_1 pg_isready -h localhost -p 5432 > /dev/null; do \
			sleep 1; \
	done
	@echo "Migrate db..."
	@docker cp db.sql fulltext-search_postgres_1:/db.sql
	@docker exec -u postgres fulltext-search_postgres_1 psql blog postgres -f /db.sql

build:
	go build -o bin/main cmd/main.go

run: build
	./bin/main