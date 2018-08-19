package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/rerost/todolist-server/api"
)

// NewTaskServiceServer creates a new TaskServiceServer instance.
func NewTaskServiceServer() interface {
	api_pb.TaskServiceServer
	grapiserver.Server
} {
	return &taskServiceServerImpl{}
}

type taskServiceServerImpl struct {
}

func (s *taskServiceServerImpl) ListTasks(ctx context.Context, req *api_pb.ListTasksRequest) (*api_pb.ListTasksResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *taskServiceServerImpl) GetTask(ctx context.Context, req *api_pb.GetTaskRequest) (*api_pb.Task, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *taskServiceServerImpl) CreateTask(ctx context.Context, req *api_pb.CreateTaskRequest) (*api_pb.Task, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *taskServiceServerImpl) UpdateTask(ctx context.Context, req *api_pb.UpdateTaskRequest) (*api_pb.Task, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *taskServiceServerImpl) DeleteTask(ctx context.Context, req *api_pb.DeleteTaskRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
