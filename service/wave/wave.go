package wave

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
)

type Cluster struct {
	ID                *string      `json:"id,omitempty"`
	ClusterIdentifier *string      `json:"clusterIdentifier,omitempty"`
	Environment       *Environment `json:"environment,omitempty"`
	Config            *Config      `json:"config,omitempty"`
	State             *string      `json:"state,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type Environment struct {
	OperatorVersion         *string `json:"operatorVersion,omitempty"`
	CertManagerDeployed     *bool   `json:"certManagerDeployed,omitempty"`
	K8sClusterProvisioned   *bool   `json:"k8sClusterProvisioned,omitempty"`
	OceanClusterProvisioned *bool   `json:"oceanClusterProvisioned,omitempty"`
	EnvironmentNamespace    *string `json:"environmentNamespace,omitempty"`
	OceanClusterId          *string `json:"oceanClusterId,omitempty"`
}

type Config struct {
	Components []*Component `json:"components,omitempty"`
}

type Component struct {
	Uid             *string            `json:"uid,omitempty"`
	Name            *string            `json:"name,omitempty"`
	OperatorVersion *string            `json:"operatorVersion,omitempty"`
	Version         *string            `json:"version,omitempty"`
	Properties      *map[string]string `json:"properties,omitempty"`
	State           *string            `json:"state,omitempty"`
}

type ListClustersInput struct{}

type ListClustersOutput struct {
	Clusters []*Cluster `json:"clusters,omitempty"`
}

func (s *ServiceOp) ListClusters(ctx context.Context, input *ListClustersInput) (*ListClustersOutput, error) {
	r := client.NewRequest(http.MethodGet, "/wave/cluster")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	clusters, err := clustersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListClustersOutput{Clusters: clusters}, nil
}

func clustersFromHttpResponse(resp *http.Response) ([]*Cluster, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return clustersFromJSON(body)
}

func clustersFromJSON(in []byte) ([]*Cluster, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Cluster, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := clusterFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func clusterFromJSON(in []byte) (*Cluster, error) {
	b := new(Cluster)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}
