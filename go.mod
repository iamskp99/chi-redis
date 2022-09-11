module gitlab.com/pragmaticreviews/golang-mux-api

go 1.12

replace github.com/keploy/go-sdk => ../go-sdk

require (
	github.com/go-chi/chi v4.0.3+incompatible
	github.com/go-redis/redis/v8 v8.11.5
	github.com/keploy/go-sdk v0.5.3
	github.com/mattn/go-sqlite3 v2.0.2+incompatible
)
