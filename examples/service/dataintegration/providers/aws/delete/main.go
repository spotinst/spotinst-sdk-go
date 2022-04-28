package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/dataintegration"
	"github.com/spotinst/spotinst-sdk-go/service/dataintegration/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
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
	svc := dataintegration.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Delete an existing data integration.
	_, err := svc.CloudProviderAWS().DeleteDataIntegration(ctx, &aws.DeleteDataIntegrationInput{
		DataIntegrationId: spotinst.String("di-12345"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to delete data integration: %v", err)
	}
}
