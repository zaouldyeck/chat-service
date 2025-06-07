package main

import (
	"context"
	"os"
	"runtime"
)

func main() {
	var log *logger.Logger

	traceIDFn := func(ctx context.Context) string {
		return "" // TODO: NEED TRACE ID
	}

	log = logger.New(os.Stdout, logger.LevelInfo, "CAP", traceIDFn)

	// ---------------------------------------------------------------------

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {

	// ----------------------------------------------------------------------
	// GOMAXPROCS

	// How many OS threads?
	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	return nil
}
