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

	err := svc.UpdateNotificationCenterPolicy(ctx, &notificationcenter.NotificationCenter{
		ID:          spotinst.String("snp-12345678"),
		Description: spotinst.String("update-notification-center-description"),
	})

	if err != nil {
		log.Fatalf("UpdateNotificationCenterPolicy returned error: %v", err)
	}
}
