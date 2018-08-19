package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api_pb "github.com/rerost/todolist-server/api"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *taskServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterTaskServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *taskServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterTaskServiceHandler(ctx, mux, conn)
}
