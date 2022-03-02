package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
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

	// List all extended resource definitions.
	out, err := svc.CloudProviderAWS().ListExtendedResourceDefinition(ctx, &aws.ListExtendedResourceDefinitionsInput{})
	if err != nil {
		log.Fatalf("spotinst: failed to list extended resource definitions: %v", err)
	}

	// Output all extended resource definitions, if any.
	if len(out.ExtendedResourceDefinitions) > 0 {
		for _, erd := range out.ExtendedResourceDefinitions {
			log.Printf("Extended Resource Definitions %q: %s",
				spotinst.StringValue(erd.ID),
				stringutil.Stringify(erd))
		}
	}
}
