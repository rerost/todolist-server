package util

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/golang/protobuf/ptypes"

	api_pb "github.com/rerost/todolist-server/api"
	"github.com/rerost/todolist-server/app/record/todolist"
	"github.com/volatiletech/null"
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
	createdAt, err := ptypes.TimestampProto(task.CreatedAt)

	if err != nil {
		// When error occured, db's time is not valid
		panic(err)
	}

	var deadline *timestamp.Timestamp
	if task.Deadline.Valid {
		deadline, err = ptypes.TimestampProto(task.Deadline.Time)

		if err != nil {
			panic(err)
		}
	}

	pb := &api_pb.Task{
		TaskId:    int64(task.ID),
		Title:     task.Title,
		CreatedAt: createdAt,
		Deadline:  deadline,
	}

	return pb
}

// PBTaskToTask is hogehoge
func PBTaskToTask(ctx context.Context, pbTask *api_pb.Task) *todolist.Task {
	if pbTask == nil {
		return nil
	}

	var createdAt time.Time
	if pbTask.CreatedAt != nil {
		t, err := ptypes.Timestamp(pbTask.CreatedAt)
		if err != nil {
			panic(err)
		}
		createdAt = t
	}

	var deadline null.Time
	if pbTask.Deadline != nil {
		t, err := ptypes.Timestamp(pbTask.Deadline)
		if err != nil {
			panic(err)
		}

		deadline = null.TimeFrom(t)
	}

	return &todolist.Task{
		ID:        int(pbTask.TaskId),
		Title:     pbTask.Title,
		CreatedAt: createdAt,
		Deadline:  deadline,
	}
}
