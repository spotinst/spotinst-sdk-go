package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

func main() {
	sess := session.New()
	svc := account.New(sess)
	ctx := context.Background()
	_, err := svc.CloudProviderAzure().SetCredentials(ctx, &azure.SetCredentialsInput{
		Credentials: &azure.Credentials{
			AccountId:      spotinst.String("accountId"),
			ClientId:       spotinst.String("clientId"),
			ClientSecret:   spotinst.String("clientSecret"),
			TenantId:       spotinst.String("tenantId"),
			SubscriptionId: spotinst.String("subscriptionId"),
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to set credential: %v", err)
	}

}
