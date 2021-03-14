package tasks

import (
	"context"
)

type ServiceImpl struct{}

var _ Service = (*ServiceImpl)(nil)

var tasks = []Task{
	{ID: 1, Content: "first", Category: "TODO"},
	{ID: 2, Content: "second", Category: "Done"},
	{ID: 3, Content: "third", Category: "TODO"},
	{ID: 4, Content: "fourth", Category: "Blocked"},
}

func (*ServiceImpl) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var ret []int
	for _, t := range tasks {
		if t.Category == req.Category {
			ret = append(ret, t.ID)
		}
	}
	return &ListResponse{IDs: ret}, nil
}

func (*ServiceImpl) BatchGet(ctx context.Context, req BatchGetRequest) (*BatchGetResponse, error) {
	var ret []Task

	for _, t := range req.IDs {
		for _, tt := range tasks {
			if tt.ID == t {
				ret = append(ret, tt)
				break
			}
		}
	}
	return &BatchGetResponse{Tasks: ret}, nil
}
