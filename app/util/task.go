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
	var title string
	if task.Title.Valid {
		title = task.Title.String
	}

	var createdAt time.Time
	if task.CreatedAt.Valid {
		createdAt = task.CreatedAt.Time
	}

	pb := &api_pb.Task{
		TaskId:    string(task.ID),
		Title:     title,
		CreatedAt: createdAt.Format(time.RFC3339),
	}

	return pb
}
