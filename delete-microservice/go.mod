module delete

go 1.21.0

require (
	example.com/database v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
)

replace example.com/database => ../database
