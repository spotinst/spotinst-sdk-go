package rightSizing

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
	serviceRightSizingRule
}

type serviceRightSizingRule interface {
	CreateRightSizingRule(context.Context, *CreateRightSizingRuleInput) (*CreateRightSizingRuleOutput, error)
	ReadRightSizingRule(context.Context, *ReadRightSizingRuleInput) (*ReadRightSizingRuleOutput, error)
	ListRightSizingRules(context.Context, *ListRightSizingRulesInput) (*ListRightSizingRulesOutput, error)
	UpdateRightSizingRule(context.Context, *UpdateRightSizingRuleInput) (*UpdateRightSizingRuleOutput, error)
	DeleteRightSizingRules(context.Context, *DeleteRightSizingRuleInput) (*DeleteRightSizingRuleOutput, error)
	AttachWorkloadsToRule(context.Context, *RightSizingAttachDetachInput) (*RightSizingAttachDetachOutput, error)
	DetachWorkloadsFromRule(context.Context, *RightSizingAttachDetachInput) (*RightSizingAttachDetachOutput, error)
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

