/*
  @author    Liran Polak
  @copyright Copyright (c) 2016, Spotinst
  @license   GPL-3.0
*/

package spotinstsdk

import (
	"net/http"
	"fmt"
)

const (
	path = "aws/ec2/group"
)

type GroupService struct {
	client *Client
}

type Group struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Capacity    GroupCapacity `json:"capacity"`
	Compute     GroupCompute `json:"compute"`
	Strategy    GroupStrategy `json:"strategy"`
	Scaling     GroupScaling `json:"scaling,omitempty"`
}

type GroupScaling struct {
	Up   []GroupScalingPolicy `json:"up,omitempty"`
	Down []GroupScalingPolicy `json:"down,omitempty"`
}

type GroupScalingPolicy struct {
	PolicyName        string `json:"policyName"`
	MetricName        string `json:"metricName"`
	Statistic         string `json:"statistic"`
	Unit              string `json:"unit"`
	Threshold         float32 `json:"threshold"`
	Adjustment        int `json:"adjustment"`
	Namespace         string `json:"namespace"`
	Period            int `json:"period"`
	EvaluationPeriods int `json:"evaluationPeriods"`
	Cooldown          int `json:"cooldown"`
	Dimensions        []map[string]string `json:"dimensions"`
}

type GroupStrategy struct {
	Risk               float32 `json:"risk"`
	AvailabilityVsCost string `json:"availabilityVsCost"`
}

type GroupCapacity struct {
	Minimum int `json:"minimum"`
	Maximum int `json:"maximum"`
	Target  int `json:"target"`
}

type GroupCompute struct {
	Product             string `json:"product"`
	InstanceTypes       GroupComputeInstanceType `json:"instanceTypes"`
	LaunchSpecification GroupComputeLaunchSpecification `json:"launchSpecification"`
	AvailabilityZones   []GroupComputeAvailabilityZone `json:"availabilityZones"`
}

type GroupComputeInstanceType struct {
	OnDemand string `json:"ondemand"`
	Spot     []string `json:"spot"`
}

type GroupComputeAvailabilityZone struct {
	Name     string `json:"name"`
	SubnetId string `json:"subnetId,omitempty"`
}

type GroupComputeLaunchSpecification struct {
	SecurityGroupIds []string `json:"securityGroupIds"`
	Monitoring       bool  `json:"monitoring"`
	ImageId          string `json:"imageId"`
	KeyPair          string `json:"keyPair"`
}

type groupWrapper struct {
	Group Group `json:"group"`
}

// Lists a specific/all groups.
func (s *GroupService) Get(id string) ([]Group, error) {
	var retval GroupResponse
	_, err := s.client.get(fmt.Sprintf("%s/%s", path, id), &retval)
	if err != nil {
		return nil, err
	}

	return retval.Response.Items, nil
}

// Creates a new group.
func (s *GroupService) Create(group Group) ([]Group, error) {
	var retval GroupResponse
	_, err := s.client.post(path, groupWrapper{Group: group}, &retval)
	if err != nil {
		return nil, err
	}

	return retval.Response.Items, nil
}

// Updates an existing group.
func (s *GroupService) Update(group Group) ([]Group, error) {
	var retval GroupResponse
	_, err := s.client.put(fmt.Sprintf("%s/%s", path, group.Id), group, &retval)
	if err != nil {
		return nil, err
	}

	return retval.Response.Items, nil
}

// Deletes an existing group.
func (s *GroupService) Delete(group Group) (*http.Response, error) {
	return s.client.delete(fmt.Sprintf("%s/%s", path, group.Id), nil)
}