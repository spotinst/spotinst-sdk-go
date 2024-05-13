package oceancd

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
	ListVerificationProviders(context.Context) (*ListVerificationProvidersOutput, error)
	CreateVerificationProvider(context.Context, *CreateVerificationProviderInput) (*CreateVerificationProviderOutput, error)
	ReadVerificationProvider(context.Context, *ReadVerificationProviderInput) (*ReadVerificationProviderOutput, error)
	UpdateVerificationProvider(context.Context, *UpdateVerificationProviderInput) (*UpdateVerificationProviderOutput, error)
	PatchVerificationProvider(context.Context, *PatchVerificationProviderInput) (*PatchVerificationProviderOutput, error)
	DeleteVerificationProvider(context.Context, *DeleteVerificationProviderInput) (*DeleteVerificationProviderOutput, error)

	ListVerificationTemplates(context.Context) (*ListVerificationTemplatesOutput, error)
	CreateVerificationTemplate(context.Context, *CreateVerificationTemplateInput) (*CreateVerificationTemplateOutput, error)
	ReadVerificationTemplate(context.Context, *ReadVerificationTemplateInput) (*ReadVerificationTemplateOutput, error)
	UpdateVerificationTemplate(context.Context, *UpdateVerificationTemplateInput) (*UpdateVerificationTemplateOutput, error)
	PatchVerificationTemplate(context.Context, *PatchVerificationTemplateInput) (*PatchVerificationTemplateOutput, error)
	DeleteVerificationTemplate(context.Context, *DeleteVerificationTemplateInput) (*DeleteVerificationTemplateOutput, error)

	ListStrategies(context.Context) (*ListStrategiesOutput, error)
	CreateStrategy(context.Context, *CreateStrategyInput) (*CreateStrategyOutput, error)
	ReadStrategy(context.Context, *ReadStrategyInput) (*ReadStrategyOutput, error)
	UpdateStrategy(context.Context, *UpdateStrategyInput) (*UpdateStrategyOutput, error)
	PatchStrategy(context.Context, *PatchStrategyInput) (*PatchStrategyOutput, error)
	DeleteStrategy(context.Context, *DeleteStrategyInput) (*DeleteStrategyOutput, error)
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
