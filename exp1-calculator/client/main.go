package main

import (
	"context"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/lmittmann/tint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"exp/common"
)

func Logger(w io.Writer, levelAsString string) *slog.Logger {
	var level slog.Level

	switch strings.ToLower(levelAsString) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "Error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	logger := slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      level,
			TimeFormat: time.TimeOnly,
		}),
	)

	return logger
}

func main() {
	ctx := context.Background()
	log := Logger(os.Stderr, os.Getenv("LOG_LEVEL"))

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("could not connect to server", tint.Err(err))
		os.Exit(1)
	}
	defer conn.Close()

	client := common.NewCalculatorClient(conn)
	// Add
	req := &common.AddRequest{N1: 5, N2: 3}
	response, err := client.Add(ctx, req)
	if err != nil {
		log.Error("could not add", tint.Err(err))
	}
	log.Info("Add", "n1", req.N1, "n2", req.N2, "result", response.N1)

	// Subtract
	reqSub := &common.SubtractRequest{N1: 5, N2: 3}
	responseSub, err := client.Subtract(ctx, reqSub)
	if err != nil {
		log.Error("could not subtract", tint.Err(err))
	}
	log.Info("Subtract", "n1", reqSub.N1, "n2", reqSub.N2, "result", responseSub.N1)

	// Multiply
	reqMul := &common.MultiplyRequest{N1: 5, N2: 3}
	responseMul, err := client.Multiply(ctx, reqMul)
	if err != nil {
		log.Error("could not multiply", tint.Err(err))
	}
	log.Info("Multiply", "n1", reqMul.N1, "n2", reqMul.N2, "result", responseMul.N1)

	// Divide
	reqDiv := &common.DivideRequest{N1: 5, N2: 3}
	responseDiv, err := client.Divide(ctx, reqDiv)
	if err != nil {
		log.Error("could not divide", tint.Err(err))
	}
	log.Info("Divide", "n1", reqDiv.N1, "n2", reqDiv.N2, "result", responseDiv.N1)

	// Divide by zero
	reqDivByZero := &common.DivideRequest{N1: 5, N2: 0}
	_, err = client.Divide(ctx, reqDivByZero)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			log.Warn("Divide", "n1", reqDivByZero.N1, "n2", reqDivByZero.N2, "message", s.Message(), "code", s.Code())
		}
		log.Error("could not divide", tint.Err(err))
	}
}
