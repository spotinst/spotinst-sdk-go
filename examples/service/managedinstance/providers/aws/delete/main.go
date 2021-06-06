package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/managedinstance"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
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
	svc := managedinstance.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Delete an existing managed instance.
	_, err := svc.CloudProviderAWS().Delete(ctx, &aws.DeleteManagedInstanceInput{
		ManagedInstanceID: spotinst.String("smi-12345"),
		DeallocationConfig: &aws.DeallocationConfig{
			ShouldDeleteImages:            spotinst.Bool(true),
			ShouldDeleteNetworkInterfaces: spotinst.Bool(true),
			ShouldDeleteSnapshots:         spotinst.Bool(true),
			ShouldDeleteVolumes:           spotinst.Bool(true),
			ShouldTerminateInstance:       spotinst.Bool(true),
		},
		AMIBackup: &aws.AMIBackup{
			ShouldDeleteImages: spotinst.Bool(true),
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to delete managed instance: %v", err)
	}
}
