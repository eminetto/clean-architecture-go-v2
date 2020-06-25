// +build dev

package config

const (
	MONGODB_HOST            = "mongodb://127.0.0.1:27017"
	MONGODB_DATABASE        = "bookmark"
	MONGODB_CONNECTION_POOL = 5
	API_PORT                = 8080
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
)
