
run:
	@go run cmd/main.go

postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=ROOT -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql -U ROOT
createdb:
	docker exec -it postgres15 createdb --username=ROOT --owner=ROOT online-store
dropdb:
	docker exec -it postgres15 dropdb
migrateup:
	migrate -path db/migrations -database "postgresql://ROOT:password@localhost:5433/online-store?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://ROOT:password@localhost:5433/online-store?sslmode=disable" -verbose down


.PHONY : build run postgresinit postgres createdb dropdb migrateup migratedown