package util

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/golang/protobuf/ptypes"

	api_pb "github.com/rerost/todolist-server/api"
	"github.com/rerost/todolist-server/app/record/todolist"
	"github.com/volatiletech/null"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	var note *wrappers.StringValue
	if task.Note.Valid {
		note = &wrappers.StringValue{Value: task.Note.String}
	}

	pb := &api_pb.Task{
		TaskId:    int64(task.ID),
		Title:     task.Title,
		CreatedAt: createdAt,
		Deadline:  deadline,
		Note:      note,
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

	var note null.String
	if pbTask.Note != nil {
		note = null.StringFrom(pbTask.Note.Value)
	}

	return &todolist.Task{
		ID:        int(pbTask.TaskId),
		Title:     pbTask.Title,
		CreatedAt: createdAt,
		Deadline:  deadline,
		Note:      note,
	}
}

// GetTasks is return api's taksks
func GetTasks(ctx context.Context, db *sql.DB, fieldMask *field_mask.FieldMask) ([]*api_pb.Task, error) {
	// TODO(@rerost) 取得するカラムを絞り込む
	dbTasks, err := todolist.Tasks().All(ctx, db)
	if err != nil {
		panic(err)
	}
	pbTasks := TasksToPB(ctx, dbTasks)

	maskedTasks, err := MaskTasks(pbTasks, fieldMask)

	return maskedTasks, err
}

func MaskTasks(tasks []*api_pb.Task, fieldMask *field_mask.FieldMask) ([]*api_pb.Task, error) {
	maskedTasks := make([]*api_pb.Task, len(tasks), len(tasks))
	for i, task := range tasks {
		var err error
		maskedTasks[i], err = MaskTask(task, fieldMask)
		if err != nil {
			return nil, err
		}
	}

	return maskedTasks, nil
}
func MaskTask(task *api_pb.Task, fieldMask *field_mask.FieldMask) (*api_pb.Task, error) {
	if fieldMask == nil {
		return task, nil
	}

	maskedTask := api_pb.Task{}
	for _, accepetedField := range fieldMask.Paths {
		// TODO(@rerost) Generate Automaticaly
		switch accepetedField {
		case "task_id":
			maskedTask.TaskId = task.TaskId
		case "title":
			maskedTask.Title = task.Title
		case "created_at":
			maskedTask.CreatedAt = task.CreatedAt
		case "deadline":
			maskedTask.Deadline = task.Deadline
		default:
			return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Received Unimplemented field: %s", accepetedField))
		}
	}
	return &maskedTask, nil
}
