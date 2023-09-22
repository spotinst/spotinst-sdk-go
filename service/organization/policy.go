package organization

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

type Policy struct {
	Description   *string        `json:"description,omitempty"`
	Name          *string        `json:"name,omitempty"`
	PolicyContent *PolicyContent `json:"policyContent,omitempty"`
	PolicyID      *string        `json:"id,omitempty"`

	// forceSendFields is a list of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string

	// nullFields is a list of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string
}

type PolicyContent struct {
	Statements []*Statement `json:"statements,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Statement struct {
	Actions   []string `json:"actions,omitempty"`
	Effect    *string  `json:"effect,omitempty"`
	Resources []string `json:"resources,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListPoliciesInput struct{}

type ListPoliciesOutput struct {
	Policies []*Policy `json:"policies,omitempty"`
}

type CreatePolicyInput struct {
	Policy *Policy `json:"policy,omitempty"`
}

type CreatePolicyOutput struct {
	Policy *Policy `json:"policy,omitempty"`
}

type ReadPolicyInput struct {
	PolicyID *string `json:"policyId,omitempty"`
}

type ReadPolicyOutput struct {
	Policy *Policy `json:"policy,omitempty"`
}

type UpdatePolicyInput struct {
	Policy *Policy `json:"policy,omitempty"`
}

type UpdatePolicyOutput struct {
	Policy *Policy `json:"policy,omitempty"`
}

type DeletePolicyInput struct {
	PolicyID *string `json:"id,omitempty"`
}

type DeletePolicyOutput struct{}

func policyFromJSON(in []byte) (*Policy, error) {
	b := new(Policy)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func policiesFromJSON(in []byte) ([]*Policy, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Policy, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := policyFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func policiesFromHttpResponse(resp *http.Response) ([]*Policy, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return policiesFromJSON(body)
}

func (s *ServiceOp) ListPolicies(ctx context.Context, input *ListPoliciesInput) (*ListPoliciesOutput, error) {
	r := client.NewRequest(http.MethodGet, "/setup/organization/policy")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := policiesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListPoliciesOutput{Policies: gs}, nil
}

func (s *ServiceOp) CreatePolicy(ctx context.Context, input *CreatePolicyInput) (*CreatePolicyOutput, error) {
	r := client.NewRequest(http.MethodPost, "/setup/access/policy")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ss, err := policiesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreatePolicyOutput)
	if len(ss) > 0 {
		output.Policy = ss[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadPolicy(ctx context.Context, input *ReadPolicyInput) (*ReadPolicyOutput, error) {

	r := client.NewRequest(http.MethodGet, "/setup/organization/policy")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := policiesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadPolicyOutput)
	if len(gs) > 0 {
		for i, value := range gs {
			if spotinst.StringValue(input.PolicyID) == spotinst.StringValue(value.PolicyID) {
				output.Policy = gs[i]
				break
			}
		}
	}

	return output, nil
}

func (s *ServiceOp) UpdatePolicy(ctx context.Context, input *UpdatePolicyInput) (*UpdatePolicyOutput, error) {
	path, err := uritemplates.Expand("/setup/access/policy/{policyId}", uritemplates.Values{
		"policyId": spotinst.StringValue(input.Policy.PolicyID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Policy.PolicyID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ss, err := policiesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdatePolicyOutput)
	if len(ss) > 0 {
		output.Policy = ss[0]
	}

	return output, nil
}

func (s *ServiceOp) DeletePolicy(ctx context.Context, input *DeletePolicyInput) (*DeletePolicyOutput, error) {
	path, err := uritemplates.Expand("/setup/access/policy/{policyId}", uritemplates.Values{
		"policyId": spotinst.StringValue(input.PolicyID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodDelete, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeletePolicyOutput{}, nil
}

// region Policy

func (o Policy) MarshalJSON() ([]byte, error) {
	type noMethod Policy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Policy) SetDescription(v *string) *Policy {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *Policy) SetName(v *string) *Policy {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Policy) SetPolicyContent(v *PolicyContent) *Policy {
	if o.PolicyContent = v; o.PolicyContent == nil {
		o.nullFields = append(o.nullFields, "PolicyContent")
	}
	return o
}

// endregion

func (o PolicyContent) MarshalJSON() ([]byte, error) {
	type noMethod PolicyContent
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *PolicyContent) SetStatements(v []*Statement) *PolicyContent {
	if o.Statements = v; o.Statements == nil {
		o.nullFields = append(o.nullFields, "Statements")
	}
	return o
}

// endregion

func (o Statement) MarshalJSON() ([]byte, error) {
	type noMethod Statement
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Statement) SetEffect(v *string) *Statement {
	if o.Effect = v; o.Effect == nil {
		o.nullFields = append(o.nullFields, "Effect")
	}
	return o
}

func (o *Statement) SetResources(v []string) *Statement {
	if o.Resources = v; o.Resources == nil {
		o.nullFields = append(o.nullFields, "Resources")
	}
	return o
}

func (o *Statement) SetActions(v []string) *Statement {
	if o.Actions = v; o.Actions == nil {
		o.nullFields = append(o.nullFields, "Actions")
	}
	return o
}

// endregion
