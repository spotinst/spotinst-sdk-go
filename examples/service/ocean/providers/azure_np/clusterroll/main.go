package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
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

	// Trigger cluster roll.
	out, err := svc.CloudProviderAzureNP().CreateRoll(ctx, &azure_np.CreateRollInput{
		Roll: &azure_np.RollSpec{
			ClusterID:                 spotinst.String("o-12345"),
			BatchSizePercentage:       spotinst.Int(20),
			Comment:                   spotinst.String("Comment to describe roll."),
			RespectPDB:                spotinst.Bool(true),
			BatchMinHealthyPercentage: spotinst.Int(100),
			NodeNames:                 []string{"node123"},
			VngIds:                    []string{"vng-123"},
			NodePoolNames:             []string{"nodepool12345", "nodepool67890"},
			RespectRestrictScaleDown:  spotinst.Bool(true),
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to roll cluster: %v", err)
	}

	if (out.Roll) != nil {
		log.Printf("Roll details: %q: %s",
			spotinst.StringValue(out.Roll.ID),
			stringutil.Stringify(out))
	}
}
