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
	out, err := svc.CloudProviderAWS().ListAccounts(ctx, &aws.ListAccountsInput{
		ListAccounts: &aws.ListAccounts{
			CloudAccountId: spotinst.String("123456789"),
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to fetch accounts: %v", err)
	}

	if len(out.ListAccounts) > 0 {
		for _, account := range out.ListAccounts {
			log.Printf("Account %q: %s",
				spotinst.StringValue(account.AccountId),
				spotinst.StringValue(account.Name))
		}
	}

}
