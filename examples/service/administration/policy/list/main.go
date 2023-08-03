package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/administration"
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
	svc := administration.New(sess)

	// Create a new context.
	ctx := context.Background()

	// List all groups.
	out, err := svc.ListPolicies(ctx, &administration.ListPoliciesInput{})
	if err != nil {
		log.Fatalf("spotinst: failed to list users: %v", err)
	}

	// Output all groups, if any.
	if len(out.Policies) > 0 {
		for _, Policy := range out.Policies {
			log.Printf("Policy %q: %s",
				spotinst.StringValue(Policy.PolicyID),
				stringutil.Stringify(Policy))
		}
	}
}
