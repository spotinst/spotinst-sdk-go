package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

func main() {
	sess := session.New()
	svc := account.New(sess)
	ctx := context.Background()
	err := svc.CloudProviderAWS().SetCredential(ctx, &aws.SetCredentialInput{
		&aws.Credentials{
			IamRole:   spotinst.String("arn:aws:iam::123456789:role/SpotRole-act-123456-1234567890"),
			AccountId: spotinst.String("act-12345"),
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to fetch accounts: %v", err)
	}

}
