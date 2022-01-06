package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
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

	// List right sizing suggestions.
	out, err := svc.CloudProviderAzure().ListResourceSuggestions(ctx, &azure.ListResourceSuggestionsInput{
		OceanID: spotinst.String("o-12345"),
		Filter: &azure.Filter{
			Attribute: &azure.Attribute{
				Key:      spotinst.String("foo"),
				Operator: spotinst.String("foo"),
				Type:     spotinst.String("foo"),
				Value:    spotinst.String("foo"),
			},
			Namespaces: []string{"foo"},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to list right sizing suggestions: %v", err)
	}

	// Output.
	if out.Suggestions != nil {
		log.Printf("Right Sizing Suggestions: %s",
			stringutil.Stringify(out.Suggestions))
	}
}
