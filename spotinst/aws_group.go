package spotinst

import (
	"fmt"
	"net/http"
)

// AwsGroupService handles communication with the AwsGroup related methods of
// the Spotinst API.
type AwsGroupService struct {
	client *Client
}

type AwsGroup struct {
	Id          string            `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Capacity    *AwsGroupCapacity `json:"capacity,omitempty"`
	Compute     *AwsGroupCompute  `json:"compute,omitempty"`
	Strategy    *AwsGroupStrategy `json:"strategy,omitempty"`
	Scaling     *AwsGroupScaling  `json:"scaling,omitempty"`
}

type AwsGroupResponse struct {
	Response struct {
		Errors []Error     `json:"errors"`
		Items  []*AwsGroup `json:"items"`
	} `json:"response"`
}

type AwsGroupScaling struct {
	Up   []*AwsGroupScalingPolicy `json:"up,omitempty"`
	Down []*AwsGroupScalingPolicy `json:"down,omitempty"`
}

type AwsGroupScalingPolicy struct {
	PolicyName        string                            `json:"policyName,omitempty"`
	MetricName        string                            `json:"metricName,omitempty"`
	Statistic         string                            `json:"statistic,omitempty"`
	Unit              string                            `json:"unit,omitempty"`
	Threshold         float64                           `json:"threshold,omitempty"`
	Adjustment        int                               `json:"adjustment,omitempty"`
	Namespace         string                            `json:"namespace,omitempty"`
	EvaluationPeriods int                               `json:"evaluationPeriods"`
	Period            int                               `json:"period"`
	Cooldown          int                               `json:"cooldown"`
	Dimensions        []*AwsGroupScalingPolicyDimension `json:"dimensions,omitempty"`
}

type AwsGroupScalingPolicyDimension struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type AwsGroupStrategy struct {
	AvailabilityVsCost string  `json:"availabilityVsCost,omitempty"`
	Risk               float64 `json:"risk"`
}

type AwsGroupCapacity struct {
	Minimum int `json:"minimum"`
	Maximum int `json:"maximum"`
	Target  int `json:"target"`
}

type AwsGroupCompute struct {
	Product             string                              `json:"product,omitempty"`
	InstanceTypes       *AwsGroupComputeInstanceType        `json:"instanceTypes,omitempty"`
	LaunchSpecification *AwsGroupComputeLaunchSpecification `json:"launchSpecification,omitempty"`
	AvailabilityZones   []*AwsGroupComputeAvailabilityZone  `json:"availabilityZones,omitempty"`
}

type AwsGroupComputeInstanceType struct {
	OnDemand string   `json:"ondemand,omitempty"`
	Spot     []string `json:"spot,omitempty"`
}

type AwsGroupComputeAvailabilityZone struct {
	Name     string `json:"name,omitempty"`
	SubnetId string `json:"subnetId,omitempty"`
}

type AwsGroupComputeLaunchSpecification struct {
	SecurityAwsGroupIds []string `json:"securityGroupIds,omitempty"`
	ImageId             string   `json:"imageId,omitempty"`
	KeyPair             string   `json:"keyPair,omitempty"`
	Monitoring          bool     `json:"monitoring"`
}

type groupWrapper struct {
	Group AwsGroup `json:"group"`
}

// Get an existing group configuration.
func (s *AwsGroupService) Get(args ...string) ([]*AwsGroup, *http.Response, error) {
	var gid string
	if len(args) > 0 {
		gid = args[0]
	}

	path := fmt.Sprintf("aws/ec2/group/%s", gid)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var retval AwsGroupResponse
	resp, err := s.client.Do(req, &retval)
	if err != nil {
		return nil, resp, err
	}

	return retval.Response.Items, resp, err
}

// Create a new group.
func (s *AwsGroupService) Create(group *AwsGroup) ([]*AwsGroup, *http.Response, error) {
	path := "aws/ec2/group"

	req, err := s.client.NewRequest("POST", path, groupWrapper{Group: *group})
	if err != nil {
		return nil, nil, err
	}

	var retval AwsGroupResponse
	resp, err := s.client.Do(req, &retval)
	if err != nil {
		return nil, resp, err
	}

	return retval.Response.Items, resp, nil
}

// Update an existing group.
func (s *AwsGroupService) Update(group *AwsGroup) ([]*AwsGroup, *http.Response, error) {
	gid := (*group).Id
	(*group).Id = ""
	path := fmt.Sprintf("aws/ec2/group/%s", gid)

	req, err := s.client.NewRequest("PUT", path, groupWrapper{Group: *group})
	if err != nil {
		return nil, nil, err
	}

	var retval AwsGroupResponse
	resp, err := s.client.Do(req, &retval)
	if err != nil {
		return nil, resp, err
	}

	return retval.Response.Items, resp, nil
}

// Delete an existing group.
func (s *AwsGroupService) Delete(group *AwsGroup) (*http.Response, error) {
	gid := (*group).Id
	(*group).Id = ""
	path := fmt.Sprintf("aws/ec2/group/%s", gid)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
