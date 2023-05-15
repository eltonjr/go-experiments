# Interface adapter

This is a simple proof-of-concept of an interface calling itself.

To avoid bloating the real implementation of anything with logs, traces, profiling etc, one could implement the interface adding those adapters just for that, leaving the real implementation cleaner.

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

    // DoSomethingAdapter has an DoSomethinger but also implements DoSomethinger
    // just to add instrumentation
    type DoSomethingAdapter struct {
        impl DoSomethinger
    }

    func (adp DoSomethingAdapter) DoSomething(ctx context.Context) {
        var span trace.Span
        ctx, span = tracer.Start(ctx, "DoSomething") // instrument with traces
        defer span.End()

        logger.Debug("at DoSomething") // instrument with log
        defer logger.Debug("finished DoSomething")

        adp.impl.DoSomething(ctx)
    }

    // RealImpl also implements DoSomethinger
    type RealImpl struct {}
    func (real RealImpl) DoSomething(ctx context.Context) {
        // just do what it is supposed to do
    }
