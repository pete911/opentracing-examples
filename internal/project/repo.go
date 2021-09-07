package project

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func GetProject(ctx context.Context, id string) (Project, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "GetProject")
	defer span.Finish()

	return Project{
		Id:   id,
		Name: "pete",
	}, nil
}
