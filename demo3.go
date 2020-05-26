package main

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/api/global"
)

//Inner trace
func demo3() {

	//Get the tracer
	tracer := global.Tracer("")

	//Start the span
	ctx, sp := tracer.Start(context.Background(), "demo3")
	//End the span
	defer sp.End()

	time.Sleep(2 * time.Second)
	demo3InnerFunc(ctx)
	// End the span

}

func demo3InnerFunc(ctx context.Context) {
	//Get the tracer
	tracer := global.Tracer("")

	//Start the span
	ctx, sp := tracer.Start(ctx, "demo3InnerFunc")
	//End the span
	defer sp.End()

	time.Sleep(4 * time.Second)
}
