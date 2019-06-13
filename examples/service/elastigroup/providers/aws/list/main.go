package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
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
	svc := elastigroup.New(sess)

	// Create a new context.
	ctx := context.Background()

	// List all groups.
	out, err := svc.CloudProviderAWS().List(ctx, &aws.ListGroupsInput{})
	if err != nil {
		log.Fatalf("spotinst: failed to list groups: %v", err)
	}

	// Output all groups, if any.
	if len(out.Groups) > 0 {
		for _, group := range out.Groups {
			log.Printf("Group %q: %s",
				spotinst.StringValue(group.ID),
				stringutil.Stringify(group))
		}
	}
}
