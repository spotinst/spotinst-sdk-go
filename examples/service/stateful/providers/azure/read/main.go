package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
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
	svc := azure.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read stateful node configuration.
	out, err := svc.Read(ctx, &azure.ReadStatefulNodeInput{
		ID: spotinst.String("ssn-12345678"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to read stateful node: %v", err)
	}

	// Output.
	if out.StatefulNode != nil {
		log.Printf("StatefulNode %q: %s",
			spotinst.StringValue(out.StatefulNode.ID),
			stringutil.Stringify(out.StatefulNode))
	}
}
