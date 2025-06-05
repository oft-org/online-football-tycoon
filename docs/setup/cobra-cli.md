go run cmd/main.go --help


run app with CLI COBRA:
go run cmd/main.go server


run migrations with CLI COBRA:
go run cmd/main.go migrations


fix dirty database:
migrate -path ./migrations -database "postgres://postgres:mysecretpassword@localhost:5432/online_football_tycoon?sslmode=disable" force 20250314160635

drop database:
task migration-down  