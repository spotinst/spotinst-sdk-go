package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"log"

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
	svc := oceancd.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Delete an existing cluster.
	_, err := svc.DeleteVerificationProvider(ctx, &oceancd.DeleteVerificationProviderInput{
		Name: spotinst.String("Foo"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to delete Verification Provider: %v", err)
	}
}
