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
	out, err := svc.CloudProviderAWS().CreateAccount(ctx, &aws.CreateAccountInput{
		Account: &aws.Account{
			Name: spotinst.String("testAcct_123"),
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to create account: %v", err)
	}

	// Output.
	if out.Account != nil {
		log.Printf("Account %q: %s",
			spotinst.StringValue(out.Account.ID),
			stringutil.Stringify(out.Account))
	}

}
