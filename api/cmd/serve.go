package cmd

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/grpcreflect"
	"github.com/ei-sugimoto/yamicheck/api/gen/job/v1/jobv1connect"
	"github.com/ei-sugimoto/yamicheck/api/internal/adapters/handler"
	"github.com/ei-sugimoto/yamicheck/api/internal/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Serve() {
	slog.Info("Starting server...")

	JobService := usecase.NewJobService()
	JobHandler := handler.NewJobHandler(JobService)

	mux := http.NewServeMux()
	path, handler := jobv1connect.NewJobServiceHandler(JobHandler)
	mux.Handle(path, handler)

	reflector := grpcreflect.NewStaticReflector(
		"job.v1.JobService",
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	slog.Error("Server stopped", "Because:", http.ListenAndServe(":8000", h2c.NewHandler(mux, &http2.Server{})))
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	slog.Info("Shutting down")
	defer func() {
		stop()
		cancel()
	}()

}
