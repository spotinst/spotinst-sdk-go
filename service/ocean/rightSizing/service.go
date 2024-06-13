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
	serviceRightsizingRule
}

type serviceRightsizingRule interface {
	CreateRightsizingRule(context.Context, *CreateRightsizingRuleInput) (*CreateRightsizingRuleOutput, error)
	ReadRightsizingRule(context.Context, *ReadRightsizingRuleInput) (*ReadRightsizingRuleOutput, error)
	ListRightsizingRules(context.Context, *ListRightsizingRulesInput) (*ListRightsizingRulesOutput, error)
	UpdateRightsizingRule(context.Context, *UpdateRightsizingRuleInput) (*UpdateRightsizingRuleOutput, error)
	DeleteRightsizingRules(context.Context, *DeleteRightsizingRuleInput) (*DeleteRightsizingRuleOutput, error)
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
