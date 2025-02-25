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

	// Get the group status.
	out, err := svc.CloudProviderAWS().Create(ctx, &aws.CreateManagedInstanceInput{
		ManagedInstance: &aws.ManagedInstance{
			Name:        spotinst.String("Test-MI"),
			Region:      spotinst.String("us-east-1"),
			Description: spotinst.String("Managed Instance"),
			Compute: &aws.Compute{
				Product:   spotinst.String("Linux/Unix"),
				VpcID:     spotinst.String("vpc-id"),
				SubnetIDs: []string{"abc1234"},
				LaunchSpecification: &aws.LaunchSpecification{
					ImageID: spotinst.String("image-id"),
					InstanceTypes: &aws.InstanceTypes{
						Types:          []string{"t3.small", "t3.medium", "t3.large", "t2.small"},
						PreferredTypes: []string{"t3.medium", "t2.small"},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to get group status: %v", err)
	}

	// Output.
	if out.ManagedInstance.ID != nil {
		log.Printf("Managed Instance Id %q:",
			spotinst.StringValue(out.ManagedInstance.ID))
	}
}
