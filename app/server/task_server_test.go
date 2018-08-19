package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/rerost/todolist-server/api"
)

func Test_TaskServiceServer_ListTasks(t *testing.T) {
	svr := NewTaskServiceServer()

	ctx := context.Background()
	req := &api_pb.ListTasksRequest{}

	resp, err := svr.ListTasks(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_TaskServiceServer_GetTask(t *testing.T) {
	svr := NewTaskServiceServer()

	ctx := context.Background()
	req := &api_pb.GetTaskRequest{}

	resp, err := svr.GetTask(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_TaskServiceServer_CreateTask(t *testing.T) {
	svr := NewTaskServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateTaskRequest{}

	resp, err := svr.CreateTask(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_TaskServiceServer_UpdateTask(t *testing.T) {
	svr := NewTaskServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateTaskRequest{}

	resp, err := svr.UpdateTask(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}

func Test_TaskServiceServer_DeleteTask(t *testing.T) {
	svr := NewTaskServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteTaskRequest{}

	resp, err := svr.DeleteTask(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		at.Error("response should not nil")
	}
}
