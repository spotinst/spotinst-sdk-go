package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/account"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"log"
)

func main() {
	sess := session.New()
	svc := account.New(sess)
	ctx := context.Background()
	out, err := svc.CloudProviderAWS().ReadAccount(ctx, &aws.ReadAccountInput{
		AccountID: spotinst.String("act-123456"),
	})

	if err != nil {
		log.Fatalf("spotinst: faccount not found: %v", err)
	}

	if out.Account != nil {
		log.Printf("Account %q: %s",
			spotinst.StringValue(out.Account.ID),
			stringutil.Stringify(out.Account))
	}

}
