package main

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/api/global"
)

//Inner trace
func demo4() {

	//Get the tracer
	tracer := global.Tracer("")

	//Start the span
	ctx, sp := tracer.Start(context.Background(), "demo4")
	//End the span
	defer sp.End()

	sp.SetAttribute("my-attribute-key", "my-attribute-value")

	time.Sleep(2 * time.Second)

	demo4InnerFunc(ctx)
	// End the span

}

func demo4InnerFunc(ctx context.Context) {
	//Get the tracer
	tracer := global.Tracer("")

	//Start the span
	ctx, sp := tracer.Start(ctx, "demo4InnerFunc")
	//End the span
	defer sp.End()

	time.Sleep(4 * time.Second)
	sp.AddEvent(ctx, "4 seconds already passed")
	time.Sleep(1 * time.Second)
}
