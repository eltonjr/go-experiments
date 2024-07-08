# Decorator pattern

This is a simple proof-of-concept of an interface calling itself, also known as a Decorator design pattern.

To avoid bloating the real implementation of anything with logs, traces, profiling etc, one could implement the interface adding those decorators just for that, leaving the real implementation cleaner.

It's not unusual to have something like this everywhere in your codebase:

    func DoSomething(ctx context.Context) {
        var span trace.Span
        ctx, span = tracer.Start(ctx, "DoSomething") // instrument with traces
        defer span.End()

        logger.Debug("at DoSomething") // instrument with log
        defer logger.Debug("finished DoSomething")
        
        // now at last do the something it is supposed to do
    }

To avoid that, one could implement the DoSomething interface to add traces and logs, and, at the end, call the DoSomething interface again, but now with the real implementation.

    type DoSomethinger interface {
        DoSomething(ctx context.Context)
    }

    // DoSomethingDecorator has an DoSomethinger but also implements DoSomethinger
    // just to add instrumentation
    type DoSomethingDecorator struct {
        impl DoSomethinger
    }

    func (deco DoSomethingDecorator) DoSomething(ctx context.Context) {
        var span trace.Span
        ctx, span = tracer.Start(ctx, "DoSomething") // instrument with traces
        defer span.End()

        logger.Debug("at DoSomething") // instrument with log
        defer logger.Debug("finished DoSomething")

        deco.impl.DoSomething(ctx)
    }

    // RealImpl also implements DoSomethinger
    type RealImpl struct {}
    func (real RealImpl) DoSomething(ctx context.Context) {
        // just do what it is supposed to do
    }

## Running

To run the project, use

    go run main.go

This will start 2 servers, one with log (using the decorator), and one without it.

Call the endpoints at :8080 to access the server without log, and :8081 to access the server that logs everything

```
curl -L -X POST 'http://localhost:8080/items' \
-H 'Content-Type: application/json' \
--data-raw '{"id": 2}'

curl -L -X GET 'http://localhost:8080/items/2'

curl -L -X POST 'http://localhost:8081/items' \
-H 'Content-Type: application/json' \
--data-raw '{"id": 3}'

curl -L -X GET 'http://localhost:8081/items/3'
``` 