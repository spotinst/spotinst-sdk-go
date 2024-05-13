package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
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
	svc := elastigroup.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read roll status.
	out, err := svc.CloudProviderAWS().RollStatus(ctx, &aws.RollStatusInput{
		GroupID: spotinst.String("sig-123456"),
		RollID:  spotinst.String("sbgd-123456"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to read roll status: %v", err)
	}

	//Output.
	if out != nil {
		log.Printf("Roll %q: %s",
			stringutil.Stringify(out.CurrentBatch),
			stringutil.Stringify(out.NumberOfBatches))

	}
}
