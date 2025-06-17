package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/notificationcenter"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"log"
)

func main() {
	sess := session.New()
	svc := notificationcenter.New(sess)
	ctx := context.Background()

	out, err := svc.CreateNotificationCenterPolicy(ctx, &notificationcenter.NotificationCenter{
		Name:         spotinst.String("Notification Center Policy"),
		Description:  spotinst.String("Spot Notification Center Policy Test"),
		PrivacyLevel: spotinst.String("public"),
		IsActive:     spotinst.Bool(true),
		RegisteredUsers: []*notificationcenter.RegisteredUsers{
			&notificationcenter.RegisteredUsers{
				UserEmail:         spotinst.String("testing@flexera.com"),
				SubscriptionTypes: []string{"email", "console"},
			},
		},
		Subscriptions: []*notificationcenter.Subscriptions{
			&notificationcenter.Subscriptions{
				Endpoint: spotinst.String("https://spotinst.example.com"),
				Type:     spotinst.String("email"),
			},
		},
		ComputePolicyConfig: &notificationcenter.ComputePolicyConfig{
			Events: []*notificationcenter.Events{
				&notificationcenter.Events{
					Event: spotinst.String("Beanstalk Missing Permissions"),
					Type:  spotinst.String("ERROR"),
				},
			},
			ShouldIncludeAllResources: spotinst.Bool(true),
		},
	})
	if err != nil {
		log.Fatalf("CreateNotificationCenter failed: %v", err)
	}

	if out.NotificationCenter != nil {
		log.Printf("CreateNotificationCenter returned NotificationCenter %v",
			spotinst.StringValue(out.NotificationCenter.ID))
	}
}
