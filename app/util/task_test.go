package util_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	api_pb "github.com/rerost/todolist-server/api"
	"github.com/rerost/todolist-server/app/util"
	"google.golang.org/genproto/protobuf/field_mask"
)

func TestMaskTask(t *testing.T) {
	type inType struct {
		task      *api_pb.Task
		fieldMask *field_mask.FieldMask
	}
	inOutPairs := []struct {
		test string
		in   inType
		out  *api_pb.Task
	}{
		{
			test: "When FieldMask is id and title",
			in: inType{
				task: &api_pb.Task{
					TaskId: 1,
					Title:  "test",
				},
				fieldMask: &field_mask.FieldMask{Paths: []string{"task_id", "title"}},
			},
			out: &api_pb.Task{
				TaskId: 1,
				Title:  "test",
			},
		},
		{
			test: "When FieldMask is empty",
			in: inType{
				task: &api_pb.Task{
					TaskId: 1,
					Title:  "test",
				},
				fieldMask: &field_mask.FieldMask{},
			},
			out: &api_pb.Task{},
		},
		{
			test: "When FieldMask is \"test_id\"",
			in: inType{
				task: &api_pb.Task{
					TaskId: 1,
					Title:  "test",
				},
				fieldMask: &field_mask.FieldMask{Paths: []string{"task_id"}},
			},
			out: &api_pb.Task{
				TaskId: 1,
			},
		},
	}
	for _, inOutPair := range inOutPairs {
		out, err := util.MaskTask(inOutPair.in.task, inOutPair.in.fieldMask)
		if err != nil {
			t.Errorf("%s: occured error: %v", inOutPair.test, err)
		}
		if !cmp.Equal(out, inOutPair.out) {
			t.Errorf("%s: %s", inOutPair.test, cmp.Diff(out, inOutPair.out))
		}
	}
}
