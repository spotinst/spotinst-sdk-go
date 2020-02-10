package aws

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

//ResourceSuggestion - single resource suggestion from Spot API
type ResourceSuggestion struct {
	DeploymentName  *string `json:"deploymentName,omitempty"`
	Namespace       *string `json:"namespace,omitempty"`
	SuggestedCPU    *int    `json:"suggestedCPU,omitempty"`
	RequestedCPU    *int    `json:"requestedCPU,omitempty"`
	SuggestedMemory *int    `json:"suggestedMemory,omitempty"`
	RequestedMemory *int    `json:"requestedMemory,omitempty"`
}

//ListResourceSuggestionsInput - Input struct required for getting Spot Right
//Sizing suggestions for an Ocean cluster
type ListResourceSuggestionsInput struct {
	OceanID *string `json:"oceanId,omitempty"`
}

//ListResourceSuggestionsOutput - output struct of suggestion array as Right Sizing
//API response with array of suggestions per Namespace & Deploymnet
type ListResourceSuggestionsOutput struct {
	Suggestions []*ResourceSuggestion `json:"suggestions,omitempty"`
}

func resourceSuggestionFromJSON(in []byte) (*ResourceSuggestion, error) {
	b := new(ResourceSuggestion)

	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}
func resourceSuggestionsFromJSON(in []byte) ([]*ResourceSuggestion, error) {
	var rw client.Response

	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ResourceSuggestion, len(rw.Response.Items))

	for i, rb := range rw.Response.Items {
		b, err := resourceSuggestionFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func resourceSuggestionsFromHTTPResponse(resp *http.Response) ([]*ResourceSuggestion, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return resourceSuggestionsFromJSON(body)
}

//ListResourceSuggestions - get all right-sizing resource suggestionsfor an Ocean cluster
func (s *ServiceOp) ListResourceSuggestions(ctx context.Context, input *ListResourceSuggestionsInput) (*ListResourceSuggestionsOutput, error) {
	path, err := uritemplates.Expand("/ocean/aws/k8s/cluster/{oceanId}/rightSizing/resourceSuggestion", uritemplates.Values{
		"oceanId": spotinst.StringValue(input.OceanID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := resourceSuggestionsFromHTTPResponse(resp)
	if err != nil {
		return nil, err
	}
	return &ListResourceSuggestionsOutput{Suggestions: gs}, nil
}
