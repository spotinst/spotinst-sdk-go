package mcs

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

type ClusterCostInput struct {
	ClusterID *string `json:"clusterId,omitempty"`
	ToDate    *string `json:"toDate,omitempty"`
	FromDate  *string `json:"fromDate,omitempty"`
}

type ClusterCostOutput struct {
	ClusterCosts []*ClusterCost `json:"clusterCosts,omitempty"`
}

type ClusterCost struct {
	Namespaces         []*Namespace  `json:"namespaces,omitempty"`
	Deployments        []*Deployment `json:"deployments,omitempty"`
	TotalCost          *float64      `json:"totalCost,omitempty"`
	TotalComputeCost   *float64      `json:"totalComputeCost,omitempty"`
	TotalEBSCost       *float64      `json:"totalEbsCost,omitempty"`
	TotalStorageCost   *float64      `json:"totalStorageCost,omitempty"`
	UnusedStorageCost  *float64      `json:"unusedStorageCost,omitempty"`
	StandAlonePodsCost *float64      `json:"standAlonePodsCost,omitempty"`
	HeadroomCost       *float64      `json:"headroomCost,omitempty"`
	IdleCost           *float64      `json:"idleCost,omitempty"`
}

type Namespace struct {
	Namespace          *string           `json:"namespace,omitempty"`
	Cost               *float64          `json:"cost,omitempty"`
	ComputeCost        *float64          `json:"computeCost,omitempty"`
	EBSCost            *float64          `json:"ebsCost,omitempty"`
	StorageCost        *float64          `json:"storageCost,omitempty"`
	Deployments        []*Resource       `json:"deployments,omitempty"`
	StatefulSets       []*Resource       `json:"statefulSets,omitempty"`
	DaemonSets         []*Resource       `json:"daemonSets,omitempty"`
	Jobs               []*Resource       `json:"jobs,omitempty"`
	StandAlonePodsCost *Resource         `json:"standAlonePodsCost,omitempty"`
	Labels             map[string]string `json:"labels,omitempty"`
	Annotations        map[string]string `json:"annotations,omitempty"`
}

// Deprecated: Use Resource instead. Kept for backward compatibility.
type Deployment struct {
	DeploymentName *string           `json:"name,omitempty"`
	Namespace      *string           `json:"namespace,omitempty"`
	Cost           *float64          `json:"cost,omitempty"`
	ComputeCost    *float64          `json:"computeCost,omitempty"`
	StorageCost    *float64          `json:"storageCost,omitempty"`
	Labels         map[string]string `json:"labels,omitempty"`
	Annotations    map[string]string `json:"annotations,omitempty"`
}

type Resource struct {
	Name        *string           `json:"name,omitempty"`
	Namespace   *string           `json:"namespace,omitempty"`
	Cost        *float64          `json:"cost,omitempty"`
	ComputeCost *float64          `json:"computeCost,omitempty"`
	EBSCost     *float64          `json:"ebsCost,omitempty"`
	StorageCost *float64          `json:"storageCost,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

func clusterCostFromJSON(in []byte) (*ClusterCost, error) {
	b := new(ClusterCost)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func clusterCostsFromJSON(in []byte) ([]*ClusterCost, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ClusterCost, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := clusterCostFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func clusterCostsFromHttpResponse(resp *http.Response) ([]*ClusterCost, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return clusterCostsFromJSON(body)
}

// GetClusterCosts accepts Kubernetes `clusterId`, `fromDate`, and `toDate` and
// returns a list of cost objects. Dates can be in the format of `yyyy-mm-dd`
// or Unix timestamp (1494751821472).
func (s *ServiceOp) GetClusterCosts(ctx context.Context, input *ClusterCostInput) (*ClusterCostOutput, error) {
	path, err := uritemplates.Expand("/mcs/kubernetes/cluster/{clusterIdentifier}/costs", uritemplates.Values{
		"clusterIdentifier": spotinst.StringValue(input.ClusterID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)

	if input.ToDate != nil {
		r.Params.Set("toDate", *input.ToDate)
	}
	if input.FromDate != nil {
		r.Params.Set("fromDate", *input.FromDate)
	}
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	costs, err := clusterCostsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ClusterCostOutput{costs}, nil
}
