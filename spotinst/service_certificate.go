package spotinst

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

// CertificateService is an interface for interfacing with the certificate
// endpoints of the Spotinst API.
type CertificateService interface {
	ListCertificates(*ListCertificatesInput) (*ListCertificatesOutput, error)
	CreateCertificate(*CreateCertificateInput) (*CreateCertificateOutput, error)
	ReadCertificate(*ReadCertificateInput) (*ReadCertificateOutput, error)
	UpdateCertificate(*UpdateCertificateInput) (*UpdateCertificateOutput, error)
	DeleteCertificate(*DeleteCertificateInput) (*DeleteCertificateOutput, error)
}

// CertificateServiceOp handles communication with the certificate related methods
// of the Spotinst API.
type CertificateServiceOp struct {
	client *Client
}

var _ CertificateService = &CertificateServiceOp{}

type Certificate struct {
	ID           *string    `json:"id,omitempty"`
	AccountID    *string    `json:"accountId,omitempty"`
	Name         *string    `json:"name,omitempty"`
	CertPEMBlock *string    `json:"certificatePemBlock,omitempty"`
	KeyPEMBlock  *string    `json:"keyPemBlock,omitempty"`
	Tags         []*Tag     `json:"tags,omitempty"`
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
}

type ListCertificatesInput struct {
	AccountID *string `json:"accountId,omitempty"`
}

type ListCertificatesOutput struct {
	Certificates []*Certificate `json:"certificates,omitempty"`
}

type CreateCertificateInput struct {
	AccountID   *string      `json:"accountId,omitempty"`
	Certificate *Certificate `json:"certificate,omitempty"`
}

type CreateCertificateOutput struct {
	Certificate *Certificate `json:"certificate,omitempty"`
}

type ReadCertificateInput struct {
	AccountID     *string `json:"accountId,omitempty"`
	CertificateID *string `json:"certificateId,omitempty"`
}

type ReadCertificateOutput struct {
	Certificate *Certificate `json:"certificate,omitempty"`
}

type UpdateCertificateInput struct {
	AccountID   *string      `json:"accountId,omitempty"`
	Certificate *Certificate `json:"certificate,omitempty"`
}

type UpdateCertificateOutput struct{}

type DeleteCertificateInput struct {
	AccountID     *string `json:"accountId,omitempty"`
	CertificateID *string `json:"certificateId,omitempty"`
}

type DeleteCertificateOutput struct{}

func certificateFromJSON(in []byte) (*Certificate, error) {
	b := new(Certificate)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func certificatesFromJSON(in []byte) ([]*Certificate, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Certificate, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rp := range rw.Response.Items {
		p, err := certificateFromJSON(rp)
		if err != nil {
			return nil, err
		}
		out[i] = p
	}
	return out, nil
}

func certificatesFromHttpResponse(resp *http.Response) ([]*Certificate, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return certificatesFromJSON(body)
}

func (c *CertificateServiceOp) ListCertificates(input *ListCertificatesInput) (*ListCertificatesOutput, error) {
	r := c.client.newRequest("GET", "/loadBalancer/certificate")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs, err := certificatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListCertificatesOutput{
		Certificates: cs,
	}, nil
}

func (c *CertificateServiceOp) CreateCertificate(input *CreateCertificateInput) (*CreateCertificateOutput, error) {
	r := c.client.newRequest("POST", "/loadBalancer/certificate")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs, err := certificatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &CreateCertificateOutput{
		Certificate: cs[0],
	}, nil
}

func (c *CertificateServiceOp) ReadCertificate(input *ReadCertificateInput) (*ReadCertificateOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/certificate/{certificateId}", map[string]string{
		"certificateId": StringValue(input.CertificateID),
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

	cs, err := certificatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ReadCertificateOutput{
		Certificate: cs[0],
	}, nil
}

func (c *CertificateServiceOp) UpdateCertificate(input *UpdateCertificateInput) (*UpdateCertificateOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/certificate/{certificateId}", map[string]string{
		"certificateId": StringValue(input.Certificate.ID),
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

	return &UpdateCertificateOutput{}, nil
}

func (c *CertificateServiceOp) DeleteCertificate(input *DeleteCertificateInput) (*DeleteCertificateOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/certificate/{certificateId}", map[string]string{
		"certificateId": StringValue(input.CertificateID),
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

	return &DeleteCertificateOutput{}, nil
}
