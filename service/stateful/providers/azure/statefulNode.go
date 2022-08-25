package azure

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
	"io/ioutil"
	"net/http"
	"time"
)

type StatefulNode struct {
	ID                *string      `json:"id,omitempty"`
	Name              *string      `json:"name,omitempty"`
	ResourceGroupName *string      `json:"resourceGroupName,omitempty"`
	Region            *string      `json:"region,omitempty"`
	Description       *string      `json:"description,omitempty"`
	Strategy          *Strategy    `json:"strategy,omitempty"`
	Compute           *Compute     `json:"compute,omitempty"`
	Persistence       *Persistence `json:"persistence,omitempty"`
	Scheduling        *Scheduling  `json:"scheduling,omitempty"`
	Health            *Health      `json:"health,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

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

type Strategy struct {
	PreferredLifecycle  *string       `json:"preferredLifecycle,omitempty"`
	Signals             []*Signal     `json:"signals,omitempty"`
	FallbackToOnDemand  *bool         `json:"fallbackToOd,omitempty"`
	DrainingTimeout     *int          `json:"drainingTimeout,omitempty"`
	RevertToSpot        *RevertToSpot `json:"revertToSpot,omitempty"`
	OptimizationWindows []string      `json:"optimizationWindows,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Signal struct {
	Type    *string `json:"type,omitempty"`
	Timeout *int    `json:"timeout,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RevertToSpot struct {
	PerformAt *string `json:"performAt,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Compute struct {
	OS                  *string              `json:"os,omitempty"`
	VMSizes             *VMSizes             `json:"vmSizes,omitempty"`
	Zones               []string             `json:"zones,omitempty"`
	PreferredZone       *string              `json:"preferredZone,omitempty"`
	LaunchSpecification *LaunchSpecification `json:"launchSpecification,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type VMSizes struct {
	OnDemandSizes      []string `json:"odSizes,omitempty"`
	SpotSizes          []string `json:"spotSizes,omitempty"`
	PreferredSpotSizes []string `json:"preferredSpotSizes,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecification struct {
	Image                    *Image                    `json:"image,omitempty"`
	Network                  *Network                  `json:"network,omitempty"`
	Login                    *Login                    `json:"login,omitempty"`
	CustomData               *string                   `json:"customData,omitempty"`
	ShutdownScript           *string                   `json:"shutdownScript,omitempty"`
	LoadBalancersConfig      *LoadBalancersConfig      `json:"loadBalancersConfig,omitempty"`
	Tags                     []*Tag                    `json:"tags,omitempty"`
	ManagedServiceIdentities []*ManagedServiceIdentity `json:"managedServiceIdentities,omitempty"`
	Extensions               []*Extension              `json:"extensions,omitempty"`
	OSDisk                   *OSDisk                   `json:"osDisk,omitempty"`
	DataDisks                []*DataDisk               `json:"dataDisks,omitempty"`
	Secrets                  []*Secret                 `json:"secrets,omitempty"`
	BootDiagnostics          *BootDiagnostics          `json:"bootDiagnostics,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Secret struct {
	SourceVault       *SourceVault        `json:"sourceVault,omitempty"`
	VaultCertificates []*VaultCertificate `json:"vaultCertificates,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type SourceVault struct {
	Name              *string `json:"name,omitempty"`
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type VaultCertificate struct {
	CertificateURL   *string `json:"certificateUrl,omitempty"`
	CertificateStore *string `json:"certificateStore,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Extension struct {
	Name                    *string                `json:"name,omitempty"`
	Type                    *string                `json:"type,omitempty"`
	Publisher               *string                `json:"publisher,omitempty"`
	APIVersion              *string                `json:"apiVersion,omitempty"`
	MinorVersionAutoUpgrade *bool                  `json:"minorVersionAutoUpgrade,omitempty"`
	ProtectedSettings       map[string]interface{} `json:"protectedSettings,omitempty"`
	PublicSettings          map[string]interface{} `json:"publicSettings,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Tag struct {
	TagKey   *string `json:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ManagedServiceIdentity struct {
	Name              *string `json:"name,omitempty"`
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Image struct {
	MarketPlace *MarketPlaceImage `json:"marketplace,omitempty"`
	Custom      *CustomImage      `json:"custom,omitempty"`
	Gallery     *Gallery          `json:"gallery,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type MarketPlaceImage struct {
	Publisher *string `json:"publisher,omitempty"`
	Offer     *string `json:"offer,omitempty"`
	SKU       *string `json:"sku,omitempty"`
	Version   *string `json:"version,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type CustomImage struct {
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`
	Name              *string `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Gallery struct {
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`
	GalleryName       *string `json:"galleryName,omitempty"`
	ImageName         *string `json:"imageName,omitempty"`
	VersionName       *string `json:"versionName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Network struct {
	VirtualNetworkName *string             `json:"virtualNetworkName,omitempty"`
	ResourceGroupName  *string             `json:"resourceGroupName,omitempty"`
	NetworkInterfaces  []*NetworkInterface `json:"networkInterfaces,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type NetworkInterface struct {
	SubnetName                 *string                      `json:"subnetName,omitempty"`
	AssignPublicIP             *bool                        `json:"assignPublicIp,omitempty"`
	IsPrimary                  *bool                        `json:"isPrimary,omitempty"`
	PublicIPSku                *string                      `json:"publicIpSku,omitempty"`
	NetworkSecurityGroup       *NetworkSecurityGroup        `json:"networkSecurityGroup,omitempty"`
	EnableIPForwarding         *bool                        `json:"enableIPForwarding,omitempty"`
	PrivateIPAddresses         []string                     `json:"privateIPAddresses,omitempty"`
	AdditionalIPConfigurations []*AdditionalIPConfiguration `json:"additionalIpConfigurations,omitempty"`
	PublicIPs                  []*PublicIP                  `json:"publicIps,omitempty"`
	ApplicationSecurityGroups  []*ApplicationSecurityGroup  `json:"applicationSecurityGroups,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type NetworkSecurityGroup struct {
	Name              *string `json:"name,omitempty"`
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AdditionalIPConfiguration struct {
	PrivateIPAddressVersion *string `json:"privateIpAddressVersion,omitempty"`
	Name                    *string `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type PublicIP struct {
	Name              *string `json:"name,omitempty"`
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ApplicationSecurityGroup struct {
	Name              *string `json:"name,omitempty"`
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Login struct {
	UserName     *string `json:"userName,omitempty"`
	SSHPublicKey *string `json:"sshPublicKey,omitempty"`
	Password     *string `json:"password,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LoadBalancersConfig struct {
	LoadBalancers []*LoadBalancer `json:"loadBalancers,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LoadBalancer struct {
	Type              *string  `json:"type,omitempty"`
	ResourceGroupName *string  `json:"resourceGroupName,omitempty"`
	Name              *string  `json:"name,omitempty"`
	SKU               *string  `json:"sku,omitempty"`
	BackendPoolNames  []string `json:"backendPoolNames,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type OSDisk struct {
	SizeGB *int    `json:"sizeGB,omitempty"`
	Type   *string `json:"type,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DataDisk struct {
	SizeGB *int    `json:"sizeGB,omitempty"`
	LUN    *int    `json:"lun,omitempty"`
	Type   *string `json:"type,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type BootDiagnostics struct {
	IsEnabled  *bool   `json:"isEnabled,omitempty"`
	Type       *string `json:"type,omitempty"`
	StorageURL *string `json:"storageUrl,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Persistence struct {
	ShouldPersistOSDisk      *bool   `json:"shouldPersistOsDisk,omitempty"`
	OSDiskPersistenceMode    *string `json:"osDiskPersistenceMode,omitempty"`
	ShouldPersistDataDisks   *bool   `json:"shouldPersistDataDisks,omitempty"`
	DataDisksPersistenceMode *string `json:"dataDisksPersistenceMode,omitempty"`
	ShouldPersistNetwork     *bool   `json:"shouldPersistNetwork,omitempty"`
	ShouldPersistVM          *bool   `json:"shouldPersistVm,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Scheduling struct {
	Tasks []*Task `json:"tasks,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Task struct {
	IsEnabled      *bool   `json:"isEnabled,omitempty"`
	Type           *string `json:"type,omitempty"`
	CronExpression *string `json:"cronExpression,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Health struct {
	HealthCheckTypes  []string `json:"healthCheckTypes,omitempty"`
	GracePeriod       *int     `json:"gracePeriod,omitempty"`
	UnhealthyDuration *int     `json:"unhealthyDuration,omitempty"`
	AutoHealing       *bool    `json:"autoHealing,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type CreateStatefulNodeInput struct {
	StatefulNode *StatefulNode `json:"statefulNode,omitempty"`
}

type CreateStatefulNodeOutput struct {
	StatefulNode *StatefulNode `json:"statefulNode,omitempty"`
}

type ReadStatefulNodeInput struct {
	ID *string `json:"id,omitempty"`
}

type ReadStatefulNodeOutput struct {
	StatefulNode *StatefulNode `json:"statefulNode,omitempty"`
}

type UpdateStatefulNodeInput struct {
	StatefulNode *StatefulNode `json:"statefulNode,omitempty"`
}

type UpdateStatefulNodeOutput struct {
	StatefulNode *StatefulNode `json:"statefulNode,omitempty"`
}

type DeleteStatefulNodeInput struct {
	ID                 *string             `json:"id,omitempty"`
	DeallocationConfig *DeallocationConfig `json:"deallocationConfig,omitempty"`
}

type DeallocationConfig struct {
	ShouldTerminateVM          *bool                       `json:"shouldTerminateVm,omitempty"`
	NetworkDeallocationConfig  *ResourceDeallocationConfig `json:"networkDeallocationConfig,omitempty"`
	DiskDeallocationConfig     *ResourceDeallocationConfig `json:"diskDeallocationConfig,omitempty"`
	SnapshotDeallocationConfig *ResourceDeallocationConfig `json:"snapshotDeallocationConfig,omitempty"`
	PublicIPDeallocationConfig *ResourceDeallocationConfig `json:"publicIpDeallocationConfig,omitempty"`
}

type ResourceDeallocationConfig struct {
	ShouldDeallocate *bool `json:"shouldDeallocate,omitempty"`
	TTLInHours       *int  `json:"ttlInHours,omitempty"`
}

type DeleteStatefulNodeOutput struct{}

type ListStatefulNodesInput struct{}

type ListStatefulNodesOutput struct {
	StatefulNodes []*StatefulNode `json:"statefulNodes,omitempty"`
}

type UpdateStatefulNodeStateInput struct {
	ID                *string `json:"id,omitempty"`
	StatefulNodeState *string `json:"state,omitempty"`
}

type UpdateStatefulNodeStateOutput struct{}

type DetachStatefulNodeDataDiskInput struct {
	ID                        *string `json:"id,omitempty"`
	DataDiskName              *string `json:"dataDiskName,omitempty"`
	DataDiskResourceGroupName *string `json:"dataDiskResourceGroupName,omitempty"`
	ShouldDeallocate          *bool   `json:"shouldDeallocate,omitempty"`
	TTLInHours                *int    `json:"ttlInHours,omitempty"`
}

type DetachStatefulNodeDataDiskOutput struct{}

type AttachStatefulNodeDataDiskInput struct {
	ID                        *string `json:"id,omitempty"`
	DataDiskName              *string `json:"dataDiskName,omitempty"`
	DataDiskResourceGroupName *string `json:"dataDiskResourceGroupName,omitempty"`
	StorageAccountType        *string `json:"storageAccountType,omitempty"`
	SizeGB                    *int    `json:"sizeGB,omitempty"`
	LUN                       *int    `json:"lun,omitempty"`
	Zone                      *string `json:"zone,omitempty"`
}

type AttachStatefulNodeDataDiskOutput struct{}

type GetStatefulNodeStateInput struct {
	ID *string `json:"id,omitempty"`
}

type GetStatefulNodeStateOutput struct {
	StatefulNodeState *StatefulNodeState `json:"statefulNodeState,omitempty"`
}

type StatefulNodeState struct {
	ID                *string `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	Region            *string `json:"region,omitempty"`
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`
	Status            *string `json:"status,omitempty"`
	VMName            *string `json:"vmName,omitempty"`
	VMSize            *string `json:"vmSize,omitempty"`
	LifeCycle         *string `json:"lifeCycle,omitempty"`
	RollbackReason    *string `json:"rollbackReason,omitempty"`
	ErrorReason       *string `json:"errorReason,omitempty"`
	PrivateIP         *string `json:"privateIP,omitempty"`
	PublicIP          *string `json:"publicIP,omitempty"`
}

type StatefulNodeImport struct {
	StatefulImportID       *string       `json:"statefulImportId,omitempty"`
	ResourceGroupName      *string       `json:"resourceGroupName,omitempty"`
	OriginalVMName         *string       `json:"originalVmName,omitempty"`
	DrainingTimeout        *int          `json:"drainingTimeout,omitempty"`
	ResourcesRetentionTime *int          `json:"resourcesRetentionTime,omitempty"`
	StatefulNode           *StatefulNode `json:"node,omitempty"`
}

type ImportVMStatefulNodeInput struct {
	StatefulNodeImport *StatefulNodeImport `json:"statefulNodeImport,omitempty"`
}

type ImportVMStatefulNodeOutput struct {
	StatefulNodeImport *StatefulNodeImport `json:"statefulNodeImport,omitempty"`
}

// region Unmarshallers

func statefulNodeFromJSON(in []byte) (*StatefulNode, error) {
	b := new(StatefulNode)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func statefulNodesFromJSON(in []byte) ([]*StatefulNode, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*StatefulNode, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := statefulNodeFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func statefulNodesFromHttpResponse(resp *http.Response) ([]*StatefulNode, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return statefulNodesFromJSON(body)
}

func statefulNodeImportFromJSON(in []byte) (*StatefulNodeImport, error) {
	b := new(StatefulNodeImport)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func statefulNodesImportFromJSON(in []byte) ([]*StatefulNodeImport, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*StatefulNodeImport, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := statefulNodeImportFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func statefulNodesImportFromHttpResponse(resp *http.Response) ([]*StatefulNodeImport, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return statefulNodesImportFromJSON(body)
}

func statefulNodeStateFromJSON(in []byte) (*StatefulNodeState, error) {
	b := new(StatefulNodeState)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func statefulNodeStatesFromJSON(in []byte) ([]*StatefulNodeState, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*StatefulNodeState, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := statefulNodeStateFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func statefulNodeStatesFromHttpResponse(resp *http.Response) ([]*StatefulNodeState, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return statefulNodeStatesFromJSON(body)
}

// endregion

// region API Requests

func (s *ServiceOp) List(ctx context.Context, input *ListStatefulNodesInput) (*ListStatefulNodesOutput, error) {
	r := client.NewRequest(http.MethodGet, "/azure/compute/statefulNode")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := statefulNodesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListStatefulNodesOutput{StatefulNodes: gs}, nil
}

func (s *ServiceOp) Create(ctx context.Context, input *CreateStatefulNodeInput) (*CreateStatefulNodeOutput, error) {
	r := client.NewRequest(http.MethodPost, "/azure/compute/statefulNode")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := statefulNodesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateStatefulNodeOutput)
	if len(gs) > 0 {
		output.StatefulNode = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) Read(ctx context.Context, input *ReadStatefulNodeInput) (*ReadStatefulNodeOutput, error) {
	path, err := uritemplates.Expand("/azure/compute/statefulNode/{statefulNodeId}", uritemplates.Values{
		"statefulNodeId": spotinst.StringValue(input.ID),
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

	gs, err := statefulNodesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadStatefulNodeOutput)
	if len(gs) > 0 {
		output.StatefulNode = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) Update(ctx context.Context, input *UpdateStatefulNodeInput) (*UpdateStatefulNodeOutput, error) {
	path, err := uritemplates.Expand("/azure/compute/statefulNode/{statefulNodeId}", uritemplates.Values{
		"statefulNodeId": spotinst.StringValue(input.StatefulNode.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the ID anymore, so let's drop it.
	input.StatefulNode.ID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := statefulNodesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateStatefulNodeOutput)
	if len(gs) > 0 {
		output.StatefulNode = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) Delete(ctx context.Context, input *DeleteStatefulNodeInput) (*DeleteStatefulNodeOutput, error) {
	path, err := uritemplates.Expand("/azure/compute/statefulNode/{statefulNodeId}", uritemplates.Values{
		"statefulNodeId": spotinst.StringValue(input.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the ID anymore, so let's drop it.
	input.ID = nil

	r := client.NewRequest(http.MethodDelete, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteStatefulNodeOutput{}, nil
}

func (s *ServiceOp) UpdateState(ctx context.Context,
	input *UpdateStatefulNodeStateInput) (*UpdateStatefulNodeStateOutput, error) {
	path, err := uritemplates.Expand("/azure/compute/statefulNode/{statefulNodeId}/state", uritemplates.Values{
		"statefulNodeId": spotinst.StringValue(input.ID),
	})
	if err != nil {
		return nil, err
	}
	// We do NOT need the ID anymore, so let's drop it.
	input.ID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateStatefulNodeStateOutput{}, nil
}

func (s *ServiceOp) DetachDataDisk(ctx context.Context,
	input *DetachStatefulNodeDataDiskInput) (*DetachStatefulNodeDataDiskOutput, error) {
	path, err := uritemplates.Expand("/azure/compute/statefulNode/{statefulNodeId}/dataDisk/detach", uritemplates.Values{
		"statefulNodeId": spotinst.StringValue(input.ID),
	})
	if err != nil {
		return nil, err
	}
	// We do NOT need the ID anymore, so let's drop it.
	input.ID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DetachStatefulNodeDataDiskOutput{}, nil
}

func (s *ServiceOp) AttachDataDisk(ctx context.Context,
	input *AttachStatefulNodeDataDiskInput) (*AttachStatefulNodeDataDiskOutput, error) {
	path, err := uritemplates.Expand("/azure/compute/statefulNode/{statefulNodeId}/dataDisk/attach", uritemplates.Values{
		"statefulNodeId": spotinst.StringValue(input.ID),
	})
	if err != nil {
		return nil, err
	}
	// We do NOT need the ID anymore, so let's drop it.
	input.ID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &AttachStatefulNodeDataDiskOutput{}, nil
}

func (s *ServiceOp) GetState(ctx context.Context, input *GetStatefulNodeStateInput) (*GetStatefulNodeStateOutput, error) {
	path, err := uritemplates.Expand("/azure/compute/statefulNode/{statefulNodeId}/status", uritemplates.Values{
		"statefulNodeId": spotinst.StringValue(input.ID),
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

	states, err := statefulNodeStatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(GetStatefulNodeStateOutput)
	if len(states) > 0 {
		output.StatefulNodeState = states[0]
	}

	return output, nil
}

func (s *ServiceOp) ImportVM(ctx context.Context, input *ImportVMStatefulNodeInput) (*ImportVMStatefulNodeOutput, error) {
	r := client.NewRequest(http.MethodPost, "/azure/compute/statefulNode/import")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := statefulNodesImportFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ImportVMStatefulNodeOutput)
	if len(gs) > 0 {
		output.StatefulNodeImport = gs[0]
	}

	return output, nil
}

// endregion

// region statefulNode

func (o StatefulNode) MarshalJSON() ([]byte, error) {
	type noMethod StatefulNode
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *StatefulNode) SetID(v *string) *StatefulNode {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *StatefulNode) SetName(v *string) *StatefulNode {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *StatefulNode) SetResourceGroupName(v *string) *StatefulNode {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

func (o *StatefulNode) SetDescription(v *string) *StatefulNode {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *StatefulNode) SetPersistence(v *Persistence) *StatefulNode {
	if o.Persistence = v; o.Persistence == nil {
		o.nullFields = append(o.nullFields, "Persistence")
	}
	return o
}

func (o *StatefulNode) SetCompute(v *Compute) *StatefulNode {
	if o.Compute = v; o.Compute == nil {
		o.nullFields = append(o.nullFields, "Compute")
	}
	return o
}

func (o *StatefulNode) SetStrategy(v *Strategy) *StatefulNode {
	if o.Strategy = v; o.Strategy == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}

func (o *StatefulNode) SetRegion(v *string) *StatefulNode {
	if o.Region = v; o.Region == nil {
		o.nullFields = append(o.nullFields, "Region")
	}
	return o
}

func (o *StatefulNode) SetScheduling(v *Scheduling) *StatefulNode {
	if o.Scheduling = v; o.Scheduling == nil {
		o.nullFields = append(o.nullFields, "Scheduling")
	}
	return o
}

func (o *StatefulNode) SetHealth(v *Health) *StatefulNode {
	if o.Health = v; o.Health == nil {
		o.nullFields = append(o.nullFields, "Health")
	}
	return o
}

// endregion

// region Strategy

func (o Strategy) MarshalJSON() ([]byte, error) {
	type noMethod Strategy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Strategy) SetDrainingTimeout(v *int) *Strategy {
	if o.DrainingTimeout = v; o.DrainingTimeout == nil {
		o.nullFields = append(o.nullFields, "DrainingTimeout")
	}
	return o
}

func (o *Strategy) SetFallbackToOnDemand(v *bool) *Strategy {
	if o.FallbackToOnDemand = v; o.FallbackToOnDemand == nil {
		o.nullFields = append(o.nullFields, "FallbackToOnDemand")
	}
	return o
}

func (o *Strategy) SetPreferredLifecycle(v *string) *Strategy {
	if o.PreferredLifecycle = v; o.PreferredLifecycle == nil {
		o.nullFields = append(o.nullFields, "PreferredLifecycle")
	}
	return o
}

func (o *Strategy) SetSignals(v []*Signal) *Strategy {
	if o.Signals = v; o.Signals == nil {
		o.nullFields = append(o.nullFields, "Signals")
	}
	return o
}

func (o *Strategy) SetRevertToSpot(v *RevertToSpot) *Strategy {
	if o.RevertToSpot = v; o.RevertToSpot == nil {
		o.nullFields = append(o.nullFields, "RevertToSpot")
	}
	return o
}

func (o *Strategy) SetOptimizationWindows(v []string) *Strategy {
	if o.OptimizationWindows = v; o.OptimizationWindows == nil {
		o.nullFields = append(o.nullFields, "OptimizationWindows")
	}
	return o
}

// endregion

// region Signal

func (o Signal) MarshalJSON() ([]byte, error) {
	type noMethod Signal
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Signal) SetType(v *string) *Signal {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *Signal) SetTimeout(v *int) *Signal {
	if o.Timeout = v; o.Timeout == nil {
		o.nullFields = append(o.nullFields, "Timeout")
	}
	return o
}

// endregion

// region RevertToSpot

func (o RevertToSpot) MarshalJSON() ([]byte, error) {
	type noMethod RevertToSpot
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RevertToSpot) SetPerformAt(v *string) *RevertToSpot {
	if o.PerformAt = v; o.PerformAt == nil {
		o.nullFields = append(o.nullFields, "PerformAt")
	}
	return o
}

// endregion

// region Persistence

func (o *Persistence) SetShouldPersistOSDisk(v *bool) *Persistence {
	if o.ShouldPersistOSDisk = v; o.ShouldPersistOSDisk == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistOsDisk")
	}
	return o
}

func (o *Persistence) SetOSDiskPersistenceMode(v *string) *Persistence {
	if o.OSDiskPersistenceMode = v; o.OSDiskPersistenceMode == nil {
		o.nullFields = append(o.nullFields, "OsDiskPersistenceMode")
	}
	return o
}

func (o *Persistence) SetShouldPersistDataDisks(v *bool) *Persistence {
	if o.ShouldPersistDataDisks = v; o.ShouldPersistDataDisks == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistDataDisks")
	}
	return o
}

func (o *Persistence) SetDataDisksPersistenceMode(v *string) *Persistence {
	if o.DataDisksPersistenceMode = v; o.DataDisksPersistenceMode == nil {
		o.nullFields = append(o.nullFields, "DataDisksPersistenceMode")
	}
	return o
}

func (o *Persistence) SetShouldPersistNetwork(v *bool) *Persistence {
	if o.ShouldPersistNetwork = v; o.ShouldPersistNetwork == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistNetwork")
	}
	return o
}

func (o *Persistence) SetShouldPersistVM(v *bool) *Persistence {
	if o.ShouldPersistVM = v; o.ShouldPersistVM == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistVm")
	}
	return o
}

// endregion

// region Scheduling

func (o Scheduling) MarshalJSON() ([]byte, error) {
	type noMethod Scheduling
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Scheduling) SetTasks(v []*Task) *Scheduling {
	if o.Tasks = v; o.Tasks == nil {
		o.nullFields = append(o.nullFields, "Tasks")
	}
	return o
}

// endregion

// region Task

func (o Task) MarshalJSON() ([]byte, error) {
	type noMethod Task
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Task) SetIsEnabled(v *bool) *Task {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *Task) SetType(v *string) *Task {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *Task) SetCronExpression(v *string) *Task {
	if o.CronExpression = v; o.CronExpression == nil {
		o.nullFields = append(o.nullFields, "CronExpression")
	}
	return o
}

// endregion

// region Health

func (o Health) MarshalJSON() ([]byte, error) {
	type noMethod Health
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Health) SetHealthCheckTypes(v []string) *Health {
	if o.HealthCheckTypes = v; o.HealthCheckTypes == nil {
		o.nullFields = append(o.nullFields, "HealthCheckTypes")
	}
	return o
}

func (o *Health) SetGracePeriod(v *int) *Health {
	if o.GracePeriod = v; o.GracePeriod == nil {
		o.nullFields = append(o.nullFields, "GracePeriod")
	}
	return o
}

func (o *Health) SetUnhealthyDuration(v *int) *Health {
	if o.UnhealthyDuration = v; o.UnhealthyDuration == nil {
		o.nullFields = append(o.nullFields, "UnhealthyDuration")
	}
	return o
}

func (o *Health) SetAutoHealing(v *bool) *Health {
	if o.AutoHealing = v; o.AutoHealing == nil {
		o.nullFields = append(o.nullFields, "AutoHealing")
	}
	return o
}

// endregion

// region Compute

func (o Compute) MarshalJSON() ([]byte, error) {
	type noMethod Compute
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Compute) SetVMSizes(v *VMSizes) *Compute {
	if o.VMSizes = v; o.VMSizes == nil {
		o.nullFields = append(o.nullFields, "VMSizes")
	}
	return o
}

func (o *Compute) SetOS(v *string) *Compute {
	if o.OS = v; o.OS == nil {
		o.nullFields = append(o.nullFields, "OS")
	}
	return o
}

func (o *Compute) SetLaunchSpecification(v *LaunchSpecification) *Compute {
	if o.LaunchSpecification = v; o.LaunchSpecification == nil {
		o.nullFields = append(o.nullFields, "LaunchSpecification")
	}
	return o
}

func (o *Compute) SetZones(v []string) *Compute {
	if o.Zones = v; o.Zones == nil {
		o.nullFields = append(o.nullFields, "Zones")
	}
	return o
}

func (o *Compute) SetPreferredZone(v *string) *Compute {
	if o.PreferredZone = v; o.PreferredZone == nil {
		o.nullFields = append(o.nullFields, "PreferredZone")
	}
	return o
}

// endregion

// region VMSizes

func (o VMSizes) MarshalJSON() ([]byte, error) {
	type noMethod VMSizes
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *VMSizes) SetOnDemandSizes(v []string) *VMSizes {
	if o.OnDemandSizes = v; o.OnDemandSizes == nil {
		o.nullFields = append(o.nullFields, "OnDemandSizes")
	}
	return o
}

func (o *VMSizes) SetSpotSizes(v []string) *VMSizes {
	if o.SpotSizes = v; o.SpotSizes == nil {
		o.nullFields = append(o.nullFields, "SpotSizes")
	}
	return o
}

func (o *VMSizes) SetPreferredSpotSizes(v []string) *VMSizes {
	if o.PreferredSpotSizes = v; o.PreferredSpotSizes == nil {
		o.nullFields = append(o.nullFields, "PreferredSpotSizes")
	}
	return o
}

// endregion

// region LaunchSpecification

func (o LaunchSpecification) MarshalJSON() ([]byte, error) {
	type noMethod LaunchSpecification
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LaunchSpecification) SetImage(v *Image) *LaunchSpecification {
	if o.Image = v; o.Image == nil {
		o.nullFields = append(o.nullFields, "Image")
	}
	return o
}

func (o *LaunchSpecification) SetNetwork(v *Network) *LaunchSpecification {
	if o.Network = v; o.Network == nil {
		o.nullFields = append(o.nullFields, "Network")
	}
	return o
}

func (o *LaunchSpecification) SetLogin(v *Login) *LaunchSpecification {
	if o.Login = v; o.Login == nil {
		o.nullFields = append(o.nullFields, "Login")
	}
	return o
}

func (o *LaunchSpecification) SetCustomData(v *string) *LaunchSpecification {
	if o.CustomData = v; o.CustomData == nil {
		o.nullFields = append(o.nullFields, "CustomData")
	}
	return o
}

func (o *LaunchSpecification) SetLoadBalancersConfig(v *LoadBalancersConfig) *LaunchSpecification {
	if o.LoadBalancersConfig = v; o.LoadBalancersConfig == nil {
		o.nullFields = append(o.nullFields, "LoadBalancersConfig")
	}
	return o
}

func (o *LaunchSpecification) SetOSDisk(v *OSDisk) *LaunchSpecification {
	if o.OSDisk = v; o.OSDisk == nil {
		o.nullFields = append(o.nullFields, "OSDisk")
	}
	return o
}

func (o *LaunchSpecification) SetDataDisks(v []*DataDisk) *LaunchSpecification {
	if o.DataDisks = v; o.DataDisks == nil {
		o.nullFields = append(o.nullFields, "DataDisks")
	}
	return o
}

func (o *LaunchSpecification) SetBootDiagnostics(v *BootDiagnostics) *LaunchSpecification {
	if o.BootDiagnostics = v; o.BootDiagnostics == nil {
		o.nullFields = append(o.nullFields, "BootDiagnostics")
	}
	return o
}

func (o *LaunchSpecification) SetShutdownScript(v *string) *LaunchSpecification {
	if o.ShutdownScript = v; o.ShutdownScript == nil {
		o.nullFields = append(o.nullFields, "ShutdownScript")
	}
	return o
}

func (o *LaunchSpecification) SetSecrets(v []*Secret) *LaunchSpecification {
	if o.Secrets = v; o.Secrets == nil {
		o.nullFields = append(o.nullFields, "Secrets")
	}
	return o
}

func (o *LaunchSpecification) SetTags(v []*Tag) *LaunchSpecification {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

func (o *LaunchSpecification) SetManagedServiceIdentities(v []*ManagedServiceIdentity) *LaunchSpecification {
	if o.ManagedServiceIdentities = v; o.ManagedServiceIdentities == nil {
		o.nullFields = append(o.nullFields, "ManagedServiceIdentities")
	}
	return o
}

func (o *LaunchSpecification) SetExtensions(v []*Extension) *LaunchSpecification {
	if o.Extensions = v; o.Extensions == nil {
		o.nullFields = append(o.nullFields, "Extensions")
	}
	return o
}

// endregion

// region Secret

func (o Secret) MarshalJSON() ([]byte, error) {
	type noMethod Secret
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Secret) SetSourceVault(v *SourceVault) *Secret {
	if o.SourceVault = v; o.SourceVault == nil {
		o.nullFields = append(o.nullFields, "SourceVault")
	}
	return o
}

func (o *Secret) SetVaultCertificates(v []*VaultCertificate) *Secret {
	if o.VaultCertificates = v; o.VaultCertificates == nil {
		o.nullFields = append(o.nullFields, "VaultCertificates")
	}
	return o
}

// endregion

// region SourceVault

func (o SourceVault) MarshalJSON() ([]byte, error) {
	type noMethod SourceVault
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *SourceVault) SetName(v *string) *SourceVault {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *SourceVault) SetResourceGroupName(v *string) *SourceVault {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

// endregion

// region VaultCertificates

func (o VaultCertificate) MarshalJSON() ([]byte, error) {
	type noMethod VaultCertificate
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *VaultCertificate) SetCertificateURL(v *string) *VaultCertificate {
	if o.CertificateURL = v; o.CertificateURL == nil {
		o.nullFields = append(o.nullFields, "CertificateUrl")
	}
	return o
}

func (o *VaultCertificate) SetCertificateStore(v *string) *VaultCertificate {
	if o.CertificateStore = v; o.CertificateStore == nil {
		o.nullFields = append(o.nullFields, "CertificateStore")
	}
	return o
}

// endregion

// region Extension

func (o Extension) MarshalJSON() ([]byte, error) {
	type noMethod Extension
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Extension) SetName(v *string) *Extension {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Extension) SetType(v *string) *Extension {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *Extension) SetPublisher(v *string) *Extension {
	if o.Publisher = v; o.Publisher == nil {
		o.nullFields = append(o.nullFields, "Publisher")
	}
	return o
}

func (o *Extension) SetAPIVersion(v *string) *Extension {
	if o.APIVersion = v; o.APIVersion == nil {
		o.nullFields = append(o.nullFields, "APIVersion")
	}
	return o
}

func (o *Extension) SetMinorVersionAutoUpgrade(v *bool) *Extension {
	if o.MinorVersionAutoUpgrade = v; o.MinorVersionAutoUpgrade == nil {
		o.nullFields = append(o.nullFields, "MinorVersionAutoUpgrade")
	}
	return o
}

func (o *Extension) SetProtectedSettings(v map[string]interface{}) *Extension {
	if o.ProtectedSettings = v; o.ProtectedSettings == nil {
		o.nullFields = append(o.nullFields, "ProtectedSettings")
	}
	return o
}

func (o *Extension) SetPublicSettings(v map[string]interface{}) *Extension {
	if o.PublicSettings = v; o.PublicSettings == nil {
		o.nullFields = append(o.nullFields, "PublicSettings")
	}
	return o
}

// endregion

// region Tag

func (o Tag) MarshalJSON() ([]byte, error) {
	type noMethod Tag
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Tag) SetTagKey(v *string) *Tag {
	if o.TagKey = v; o.TagKey == nil {
		o.nullFields = append(o.nullFields, "TagKey")
	}
	return o
}

func (o *Tag) SetTagValue(v *string) *Tag {
	if o.TagValue = v; o.TagValue == nil {
		o.nullFields = append(o.nullFields, "TagValue")
	}
	return o
}

// endregion

// region ManagedServiceIdentity

func (o ManagedServiceIdentity) MarshalJSON() ([]byte, error) {
	type noMethod ManagedServiceIdentity
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ManagedServiceIdentity) SetName(v *string) *ManagedServiceIdentity {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *ManagedServiceIdentity) SetResourceGroupName(v *string) *ManagedServiceIdentity {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

// endregion

// region Image

func (o Image) MarshalJSON() ([]byte, error) {
	type noMethod Image
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Image) SetMarketPlaceImage(v *MarketPlaceImage) *Image {
	if o.MarketPlace = v; o.MarketPlace == nil {
		o.nullFields = append(o.nullFields, "MarketPlace")
	}
	return o
}

func (o *Image) SetCustom(v *CustomImage) *Image {
	if o.Custom = v; o.Custom == nil {
		o.nullFields = append(o.nullFields, "Custom")
	}
	return o
}

func (o *Image) SetGallery(v *Gallery) *Image {
	if o.Gallery = v; o.Gallery == nil {
		o.nullFields = append(o.nullFields, "Gallery")
	}
	return o
}

// endregion

// region MarketPlaceImage

func (o MarketPlaceImage) MarshalJSON() ([]byte, error) {
	type noMethod MarketPlaceImage
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *MarketPlaceImage) SetPublisher(v *string) *MarketPlaceImage {
	if o.Publisher = v; o.Publisher == nil {
		o.nullFields = append(o.nullFields, "Publisher")
	}
	return o
}

func (o *MarketPlaceImage) SetOffer(v *string) *MarketPlaceImage {
	if o.Offer = v; o.Offer == nil {
		o.nullFields = append(o.nullFields, "Offer")
	}
	return o
}

func (o *MarketPlaceImage) SetSKU(v *string) *MarketPlaceImage {
	if o.SKU = v; o.SKU == nil {
		o.nullFields = append(o.nullFields, "SKU")
	}
	return o
}

func (o *MarketPlaceImage) SetVersion(v *string) *MarketPlaceImage {
	if o.Version = v; o.Version == nil {
		o.nullFields = append(o.nullFields, "Version")
	}
	return o
}

// endregion

// region CustomImage

func (o CustomImage) MarshalJSON() ([]byte, error) {
	type noMethod CustomImage
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *CustomImage) SetResourceGroupName(v *string) *CustomImage {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

func (o *CustomImage) SetName(v *string) *CustomImage {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

// endregion

// region Gallery

func (o Gallery) MarshalJSON() ([]byte, error) {
	type noMethod Gallery
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Gallery) SetResourceGroupName(v *string) *Gallery {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

func (o *Gallery) SetGalleryName(v *string) *Gallery {
	if o.GalleryName = v; o.GalleryName == nil {
		o.nullFields = append(o.nullFields, "GalleryName")
	}
	return o
}

func (o *Gallery) SetImageName(v *string) *Gallery {
	if o.ImageName = v; o.ImageName == nil {
		o.nullFields = append(o.nullFields, "ImageName")
	}
	return o
}

func (o *Gallery) SetVersionName(v *string) *Gallery {
	if o.VersionName = v; o.VersionName == nil {
		o.nullFields = append(o.nullFields, "VersionName")
	}
	return o
}

// endregion

// region Network

func (o Network) MarshalJSON() ([]byte, error) {
	type noMethod Network
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Network) SetVirtualNetworkName(v *string) *Network {
	if o.VirtualNetworkName = v; o.VirtualNetworkName == nil {
		o.nullFields = append(o.nullFields, "VirtualNetworkName")
	}
	return o
}

func (o *Network) SetResourceGroupName(v *string) *Network {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

func (o *Network) SetNetworkInterfaces(v []*NetworkInterface) *Network {
	if o.NetworkInterfaces = v; o.NetworkInterfaces == nil {
		o.nullFields = append(o.nullFields, "NetworkInterfaces")
	}
	return o
}

// endregion

// region NetworkInterface

func (o NetworkInterface) MarshalJSON() ([]byte, error) {
	type noMethod NetworkInterface
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *NetworkInterface) SetSubnetName(v *string) *NetworkInterface {
	if o.SubnetName = v; o.SubnetName == nil {
		o.nullFields = append(o.nullFields, "SubnetName")
	}
	return o
}

func (o *NetworkInterface) SetAssignPublicIP(v *bool) *NetworkInterface {
	if o.AssignPublicIP = v; o.AssignPublicIP == nil {
		o.nullFields = append(o.nullFields, "AssignPublicIp")
	}
	return o
}

func (o *NetworkInterface) SetIsPrimary(v *bool) *NetworkInterface {
	if o.IsPrimary = v; o.IsPrimary == nil {
		o.nullFields = append(o.nullFields, "IsPrimary")
	}
	return o
}

func (o *NetworkInterface) SetPublicIPSku(v *string) *NetworkInterface {
	if o.PublicIPSku = v; o.PublicIPSku == nil {
		o.nullFields = append(o.nullFields, "PublicIpSku")
	}
	return o
}

func (o *NetworkInterface) SetNetworkSecurityGroup(v *NetworkSecurityGroup) *NetworkInterface {
	if o.NetworkSecurityGroup = v; o.NetworkSecurityGroup == nil {
		o.nullFields = append(o.nullFields, "NetworkSecurityGroup")
	}
	return o
}

func (o *NetworkInterface) SetEnableIPForwarding(v *bool) *NetworkInterface {
	if o.EnableIPForwarding = v; o.EnableIPForwarding == nil {
		o.nullFields = append(o.nullFields, "EnableIPForwarding")
	}
	return o
}

func (o *NetworkInterface) SetPrivateIPAddresses(v []string) *NetworkInterface {
	if o.PrivateIPAddresses = v; o.PrivateIPAddresses == nil {
		o.nullFields = append(o.nullFields, "PrivateIPAddresses")
	}
	return o
}

func (o *NetworkInterface) SetAdditionalIPConfigurations(v []*AdditionalIPConfiguration) *NetworkInterface {
	if o.AdditionalIPConfigurations = v; o.AdditionalIPConfigurations == nil {
		o.nullFields = append(o.nullFields, "AdditionalIpConfigurations")
	}
	return o
}

func (o *NetworkInterface) SetPublicIPs(v []*PublicIP) *NetworkInterface {
	if o.PublicIPs = v; o.PublicIPs == nil {
		o.nullFields = append(o.nullFields, "PublicIPs")
	}
	return o
}

func (o *NetworkInterface) SetApplicationSecurityGroups(v []*ApplicationSecurityGroup) *NetworkInterface {
	if o.ApplicationSecurityGroups = v; o.ApplicationSecurityGroups == nil {
		o.nullFields = append(o.nullFields, "ApplicationSecurityGroups")
	}
	return o
}

// endregion

// region NetworkSecurityGroup

func (o NetworkSecurityGroup) MarshalJSON() ([]byte, error) {
	type noMethod NetworkSecurityGroup
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *NetworkSecurityGroup) SetName(v *string) *NetworkSecurityGroup {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *NetworkSecurityGroup) SetResourceGroupName(v *string) *NetworkSecurityGroup {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

// endregion

// region AdditionalIpConfiguration

func (o AdditionalIPConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod AdditionalIPConfiguration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AdditionalIPConfiguration) SetName(v *string) *AdditionalIPConfiguration {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AdditionalIPConfiguration) SetPrivateIPAddressVersion(v *string) *AdditionalIPConfiguration {
	if o.PrivateIPAddressVersion = v; o.PrivateIPAddressVersion == nil {
		o.nullFields = append(o.nullFields, "PrivateIpAddressVersion")
	}
	return o
}

// endregion

// region PublicIP

func (o PublicIP) MarshalJSON() ([]byte, error) {
	type noMethod PublicIP
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *PublicIP) SetName(v *string) *PublicIP {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *PublicIP) SetResourceGroupName(v *string) *PublicIP {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

// endregion

// region ApplicationSecurityGroup

func (o ApplicationSecurityGroup) MarshalJSON() ([]byte, error) {
	type noMethod ApplicationSecurityGroup
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ApplicationSecurityGroup) SetName(v *string) *ApplicationSecurityGroup {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *ApplicationSecurityGroup) SetResourceGroupName(v *string) *ApplicationSecurityGroup {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

// endregion

// region Login

func (o Login) MarshalJSON() ([]byte, error) {
	type noMethod Login
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Login) SetUserName(v *string) *Login {
	if o.UserName = v; o.UserName == nil {
		o.nullFields = append(o.nullFields, "UserName")
	}
	return o
}

func (o *Login) SetSSHPublicKey(v *string) *Login {
	if o.SSHPublicKey = v; o.SSHPublicKey == nil {
		o.nullFields = append(o.nullFields, "SSHPublicKey")
	}
	return o
}

func (o *Login) SetPassword(v *string) *Login {
	if o.Password = v; o.Password == nil {
		o.nullFields = append(o.nullFields, "Password")
	}
	return o
}

// endregion

// region LoadBalancersConfig

func (o LoadBalancersConfig) MarshalJSON() ([]byte, error) {
	type noMethod LoadBalancersConfig
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LoadBalancersConfig) SetLoadBalancers(v []*LoadBalancer) *LoadBalancersConfig {
	if o.LoadBalancers = v; o.LoadBalancers == nil {
		o.nullFields = append(o.nullFields, "LoadBalancers")
	}
	return o
}

// endregion

// region LoadBalancer

func (o LoadBalancer) MarshalJSON() ([]byte, error) {
	type noMethod LoadBalancer
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LoadBalancer) SetType(v *string) *LoadBalancer {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *LoadBalancer) SetResourceGroupName(v *string) *LoadBalancer {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

func (o *LoadBalancer) SetName(v *string) *LoadBalancer {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *LoadBalancer) SetSKU(v *string) *LoadBalancer {
	if o.SKU = v; o.SKU == nil {
		o.nullFields = append(o.nullFields, "SKU")
	}
	return o
}

func (o *LoadBalancer) SeBackendPoolNames(v []string) *LoadBalancer {
	if o.BackendPoolNames = v; o.BackendPoolNames == nil {
		o.nullFields = append(o.nullFields, "BackendPoolNames")
	}
	return o
}

// endregion

// region OSDisk

func (o OSDisk) MarshalJSON() ([]byte, error) {
	type noMethod OSDisk
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *OSDisk) SetSizeGB(v *int) *OSDisk {
	if o.SizeGB = v; o.SizeGB == nil {
		o.nullFields = append(o.nullFields, "SizeGB")
	}
	return o
}

func (o *OSDisk) SetType(v *string) *OSDisk {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

// endregion

// region DataDisks

func (o DataDisk) MarshalJSON() ([]byte, error) {
	type noMethod DataDisk
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DataDisk) SetSizeGB(v *int) *DataDisk {
	if o.SizeGB = v; o.SizeGB == nil {
		o.nullFields = append(o.nullFields, "SizeGB")
	}
	return o
}

func (o *DataDisk) SetLUN(v *int) *DataDisk {
	if o.LUN = v; o.LUN == nil {
		o.nullFields = append(o.nullFields, "Lun")
	}
	return o
}

func (o *DataDisk) SetType(v *string) *DataDisk {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

// endregion

// region BootDiagnostics

func (o BootDiagnostics) MarshalJSON() ([]byte, error) {
	type noMethod BootDiagnostics
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *BootDiagnostics) SetIsEnabled(v *bool) *BootDiagnostics {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *BootDiagnostics) SetType(v *string) *BootDiagnostics {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *BootDiagnostics) SetStorageURL(v *string) *BootDiagnostics {
	if o.StorageURL = v; o.StorageURL == nil {
		o.nullFields = append(o.nullFields, "StorageUrl")
	}
	return o
}

// endregion
