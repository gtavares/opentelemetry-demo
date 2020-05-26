package main

import (
	"context"

	"go.opentelemetry.io/otel/api/global"
)

//Simple trace
func demo1() {

	//Get the tracer
	tracer := global.Tracer("")

	//Start the span
	_, sp := tracer.Start(context.Background(), "myFunc")
	defer sp.End()

	println("Hello")

	// End the span
}
