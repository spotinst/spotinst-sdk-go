package right_sizing_cluster_config

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

type RightsizingClusterConfiguration struct {
	AdjustLimitOnDownsize           *bool `json:"adjustLimitOnDownsize,omitempty"`
	DownsideOnly                    *bool `json:"downsideOnly,omitempty"`
	RecommendationsCpuPercentile    *int  `json:"recommendationsCpuPercentile,omitempty"`
	RecommendationsMemoryPercentile *int  `json:"recommendationsMemoryPercentile,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RightsizingClusterConfigurationInput struct {
	OceanId           *string                          `json:"oceanId,omitempty"`
	ClusterIdentifier *string                          `json:"clusterIdentifier,omitempty"`
	Config            *RightsizingClusterConfiguration `json:"config,omitempty"`
}

type RightsizingClusterConfigurationOutput struct {
	ClusterConfiguration *RightsizingClusterConfiguration `json:"config,omitempty"`
}

type ReadRightsizingClusterConfigurationInput struct {
	OceanId           *string `json:"oceanId,omitempty"`
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty"`
}

type ReadRightsizingClusterConfigurationOutput struct {
	ClusterConfiguration *RightsizingClusterConfiguration `json:"config,omitempty"`
}

type clusterConfigurationWrapper struct {
	Config *RightsizingClusterConfiguration `json:"config,omitempty"`
}

func clusterConfigurationWrapperFromJSON(in []byte) (*clusterConfigurationWrapper, error) {
	b := new(clusterConfigurationWrapper)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func clusterConfigurationFromJSON(in []byte) (*RightsizingClusterConfiguration, error) {
	b := new(RightsizingClusterConfiguration)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func clusterConfigurationsFromJSON(in []byte) ([]*RightsizingClusterConfiguration, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}

	out := make([]*RightsizingClusterConfiguration, len(rw.Response.Items))
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

func clusterConfigurationsFromHttpResponse(resp *http.Response) ([]*RightsizingClusterConfiguration, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return clusterConfigurationsFromJSON(body)
}

func (s *ServiceOp) PostClusterConfiguration(ctx context.Context, input *RightsizingClusterConfigurationInput) (*RightsizingClusterConfigurationOutput, error) {
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

	output := new(RightsizingClusterConfigurationOutput)
	if len(cfgs) > 0 {
		output.ClusterConfiguration = cfgs[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadClusterConfiguration(ctx context.Context, input *ReadRightsizingClusterConfigurationInput) (*ReadRightsizingClusterConfigurationOutput, error) {
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

	output := new(ReadRightsizingClusterConfigurationOutput)
	if len(cfgs) > 0 {
		output.ClusterConfiguration = cfgs[0]
	}

	return output, nil
}

// region ClusterConfiguration

func (o RightsizingClusterConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod RightsizingClusterConfiguration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RightsizingClusterConfiguration) SetAdjustLimitOnDownsize(v *bool) *RightsizingClusterConfiguration {
	if o.AdjustLimitOnDownsize = v; o.AdjustLimitOnDownsize == nil {
		o.nullFields = append(o.nullFields, "AdjustLimitOnDownsize")
	}
	return o
}

func (o *RightsizingClusterConfiguration) SetDownsideOnly(v *bool) *RightsizingClusterConfiguration {
	if o.DownsideOnly = v; o.DownsideOnly == nil {
		o.nullFields = append(o.nullFields, "DownsideOnly")
	}
	return o
}

func (o *RightsizingClusterConfiguration) SetRecommendationsCpuPercentile(v *int) *RightsizingClusterConfiguration {
	if o.RecommendationsCpuPercentile = v; o.RecommendationsCpuPercentile == nil {
		o.nullFields = append(o.nullFields, "RecommendationsCpuPercentile")
	}
	return o
}

func (o *RightsizingClusterConfiguration) SetRecommendationsMemoryPercentile(v *int) *RightsizingClusterConfiguration {
	if o.RecommendationsMemoryPercentile = v; o.RecommendationsMemoryPercentile == nil {
		o.nullFields = append(o.nullFields, "RecommendationsMemoryPercentile")
	}
	return o
}
