package notificationcenter

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
	"io/ioutil"
	"net/http"
)

type NotificationCenter struct {
	Name                *string              `json:"name,omitempty"`
	Description         *string              `json:"description,omitempty"`
	PrivacyLevel        *string              `json:"privacyLevel,omitempty"`
	IsActive            *bool                `json:"isActive,omitempty"`
	RegisteredUsers     []*RegisteredUsers   `json:"registeredUsers,omitempty"`
	Subscriptions       []*Subscriptions     `json:"subscriptions,omitempty"`
	ComputePolicyConfig *ComputePolicyConfig `json:"computePolicyConfig,omitempty"`
	ID                  *string              `json:"id,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RegisteredUsers struct {
	UserEmail         *string  `json:"userEmail,omitempty"`
	SubscriptionTypes []string `json:"subscriptionTypes,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Subscriptions struct {
	Endpoint *string `json:"endpoint,omitempty"`
	Type     *string `json:"type,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ComputePolicyConfig struct {
	Events                    []*Events       `json:"events,omitempty"`
	ShouldIncludeAllResources *bool           `json:"shouldIncludeAllResources,omitempty"`
	ResourceIds               []string        `json:"resourceIds,omitempty"`
	DynamicRules              []*DynamicRules `json:"dynamicRules,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Events struct {
	Event *string `json:"event,omitempty"`
	Type  *string `json:"type,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DynamicRules struct {
	FilterConditions []*FilterConditions `json:"filterConditions,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type FilterConditions struct {
	Identifier *string `json:"identifier,omitempty"`
	Operator   *string `json:"operator,omitempty"`
	Expression *string `json:"expression,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListNotificationCenterPolicyInput struct{}

type ListNotificationCenterPolicyOutput struct {
	NotificationCenter []*NotificationCenter `json:"notificationCenter,omitempty"`
}

type CreateNotificationCenterPolicyOutput struct {
	NotificationCenter *NotificationCenter `json:"notificationCenter,omitempty"`
}

type ReadNotificationCenterPolicyInput struct {
	PolicyId *string `json:"policyId,omitempty"`
}

type ReadNotificationCenterPolicyOutput struct {
	NotificationCenter *NotificationCenter `json:"notificationCenter,omitempty"`
}

type UpdateNotificationCenterPolicyInput struct {
	NotificationCenter *NotificationCenter `json:"notificationCenter,omitempty"`
}
type UpdateNotificationCenterPolicyOutput struct{}

type DeleteNotificationCenterPolicyInput struct {
	PolicyId *string `json:"policyId,omitempty"`
}

type DeleteNotificationCenterPolicyOutput struct{}

func notificationFromJSON(data []byte) (*NotificationCenter, error) {
	b := new(NotificationCenter)
	if err := json.Unmarshal(data, &b); err != nil {
		return nil, err
	}
	return b, nil
}

func notificationsFromJSON(data []byte) ([]*NotificationCenter, error) {
	var rw client.Response
	if err := json.Unmarshal(data, &rw); err != nil {
		return nil, err
	}
	out := make([]*NotificationCenter, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := notificationFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func notificationsFromHttpResponse(resp *http.Response) ([]*NotificationCenter, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return notificationsFromJSON(body)
}

func (s *ServiceOp) ListNotificationCenterPolicy(ctx context.Context, input *ListNotificationCenterPolicyInput) (*ListNotificationCenterPolicyOutput, error) {
	r := client.NewRequest(http.MethodGet, "/notificationCenter/policy")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	gs, err := notificationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}
	return &ListNotificationCenterPolicyOutput{NotificationCenter: gs}, nil
}

func (s *ServiceOp) CreateNotificationCenterPolicy(ctx context.Context, input *NotificationCenter) (*CreateNotificationCenterPolicyOutput, error) {
	r := client.NewRequest(http.MethodPost, "/notificationCenter/policy")
	r.Obj = input
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	gs, err := notificationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}
	output := new(CreateNotificationCenterPolicyOutput)
	if len(gs) > 0 {
		output.NotificationCenter = gs[0]
	}
	return output, nil
}

func (s *ServiceOp) ReadNotificationCenterPolicy(ctx context.Context, input *ReadNotificationCenterPolicyInput) (*ReadNotificationCenterPolicyOutput, error) {
	path, err := uritemplates.Expand("/notificationCenter/policy/{policyId}", uritemplates.Values{
		"policyId": spotinst.StringValue(input.PolicyId),
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
	gs, err := notificationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}
	output := new(ReadNotificationCenterPolicyOutput)
	if len(gs) > 0 {
		output.NotificationCenter = gs[0]
	}
	return output, nil
}

func (s *ServiceOp) UpdateNotificationCenterPolicy(ctx context.Context, input *UpdateNotificationCenterPolicyInput) (*UpdateNotificationCenterPolicyOutput, error) {
	path, err := uritemplates.Expand("/notificationCenter/policy/{policyId}", uritemplates.Values{
		"policyId": spotinst.StringValue(input.NotificationCenter.ID),
	})
	if err != nil {
		return nil, err
	}
	input.NotificationCenter.ID = nil
	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	/*gs, err := notificationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}
	output := new(UpdateNotificationCenterPolicyOutput)
	if len(gs) > 0 {
	output.NotificationCenter = gs[0]
	}
	return output, nil
	*/
	return &UpdateNotificationCenterPolicyOutput{}, nil
}

func (s *ServiceOp) DeleteNotificationCenterPolicy(ctx context.Context, input *DeleteNotificationCenterPolicyInput) (*DeleteNotificationCenterPolicyOutput, error) {
	path, err := uritemplates.Expand("/notificationCenter/policy/{policyId}", uritemplates.Values{
		"policyId": spotinst.StringValue(input.PolicyId),
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
	return &DeleteNotificationCenterPolicyOutput{}, nil
}

func (o NotificationCenter) MarshalJSON() ([]byte, error) {
	type noMethod NotificationCenter
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *NotificationCenter) SetName(v *string) *NotificationCenter {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *NotificationCenter) SetDescription(v *string) *NotificationCenter {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *NotificationCenter) SetPrivacyLevel(v *string) *NotificationCenter {
	if o.PrivacyLevel = v; o.PrivacyLevel == nil {
		o.nullFields = append(o.nullFields, "PrivacyLevel")
	}
	return o
}

func (o *NotificationCenter) SetIsActive(v *bool) *NotificationCenter {
	if o.IsActive = v; o.IsActive == nil {
		o.nullFields = append(o.nullFields, "IsActive")
	}
	return o
}

func (o *NotificationCenter) SetID(v *string) *NotificationCenter {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *NotificationCenter) SetRegisteredUsers(v []*RegisteredUsers) *NotificationCenter {
	if o.RegisteredUsers = v; o.RegisteredUsers == nil {
		o.nullFields = append(o.nullFields, "RegisteredUsers")
	}
	return o
}

func (o *NotificationCenter) SetSubscriptions(v []*Subscriptions) *NotificationCenter {
	if o.Subscriptions = v; o.Subscriptions == nil {
		o.nullFields = append(o.nullFields, "Subscriptions")
	}
	return o
}

func (o *NotificationCenter) SetComputePolicyConfig(v *ComputePolicyConfig) *NotificationCenter {
	if o.ComputePolicyConfig = v; o.ComputePolicyConfig == nil {
		o.nullFields = append(o.nullFields, "ComputePolicyConfig")
	}
	return o
}

func (o RegisteredUsers) MarshalJSON() ([]byte, error) {
	type noMethod RegisteredUsers
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RegisteredUsers) SetUserEmail(v *string) *RegisteredUsers {
	if o.UserEmail = v; o.UserEmail == nil {
		o.nullFields = append(o.nullFields, "UserEmail")
	}
	return o
}

func (o *RegisteredUsers) SetSubscriptionTypes(v []string) *RegisteredUsers {
	if o.SubscriptionTypes = v; o.SubscriptionTypes == nil {
		o.nullFields = append(o.nullFields, "SubscriptionTypes")
	}
	return o
}

func (o Subscriptions) MarshalJSON() ([]byte, error) {
	type noMethod Subscriptions
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Subscriptions) SetEndpoint(v *string) *Subscriptions {
	if o.Endpoint = v; o.Endpoint == nil {
		o.nullFields = append(o.nullFields, "Endpoint")
	}
	return o
}

func (o *Subscriptions) SetType(v *string) *Subscriptions {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o ComputePolicyConfig) MarshalJSON() ([]byte, error) {
	type noMethod ComputePolicyConfig
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ComputePolicyConfig) SetEvents(v []*Events) *ComputePolicyConfig {
	if o.Events = v; o.Events == nil {
		o.nullFields = append(o.nullFields, "Events")
	}
	return o
}

func (o *ComputePolicyConfig) SetShouldIncludeAllResources(v *bool) *ComputePolicyConfig {
	if o.ShouldIncludeAllResources = v; o.ShouldIncludeAllResources == nil {
		o.nullFields = append(o.nullFields, "ShouldIncludeAllResources")
	}
	return o
}

func (o *ComputePolicyConfig) SetResourceIds(v []string) *ComputePolicyConfig {
	if o.ResourceIds = v; o.ResourceIds == nil {
		o.nullFields = append(o.nullFields, "ResourceIds")
	}
	return o
}

func (o *ComputePolicyConfig) SetDynamicRules(v []*DynamicRules) *ComputePolicyConfig {
	if o.DynamicRules = v; o.DynamicRules == nil {
		o.nullFields = append(o.nullFields, "DynamicRules")
	}
	return o
}

func (o Events) MarshalJSON() ([]byte, error) {
	type noMethod Events
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Events) SetEvent(v *string) *Events {
	if o.Event = v; o.Event == nil {
		o.nullFields = append(o.nullFields, "Event")
	}
	return o
}

func (o *Events) SetType(v *string) *Events {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o DynamicRules) MarshalJSON() ([]byte, error) {
	type noMethod DynamicRules
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DynamicRules) SetFilterConditions(v []*FilterConditions) *DynamicRules {
	if o.FilterConditions = v; o.FilterConditions == nil {
		o.nullFields = append(o.nullFields, "FilterConditions")
	}
	return o
}

func (o FilterConditions) MarshalJSON() ([]byte, error) {
	type noMethod FilterConditions
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *FilterConditions) SetIdentifier(v *string) *FilterConditions {
	if o.Identifier = v; o.Identifier == nil {
		o.nullFields = append(o.nullFields, "Identifier")
	}
	return o
}

func (o *FilterConditions) SetOperator(v *string) *FilterConditions {
	if o.Operator = v; o.Operator == nil {
		o.nullFields = append(o.nullFields, "Operator")
	}
	return o
}

func (o *FilterConditions) SetExpression(v *string) *FilterConditions {
	if o.Expression = v; o.Expression == nil {
		o.nullFields = append(o.nullFields, "Expression")
	}
	return o
}
