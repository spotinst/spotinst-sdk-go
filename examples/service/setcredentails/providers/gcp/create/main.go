package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

func main() {
	sess := session.New()
	svc := account.New(sess)
	ctx := context.Background()

	_, err := svc.CloudProviderGCP().SetServiceAccount(ctx, &gcp.SetServiceAccountsInput{
		ServiceAccounts: &gcp.ServiceAccounts{
			AccountId:               spotinst.String("act-12345"),
			Type:                    spotinst.String("service_account"),
			ProjectId:               spotinst.String("demo_labs"),
			PrivateKeyId:            spotinst.String("1234567890"),
			PrivateKey:              spotinst.String("-----BEGIN PRIVATE KEY-----\n123456ygryyfderyyrfgg-----END PRIVATE KEY-----\n"),
			ClientEmail:             spotinst.String("demo_role"),
			ClientId:                spotinst.String("1234567890"),
			AuthUri:                 spotinst.String("authURI"),
			TokenUri:                spotinst.String("tokenURI"),
			AuthProviderX509CertUrl: spotinst.String("authProviderX509CertUrl"),
			ClientX509CertUrl:       spotinst.String("clientCertUrl"),
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to set credential: %v", err)
	}

}
