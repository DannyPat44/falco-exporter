module github.com/falcosecurity/falco-exporter

go 1.13

require (
	// todo(leogr): currently pointing to https://github.com/falcosecurity/client-go/pull/44
	//              update before merging
	github.com/falcosecurity/client-go v0.1.1-0.20200526171547-e0eacfb82cdc
	github.com/prometheus/client_golang v1.1.0
	github.com/spf13/pflag v1.0.5
	google.golang.org/grpc v1.28.0
)
