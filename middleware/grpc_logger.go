package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// GRPCLogger is the middleware for logging gRPC requests
func GRPCLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	// Log the incoming request
	log.Printf("gRPC call - Method: %s", info.FullMethod)

	// Call the handler
	resp, err := handler(ctx, req)

	// Log completion with status and elapsed time
	st, _ := status.FromError(err)
	log.Printf("Completed gRPC call - Method: %s, Status: %s, Elapsed time: %v", info.FullMethod, st.Code(), time.Since(start))

	return resp, err
}
