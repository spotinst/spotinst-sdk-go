package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/organization"
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
	svc := organization.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create a new group.
	err := svc.UpdateUserMappingOfUserGroup(ctx, &organization.UpdateUserMappingOfUserGroupInput{
		UserGroupId: spotinst.String("ugr-abcd1234"),
		UserIds: []string{
			"pu-abcd1234",
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to update user group: %v", err)
	}

}
