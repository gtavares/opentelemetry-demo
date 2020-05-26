package main

import (
	"context"

	"go.opentelemetry.io/otel/api/global"
)

//Simple trace
func demo2() {
	tracer := global.Tracer("")

	tracer.WithSpan(context.Background(), "demo2", func(ctx context.Context) error {
		println("Hello")
		return nil
	})
}
