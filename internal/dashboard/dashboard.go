package dashboard

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/pete911/opentracing-examples/internal/project"
	"github.com/pete911/opentracing-examples/internal/user"
)

type Dashboard struct {
	Project project.Project `json:"name"`
	User    user.User       `json:"user"`
}

func GetDashboard(ctx context.Context, userId, projectId string) (Dashboard, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "GetDashboard")
	defer span.Finish()

	user, err := user.GetUser(ctx, userId)
	if err != nil {
		return Dashboard{}, fmt.Errorf("get user: %w", err)
	}

	project, err := project.GetProject(ctx, projectId)
	if err != nil {
		return Dashboard{}, fmt.Errorf("get project: %w", err)
	}

	return Dashboard{
		Project: project,
		User:    user,
	}, nil
}
