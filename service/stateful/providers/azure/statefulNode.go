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
	Strategy          *Strategy    `json:"strategy,omitempty"`
	Compute           *Compute     `json:"compute,omitempty"`
	Persistence       *Persistence `json:"persistence,omitempty"`

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
	DrainingTimeout    *int  `json:"drainingTimeout,omitempty"`
	FallbackToOnDemand *bool `json:"fallbackToOd,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Compute struct {
	OS                  *string              `json:"os,omitempty"`
	VMSizes             *VMSizes             `json:"vmSizes,omitempty"`
	LaunchSpecification *LaunchSpecification `json:"launchSpecification,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type VMSizes struct {
	OnDemandSizes []string `json:"odSizes,omitempty"`
	SpotSizes     []string `json:"spotSizes,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecification struct {
	Image               *Image               `json:"image,omitempty"`
	Network             *Network             `json:"network,omitempty"`
	Login               *Login               `json:"login,omitempty"`
	CustomData          *string              `json:"customData,omitempty"`
	LoadBalancersConfig *LoadBalancersConfig `json:"loadBalancersConfig,omitempty"`
	OSDisk              *OSDisk              `json:"osDisk,omitempty"`
	DataDisks           *DataDisks           `json:"dataDisks,omitempty"`
	BootDiagnostics     *BootDiagnostics     `json:"bootDiagnostics,omitempty"`

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
	SubnetName     *string `json:"subnetName,omitempty"`
	AssignPublicIP *bool   `json:"assignPublicIp,omitempty"`
	IsPrimary      *bool   `json:"isPrimary,omitempty"`
	PublicIPSku    *string `json:"publicIpSku,omitempty"`

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

type DataDisks struct {
	SizeGB *int    `json:"sizeGB,omitempty"`
	Lun    *int    `json:"lun,omitempty"`
	Type   *string `json:"type,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type BootDiagnostics struct {
	IsEnabled  *bool   `json:"isEnabled,omitempty"`
	Type       *string `json:"type,omitempty"`
	StorageUrl *string `json:"storageUrl,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Persistence struct {
	ShouldPersistOsDisk     *bool   `json:"shouldPersistOsDisk,omitempty"`
	OsDiskPersistenceMode   *string `json:"osDiskPersistenceMode"`
	ShouldPersistDataDisk   *bool   `json:"shouldPersistDataDisk,omitempty"`
	DataDiskPersistenceMode *string `json:"dataDiskPersistenceMode"`
	ShouldPersistNetwork    *bool   `json:"shouldPersistNetwork"`
	ShouldPersistVm         *bool   `json:"shouldPersistVm"`

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
	ID *string `json:"id,omitempty"`
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
	DataDiskResourceGroupName *string `json:"dataDiskResourceGroupName"`
	ShouldDeallocate          *bool   `json:"shouldDeallocate"`
}

type DetachStatefulNodeDataDiskOutput struct{}

type AttachStatefulNodeDataDiskInput struct {
	ID                        *string `json:"id,omitempty"`
	DataDiskName              *string `json:"dataDiskName,omitempty"`
	DataDiskResourceGroupName *string `json:"dataDiskResourceGroupName"`
	SizeGB                    *int    `json:"sizeGB,omitempty"`
}

type AttachStatefulNodeDataDiskOutput struct{}

type StatefulNodeImport struct {
	ResourceGroupName *string       `json:"resourceGroupName,omitempty"`
	OriginalVMName    *string       `json:"originalVmName"`
	StatefulNode      *StatefulNode `json:"node,omitempty"`
}

type ImportVMStatefulNodeInput struct {
	StatefulNodeImport *StatefulNodeImport `json:"statefulNodeImport"`
}

type ImportVMStatefulNodeOutput struct {
	StatefulNode *StatefulNode `json:"statefulNode,omitempty"`
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

	r := client.NewRequest(http.MethodDelete, path)
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

func (s *ServiceOp) ImportVM(ctx context.Context, input *ImportVMStatefulNodeInput) (*ImportVMStatefulNodeOutput, error) {
	r := client.NewRequest(http.MethodPost, "/azure/compute/statefulNode/import")
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

	output := new(ImportVMStatefulNodeOutput)
	if len(gs) > 0 {
		output.StatefulNode = gs[0]
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

func (o *StatefulNode) SetCapacity(v *Persistence) *StatefulNode {
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

// endregion

// region Persistence

func (o *Persistence) SetShouldPersistOsDisk(v *bool) *Persistence {
	if o.ShouldPersistOsDisk = v; o.ShouldPersistOsDisk == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistOsDisk")
	}
	return o
}

func (o *Persistence) SetOsDiskPersistenceMode(v *string) *Persistence {
	if o.OsDiskPersistenceMode = v; o.OsDiskPersistenceMode == nil {
		o.nullFields = append(o.nullFields, "OsDiskPersistenceMode")
	}
	return o
}

func (o *Persistence) SetShouldPersistDataDisk(v *bool) *Persistence {
	if o.ShouldPersistDataDisk = v; o.ShouldPersistDataDisk == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistDataDisk")
	}
	return o
}

func (o *Persistence) SetDataDiskPersistenceMode(v *string) *Persistence {
	if o.DataDiskPersistenceMode = v; o.DataDiskPersistenceMode == nil {
		o.nullFields = append(o.nullFields, "DataDiskPersistenceMode")
	}
	return o
}

func (o *Persistence) SetShouldPersistNetwork(v *bool) *Persistence {
	if o.ShouldPersistNetwork = v; o.ShouldPersistNetwork == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistNetwork")
	}
	return o
}

func (o *Persistence) SetShouldPersistVm(v *bool) *Persistence {
	if o.ShouldPersistVm = v; o.ShouldPersistVm == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistVm")
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

func (o *LaunchSpecification) SetDataDisks(v *DataDisks) *LaunchSpecification {
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

func (o DataDisks) MarshalJSON() ([]byte, error) {
	type noMethod DataDisks
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DataDisks) SetSizeGB(v *int) *DataDisks {
	if o.SizeGB = v; o.SizeGB == nil {
		o.nullFields = append(o.nullFields, "SizeGB")
	}
	return o
}

func (o *DataDisks) SetLun(v *int) *DataDisks {
	if o.Lun = v; o.Lun == nil {
		o.nullFields = append(o.nullFields, "Lun")
	}
	return o
}

func (o *DataDisks) SetType(v *string) *DataDisks {
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

func (o *BootDiagnostics) SetStorageUrl(v *string) *BootDiagnostics {
	if o.StorageUrl = v; o.StorageUrl == nil {
		o.nullFields = append(o.nullFields, "StorageUrl")
	}
	return o
}

// endregion
