package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
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

	// Delete existing Right Sizing rules
	_, err := svc.CloudProviderAWS().DeleteRightSizingRules(ctx, &aws.DeleteRightSizingRuleInput{
		OceanId:   spotinst.String("o-1234abcd"),
		RuleNames: []string{"tf-rule-1", "tf-rule-2"},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to delete right sizing rule: %v", err)
	}
}
