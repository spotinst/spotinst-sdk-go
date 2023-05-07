package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

func main() {
	sess := session.New()
	svc := account.New(sess)
	ctx := context.Background()

	out_extId, err := svc.CloudProviderAWS().CreateAwsAccountExternalId(ctx, "act-12345")
	if err != nil {
		log.Fatal("spotinst: external Id creation failed")
	}
	if out_extId != nil {
		log.Println(spotinst.StringValue(out_extId.AWSAccountExternalId.ExternalId))
	}

}
