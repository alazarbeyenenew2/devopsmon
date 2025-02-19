migrate-down:
	- migrate -database postgres://postgres:password@localhost:5432/devopsmon?sslmode=disable -path internal/constant/query/schemas -verbose down
migrate-up:
	- migrate -database postgres://postgres:password@localhost:5432/devopsmon?sslmode=disable -path internal/constant/query/schemas -verbose up
migrate-create:
	- migrate create -ext sql -dir internal/constant/query/schemas -tz "UTC" $(args)