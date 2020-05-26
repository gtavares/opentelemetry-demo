package main

import (
	"log"

	"github.com/honeycombio/opentelemetry-exporter-go/honeycomb"
	"go.opentelemetry.io/otel/api/global"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracerHoneycomb(service string) func() {

	//Init honeycomb exporter
	exporter, _ := honeycomb.NewExporter(
		honeycomb.Config{
			APIKey: "your-api-key",
		},
		honeycomb.TargetingDataset("goliath"),
		honeycomb.WithServiceName(service),
	)

	tp, err := sdktrace.NewProvider(sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithSyncer(exporter))

	if err != nil {
		log.Fatal(err)
	}

	global.SetTraceProvider(tp)

	return exporter.Close
}
