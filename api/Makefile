DB_URL=postgresql://postgres:1412@localhost:5432/robbo?sslmode=disable

build:
	docker build . -t api

run:
	docker compose up 
	make migrateup

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up


migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down


