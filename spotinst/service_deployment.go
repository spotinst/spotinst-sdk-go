package spotinst

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

// DeploymentService is an interface for interfacing with the deployment
// endpoints of the Spotinst API.
type DeploymentService interface {
	ListDeployments(*ListDeploymentsInput) (*ListDeploymentsOutput, error)
	CreateDeployment(*CreateDeploymentInput) (*CreateDeploymentOutput, error)
	ReadDeployment(*ReadDeploymentInput) (*ReadDeploymentOutput, error)
	UpdateDeployment(*UpdateDeploymentInput) (*UpdateDeploymentOutput, error)
	DeleteDeployment(*DeleteDeploymentInput) (*DeleteDeploymentOutput, error)
}

// DeploymentServiceOp handles communication with the deployment related methods
// of the Spotinst API.
type DeploymentServiceOp struct {
	client *Client
}

var _ DeploymentService = &DeploymentServiceOp{}

type Deployment struct {
	ID        *string    `json:"id,omitempty"`
	AccountID *string    `json:"accountId,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Tags      []*Tag     `json:"tags,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type ListDeploymentsInput struct {
	AccountID *string `json:"accountId,omitempty"`
}

type ListDeploymentsOutput struct {
	Deployments []*Deployment `json:"deployments,omitempty"`
}

type CreateDeploymentInput struct {
	AccountID  *string     `json:"accountId,omitempty"`
	Deployment *Deployment `json:"deployment,omitempty"`
}

type CreateDeploymentOutput struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

type ReadDeploymentInput struct {
	AccountID    *string `json:"accountId,omitempty"`
	DeploymentID *string `json:"deploymentId,omitempty"`
}

type ReadDeploymentOutput struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

type UpdateDeploymentInput struct {
	AccountID  *string     `json:"accountId,omitempty"`
	Deployment *Deployment `json:"deployment,omitempty"`
}

type UpdateDeploymentOutput struct{}

type DeleteDeploymentInput struct {
	AccountID    *string `json:"accountId,omitempty"`
	DeploymentID *string `json:"deployment,omitempty"`
}

type DeleteDeploymentOutput struct{}

func deploymentFromJSON(in []byte) (*Deployment, error) {
	b := new(Deployment)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func deploymentsFromJSON(in []byte) ([]*Deployment, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Deployment, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rp := range rw.Response.Items {
		p, err := deploymentFromJSON(rp)
		if err != nil {
			return nil, err
		}
		out[i] = p
	}
	return out, nil
}

func deploymentsFromHttpResponse(resp *http.Response) ([]*Deployment, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return deploymentsFromJSON(body)
}

func (c *DeploymentServiceOp) ListDeployments(input *ListDeploymentsInput) (*ListDeploymentsOutput, error) {
	r := c.client.newRequest("GET", "/loadBalancer/deployment")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ds, err := deploymentsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListDeploymentsOutput{
		Deployments: ds,
	}, nil
}

func (c *DeploymentServiceOp) CreateDeployment(input *CreateDeploymentInput) (*CreateDeploymentOutput, error) {
	r := c.client.newRequest("POST", "/loadBalancer/deployment")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ds, err := deploymentsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &CreateDeploymentOutput{
		Deployment: ds[0],
	}, nil
}

func (c *DeploymentServiceOp) ReadDeployment(input *ReadDeploymentInput) (*ReadDeploymentOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/deployment/{deploymentId}", map[string]string{
		"deploymentId": StringValue(input.DeploymentID),
	})
	if err != nil {
		return nil, err
	}

	r := c.client.newRequest("GET", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ds, err := deploymentsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ReadDeploymentOutput{
		Deployment: ds[0],
	}, nil
}

func (c *DeploymentServiceOp) UpdateDeployment(input *UpdateDeploymentInput) (*UpdateDeploymentOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/deployment/{deploymentId}", map[string]string{
		"deploymentId": StringValue(input.Deployment.ID),
	})
	if err != nil {
		return nil, err
	}

	r := c.client.newRequest("PUT", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateDeploymentOutput{}, nil
}

func (c *DeploymentServiceOp) DeleteDeployment(input *DeleteDeploymentInput) (*DeleteDeploymentOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/deployment/{deploymentId}", map[string]string{
		"deploymentId": StringValue(input.DeploymentID),
	})
	if err != nil {
		return nil, err
	}

	r := c.client.newRequest("DELETE", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteDeploymentOutput{}, nil
}
