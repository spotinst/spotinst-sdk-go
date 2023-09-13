package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/account"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"log"
)

func main() {
	sess := session.New()
	svc := account.New(sess)
	ctx := context.Background()
	out, err := svc.CloudProviderAWS().ReadAWSAccountExternalId(ctx, &aws.ReadAWSAccountExternalIdInput{
		AccountID: spotinst.String("act-123456"),
	})

	if err != nil {
		log.Fatalf("spotinst: failed to fetch account: %v", err)
	}

	if out != nil {
		log.Printf("externalId: %s",
			spotinst.StringValue(out.AwsAccountExternalId.ExternalId))
	}

}
