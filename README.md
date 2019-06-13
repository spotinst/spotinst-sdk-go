# Spotinst SDK Go

The official Spotinst SDK for the Go programming language.

You can view Spotinst API docs [here](http://help.spotinst.com/api/).

## Installing

If you are using Go 1.5 with the `GO15VENDOREXPERIMENT=1` vendoring flag, or 1.6 and higher you can use the following command to retrieve the SDK. The SDK's non-testing dependencies will be included and are vendored in the `vendor` folder.

    go get -u github.com/spotinst/spotinst-sdk-go

Otherwise if your Go environment does not have vendoring support enabled, or you do not want to include the vendored SDK's dependencies you can use the following command to retrieve the SDK and its non-testing dependencies using `go get`.

    go get -u github.com/spotinst/spotinst-sdk-go/spotinst/...
    go get -u github.com/spotinst/spotinst-sdk-go/service/...

If you're looking to retrieve just the SDK without any dependencies use the following command.

    go get -d github.com/spotinst/spotinst-sdk-go/

These two processes will still include the `vendor` folder and it should be deleted if its not going to be used by your environment.

    rm -rf $GOPATH/src/github.com/spotinst/spotinst-sdk-go/vendor

### Authentication

Set a `ChainProvider` that will search for a provider which returns credentials.

The `ChainProvider` provides a way of chaining multiple providers together
which will pick the first available using priority order of the Providers
in the list. If none of the Providers retrieve valid credentials, `ChainProvider`'s
`Retrieve()` will return the error `ErrNoValidProvidersFoundInChain`. If a Provider
is found which returns valid credentials `ChainProvider` will cache that Provider
for all calls until `Retrieve` is called again.

Example of `ChainProvider` to be used with an `EnvCredentialsProvider` and
`FileCredentialsProvider`. In this example `EnvProvider` will first check if
any credentials are available via the SPOTINST_TOKEN and SPOTINST_ACCOUNT environment variables. If there are
none `ChainProvider` will check the next `Provider` in the list, `FileProvider`
in this case. If `FileCredentialsProvider` does not return any credentials
`ChainProvider` will return the error `ErrNoValidProvidersFoundInChain`.

```go
// Initial credentials loaded from SDK's default credential chain. Such as
// the environment, shared credentials (~/.spotinst/credentials), etc.
sess := session.New()

// Create the chain credentials.
creds := credentials.NewChainCredentials(
    new(credentials.FileProvider),
    new(credentials.EnvProvider),
)

// Create service client value configured for credentials
// from the chain.
svc := elastigroup.New(sess, &spotinst.Config{Credentials: creds})
```

## Complete SDK Example

```go
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

	// Read group configuration.
	out, err := svc.CloudProviderAWS().Read(ctx, &aws.ReadGroupInput{
		GroupID: spotinst.String("sig-12345"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to read group: %v", err)
	}

	// Output.
	if out.Group != nil {
		log.Printf("Group %q: %s",
			spotinst.StringValue(out.Group.ID),
			stringutil.Stringify(out.Group))
	}
}
```

## Documentation

For a comprehensive documentation, check out the [API documentation](http://api.spotinst.com/).

## Examples

For a list of examples, check out the [examples](https://github.com/spotinst/spotinst-sdk-go/tree/master/examples) directory.

## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).
