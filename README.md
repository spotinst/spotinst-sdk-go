# Spotinst SDK Go

A Go client library for accessing the Spotinst API.

You can view Spotinst API docs [here](https://spotinst.atlassian.net/wiki/display/API).

## Usage

```go
import "github.com/spotinst/spotinst-sdk-go/spotinst"
```

Create a new Spotinst client, then use the exposed services to
access different parts of the Spotinst API.

### Authentication

```go
client, _ := spotinst.NewClient(
    spotinst.SetToken("foo"),
)
```

## Examples

To create a new Elastigroup:

```go
logger := log.New(os.Stderr, "", 0)

clientOpts := []spotinst.ClientOptionFunc{
    spotinst.SetToken("foo"),
    spotinst.SetInfoLog(logger),
    spotinst.SetErrorLog(logger),
}
client, err := spotinst.NewClient(clientOpts...)
if err != nil {
    panic(err)
}

resp, err := client.AwsGroupService.List(nil)
if err != nil {
    panic(err)
}

if len(resp.Groups) > 0 {
    for _, g := range resp.Groups {
        b, _ := json.MarshalIndent(g, "", "  ")
        log.Infof(context.TODO(), string(b))
    }
}
```

## Documentation

For a comprehensive list of examples, check out the [API documentation](https://spotinst.atlassian.net/wiki/display/API).

## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).
