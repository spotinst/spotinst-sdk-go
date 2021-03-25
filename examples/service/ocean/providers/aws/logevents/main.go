package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
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
	svc := ocean.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Get log events.
	out, err := svc.CloudProviderAWS().GetLogEvents(ctx, &aws.GetLogEventsInput{
		ClusterID: spotinst.String("o-12345"),
		FromDate:  spotinst.String("yyyy-mm-dd"),
		ToDate:    spotinst.String("yyyy-mm-dd"),
		// Severity:   spotinst.String("INFO"),    // +optional
		// ResourceID: spotinst.String("i-12345"), // +optional
		// Limit:      spotinst.Int(100),          // +optional
	})
	if err != nil {
		log.Fatalf("spotinst: failed to get log events: %v", err)
	}

	// Output log events, if any.
	if len(out.Events) > 0 {
		for _, event := range out.Events {
			fmt.Printf("%s [%s] %s\n",
				spotinst.TimeValue(event.CreatedAt).Format(time.RFC3339),
				spotinst.StringValue(event.Severity),
				spotinst.StringValue(event.Message))
		}
	}
}
