package administration

import (
	"context"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

// Service provides the API operation methods for making requests to endpoints
// of the Spotinst API. See this package's package overview docs for details on
// the service.
type Service interface {
	ListUsers(context.Context, *ListUsersInput) (*ListUsersOutput, error)
	CreateUser(context.Context, *User, *bool) (*CreateUserOutput, error)
	//CreateProgUser(context.Context, *ProgrammaticUser) (*CreateProgrammaticUserOutput, error)
	ReadUser(context.Context, *ReadUserInput) (*ReadUserOutput, error)
	//Update(context.Context, *UpdateUserInput) (*UpdateUserOutput, error)
	DeleteUser(context.Context, *DeleteUserInput) (*DeleteUserOutput, error)

	ListPolicies(context.Context, *ListPoliciesInput) (*ListPoliciesOutput, error)
	CreatePolicy(context.Context, *CreatePolicyInput) (*CreatePolicyOutput, error)
	ReadPolicy(context.Context, *ReadPolicyInput) (*ReadPolicyOutput, error)
	UpdatePolicy(context.Context, *UpdatePolicyInput) (*UpdatePolicyOutput, error)
	DeletePolicy(context.Context, *DeletePolicyInput) (*DeletePolicyOutput, error)

	ListUserGroups(context.Context, *ListUserGroupsInput) (*ListUserGroupsOutput, error)
	CreateUserGroup(context.Context, *UserGroup) (*CreateUserGroupOutput, error)
	ReadUserGroup(context.Context, *ReadUserGroupInput) (*ReadUserGroupOutput, error)
	UpdateUserGroup(context.Context, *UserGroup) (*UpdateUserGroupOutput, error)
	DeleteUserGroup(context.Context, *DeleteUserGroupInput) (*DeleteUserGroupOutput, error)
}

type ServiceOp struct {
	Client *client.Client
}

var _ Service = &ServiceOp{}

func New(sess *session.Session, cfgs ...*spotinst.Config) *ServiceOp {
	cfg := &spotinst.Config{}
	cfg.Merge(sess.Config)
	cfg.Merge(cfgs...)

	return &ServiceOp{
		Client: client.New(sess.Config),
	}
}
