package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
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
	svc := ocean.New(sess)

	// Create a new context.
	ctx := context.Background()

	// List all right sizing rules in an ocean cluster
	out, err := svc.CloudProviderAWS().ListRightSizingRules(ctx, &ocean.ListRightSizingRulesInput{
		OceanId: spotinst.String("o-1234abcd"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create right sizing rule: %v", err)
	}

	// Output all rules, if any.
	if len(out.RightSizingRules) > 0 {
		for _, rule := range out.RightSizingRules {
			log.Printf("RightSizing Rule %q: %s",
				spotinst.StringValue(rule.Name),
				stringutil.Stringify(rule))
		}
	}
}
