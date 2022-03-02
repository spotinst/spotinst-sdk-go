package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"log"
)

func main() {

	provider := credentials.NewEnvCredentials()

	_, err := provider.Get()
	if err != nil {
		panic(err)
	}
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

	// Create a new extended resource definition.
	out, err := svc.CloudProviderAWS().CreateExtendedResourceDefinition(ctx, &aws.CreateExtendedResourceDefinitionInput{
		ExtendedResourceDefinition: &aws.ExtendedResourceDefinition{
			Name: spotinst.String("example.com/foo"),
			Mapping: map[string]interface{}{
				"c3.large":  "4Ki",
				"c3.xlarge": "4Ki"},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create extended resource definition: %v", err)
	}

	// Output.
	if out.ExtendedResourceDefinition != nil {
		log.Printf("Extended Resource Definition %q: %s",
			spotinst.StringValue(out.ExtendedResourceDefinition.ID),
			stringutil.Stringify(out.ExtendedResourceDefinition))
	}
}
