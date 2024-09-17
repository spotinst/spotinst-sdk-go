package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/right_sizing"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"log"
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

	// Read an existing right sizing rule
	out, err := svc.RightSizing().GetWorkloadsAttached(ctx, &right_sizing.GetWorkloadsAttachedInput{
		RuleName: spotinst.String("demo-rule"),
		OceanId:  spotinst.String("o-123456"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to read attached workload: %v", err)
	}

	// Output.
	if out.Workloads != nil {
		log.Printf("Workloads  %q",
			stringutil.Stringify(out.Workloads))

	}
}
