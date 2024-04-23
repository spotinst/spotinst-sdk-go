package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"log"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

func main() {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as account and credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.New()

	// Create a new instance of the service's client with a Session.
	// Optional spotinst.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// service specific configuration.
	svc := oceancd.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Update Verification Provider configuration.
	out, err := svc.PatchVerificationProvider(ctx, &oceancd.PatchVerificationProviderInput{
		VerificationProvider: &oceancd.VerificationProvider{
			Name:       spotinst.String("Foo"),
			ClusterIDs: []string{"Test-Rollout"},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to patch Verification Provider: %v", err)
	}

	// Output.
	if out.VerificationProvider != nil {
		log.Printf("Verification Provider %q: %s",
			spotinst.StringValue(out.VerificationProvider.Name),
			stringutil.Stringify(out.VerificationProvider))
	}
}
