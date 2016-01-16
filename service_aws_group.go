/*
  @author    Liran Polak
  @copyright Copyright (c) 2016, Spotinst
  @license   GPL-3.0
*/

package spotinstsdk

import (
	"fmt"
	"net/http"
)

const (
	serviceName = "group"
	endpoint    = "aws/ec2/group"
	apiVersion  = "v1"
)

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
		Errors []Error    `json:"errors"`
		Items  []AwsGroup `json:"items"`
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

// Lists a specific/all groups.
func (s *AwsGroupService) Get(id ...string) ([]AwsGroup, error) {
	var (
		retval AwsGroupResponse
		gid    string
	)
	if len(id) > 0 {
		gid = id[0]
	}
	_, err := s.client.get(fmt.Sprintf("%s/%s", endpoint, gid), &retval)
	if err != nil {
		return nil, err
	}

	return retval.Response.Items, nil
}

// Creates a new group.
func (s *AwsGroupService) Create(group AwsGroup) ([]AwsGroup, error) {
	var retval AwsGroupResponse
	_, err := s.client.post(endpoint, groupWrapper{Group: group}, &retval)
	if err != nil {
		return nil, err
	}

	return retval.Response.Items, nil
}

// Updates an existing group.
func (s *AwsGroupService) Update(group AwsGroup) ([]AwsGroup, error) {
	var retval AwsGroupResponse
	var gid = group.Id
	group.Id = ""
	_, err := s.client.put(fmt.Sprintf("%s/%s", endpoint, gid), groupWrapper{Group: group}, &retval)
	if err != nil {
		return nil, err
	}

	return retval.Response.Items, nil
}

// Deletes an existing group.
func (s *AwsGroupService) Delete(group AwsGroup) (*http.Response, error) {
	return s.client.delete(fmt.Sprintf("%s/%s", endpoint, group.Id), nil)
}
