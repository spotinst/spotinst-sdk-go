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
	serviceAccount
	serviceCredential
	serviceAwsAccountExternalId
}

type serviceAccount interface {
	CreateAccount(context.Context, *CreateAccountInput) (*CreateAccountOutput, error)
	DeleteAccount(context.Context, *DeleteAccountInput) (*DeleteAccountOutput, error)
	ReadAccount(context.Context, *ReadAccountInput) (*ReadAccountOutput, error)
}
type serviceCredential interface {
	SetCredential(context.Context, *SetCredentialInput) (*SetCredentialOutput, error)
	ReadCredential(context.Context, *ReadCredentialInput) (*ReadCredentialOutput, error)
}
type serviceAwsAccountExternalId interface {
	CreateAWSAccountExternalId(context.Context, *CreateAWSAccountExternalIdInput) (*CreateAWSAccountExternalIdOutput, error)
	ReadAWSAccountExternalId(context.Context, *ReadAWSAccountExternalIdInput) (*ReadAWSAccountExternalIdOutput, error)
}
type ServiceOp struct {
	Client *client.Client
}

func New(sess *session.Session, cfgs ...*spotinst.Config) *ServiceOp {
	cfg := &spotinst.Config{}
	cfg.Merge(sess.Config)
	cfg.Merge(cfgs...)

	return &ServiceOp{
		Client: client.New(sess.Config),
	}
}
