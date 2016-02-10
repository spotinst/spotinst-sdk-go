# Spotinst SDK GO

A Go client library for accessing the Spotinst API.

You can view Spotinst API docs here: [https://spotinst.atlassian.net/wiki/display/API](https://spotinst.atlassian.net/wiki/display/API)


## Usage

```go
import "github.com/spotinst/spotinst-sdk-go/spotinst"
```

Create a new Spotinst client, then use the exposed services to
access different parts of the Spotinst API.

### Authentication

You can use your credentials to create a new client:

```go
creds := &spotinst.Credentials{
  Email: "foo@bar.com",
  Password: "p@ssw0rd",
  ClientID: "CJZzWvP3e5vefCgt",
  ClientSecret: "EluFEJCfje78eQxP3u6X0cyH2scw6ZIP",
}
client := spotinst.NewClient(creds)
```

Or, you can use your Personal API Access Token to create a new client:

```go
creds := &spotinst.Credentials{
  Token: "aaaaaaaaaa.bbbbbbbbbbb.cccccccccccc",
}
client := spotinst.NewClient(creds)
```

## Examples

To create a new Elastigroup:

```go
group := &spotinst.AwsGroup{
  Name: "foo",
  Desciption: "bar",
  Strategy: &spotinst.AwsGroupStrategy{
    Risk: 100,
  },
  Capacity: &spotinst.AwsGroupCapacity{
    Target: 75,
    Minimum: 50,
    Maximum: 100,
  },
  Compute: &spotinst.AwsGroupCompute{
    Product: "Linux/UNIX",
    InstanceTypes: &spotinst.AwsGroupComputeInstanceType{
      OnDemand: "c4.large",
      Spot: ["c3.large", "c4.large"],
    },
    AvailabilityZones: []*spotinst.AwsGroupComputeAvailabilityZone{
      &spotinst.AwsGroupComputeAvailabilityZone{
        Name: "us-east-1a",
        SubnetID: "subnet-foo"
      },
      &spotinst.AwsGroupComputeAvailabilityZone{
        Name: "us-east-1b",
        SubnetID: "subnet-foo"
      }
    },
    LaunchSpecification: &spotinst.AwsGroupComputeLaunchSpecification{
      Monitoring: false,
      ImageID: "ami-",
      KeyPair: "pemfile",
      SecurityGroupIDs: ["default"],
      LoadBalancerNames: ["aws-elb-prod"]
    }
  }
}

group, _, err := client.AwsGroup.Create(group)

if err != nil {
    fmt.Printf("Something bad happened: %s\n", err)
    return err
}
```

## Documentation

For a comprehensive list of examples, check out the [API documentation](https://spotinst.atlassian.net/wiki/display/API).

## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).
