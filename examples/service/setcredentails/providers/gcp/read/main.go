package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

func main() {
	sess := session.New()
	svc := account.New(sess)
	ctx := context.Background()

	out, err := svc.CloudProviderGCP().ReadServiceAccount(ctx, &gcp.ReadServiceAccountsInput{
		AccountId: spotinst.String("act-123456"),
	})

	if err != nil {
		log.Fatalf("spotinst: failed to read credential: %v", err)
	}

	if out != nil {
		log.Printf("serviceAccount %q: %s",
			spotinst.StringValue(out.ServiceAccounts.Type),
			stringutil.Stringify(out.ServiceAccounts.TokenUri))
	}

}
