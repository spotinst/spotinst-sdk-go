package administration

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

type UserGroup struct {
	Description *string          `json:"description,omitempty"`
	Name        *string          `json:"name,omitempty"`
	Policies    []*UserPolicy    `json:"policies,omitempty"`
	UserIds     []string         `json:"userIds,omitempty"`
	UserGroupId *string          `json:"id,omitempty"`
	Users       []*UserGroupUser `json:"users,omitempty"`

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

type UserGroupPolicy struct {
	AccountIds []string `json:"accountIds,omitempty"`
	PolicyId   *string  `json:"policyId,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type UserGroupUser struct {
	Type     *string `json:"type,omitempty"`
	UserId   *string `json:"userId,omitempty"`
	UserName *string `json:"userName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListUserGroupsInput struct{}

type ListUserGroupsOutput struct {
	UserGroups []*UserGroup `json:"userGroups,omitempty"`
}

type CreateUserGroupOutput struct {
	UserGroup *UserGroup `json:"userGroup,omitempty"`
}

type ReadUserGroupInput struct {
	UserGroupID *string `json:"id,omitempty"`
}

type ReadUserGroupOutput struct {
	UserGroup *UserGroup `json:"userGroup,omitempty"`
}

type UpdateUserGroupInput struct {
	UserGroupID *string `json:"id,omitempty"`
}

type UpdateUserGroupOutput struct {
	UserGroup *UserGroup `json:"userGroup,omitempty"`
}

type DeleteUserGroupInput struct {
	UserGroupID *string `json:"id,omitempty"`
}

type DeleteUserGroupOutput struct{}

func userGroupFromJSON(in []byte) (*UserGroup, error) {
	b := new(UserGroup)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func userGroupsFromJSON(in []byte) ([]*UserGroup, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*UserGroup, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := userGroupFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func userGroupsFromHttpResponse(resp *http.Response) ([]*UserGroup, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return userGroupsFromJSON(body)
}

func (s *ServiceOp) ListUserGroups(ctx context.Context, input *ListUserGroupsInput) (*ListUserGroupsOutput, error) {
	r := client.NewRequest(http.MethodGet, "/setup/access/userGroup")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := userGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListUserGroupsOutput{UserGroups: gs}, nil
}

func (s *ServiceOp) CreateUserGroup(ctx context.Context, input *UserGroup) (*CreateUserGroupOutput, error) {
	r := client.NewRequest(http.MethodPost, "/setup/access/userGroup")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ss, err := userGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateUserGroupOutput)
	if len(ss) > 0 {
		output.UserGroup = ss[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadUserGroup(ctx context.Context, input *ReadUserGroupInput) (*ReadUserGroupOutput, error) {
	path, err := uritemplates.Expand("/setup/access/userGroup/{userGroupId}", uritemplates.Values{
		"userGroupId": spotinst.StringValue(input.UserGroupID),
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

	ss, err := userGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadUserGroupOutput)
	if len(ss) > 0 {
		output.UserGroup = ss[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateUserGroup(ctx context.Context, input *UserGroup) (*UpdateUserGroupOutput, error) {
	path, err := uritemplates.Expand("/setup/access/userGroup/{userGroupId}", uritemplates.Values{
		"userGroupId": spotinst.StringValue(input.UserGroupId),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.UserGroupId = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ss, err := userGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateUserGroupOutput)
	if len(ss) > 0 {
		output.UserGroup = ss[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteUserGroup(ctx context.Context, input *DeleteUserGroupInput) (*DeleteUserGroupOutput, error) {
	path, err := uritemplates.Expand("/setup/access/userGroup/{userGroupId}", uritemplates.Values{
		"userGroupId": spotinst.StringValue(input.UserGroupID),
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

	return &DeleteUserGroupOutput{}, nil
}

// region User

func (o UserGroup) MarshalJSON() ([]byte, error) {
	type noMethod UserGroup
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *UserGroup) SetDescription(v *string) *UserGroup {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *UserGroup) SetName(v *string) *UserGroup {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *UserGroup) SetPolicies(v []*UserPolicy) *UserGroup {
	if o.Policies = v; o.Policies == nil {
		o.nullFields = append(o.nullFields, "Policies")
	}
	return o
}

func (o *UserGroup) SetUserIds(v []string) *UserGroup {
	if o.UserIds = v; o.UserIds == nil {
		o.nullFields = append(o.nullFields, "UserIds")
	}
	return o
}

// endregion

func (o UserGroupPolicy) MarshalJSON() ([]byte, error) {
	type noMethod UserGroupPolicy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *UserGroupPolicy) SetPolicyId(v *string) *UserGroupPolicy {
	if o.PolicyId = v; o.PolicyId == nil {
		o.nullFields = append(o.nullFields, "PolicyId")
	}
	return o
}

func (o *UserGroupPolicy) SetAccountIds(v []string) *UserGroupPolicy {
	if o.AccountIds = v; o.AccountIds == nil {
		o.nullFields = append(o.nullFields, "AccountIds")
	}
	return o
}

//end region
