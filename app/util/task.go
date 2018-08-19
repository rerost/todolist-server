package util

import (
	"context"
	"fmt"
	"strconv"
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

// PBTaskToTask is hogehoge
func PBTaskToTask(ctx context.Context, pbTask *api_pb.Task) *todolist.Task {
	var ID int
	if pbTask.TaskId != "" {
		_ID, err := strconv.Atoi(pbTask.TaskId)
		check(err)
		ID = _ID
	}

	var CreatedAt time.Time
	if pbTask.TaskId != "" {
		_CreatedAt, err := time.Parse(time.RFC3339, pbTask.CreatedAt)
		check(err)
		CreatedAt = _CreatedAt
	}

	return &todolist.Task{
		ID:        ID,
		Title:     pbTask.Title,
		CreatedAt: CreatedAt,
	}
}

// TODO: Replace
func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
