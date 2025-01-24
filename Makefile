createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres hotel_manager

dropdb:
	docker exec -it postgres dropdb hotel_manager

migrateup:
	migrate -path migrations --database "postgresql://postgres:postgres@localhost:5432/hotel_manager?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations --database "postgresql://postgres:postgres@localhost:5432/hotel_manager?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb