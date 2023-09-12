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
	_, err := svc.CloudProviderAWS().DeleteAccount(ctx, &aws.DeleteAccountInput{
		spotinst.String("act-7926f067"),
	})

	if err != nil {
		log.Fatalf("spotinst: failed to delete account: %v", err)
	}

}
