package server

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/volatiletech/sqlboiler/boil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	_ "github.com/bmizerany/pq"
	api_pb "github.com/rerost/todolist-server/api"
	"github.com/rerost/todolist-server/app/util"
)

var dbhost string

func init() {
	dbhost = os.Getenv("DB_HOST")
}

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
	db, err := sql.Open("postgres", fmt.Sprintf(`dbname=todolist host=%s user=liam sslmode=disable`, dbhost))
	if err != nil {
		panic(err)
	}

	tasks, err := util.GetTasks(ctx, db, req.Fields)
	return &api_pb.ListTasksResponse{Tasks: tasks}, err
}

func (s *taskServiceServerImpl) GetTask(ctx context.Context, req *api_pb.GetTaskRequest) (*api_pb.Task, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *taskServiceServerImpl) CreateTask(ctx context.Context, req *api_pb.CreateTaskRequest) (*api_pb.Task, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(`dbname=todolist host=%s user=liam sslmode=disable`, dbhost))
	if err != nil {
		panic(err)
	}

	task := util.PBTaskToTask(ctx, req.Task)

	err = task.Insert(ctx, db, boil.Infer())
	if err != nil {
		panic(err)
	}

	return util.TaskToPB(ctx, task), nil
}

func (s *taskServiceServerImpl) UpdateTask(ctx context.Context, req *api_pb.UpdateTaskRequest) (*api_pb.Task, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *taskServiceServerImpl) DeleteTask(ctx context.Context, req *api_pb.DeleteTaskRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *taskServiceServerImpl) StreamTask(req *api_pb.CreateTaskRequest, stream api_pb.TaskService_StreamTaskServer) error {
	db, err := sql.Open("postgres", fmt.Sprintf(`dbname=todolist host=%s user=liam sslmode=disable`, dbhost))
	if err != nil {
		panic(err)
	}

	tasks, err := util.GetTasks(context.Background(), db, nil)
	for _, t := range tasks {
		stream.Send(t)
	}
	return nil
}
