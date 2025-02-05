module _examples

go 1.21

toolchain go1.22.3

replace github.com/elastic/go-elasticsearch/v8 => ../

require github.com/elastic/go-elasticsearch/v8 v8.0.0-00010101000000-000000000000

require (
	github.com/elastic/elastic-transport-go/v8 v8.5.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
)
