package user

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func GetUser(ctx context.Context, id string) (User, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "GetUser")
	defer span.Finish()

	return User{
		Id:   id,
		Name: "pete",
	}, nil
}
