package azure

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
	Create(context.Context, *CreateStatefulNodeInput) (*CreateStatefulNodeOutput, error)
	Read(context.Context, *ReadStatefulNodeInput) (*ReadStatefulNodeOutput, error)
	Update(context.Context, *UpdateStatefulNodeInput) (*UpdateStatefulNodeOutput, error)
	Delete(context.Context, *DeleteStatefulNodeInput) (*DeleteStatefulNodeOutput, error)
	List(context.Context, *ListStatefulNodesInput) (*ListStatefulNodesOutput, error)
	UpdateState(context.Context, *UpdateStatefulNodeStateInput) (*UpdateStatefulNodeStateOutput, error)
	DetachDataDisk(context.Context, *DetachStatefulNodeDataDiskInput) (*DetachStatefulNodeDataDiskOutput, error)
	AttachDataDisk(context.Context, *AttachStatefulNodeDataDiskInput) (*AttachStatefulNodeDataDiskOutput, error)
	ImportVM(context.Context, *ImportVMStatefulNodeInput) (*ImportVMStatefulNodeOutput, error)
	GetState(context.Context, *GetStatefulNodeStateInput) (*GetStatefulNodeStateOutput, error)
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
