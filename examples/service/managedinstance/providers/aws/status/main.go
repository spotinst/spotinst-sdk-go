package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/managedinstance"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
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
	svc := managedinstance.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Get the group status.
	out, err := svc.CloudProviderAWS().Status(ctx, &aws.StatusManagedInstanceInput{
		ManagedInstanceID: spotinst.String("smi-89ba0299"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to get group status: %v", err)
	}

	// Output.
	if out.ID != nil {
		log.Printf("Managed Instance Id %q: and IPV6Address %s",
			spotinst.StringValue(out.ID),
			stringutil.Stringify(out.IPv6Address))
	}
}
