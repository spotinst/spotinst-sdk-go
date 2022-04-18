package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/dataintegration"
	"github.com/spotinst/spotinst-sdk-go/service/dataintegration/providers/aws"
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
	svc := dataintegration.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create a new Data Integration.
	out, err := svc.CloudProviderAWS().CreateDataIntegration(ctx, &aws.CreateDataIntegrationInput{
		DataIntegration: &aws.DataIntegration{
			Name:   spotinst.String("my-s3-integration-check"),
			Vendor: spotinst.String("s3"),
			Config: &aws.Config{
				BucketName: spotinst.String("foo"),
				SubDir:     spotinst.String("foo"),
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create data integration: %v", err)
	}

	// Output.
	if out.DataIntegration != nil {
		log.Printf("Data Integration %q: %s",
			spotinst.StringValue(out.DataIntegration.ID),
			stringutil.Stringify(out.DataIntegration))
	}
}
