# Spotinst SDK Go

A Go client library for accessing the Spotinst API.

You can view Spotinst API docs [here](http://help.spotinst.com/api/).

## Usage

```go
import "github.com/spotinst/spotinst-sdk-go/spotinst"
```

Create a new Spotinst client, then use the exposed services to
access different parts of the Spotinst API.

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
any credentials are available via the environment variables. If there are
none `ChainProvider` will check the next `Provider` in the list, `FileProvider`
in this case. If `FileCredentialsProvider` does not return any credentials
`ChainProvider` will return the error `ErrNoValidProvidersFoundInChain`.

```go
creds := credentials.NewChainCredentials(
    new(credentials.FileProvider),
    new(credentials.EnvProvider),
)
clientOpts := []spotinst.ClientOption{
    spotinst.SetCredentials(creds),
}
client := spotinst.NewClient(clientOpts...)
```

## Examples

```go
logger := log.New(os.Stderr, "[spotinst] ", 0)

clientOpts := []spotinst.ClientOptionFunc{
    spotinst.SetTraceLog(logger),
}
client, err := spotinst.NewClient(clientOpts...)
if err != nil {
    // do something with err
}

client := spotinst.NewClient()
providerAWS := client.GroupService.CloudProviderAWS()

input := &spotinst.ReadAWSGroupInput{
    GroupID: spotinst.String("sig-foo"),
}
output, err := providerAWS.Read(context.TODO(), input)
if err != nil {
    // do something with err
}

// do something with output.Group
```

## Documentation

For a comprehensive list of examples, check out the [API documentation](http://help.spotinst.com/api/).

## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).
