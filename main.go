package main

//xray "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"

func main() {
	//close := initTracerHoneycomb("dummy-service")
	//defer close()
	initTracerStdOut()
	demo1()
}
