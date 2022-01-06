package configs

import (
	"os"
)

var PG_HOST = os.Getenv("PG_HOST")
var PG_PORT = os.Getenv("PG_PORT")
var POSTGRES_DB = os.Getenv("POSTGRES_DB")
var POSTGRES_USER = os.Getenv("POSTGRES_USER")
var POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var GRPC_PORT = os.Getenv("PROBLEMS_GRPC_PORT")
