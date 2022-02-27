package extendedresourcedefinition

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
	List(context.Context, *ListExtendedResourceDefinitionsInput) (*ListExtendedResourceDefinitionsOutput, error)
	Create(context.Context, *CreateExtendedResourceDefinitionInput) (*CreateExtendedResourceDefinitionOutput, error)
	Read(context.Context, *ReadExtendedResourceDefinitionInput) (*ReadExtendedResourceDefinitionOutput, error)
	Update(context.Context, *UpdateExtendedResourceDefinitionInput) (*UpdateExtendedResourceDefinitionOutput, error)
	Delete(context.Context, *DeleteExtendedResourceDefinitionInput) (*DeleteExtendedResourceDefinitionOutput, error)
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
