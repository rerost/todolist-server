package util

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"

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
	createdAt, err := ptypes.TimestampProto(task.CreatedAt)

	if err != nil {
		// When error occured, db's time is not valid
		panic(err)
	}
	pb := &api_pb.Task{
		TaskId:    int64(task.ID),
		Title:     task.Title,
		CreatedAt: createdAt,
	}

	return pb
}

// PBTaskToTask is hogehoge
func PBTaskToTask(ctx context.Context, pbTask *api_pb.Task) *todolist.Task {
	createdAt, err := ptypes.Timestamp(pbTask.CreatedAt)
	if err != nil {
		panic(err)
	}

	return &todolist.Task{
		ID:        int(pbTask.TaskId),
		Title:     pbTask.Title,
		CreatedAt: createdAt,
	}
}

// TODO: Replace
func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
