package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
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

	// List all Verification Providers.
	out, err := svc.ListVerificationProviders(ctx)
	if err != nil {
		log.Fatalf("spotinst: failed to list verification providers: %v", err)
	}

	// Output all Verification Providers, if any.
	if len(out.VerificationProviders) > 0 {
		for _, verificationProvider := range out.VerificationProviders {
			log.Printf("Verification Providers %q: %s",
				spotinst.StringValue(verificationProvider.Name),
				stringutil.Stringify(verificationProvider))
		}
	}
}
