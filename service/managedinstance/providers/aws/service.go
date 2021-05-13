package aws

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
	List(context.Context, *ListManagedInstancesInput) (*ListManagedInstancesOutput, error)
	Create(context.Context, *CreateManagedInstanceInput) (*CreateManagedInstanceOutput, error)
	Read(context.Context, *ReadManagedInstanceInput) (*ReadManagedInstanceOutput, error)
	Update(context.Context, *UpdateManagedInstanceInput) (*UpdateManagedInstanceOutput, error)
	Delete(context.Context, *DeleteManagedInstanceInput) (*DeleteManagedInstanceOutput, error)
	Status(context.Context, *StatusManagedInstanceInput) (*StatusManagedInstanceOutput, error)
	Costs(context.Context, *CostsManagedInstanceInput) (*CostsManagedInstanceOutput, error)
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
