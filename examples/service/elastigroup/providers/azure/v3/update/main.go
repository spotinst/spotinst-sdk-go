package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
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
	svc := elastigroup.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Update group configuration.
	out, err := svc.CloudProviderAzureV3().Update(ctx, &azurev3.UpdateGroupInput{
		Group: &azurev3.Group{
			ID: spotinst.String("sig-01245678"),
			Capacity: &azurev3.Capacity{
				Target: spotinst.Int(2),
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to update group: %v", err)
	}

	// Output.
	if out.Group != nil {
		log.Printf("Group %q: %s",
			spotinst.StringValue(out.Group.ID),
			stringutil.Stringify(out.Group))
	}
}
