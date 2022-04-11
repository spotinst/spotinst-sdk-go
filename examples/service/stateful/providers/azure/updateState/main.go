package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"log"
	"os"
)

func main() {

	os.Setenv(credentials.EnvCredentialsVarToken, "5ca74c5c6fe2cd5a3827eb63f8c342c873feaac8c1b3b9d43ea046651cd6f177")
	os.Setenv(credentials.EnvCredentialsVarAccount, "act-97b049d6")

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
	_, err := svc.UpdateState(ctx, &azure.UpdateStatefulNodeStateInput{
		ID:                spotinst.String("sig-01245678"), //TODO - change
		StatefulNodeState: spotinst.String("pause"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to update stateful node state: %v", err)
	}
}
