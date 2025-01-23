package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/azure"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

func main() {
	sess := session.New()
	svc := account.New(sess)
	ctx := context.Background()
	_, err := svc.CloudProviderAzure().ReadCredentials(ctx, &azure.ReadCredentialsInput{
		AccountId: spotinst.String("act-123456"),
	})

	if err != nil {
		log.Fatalf("spotinst: failed to fetch credential: %v", err)
	}

}
