module api-server

go 1.24.3

require (
	compiler v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1 // direct
	github.com/rs/cors v1.10.1 // direct
)

replace compiler => ../compiler
