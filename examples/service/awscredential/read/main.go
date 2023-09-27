package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
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
	out, err := svc.CloudProviderAWS().ReadCredential(ctx, &aws.ReadCredentialInput{
		AccountId: spotinst.String("act-12345"),
	})

	if err != nil {
		log.Fatalf("spotinst: failed to fetch credential: %v", err)
	}
	if out != nil {
		log.Printf("credential %q: %s",
			spotinst.StringValue(out.Credential.AccountId),
			stringutil.Stringify(out.Credential.IamRole))
	}

}
