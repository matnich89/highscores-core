postgres:
	docker run --name some-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

createdb:
	docker exec -it some-postgres createdb --username=postgres --owner=postgres highscores_core

dropdb:
	docker exec -it some-postgres dropdb --username=postgres highscores_core

migrateup:
	migrate -path internal/db/migration -database "postgresql://postgres:postgres@localhost:5432/highscores_core?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://postgres:postgres@localhost:5432/highscores_core?sslmode=disable" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
