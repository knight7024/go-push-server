package project

import (
	"context"
	"github.com/google/uuid"
	"github.com/knight7024/go-push-server/common/mysql"
	"github.com/knight7024/go-push-server/common/util"
	"github.com/knight7024/go-push-server/ent"
	predicate "github.com/knight7024/go-push-server/ent/project"
	"strings"
)

type Project struct {
	ID          int           `json:"id,omitempty" swaggerignore:"true"`
	ProjectName string        `json:"project_name,omitempty"`
	ProjectID   string        `json:"project_id,omitempty"`
	Credentials []byte        `json:"credentials,omitempty" swaggertype:"string" format:"base64"`
	ClientKey   string        `json:"client_key,omitempty" swaggerignore:"true"`
	UserID      int           `json:"user_id,omitempty" swaggerignore:"true"`
	CreatedAt   util.Datetime `json:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt   util.Datetime `json:"updated_at,omitempty" swaggerignore:"true"`
}

func CreateOne(ctx context.Context, uid int, project *Project) (*ent.Project, error) {
	return mysql.Connection.Project.Create().
		SetProjectName(project.ProjectName).
		SetProjectID(project.ProjectID).
		SetCredentials(project.Credentials).
		SetUserID(uid).
		Save(ctx)
}

func ReadOneByUserID(ctx context.Context, uid int, pid int) (*ent.Project, error) {
	return mysql.Connection.Project.Query().
		Where(predicate.And(predicate.ID(pid), predicate.UserID(uid))).
		Only(ctx)
}

func ReadOneByClientKey(ctx context.Context, clientKey string) (*ent.Project, error) {
	return mysql.Connection.Project.Query().
		Where(predicate.And(predicate.ClientKey(clientKey))).
		Only(ctx)
}

func CountByProjectID(ctx context.Context, projectID string) (int, error) {
	return mysql.Connection.Project.Query().
		Where(predicate.And(predicate.ProjectID(projectID))).
		Count(ctx)
}

func ReadAllByUserID(ctx context.Context, uid int) ([]*ent.Project, error) {
	return mysql.Connection.Project.Query().
		Where(predicate.And(predicate.UserID(uid))).
		All(ctx)
}

func UpdateOneByUserID(ctx context.Context, uid int, pid int, project *Project) error {
	update := mysql.Connection.Project.Update()
	if strings.TrimSpace(project.ProjectName) != "" {
		update.SetProjectName(project.ProjectName)
	}
	if strings.TrimSpace(project.ProjectID) != "" {
		update.SetProjectID(project.ProjectID)
	}
	if len(project.Credentials) != 0 {
		update.SetCredentials(project.Credentials)
	}
	if count, err := update.
		Where(predicate.And(predicate.ID(pid), predicate.UserID(uid))).
		Save(ctx); err != nil {
		return err
	} else if count == 0 {
		return new(ent.NotFoundError)
	}
	return nil
}

func UpdateClientKeyByUserID(ctx context.Context, uid int, pid int) (string, error) {
	clientKey := strings.ToUpper(strings.ReplaceAll(uuid.NewString(), "-", ""))
	update := mysql.Connection.Project.Update()
	update.SetClientKey(clientKey)
	if count, err := update.
		Where(predicate.And(predicate.ID(pid), predicate.UserID(uid))).
		Save(ctx); err != nil {
		return "", err
	} else if count == 0 {
		return "", new(ent.NotFoundError)
	}
	return clientKey, nil
}

func DeleteOneByUserID(ctx context.Context, pid int) error {
	return mysql.Connection.Project.DeleteOneID(pid).
		Exec(ctx)
}
