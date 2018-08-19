package util

import (
	"context"
	"time"

	api_pb "github.com/rerost/todolist-server/api"
	"github.com/rerost/todolist-server/app/record/todolist"
)

//hogehoge
func TasksToPB(ctx context.Context, tasks []*todolist.Task) []*api_pb.Task {
	pbs := make([]*api_pb.Task, len(tasks), len(tasks))

	for i, m := range tasks {
		pbs[i] = TaskToPB(ctx, m)
	}

	return pbs
}

//hogehoge
func TaskToPB(ctx context.Context, task *todolist.Task) *api_pb.Task {
	pb := &api_pb.Task{
		TaskId:    string(task.ID),
		Title:     task.Title,
		CreatedAt: task.CreatedAt.Format(time.RFC3339),
	}

	return pb
}
