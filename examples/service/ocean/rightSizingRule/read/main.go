package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/rightSizing"
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
	svc := rightSizing.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read an existing right sizing rule
	out, err := svc.ReadRightsizingRule(ctx, &rightSizing.ReadRightsizingRuleInput{
		RuleName: spotinst.String("tf-rule-6"),
		OceanId:  spotinst.String("o-9a8a856c"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create right sizing rule: %v", err)
	}

	// Output.
	if out.RightsizingRule != nil {
		log.Printf("RightSizing  Rule %q: %s",
			spotinst.StringValue(out.RightsizingRule.Name),
			stringutil.Stringify(out.RightsizingRule))
	}
}
