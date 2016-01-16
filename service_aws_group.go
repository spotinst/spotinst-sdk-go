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

type GroupService struct {
	client *Client
}

type Group struct {
	Id          string         `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Capacity    *GroupCapacity `json:"capacity,omitempty"`
	Compute     *GroupCompute  `json:"compute,omitempty"`
	Strategy    *GroupStrategy `json:"strategy,omitempty"`
	Scaling     *GroupScaling  `json:"scaling,omitempty"`
}

type GroupScaling struct {
	Up   []*GroupScalingPolicy `json:"up,omitempty"`
	Down []*GroupScalingPolicy `json:"down,omitempty"`
}

type GroupScalingPolicy struct {
	PolicyName        string                         `json:"policyName,omitempty"`
	MetricName        string                         `json:"metricName,omitempty"`
	Statistic         string                         `json:"statistic,omitempty"`
	Unit              string                         `json:"unit,omitempty"`
	Threshold         float64                        `json:"threshold,omitempty"`
	Adjustment        int                            `json:"adjustment,omitempty"`
	Namespace         string                         `json:"namespace,omitempty"`
	EvaluationPeriods int                            `json:"evaluationPeriods"`
	Period            int                            `json:"period"`
	Cooldown          int                            `json:"cooldown"`
	Dimensions        []*GroupScalingPolicyDimension `json:"dimensions,omitempty"`
}

type GroupScalingPolicyDimension struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type GroupStrategy struct {
	AvailabilityVsCost string  `json:"availabilityVsCost,omitempty"`
	Risk               float64 `json:"risk"`
}

type GroupCapacity struct {
	Minimum int `json:"minimum"`
	Maximum int `json:"maximum"`
	Target  int `json:"target"`
}

type GroupCompute struct {
	Product             string                           `json:"product,omitempty"`
	InstanceTypes       *GroupComputeInstanceType        `json:"instanceTypes,omitempty"`
	LaunchSpecification *GroupComputeLaunchSpecification `json:"launchSpecification,omitempty"`
	AvailabilityZones   []*GroupComputeAvailabilityZone  `json:"availabilityZones,omitempty"`
}

type GroupComputeInstanceType struct {
	OnDemand string   `json:"ondemand,omitempty"`
	Spot     []string `json:"spot,omitempty"`
}

type GroupComputeAvailabilityZone struct {
	Name     string `json:"name,omitempty"`
	SubnetId string `json:"subnetId,omitempty"`
}

type GroupComputeLaunchSpecification struct {
	SecurityGroupIds []string `json:"securityGroupIds,omitempty"`
	ImageId          string   `json:"imageId,omitempty"`
	KeyPair          string   `json:"keyPair,omitempty"`
	Monitoring       bool     `json:"monitoring"`
}

type groupWrapper struct {
	Group Group `json:"group"`
}

// Lists a specific/all groups.
func (s *GroupService) Get(id ...string) ([]Group, error) {
	var (
		retval GroupResponse
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
func (s *GroupService) Create(group Group) ([]Group, error) {
	var retval GroupResponse
	_, err := s.client.post(endpoint, groupWrapper{Group: group}, &retval)
	if err != nil {
		return nil, err
	}

	return retval.Response.Items, nil
}

// Updates an existing group.
func (s *GroupService) Update(group Group) ([]Group, error) {
	var retval GroupResponse
	var gid = group.Id
	group.Id = ""
	_, err := s.client.put(fmt.Sprintf("%s/%s", endpoint, gid), groupWrapper{Group: group}, &retval)
	if err != nil {
		return nil, err
	}

	return retval.Response.Items, nil
}

// Deletes an existing group.
func (s *GroupService) Delete(group Group) (*http.Response, error) {
	return s.client.delete(fmt.Sprintf("%s/%s", endpoint, group.Id), nil)
}
