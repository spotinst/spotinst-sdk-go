package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

func main() {
	sess := session.New()
	svc := account.New(sess)
	ctx := context.Background()
	_, err := svc.CloudProviderAWS().Credentials(ctx, &aws.SetCredentialsInput{
		Credentials: &aws.Credentials{
			AccountId: spotinst.String("act-12345"),
			IamRole:   spotinst.String("arn:aws:iam::12345:role/test-role"),
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to set credential: %v", err)
	}

}
