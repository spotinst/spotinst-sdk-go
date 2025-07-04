package notificationcenter

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

type Service interface {
	ListNotificationCenterPolicy(context.Context, *ListNotificationCenterPolicyInput) (*ListNotificationCenterPolicyOutput, error)
	CreateNotificationCenterPolicy(context.Context, *NotificationCenter) (*CreateNotificationCenterPolicyOutput, error)
	ReadNotificationCenterPolicy(context.Context, *ReadNotificationCenterPolicyInput) (*ReadNotificationCenterPolicyOutput, error)
	UpdateNotificationCenterPolicy(context.Context, *NotificationCenter) error
	DeleteNotificationCenterPolicy(context.Context, *DeleteNotificationCenterPolicyInput) (*DeleteNotificationCenterPolicyOutput, error)
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
