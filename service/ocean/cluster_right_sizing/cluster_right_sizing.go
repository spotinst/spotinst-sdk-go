package cluster_right_sizing

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

type ClusterConfiguration struct {
	AdjustLimitOnDownsize           *bool `json:"adjustLimitOnDownsize,omitempty"`
	DownsideOnly                    *bool `json:"downsideOnly,omitempty"`
	RecommendationsCpuPercentile    *int  `json:"recommendationsCpuPercentile,omitempty"`
	RecommendationsMemoryPercentile *int  `json:"recommendationsMemoryPercentile,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type PostClusterConfigurationInput struct {
	OceanId           *string               `json:"oceanId,omitempty"`
	ClusterIdentifier *string               `json:"clusterIdentifier,omitempty"`
	Config            *ClusterConfiguration `json:"config,omitempty"`
}

type PostClusterConfigurationOutput struct {
	ClusterConfiguration *ClusterConfiguration `json:"config,omitempty"`
}

type ReadClusterConfigurationInput struct {
	OceanId           *string `json:"oceanId,omitempty"`
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty"`
}

type ReadClusterConfigurationOutput struct {
	ClusterConfiguration *ClusterConfiguration `json:"config,omitempty"`
}

type clusterConfigurationWrapper struct {
	Config *ClusterConfiguration `json:"config,omitempty"`
}

func clusterConfigurationWrapperFromJSON(in []byte) (*clusterConfigurationWrapper, error) {
	b := new(clusterConfigurationWrapper)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func clusterConfigurationFromJSON(in []byte) (*ClusterConfiguration, error) {
	b := new(ClusterConfiguration)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func clusterConfigurationsFromJSON(in []byte) ([]*ClusterConfiguration, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}

	out := make([]*ClusterConfiguration, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}

	for i, rb := range rw.Response.Items {
		wrapper, err := clusterConfigurationWrapperFromJSON(rb)
		if err != nil {
			return nil, err
		}

		if wrapper.Config != nil {
			out[i] = wrapper.Config
			continue
		}

		cfg, err := clusterConfigurationFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = cfg
	}

	return out, nil
}

func clusterConfigurationsFromHttpResponse(resp *http.Response) ([]*ClusterConfiguration, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return clusterConfigurationsFromJSON(body)
}

func (s *ServiceOp) PostClusterConfiguration(ctx context.Context, input *PostClusterConfigurationInput) (*PostClusterConfigurationOutput, error) {
	path, err := uritemplates.Expand("/ocean/{oceanId}/rightSizing/cluster/configuration", uritemplates.Values{
		"oceanId": spotinst.StringValue(input.OceanId),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodPost, path)
	if input.ClusterIdentifier != nil {
		r.Params.Set("clusterIdentifier", spotinst.StringValue(input.ClusterIdentifier))
	}

	input.OceanId = nil
	input.ClusterIdentifier = nil
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cfgs, err := clusterConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(PostClusterConfigurationOutput)
	if len(cfgs) > 0 {
		output.ClusterConfiguration = cfgs[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadClusterConfiguration(ctx context.Context, input *ReadClusterConfigurationInput) (*ReadClusterConfigurationOutput, error) {
	path, err := uritemplates.Expand("/ocean/{oceanId}/rightSizing/cluster/configuration", uritemplates.Values{
		"oceanId": spotinst.StringValue(input.OceanId),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	if input.ClusterIdentifier != nil {
		r.Params.Set("clusterIdentifier", spotinst.StringValue(input.ClusterIdentifier))
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cfgs, err := clusterConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadClusterConfigurationOutput)
	if len(cfgs) > 0 {
		output.ClusterConfiguration = cfgs[0]
	}

	return output, nil
}

// region ClusterConfiguration

func (o ClusterConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod ClusterConfiguration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ClusterConfiguration) SetAdjustLimitOnDownsize(v *bool) *ClusterConfiguration {
	if o.AdjustLimitOnDownsize = v; o.AdjustLimitOnDownsize == nil {
		o.nullFields = append(o.nullFields, "AdjustLimitOnDownsize")
	}
	return o
}

func (o *ClusterConfiguration) SetDownsideOnly(v *bool) *ClusterConfiguration {
	if o.DownsideOnly = v; o.DownsideOnly == nil {
		o.nullFields = append(o.nullFields, "DownsideOnly")
	}
	return o
}

func (o *ClusterConfiguration) SetRecommendationsCpuPercentile(v *int) *ClusterConfiguration {
	if o.RecommendationsCpuPercentile = v; o.RecommendationsCpuPercentile == nil {
		o.nullFields = append(o.nullFields, "RecommendationsCpuPercentile")
	}
	return o
}

func (o *ClusterConfiguration) SetRecommendationsMemoryPercentile(v *int) *ClusterConfiguration {
	if o.RecommendationsMemoryPercentile = v; o.RecommendationsMemoryPercentile == nil {
		o.nullFields = append(o.nullFields, "RecommendationsMemoryPercentile")
	}
	return o
}
