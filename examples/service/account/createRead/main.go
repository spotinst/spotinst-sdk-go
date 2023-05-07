package main

import (
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

func main() {
	sess := session.New()
	svc := account.New(sess)
	ctx := context.Background()
	out, err := svc.CloudProviderAWS().CreateAccount(ctx, &aws.CreateAccountInput{
		Account: &aws.Account{
			Name: spotinst.String("testTerraformAcct_123"),
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
	println(spotinst.StringValue(out.Account.ExternalId))

	out_read, err := svc.CloudProviderAWS().ReadAccount(ctx, &aws.ReadAccountInput{
		AccountID: out.Account.ID,
	})
	if err != nil {
		log.Fatalf("spotinst: failed to read account: %v", err)
	}

	// Output.
	if out_read.Account != nil {
		log.Printf("Account %q: %s",
			spotinst.StringValue(out_read.Account.AccountId),
			stringutil.Stringify(out.Account))
	}

}
